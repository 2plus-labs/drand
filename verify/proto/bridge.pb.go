// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: bridge.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// mint pegged token
// triggered by deposit at the original token vault
type Mint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Account []byte `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Amount  []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// depositor defines the account address that made deposit at the original token chain,
	// or the account address that burned tokens at another remote chain
	// Not applicable to governance-triggered mints.
	Depositor []byte `protobuf:"bytes,4,opt,name=depositor,proto3" json:"depositor,omitempty"`
	// ref_chain_id defines the reference chain ID, taking values of:
	// 1. The common case: the chain ID on which the corresponding deposit or burn happened;
	// 2. Governance-triggered mint: the chain ID on which the minting will happen.
	RefChainId uint64 `protobuf:"varint,5,opt,name=ref_chain_id,json=refChainId,proto3" json:"ref_chain_id,omitempty"`
	// ref_id defines a unique reference ID, taking values of:
	// 1. The common case of deposit/burn-mint: the deposit or burn ID;
	// 2. Governance-triggered mint: ID as needed.
	RefId []byte `protobuf:"bytes,6,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
}

func (x *Mint) Reset() {
	*x = Mint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mint) ProtoMessage() {}

func (x *Mint) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mint.ProtoReflect.Descriptor instead.
func (*Mint) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *Mint) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Mint) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Mint) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Mint) GetDepositor() []byte {
	if x != nil {
		return x.Depositor
	}
	return nil
}

func (x *Mint) GetRefChainId() uint64 {
	if x != nil {
		return x.RefChainId
	}
	return 0
}

func (x *Mint) GetRefId() []byte {
	if x != nil {
		return x.RefId
	}
	return nil
}

// withdraw locked original tokens
// triggered by burn at the pegged token bridge
type Withdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Receiver []byte `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount   []byte `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// burn_account defines the account that burned the pegged token.
	// Not applicable to fee claims and governance-triggered withdrawals.
	BurnAccount []byte `protobuf:"bytes,4,opt,name=burn_account,json=burnAccount,proto3" json:"burn_account,omitempty"`
	// ref_chain_id defines the reference chain ID, taking values of:
	// 1. The common case of burn-withdraw: the chain ID on which the corresponding burn happened;
	// 2. Pegbridge fee claim: zero / Not applicable;
	// 3. Other governance-triggered withdrawals: the chain ID on which the withdrawal will happen.
	RefChainId uint64 `protobuf:"varint,5,opt,name=ref_chain_id,json=refChainId,proto3" json:"ref_chain_id,omitempty"`
	// ref_id defines a unique reference ID, taking values of:
	// 1. The common case of burn-withdraw: the burn ID;
	// 2. Pegbridge fee claim: a per-account nonce;
	// 3. Refund for wrong deposit: the deposit ID;
	// 4. Governance-triggered withdrawal: ID as needed.
	RefId []byte `protobuf:"bytes,6,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
}

func (x *Withdraw) Reset() {
	*x = Withdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdraw) ProtoMessage() {}

func (x *Withdraw) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdraw.ProtoReflect.Descriptor instead.
func (*Withdraw) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *Withdraw) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Withdraw) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *Withdraw) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Withdraw) GetBurnAccount() []byte {
	if x != nil {
		return x.BurnAccount
	}
	return nil
}

func (x *Withdraw) GetRefChainId() uint64 {
	if x != nil {
		return x.RefChainId
	}
	return 0
}

func (x *Withdraw) GetRefId() []byte {
	if x != nil {
		return x.RefId
	}
	return nil
}

var file_bridge_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1004,
		Name:          "pegged.soltype",
		Tag:           "bytes,1004,opt,name=soltype",
		Filename:      "bridge.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string soltype = 1004;
	E_Soltype = &file_bridge_proto_extTypes[0]
)

var File_bridge_proto protoreflect.FileDescriptor

var file_bridge_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x65, 0x67, 0x67, 0x65, 0x64, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01, 0x0a, 0x04, 0x4d, 0x69, 0x6e,
	0x74, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x24, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x75, 0x69,
	0x6e, 0x74, 0x32, 0x35, 0x36, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x09, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x5f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72,
	0x65, 0x66, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x33, 0x32, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x22, 0xec, 0x01, 0x0a,
	0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2,
	0x3e, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x75, 0x69, 0x6e, 0x74, 0x32, 0x35, 0x36, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x62, 0x75, 0x72, 0x6e, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2,
	0x3e, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x62, 0x75, 0x72, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65,
	0x66, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x66, 0x5f,
	0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xe2, 0x3e, 0x07, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x33, 0x32, 0x52, 0x05, 0x72, 0x65, 0x66, 0x49, 0x64, 0x3a, 0x38, 0x0a, 0x07, 0x73,
	0x6f, 0x6c, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f,
	0x6c, 0x74, 0x79, 0x70, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f, 0x64, 0x72, 0x61, 0x6e, 0x64, 0x2f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bridge_proto_rawDescOnce sync.Once
	file_bridge_proto_rawDescData = file_bridge_proto_rawDesc
)

func file_bridge_proto_rawDescGZIP() []byte {
	file_bridge_proto_rawDescOnce.Do(func() {
		file_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_bridge_proto_rawDescData)
	})
	return file_bridge_proto_rawDescData
}

var file_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bridge_proto_goTypes = []interface{}{
	(*Mint)(nil),                      // 0: pegged.Mint
	(*Withdraw)(nil),                  // 1: pegged.Withdraw
	(*descriptorpb.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
}
var file_bridge_proto_depIdxs = []int32{
	2, // 0: pegged.soltype:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bridge_proto_init() }
func file_bridge_proto_init() {
	if File_bridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mint); i {
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
		file_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdraw); i {
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
			RawDescriptor: file_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_bridge_proto_goTypes,
		DependencyIndexes: file_bridge_proto_depIdxs,
		MessageInfos:      file_bridge_proto_msgTypes,
		ExtensionInfos:    file_bridge_proto_extTypes,
	}.Build()
	File_bridge_proto = out.File
	file_bridge_proto_rawDesc = nil
	file_bridge_proto_goTypes = nil
	file_bridge_proto_depIdxs = nil
}
