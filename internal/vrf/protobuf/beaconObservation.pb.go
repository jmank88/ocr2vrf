package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sig []byte `protobuf:"bytes,1,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Signature) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{0}
}

func (x *Signature) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Callback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        uint64 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	NumWords         uint32 `protobuf:"varint,2,opt,name=numWords,proto3" json:"numWords,omitempty"`
	Requester        []byte `protobuf:"bytes,3,opt,name=requester,proto3" json:"requester,omitempty"`
	Arguments        []byte `protobuf:"bytes,4,opt,name=arguments,proto3" json:"arguments,omitempty"`
	SubscriptionID   uint64 `protobuf:"varint,5,opt,name=subscriptionID,proto3" json:"subscriptionID,omitempty"`
	Height           uint64 `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	ConfDelay        uint32 `protobuf:"varint,7,opt,name=confDelay,proto3" json:"confDelay,omitempty"`
	RequestHeight    uint64 `protobuf:"varint,8,opt,name=requestHeight,proto3" json:"requestHeight,omitempty"`
	RequestBlockHash []byte `protobuf:"bytes,9,opt,name=requestBlockHash,proto3" json:"requestBlockHash,omitempty"`
}

func (x *Callback) Reset() {
	*x = Callback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Callback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Callback) ProtoMessage() {}

func (x *Callback) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Callback) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{1}
}

func (x *Callback) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Callback) GetNumWords() uint32 {
	if x != nil {
		return x.NumWords
	}
	return 0
}

func (x *Callback) GetRequester() []byte {
	if x != nil {
		return x.Requester
	}
	return nil
}

func (x *Callback) GetArguments() []byte {
	if x != nil {
		return x.Arguments
	}
	return nil
}

func (x *Callback) GetSubscriptionID() uint64 {
	if x != nil {
		return x.SubscriptionID
	}
	return 0
}

func (x *Callback) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Callback) GetConfDelay() uint32 {
	if x != nil {
		return x.ConfDelay
	}
	return 0
}

func (x *Callback) GetRequestHeight() uint64 {
	if x != nil {
		return x.RequestHeight
	}
	return 0
}

func (x *Callback) GetRequestBlockHash() []byte {
	if x != nil {
		return x.RequestBlockHash
	}
	return nil
}

type CostedCallback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Callback     *Callback `protobuf:"bytes,1,opt,name=callback,proto3" json:"callback,omitempty"`
	Price        []byte    `protobuf:"bytes,2,opt,name=price,proto3" json:"price,omitempty"`
	GasAllowance []byte    `protobuf:"bytes,3,opt,name=gasAllowance,proto3" json:"gasAllowance,omitempty"`
}

func (x *CostedCallback) Reset() {
	*x = CostedCallback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostedCallback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostedCallback) ProtoMessage() {}

func (x *CostedCallback) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CostedCallback) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{2}
}

func (x *CostedCallback) GetCallback() *Callback {
	if x != nil {
		return x.Callback
	}
	return nil
}

func (x *CostedCallback) GetPrice() []byte {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *CostedCallback) GetGasAllowance() []byte {
	if x != nil {
		return x.GasAllowance
	}
	return nil
}

type RecentBlockAndHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Blockhash []byte `protobuf:"bytes,2,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
}

func (x *RecentBlockAndHash) Reset() {
	*x = RecentBlockAndHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecentBlockAndHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecentBlockAndHash) ProtoMessage() {}

func (x *RecentBlockAndHash) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*RecentBlockAndHash) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{3}
}

func (x *RecentBlockAndHash) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *RecentBlockAndHash) GetBlockhash() []byte {
	if x != nil {
		return x.Blockhash
	}
	return nil
}

type VRFResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height    uint64     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Delay     uint32     `protobuf:"varint,2,opt,name=delay,proto3" json:"delay,omitempty"`
	Blockhash []byte     `protobuf:"bytes,3,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	Sig       *Signature `protobuf:"bytes,4,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (x *VRFResponse) Reset() {
	*x = VRFResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VRFResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VRFResponse) ProtoMessage() {}

func (x *VRFResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*VRFResponse) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{4}
}

func (x *VRFResponse) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VRFResponse) GetDelay() uint32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *VRFResponse) GetBlockhash() []byte {
	if x != nil {
		return x.Blockhash
	}
	return nil
}

func (x *VRFResponse) GetSig() *Signature {
	if x != nil {
		return x.Sig
	}
	return nil
}

type Observation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JuelsPerFeeCoin   []byte                `protobuf:"bytes,1,opt,name=juelsPerFeeCoin,proto3" json:"juelsPerFeeCoin,omitempty"`
	RecentBlockHashes []*RecentBlockAndHash `protobuf:"bytes,2,rep,name=recentBlockHashes,proto3" json:"recentBlockHashes,omitempty"`
	Proofs            []*VRFResponse        `protobuf:"bytes,4,rep,name=proofs,proto3" json:"proofs,omitempty"`
	Callbacks         []*CostedCallback     `protobuf:"bytes,6,rep,name=callbacks,proto3" json:"callbacks,omitempty"`
}

func (x *Observation) Reset() {
	*x = Observation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beaconObservation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_beaconObservation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Observation) Descriptor() ([]byte, []int) {
	return file_beaconObservation_proto_rawDescGZIP(), []int{5}
}

func (x *Observation) GetJuelsPerFeeCoin() []byte {
	if x != nil {
		return x.JuelsPerFeeCoin
	}
	return nil
}

func (x *Observation) GetRecentBlockHashes() []*RecentBlockAndHash {
	if x != nil {
		return x.RecentBlockHashes
	}
	return nil
}

func (x *Observation) GetProofs() []*VRFResponse {
	if x != nil {
		return x.Proofs
	}
	return nil
}

func (x *Observation) GetCallbacks() []*CostedCallback {
	if x != nil {
		return x.Callbacks
	}
	return nil
}

var File_beaconObservation_proto protoreflect.FileDescriptor

var file_beaconObservation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x1d, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22,
	0xb0, 0x02, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x75,
	0x6d, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6e, 0x75,
	0x6d, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79,
	0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x77, 0x0a, 0x0e, 0x63, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x67, 0x61, 0x73, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x67,
	0x61, 0x73, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x12, 0x72,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x68, 0x61, 0x73, 0x68, 0x22, 0x7d, 0x0a, 0x0b, 0x56, 0x52, 0x46, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x68, 0x61,
	0x73, 0x68, 0x12, 0x22, 0x0a, 0x03, 0x73, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x03, 0x73, 0x69, 0x67, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6a, 0x75, 0x65, 0x6c, 0x73, 0x50,
	0x65, 0x72, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0f, 0x6a, 0x75, 0x65, 0x6c, 0x73, 0x50, 0x65, 0x72, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x12, 0x47, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x48, 0x61, 0x73, 0x68, 0x52, 0x11, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x52, 0x46, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x63, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52,
	0x09, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beaconObservation_proto_rawDescOnce sync.Once
	file_beaconObservation_proto_rawDescData = file_beaconObservation_proto_rawDesc
)

func file_beaconObservation_proto_rawDescGZIP() []byte {
	file_beaconObservation_proto_rawDescOnce.Do(func() {
		file_beaconObservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_beaconObservation_proto_rawDescData)
	})
	return file_beaconObservation_proto_rawDescData
}

var file_beaconObservation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_beaconObservation_proto_goTypes = []interface{}{
	(*Signature)(nil),
	(*Callback)(nil),
	(*CostedCallback)(nil),
	(*RecentBlockAndHash)(nil),
	(*VRFResponse)(nil),
	(*Observation)(nil),
}
var file_beaconObservation_proto_depIdxs = []int32{
	1,
	0,
	3,
	4,
	2,
	5,
	5,
	5,
	5,
	0,
}

func init() { file_beaconObservation_proto_init() }
func file_beaconObservation_proto_init() {
	if File_beaconObservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beaconObservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_beaconObservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Callback); i {
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
		file_beaconObservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CostedCallback); i {
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
		file_beaconObservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecentBlockAndHash); i {
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
		file_beaconObservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VRFResponse); i {
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
		file_beaconObservation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observation); i {
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
			RawDescriptor: file_beaconObservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beaconObservation_proto_goTypes,
		DependencyIndexes: file_beaconObservation_proto_depIdxs,
		MessageInfos:      file_beaconObservation_proto_msgTypes,
	}.Build()
	File_beaconObservation_proto = out.File
	file_beaconObservation_proto_rawDesc = nil
	file_beaconObservation_proto_goTypes = nil
	file_beaconObservation_proto_depIdxs = nil
}
