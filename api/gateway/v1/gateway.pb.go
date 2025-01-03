// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: api/gateway/v1/gateway.proto

package v1

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

// The request message containing the user's name.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The request message containing the user's name.
type InnerGatewayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnId uint64 `protobuf:"varint,1,opt,name=conn_id,json=connId,proto3" json:"conn_id,omitempty"`
	Data   []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *InnerGatewayRequest) Reset() {
	*x = InnerGatewayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerGatewayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerGatewayRequest) ProtoMessage() {}

func (x *InnerGatewayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerGatewayRequest.ProtoReflect.Descriptor instead.
func (*InnerGatewayRequest) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *InnerGatewayRequest) GetConnId() uint64 {
	if x != nil {
		return x.ConnId
	}
	return 0
}

func (x *InnerGatewayRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// The response message containing the greetings
type InnerGatewayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InnerGatewayResponse) Reset() {
	*x = InnerGatewayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_gateway_v1_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InnerGatewayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InnerGatewayResponse) ProtoMessage() {}

func (x *InnerGatewayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_gateway_v1_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InnerGatewayResponse.ProtoReflect.Descriptor instead.
func (*InnerGatewayResponse) Descriptor() ([]byte, []int) {
	return file_api_gateway_v1_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *InnerGatewayResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *InnerGatewayResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_gateway_v1_gateway_proto protoreflect.FileDescriptor

var file_api_gateway_v1_gateway_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x22, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x42, 0x0a,
	0x13, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x3c, 0x0a, 0x14, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32,
	0x4a, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x3f, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x32, 0xba, 0x01, 0x0a, 0x0c,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x54, 0x0a, 0x07,
	0x44, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x54, 0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x23, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x41, 0x0a, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x6f, 0x79, 0x73, 0x69, 0x63, 0x73,
	0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_gateway_v1_gateway_proto_rawDescOnce sync.Once
	file_api_gateway_v1_gateway_proto_rawDescData = file_api_gateway_v1_gateway_proto_rawDesc
)

func file_api_gateway_v1_gateway_proto_rawDescGZIP() []byte {
	file_api_gateway_v1_gateway_proto_rawDescOnce.Do(func() {
		file_api_gateway_v1_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_gateway_v1_gateway_proto_rawDescData)
	})
	return file_api_gateway_v1_gateway_proto_rawDescData
}

var file_api_gateway_v1_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_gateway_v1_gateway_proto_goTypes = []any{
	(*Message)(nil),              // 0: api.gateway.v1.Message
	(*InnerGatewayRequest)(nil),  // 1: api.gateway.v1.InnerGatewayRequest
	(*InnerGatewayResponse)(nil), // 2: api.gateway.v1.InnerGatewayResponse
}
var file_api_gateway_v1_gateway_proto_depIdxs = []int32{
	0, // 0: api.gateway.v1.Gateway.Connect:input_type -> api.gateway.v1.Message
	1, // 1: api.gateway.v1.InnerGateway.DelConn:input_type -> api.gateway.v1.InnerGatewayRequest
	1, // 2: api.gateway.v1.InnerGateway.PushMsg:input_type -> api.gateway.v1.InnerGatewayRequest
	0, // 3: api.gateway.v1.Gateway.Connect:output_type -> api.gateway.v1.Message
	2, // 4: api.gateway.v1.InnerGateway.DelConn:output_type -> api.gateway.v1.InnerGatewayResponse
	2, // 5: api.gateway.v1.InnerGateway.PushMsg:output_type -> api.gateway.v1.InnerGatewayResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_gateway_v1_gateway_proto_init() }
func file_api_gateway_v1_gateway_proto_init() {
	if File_api_gateway_v1_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_gateway_v1_gateway_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_api_gateway_v1_gateway_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InnerGatewayRequest); i {
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
		file_api_gateway_v1_gateway_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InnerGatewayResponse); i {
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
			RawDescriptor: file_api_gateway_v1_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_gateway_v1_gateway_proto_goTypes,
		DependencyIndexes: file_api_gateway_v1_gateway_proto_depIdxs,
		MessageInfos:      file_api_gateway_v1_gateway_proto_msgTypes,
	}.Build()
	File_api_gateway_v1_gateway_proto = out.File
	file_api_gateway_v1_gateway_proto_rawDesc = nil
	file_api_gateway_v1_gateway_proto_goTypes = nil
	file_api_gateway_v1_gateway_proto_depIdxs = nil
}
