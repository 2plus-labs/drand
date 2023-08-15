//
// This protobuf file contains the services and message definitions of all
// methods used by drand nodes to produce distributed randomness.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: drand/protocol.proto

package drand

import (
	common "github.com/drand/drand/protobuf/common"
	dkg "github.com/drand/drand/protobuf/crypto/dkg"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *IdentityRequest) Reset() {
	*x = IdentityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityRequest) ProtoMessage() {}

func (x *IdentityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityRequest.ProtoReflect.Descriptor instead.
func (*IdentityRequest) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type IdentityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key     []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Tls     bool   `protobuf:"varint,3,opt,name=tls,proto3" json:"tls,omitempty"`
	// BLS signature over the identity to prove possession of the private key
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	// --------------
	Metadata *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// We need to specify the scheme name to make sure the key is getting probably decoded on the client side
	SchemeName string `protobuf:"bytes,6,opt,name=schemeName,proto3" json:"schemeName,omitempty"`
}

func (x *IdentityResponse) Reset() {
	*x = IdentityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityResponse) ProtoMessage() {}

func (x *IdentityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityResponse.ProtoReflect.Descriptor instead.
func (*IdentityResponse) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IdentityResponse) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *IdentityResponse) GetTls() bool {
	if x != nil {
		return x.Tls
	}
	return false
}

func (x *IdentityResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *IdentityResponse) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *IdentityResponse) GetSchemeName() string {
	if x != nil {
		return x.SchemeName
	}
	return ""
}

// SignalDKGPacket is the packet nodes send to a coordinator that collects all
// keys and setups the group and sends them back to the nodes such that they can
// start the DKG automatically.
type SignalDKGPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node        *Identity `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	SecretProof []byte    `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// In resharing cases, previous_group_hash is the hash of the previous group.
	// It is to make sure the nodes build on top of the correct previous group.
	PreviousGroupHash []byte           `protobuf:"bytes,3,opt,name=previous_group_hash,json=previousGroupHash,proto3" json:"previous_group_hash,omitempty"`
	Metadata          *common.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SignalDKGPacket) Reset() {
	*x = SignalDKGPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalDKGPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalDKGPacket) ProtoMessage() {}

func (x *SignalDKGPacket) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalDKGPacket.ProtoReflect.Descriptor instead.
func (*SignalDKGPacket) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{2}
}

func (x *SignalDKGPacket) GetNode() *Identity {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *SignalDKGPacket) GetSecretProof() []byte {
	if x != nil {
		return x.SecretProof
	}
	return nil
}

func (x *SignalDKGPacket) GetPreviousGroupHash() []byte {
	if x != nil {
		return x.PreviousGroupHash
	}
	return nil
}

func (x *SignalDKGPacket) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// PushDKGInfor is the packet the coordinator sends that contains the group over
// which to run the DKG on, the secret proof (to prove it's he's part of the
// expected group, and it's not a random packet) and as well the time at which
// every node should start the DKG.
type DKGInfoPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewGroup    *GroupPacket `protobuf:"bytes,1,opt,name=new_group,json=newGroup,proto3" json:"new_group,omitempty"`
	SecretProof []byte       `protobuf:"bytes,2,opt,name=secret_proof,json=secretProof,proto3" json:"secret_proof,omitempty"`
	// timeout in seconds
	DkgTimeout uint32 `protobuf:"varint,3,opt,name=dkg_timeout,json=dkgTimeout,proto3" json:"dkg_timeout,omitempty"`
	// signature from the coordinator to prove he is the one sending that group
	// file.
	Signature []byte           `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Metadata  *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DKGInfoPacket) Reset() {
	*x = DKGInfoPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DKGInfoPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DKGInfoPacket) ProtoMessage() {}

func (x *DKGInfoPacket) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DKGInfoPacket.ProtoReflect.Descriptor instead.
func (*DKGInfoPacket) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{3}
}

func (x *DKGInfoPacket) GetNewGroup() *GroupPacket {
	if x != nil {
		return x.NewGroup
	}
	return nil
}

func (x *DKGInfoPacket) GetSecretProof() []byte {
	if x != nil {
		return x.SecretProof
	}
	return nil
}

func (x *DKGInfoPacket) GetDkgTimeout() uint32 {
	if x != nil {
		return x.DkgTimeout
	}
	return 0
}

func (x *DKGInfoPacket) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *DKGInfoPacket) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type PartialBeaconPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Round is the round for which the beacon will be created from the partial
	// signatures
	Round uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	// signature of the previous round - could be removed at some point but now
	// is used to verify the signature even before accessing the store
	PreviousSignature []byte `protobuf:"bytes,2,opt,name=previous_signature,json=previousSignature,proto3" json:"previous_signature,omitempty"`
	// partial signature - a threshold of them needs to be aggregated to produce
	// the final beacon at the given round.
	PartialSig []byte `protobuf:"bytes,3,opt,name=partial_sig,json=partialSig,proto3" json:"partial_sig,omitempty"`
	// message is msg use to sign by multi-party in case beacon no period
	Message  []byte           `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Metadata *common.Metadata `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *PartialBeaconPacket) Reset() {
	*x = PartialBeaconPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartialBeaconPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartialBeaconPacket) ProtoMessage() {}

func (x *PartialBeaconPacket) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartialBeaconPacket.ProtoReflect.Descriptor instead.
func (*PartialBeaconPacket) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{4}
}

func (x *PartialBeaconPacket) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *PartialBeaconPacket) GetPreviousSignature() []byte {
	if x != nil {
		return x.PreviousSignature
	}
	return nil
}

func (x *PartialBeaconPacket) GetPartialSig() []byte {
	if x != nil {
		return x.PartialSig
	}
	return nil
}

func (x *PartialBeaconPacket) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PartialBeaconPacket) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// DKGPacket is the packet that nodes send to others nodes as part of the
// broadcasting protocol.
type DKGPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dkg      *dkg.Packet      `protobuf:"bytes,1,opt,name=dkg,proto3" json:"dkg,omitempty"`
	Metadata *common.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *DKGPacket) Reset() {
	*x = DKGPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DKGPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DKGPacket) ProtoMessage() {}

func (x *DKGPacket) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DKGPacket.ProtoReflect.Descriptor instead.
func (*DKGPacket) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{5}
}

func (x *DKGPacket) GetDkg() *dkg.Packet {
	if x != nil {
		return x.Dkg
	}
	return nil
}

func (x *DKGPacket) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// SyncRequest is from a node that needs to sync up with the current head of the
// chain
type SyncRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromRound uint64           `protobuf:"varint,1,opt,name=from_round,json=fromRound,proto3" json:"from_round,omitempty"`
	Metadata  *common.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SyncRequest) Reset() {
	*x = SyncRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRequest) ProtoMessage() {}

func (x *SyncRequest) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRequest.ProtoReflect.Descriptor instead.
func (*SyncRequest) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{6}
}

func (x *SyncRequest) GetFromRound() uint64 {
	if x != nil {
		return x.FromRound
	}
	return 0
}

func (x *SyncRequest) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type BeaconPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousSignature []byte           `protobuf:"bytes,1,opt,name=previous_signature,json=previousSignature,proto3" json:"previous_signature,omitempty"`
	Round             uint64           `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Signature         []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Metadata          *common.Metadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *BeaconPacket) Reset() {
	*x = BeaconPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drand_protocol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconPacket) ProtoMessage() {}

func (x *BeaconPacket) ProtoReflect() protoreflect.Message {
	mi := &file_drand_protocol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconPacket.ProtoReflect.Descriptor instead.
func (*BeaconPacket) Descriptor() ([]byte, []int) {
	return file_drand_protocol_proto_rawDescGZIP(), []int{7}
}

func (x *BeaconPacket) GetPreviousSignature() []byte {
	if x != nil {
		return x.PreviousSignature
	}
	return nil
}

func (x *BeaconPacket) GetRound() uint64 {
	if x != nil {
		return x.Round
	}
	return 0
}

func (x *BeaconPacket) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BeaconPacket) GetMetadata() *common.Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_drand_protocol_proto protoreflect.FileDescriptor

var file_drand_protocol_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x14, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x64, 0x6b, 0x67, 0x2f, 0x64, 0x6b, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0f,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xbc, 0x01,
	0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x74, 0x6c, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb7, 0x01, 0x0a,
	0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x4b, 0x47, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x23, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x76,
	0x69, 0x6f, 0x75, 0x73, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd0, 0x01, 0x0a, 0x0d, 0x44, 0x4b, 0x47, 0x49, 0x6e,
	0x66, 0x6f, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x72,
	0x61, 0x6e, 0x64, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x08, 0x6e, 0x65, 0x77, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x6b, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x64, 0x6b, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc3, 0x01, 0x0a, 0x13, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69,
	0x6f, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x58, 0x0a, 0x09, 0x44, 0x4b, 0x47, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x03,
	0x64, 0x6b, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x6b, 0x67, 0x2e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x03, 0x64, 0x6b, 0x67, 0x12, 0x2c, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5a, 0x0a, 0x0b, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x9f, 0x01, 0x0a, 0x0c, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x97, 0x03, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x72,
	0x61, 0x6e, 0x64, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x14, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x4b,
	0x47, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x64,
	0x72, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x44, 0x4b, 0x47, 0x50, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x31, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x44, 0x4b, 0x47, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x4b, 0x47, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x44, 0x4b, 0x47, 0x12, 0x10, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x44, 0x4b,
	0x47, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0d, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x1a, 0x0c, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x36, 0x0a, 0x09, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x2e,
	0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x64, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_drand_protocol_proto_rawDescOnce sync.Once
	file_drand_protocol_proto_rawDescData = file_drand_protocol_proto_rawDesc
)

func file_drand_protocol_proto_rawDescGZIP() []byte {
	file_drand_protocol_proto_rawDescOnce.Do(func() {
		file_drand_protocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_drand_protocol_proto_rawDescData)
	})
	return file_drand_protocol_proto_rawDescData
}

var file_drand_protocol_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_drand_protocol_proto_goTypes = []interface{}{
	(*IdentityRequest)(nil),     // 0: drand.IdentityRequest
	(*IdentityResponse)(nil),    // 1: drand.IdentityResponse
	(*SignalDKGPacket)(nil),     // 2: drand.SignalDKGPacket
	(*DKGInfoPacket)(nil),       // 3: drand.DKGInfoPacket
	(*PartialBeaconPacket)(nil), // 4: drand.PartialBeaconPacket
	(*DKGPacket)(nil),           // 5: drand.DKGPacket
	(*SyncRequest)(nil),         // 6: drand.SyncRequest
	(*BeaconPacket)(nil),        // 7: drand.BeaconPacket
	(*common.Metadata)(nil),     // 8: common.Metadata
	(*Identity)(nil),            // 9: drand.Identity
	(*GroupPacket)(nil),         // 10: drand.GroupPacket
	(*dkg.Packet)(nil),          // 11: dkg.Packet
	(*StatusRequest)(nil),       // 12: drand.StatusRequest
	(*Empty)(nil),               // 13: drand.Empty
	(*StatusResponse)(nil),      // 14: drand.StatusResponse
}
var file_drand_protocol_proto_depIdxs = []int32{
	8,  // 0: drand.IdentityRequest.metadata:type_name -> common.Metadata
	8,  // 1: drand.IdentityResponse.metadata:type_name -> common.Metadata
	9,  // 2: drand.SignalDKGPacket.node:type_name -> drand.Identity
	8,  // 3: drand.SignalDKGPacket.metadata:type_name -> common.Metadata
	10, // 4: drand.DKGInfoPacket.new_group:type_name -> drand.GroupPacket
	8,  // 5: drand.DKGInfoPacket.metadata:type_name -> common.Metadata
	8,  // 6: drand.PartialBeaconPacket.metadata:type_name -> common.Metadata
	11, // 7: drand.DKGPacket.dkg:type_name -> dkg.Packet
	8,  // 8: drand.DKGPacket.metadata:type_name -> common.Metadata
	8,  // 9: drand.SyncRequest.metadata:type_name -> common.Metadata
	8,  // 10: drand.BeaconPacket.metadata:type_name -> common.Metadata
	0,  // 11: drand.Protocol.GetIdentity:input_type -> drand.IdentityRequest
	2,  // 12: drand.Protocol.SignalDKGParticipant:input_type -> drand.SignalDKGPacket
	3,  // 13: drand.Protocol.PushDKGInfo:input_type -> drand.DKGInfoPacket
	5,  // 14: drand.Protocol.BroadcastDKG:input_type -> drand.DKGPacket
	4,  // 15: drand.Protocol.PartialBeacon:input_type -> drand.PartialBeaconPacket
	6,  // 16: drand.Protocol.SyncChain:input_type -> drand.SyncRequest
	12, // 17: drand.Protocol.Status:input_type -> drand.StatusRequest
	1,  // 18: drand.Protocol.GetIdentity:output_type -> drand.IdentityResponse
	13, // 19: drand.Protocol.SignalDKGParticipant:output_type -> drand.Empty
	13, // 20: drand.Protocol.PushDKGInfo:output_type -> drand.Empty
	13, // 21: drand.Protocol.BroadcastDKG:output_type -> drand.Empty
	13, // 22: drand.Protocol.PartialBeacon:output_type -> drand.Empty
	7,  // 23: drand.Protocol.SyncChain:output_type -> drand.BeaconPacket
	14, // 24: drand.Protocol.Status:output_type -> drand.StatusResponse
	18, // [18:25] is the sub-list for method output_type
	11, // [11:18] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_drand_protocol_proto_init() }
func file_drand_protocol_proto_init() {
	if File_drand_protocol_proto != nil {
		return
	}
	file_drand_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_drand_protocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalDKGPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DKGInfoPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartialBeaconPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DKGPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drand_protocol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drand_protocol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drand_protocol_proto_goTypes,
		DependencyIndexes: file_drand_protocol_proto_depIdxs,
		MessageInfos:      file_drand_protocol_proto_msgTypes,
	}.Build()
	File_drand_protocol_proto = out.File
	file_drand_protocol_proto_rawDesc = nil
	file_drand_protocol_proto_goTypes = nil
	file_drand_protocol_proto_depIdxs = nil
}
