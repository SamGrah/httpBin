//
//Execute generate code from this protofile with the following command line:
//protoc --go_out=./generated --go_opt=paths=source_relative --go-grpc_out=./generated --go-grpc_opt=paths=source_relative adapters/bin.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: pkg/bin.proto

package binManager

import (
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

type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{0}
}

type NewBinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinId string `protobuf:"bytes,1,opt,name=binId,proto3" json:"binId,omitempty"`
}

func (x *NewBinResponse) Reset() {
	*x = NewBinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBinResponse) ProtoMessage() {}

func (x *NewBinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBinResponse.ProtoReflect.Descriptor instead.
func (*NewBinResponse) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{1}
}

func (x *NewBinResponse) GetBinId() string {
	if x != nil {
		return x.BinId
	}
	return ""
}

type LogRequestParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinId        string            `protobuf:"bytes,1,opt,name=binId,proto3" json:"binId,omitempty"`
	RequestToLog map[string]string `protobuf:"bytes,2,rep,name=requestToLog,proto3" json:"requestToLog,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LogRequestParams) Reset() {
	*x = LogRequestParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequestParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequestParams) ProtoMessage() {}

func (x *LogRequestParams) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequestParams.ProtoReflect.Descriptor instead.
func (*LogRequestParams) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{2}
}

func (x *LogRequestParams) GetBinId() string {
	if x != nil {
		return x.BinId
	}
	return ""
}

func (x *LogRequestParams) GetRequestToLog() map[string]string {
	if x != nil {
		return x.RequestToLog
	}
	return nil
}

type LogRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogRequestResponse) Reset() {
	*x = LogRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequestResponse) ProtoMessage() {}

func (x *LogRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequestResponse.ProtoReflect.Descriptor instead.
func (*LogRequestResponse) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{3}
}

type FetchBinContentsParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinId string `protobuf:"bytes,1,opt,name=binId,proto3" json:"binId,omitempty"`
}

func (x *FetchBinContentsParams) Reset() {
	*x = FetchBinContentsParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBinContentsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBinContentsParams) ProtoMessage() {}

func (x *FetchBinContentsParams) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBinContentsParams.ProtoReflect.Descriptor instead.
func (*FetchBinContentsParams) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{4}
}

func (x *FetchBinContentsParams) GetBinId() string {
	if x != nil {
		return x.BinId
	}
	return ""
}

type HttpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contents map[string]string `protobuf:"bytes,1,rep,name=contents,proto3" json:"contents,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HttpRequest) Reset() {
	*x = HttpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRequest) ProtoMessage() {}

func (x *HttpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRequest.ProtoReflect.Descriptor instead.
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{5}
}

func (x *HttpRequest) GetContents() map[string]string {
	if x != nil {
		return x.Contents
	}
	return nil
}

type FetchBinContentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BinContents []*HttpRequest `protobuf:"bytes,1,rep,name=binContents,proto3" json:"binContents,omitempty"`
}

func (x *FetchBinContentsResponse) Reset() {
	*x = FetchBinContentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_bin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchBinContentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchBinContentsResponse) ProtoMessage() {}

func (x *FetchBinContentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_bin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchBinContentsResponse.ProtoReflect.Descriptor instead.
func (*FetchBinContentsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_bin_proto_rawDescGZIP(), []int{6}
}

func (x *FetchBinContentsResponse) GetBinContents() []*HttpRequest {
	if x != nil {
		return x.BinContents
	}
	return nil
}

var File_pkg_bin_proto protoreflect.FileDescriptor

var file_pkg_bin_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x6b, 0x67, 0x2f, 0x62, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0x08, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xbd, 0x01,
	0x0a, 0x10, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x62, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x1a, 0x3f, 0x0a, 0x11,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x14, 0x0a,
	0x12, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2e, 0x0a, 0x16, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x69,
	0x6e, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x18, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0b, 0x62, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x83, 0x02, 0x0a, 0x0a, 0x42,
	0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x42, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x62, 0x69,
	0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x77,
	0x42, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a,
	0x0f, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x6f, 0x42, 0x69, 0x6e,
	0x12, 0x1c, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1e,
	0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x24, 0x2e, 0x62, 0x69, 0x6e, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x42, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2f, 0x62, 0x69, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_bin_proto_rawDescOnce sync.Once
	file_pkg_bin_proto_rawDescData = file_pkg_bin_proto_rawDesc
)

func file_pkg_bin_proto_rawDescGZIP() []byte {
	file_pkg_bin_proto_rawDescOnce.Do(func() {
		file_pkg_bin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_bin_proto_rawDescData)
	})
	return file_pkg_bin_proto_rawDescData
}

var file_pkg_bin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_bin_proto_goTypes = []interface{}{
	(*Params)(nil),                   // 0: binManager.Params
	(*NewBinResponse)(nil),           // 1: binManager.NewBinResponse
	(*LogRequestParams)(nil),         // 2: binManager.LogRequestParams
	(*LogRequestResponse)(nil),       // 3: binManager.LogRequestResponse
	(*FetchBinContentsParams)(nil),   // 4: binManager.FetchBinContentsParams
	(*HttpRequest)(nil),              // 5: binManager.HttpRequest
	(*FetchBinContentsResponse)(nil), // 6: binManager.FetchBinContentsResponse
	nil,                              // 7: binManager.LogRequestParams.RequestToLogEntry
	nil,                              // 8: binManager.HttpRequest.ContentsEntry
}
var file_pkg_bin_proto_depIdxs = []int32{
	7, // 0: binManager.LogRequestParams.requestToLog:type_name -> binManager.LogRequestParams.RequestToLogEntry
	8, // 1: binManager.HttpRequest.contents:type_name -> binManager.HttpRequest.ContentsEntry
	5, // 2: binManager.FetchBinContentsResponse.binContents:type_name -> binManager.HttpRequest
	0, // 3: binManager.BinManager.GenerateNewBin:input_type -> binManager.Params
	2, // 4: binManager.BinManager.LogRequestToBin:input_type -> binManager.LogRequestParams
	4, // 5: binManager.BinManager.FetchBinContents:input_type -> binManager.FetchBinContentsParams
	1, // 6: binManager.BinManager.GenerateNewBin:output_type -> binManager.NewBinResponse
	3, // 7: binManager.BinManager.LogRequestToBin:output_type -> binManager.LogRequestResponse
	6, // 8: binManager.BinManager.FetchBinContents:output_type -> binManager.FetchBinContentsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_bin_proto_init() }
func file_pkg_bin_proto_init() {
	if File_pkg_bin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_bin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
		file_pkg_bin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBinResponse); i {
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
		file_pkg_bin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequestParams); i {
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
		file_pkg_bin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequestResponse); i {
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
		file_pkg_bin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBinContentsParams); i {
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
		file_pkg_bin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpRequest); i {
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
		file_pkg_bin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchBinContentsResponse); i {
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
			RawDescriptor: file_pkg_bin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_bin_proto_goTypes,
		DependencyIndexes: file_pkg_bin_proto_depIdxs,
		MessageInfos:      file_pkg_bin_proto_msgTypes,
	}.Build()
	File_pkg_bin_proto = out.File
	file_pkg_bin_proto_rawDesc = nil
	file_pkg_bin_proto_goTypes = nil
	file_pkg_bin_proto_depIdxs = nil
}