package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/drand/drand/chain"
	"github.com/drand/drand/chain/beacon"
	"github.com/drand/drand/crypto"
	"github.com/drand/drand/net"
	"github.com/drand/drand/protobuf/drand"
)

// PartialBeacon receives a beacon generation request and answers
// with the partial signature from this drand node.
func (bp *BeaconProcess) PartialBeacon(c context.Context, in *drand.PartialBeaconPacket) (*drand.Empty, error) {
	bp.state.RLock()
	// we need to defer unlock here to avoid races during the partial processing
	defer bp.state.RUnlock()
	inst := bp.beacon
	if inst == nil || len(bp.chainHash) == 0 {
		return nil, errors.New("DKG not finished yet")
	}

	_, err := inst.ProcessPartialBeacon(c, in)
	return &drand.Empty{Metadata: bp.newMetadata()}, err
}

// PublicRand returns a public random beacon according to the request. If the Round
// field is 0, then it returns the last one generated.
func (bp *BeaconProcess) PublicRand(ctx context.Context, in *drand.PublicRandRequest) (*drand.PublicRandResponse, error) {
	var addr = net.RemoteAddress(ctx)

	bp.state.RLock()
	defer bp.state.RUnlock()

	if bp.beacon == nil || len(bp.chainHash) == 0 {
		return nil, errors.New("drand: beacon generation not started yet")
	}
	var beaconResp *chain.Beacon
	var err error
	if in.GetRound() == 0 {
		beaconResp, err = bp.beacon.Store().Last(ctx)
	} else {
		// fetch the correct entry or the next one if not found
		beaconResp, err = bp.beacon.Store().Get(ctx, in.GetRound())
	}
	if err != nil || beaconResp == nil {
		bp.log.Debugw("", "public_rand", "unstored_beacon", "round", in.GetRound(), "from", addr)
		return nil, fmt.Errorf("can't retrieve beacon: %w", err)
	}
	bp.log.Debugw("", "public_rand", addr, "round", beaconResp.Round, "reply", beaconResp.String())

	response := beaconToProto(beaconResp)
	response.Metadata = bp.newMetadata()

	return response, nil
}

// a proxy type so public streaming request can use the same logic as in priate
// / protocol syncing request, even though the types differ, so it prevents
// changing the protobuf structs.
type proxyRequest struct {
	*drand.PublicRandRequest
}

func (p *proxyRequest) GetFromRound() uint64 {
	return p.PublicRandRequest.GetRound()
}

type proxyStream struct {
	drand.Public_PublicRandStreamServer
}

func (p *proxyStream) Send(b *drand.BeaconPacket) error {
	return p.Public_PublicRandStreamServer.Send(&drand.PublicRandResponse{
		Round:             b.Round,
		Signature:         b.Signature,
		PreviousSignature: b.PreviousSignature,
		Randomness:        crypto.RandomnessFromSignature(b.Signature),
		Metadata:          b.Metadata,
	})
}

// PublicRandStream exports a stream of new beacons as they are generated over gRPC
func (bp *BeaconProcess) PublicRandStream(req *drand.PublicRandRequest, stream drand.Public_PublicRandStreamServer) error {
	bp.state.RLock()
	if bp.beacon == nil || len(bp.chainHash) == 0 {
		bp.state.RUnlock()
		return errors.New("beacon has not started on this node yet")
	}
	bp.state.RUnlock()

	store := bp.beacon.Store()
	proxyReq := &proxyRequest{
		req,
	}
	// make sure we have the correct metadata
	proxyReq.Metadata = bp.newMetadata()
	proxyStr := &proxyStream{stream}
	return beacon.SyncChain(bp.log.Named("PublicRand"), store, proxyReq, proxyStr)
}

// Home provides the address the local node is listening
func (bp *BeaconProcess) Home(c context.Context, _ *drand.HomeRequest) (*drand.HomeResponse, error) {
	bp.log.With("module", "public").Infow("", "home", net.RemoteAddress(c))

	return &drand.HomeResponse{
		Status: fmt.Sprintf("drand up and running on %s",
			bp.priv.Public.Address()),
		Metadata: bp.newMetadata(),
	}, nil
}

// ChainInfo replies with the chain information this node participates to
func (bp *BeaconProcess) ChainInfo(context.Context, *drand.ChainInfoRequest) (*drand.ChainInfoPacket, error) {
	bp.state.RLock()
	group := bp.group
	chainHash := bp.chainHash
	bp.state.RUnlock()
	if group == nil || len(chainHash) == 0 {
		return nil, errors.New("no dkg group setup yet")
	}

	response := chain.NewChainInfoWithLogger(bp.log, group).ToProto(bp.newMetadata())

	return response, nil
}

// SyncChain is an inter-node protocol that replies to a syncing request from a
// given round
func (bp *BeaconProcess) SyncChain(req *drand.SyncRequest, stream drand.Protocol_SyncChainServer) error {
	bp.state.RLock()
	logger := bp.log.Named("SyncChain")
	b := bp.beacon
	c := bp.chainHash
	if b == nil || len(c) == 0 {
		logger.Errorw("Received a SyncRequest, but no beacon handler is set yet", "request", req)
		bp.state.RUnlock()
		return fmt.Errorf("no beacon handler available")
	}
	store := b.Store()
	// we cannot just defer Unlock because beacon.SyncChain can run for a long time
	bp.state.RUnlock()

	return beacon.SyncChain(logger, store, req, stream)
}

// GetIdentity returns the identity of this drand node
func (bp *BeaconProcess) GetIdentity(context.Context, *drand.IdentityRequest) (*drand.IdentityResponse, error) {
	i := bp.priv.Public.ToProto()

	response := &drand.IdentityResponse{
		Address:    i.Address,
		Key:        i.Key,
		Tls:        i.Tls,
		Signature:  i.Signature,
		Metadata:   bp.newMetadata(),
		SchemeName: bp.priv.Scheme().String(),
	}
	return response, nil
}
