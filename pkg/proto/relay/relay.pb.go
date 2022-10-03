// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.7
// source: relay/relay.proto

package pbRelay

import (
	sdk_ws "GoIM/pkg/proto/sdk_ws"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OnlinePushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationID  string          `protobuf:"bytes,1,opt,name=OperationID,proto3" json:"OperationID,omitempty"`
	MsgData      *sdk_ws.MsgData `protobuf:"bytes,2,opt,name=msgData,proto3" json:"msgData,omitempty"`
	PushToUserID string          `protobuf:"bytes,3,opt,name=pushToUserID,proto3" json:"pushToUserID,omitempty"`
}

func (x *OnlinePushMsgReq) Reset() {
	*x = OnlinePushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgReq) ProtoMessage() {}

func (x *OnlinePushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgReq.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgReq) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{0}
}

func (x *OnlinePushMsgReq) GetOperationID() string {
	if x != nil {
		return x.OperationID
	}
	return ""
}

func (x *OnlinePushMsgReq) GetMsgData() *sdk_ws.MsgData {
	if x != nil {
		return x.MsgData
	}
	return nil
}

func (x *OnlinePushMsgReq) GetPushToUserID() string {
	if x != nil {
		return x.PushToUserID
	}
	return ""
}

type OnlinePushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp []*SingleMsgToUserPlatform `protobuf:"bytes,1,rep,name=resp,proto3" json:"resp,omitempty"`
}

func (x *OnlinePushMsgResp) Reset() {
	*x = OnlinePushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnlinePushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnlinePushMsgResp) ProtoMessage() {}

func (x *OnlinePushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnlinePushMsgResp.ProtoReflect.Descriptor instead.
func (*OnlinePushMsgResp) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{1}
}

func (x *OnlinePushMsgResp) GetResp() []*SingleMsgToUserPlatform {
	if x != nil {
		return x.Resp
	}
	return nil
}

type SingleMsgToUserPlatform struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode     int64  `protobuf:"varint,1,opt,name=ResultCode,proto3" json:"ResultCode,omitempty"`
	RecvID         string `protobuf:"bytes,2,opt,name=RecvID,proto3" json:"RecvID,omitempty"`
	RecvPlatFormID int32  `protobuf:"varint,3,opt,name=RecvPlatFormID,proto3" json:"RecvPlatFormID,omitempty"`
}

func (x *SingleMsgToUserPlatform) Reset() {
	*x = SingleMsgToUserPlatform{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleMsgToUserPlatform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleMsgToUserPlatform) ProtoMessage() {}

func (x *SingleMsgToUserPlatform) ProtoReflect() protoreflect.Message {
	mi := &file_relay_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleMsgToUserPlatform.ProtoReflect.Descriptor instead.
func (*SingleMsgToUserPlatform) Descriptor() ([]byte, []int) {
	return file_relay_relay_proto_rawDescGZIP(), []int{2}
}

func (x *SingleMsgToUserPlatform) GetResultCode() int64 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *SingleMsgToUserPlatform) GetRecvID() string {
	if x != nil {
		return x.RecvID
	}
	return ""
}

func (x *SingleMsgToUserPlatform) GetRecvPlatFormID() int32 {
	if x != nil {
		return x.RecvPlatFormID
	}
	return 0
}

var File_relay_relay_proto protoreflect.FileDescriptor

var file_relay_relay_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x1a, 0x1e, 0x47, 0x6f, 0x49, 0x4d,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x64, 0x6b, 0x5f, 0x77,
	0x73, 0x2f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x5f,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x75, 0x73, 0x68, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x75, 0x73, 0x68, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x11, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x32, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73, 0x67,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x22, 0x79, 0x0a, 0x17, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4d, 0x73,
	0x67, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x65, 0x63, 0x76, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x76, 0x50,
	0x6c, 0x61, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x52, 0x65, 0x63, 0x76, 0x50, 0x6c, 0x61, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x49, 0x44, 0x32,
	0x4b, 0x0a, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x12, 0x17, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x11, 0x5a, 0x0f,
	0x2e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x3b, 0x70, 0x62, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_relay_proto_rawDescOnce sync.Once
	file_relay_relay_proto_rawDescData = file_relay_relay_proto_rawDesc
)

func file_relay_relay_proto_rawDescGZIP() []byte {
	file_relay_relay_proto_rawDescOnce.Do(func() {
		file_relay_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_relay_proto_rawDescData)
	})
	return file_relay_relay_proto_rawDescData
}

var file_relay_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_relay_relay_proto_goTypes = []interface{}{
	(*OnlinePushMsgReq)(nil),        // 0: relay.OnlinePushMsgReq
	(*OnlinePushMsgResp)(nil),       // 1: relay.OnlinePushMsgResp
	(*SingleMsgToUserPlatform)(nil), // 2: relay.SingleMsgToUserPlatform
	(*sdk_ws.MsgData)(nil),          // 3: server_api_params.MsgData
}
var file_relay_relay_proto_depIdxs = []int32{
	3, // 0: relay.OnlinePushMsgReq.msgData:type_name -> server_api_params.MsgData
	2, // 1: relay.OnlinePushMsgResp.resp:type_name -> relay.SingleMsgToUserPlatform
	0, // 2: relay.relay.OnlinePushMsg:input_type -> relay.OnlinePushMsgReq
	1, // 3: relay.relay.OnlinePushMsg:output_type -> relay.OnlinePushMsgResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_relay_relay_proto_init() }
func file_relay_relay_proto_init() {
	if File_relay_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMsgReq); i {
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
		file_relay_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnlinePushMsgResp); i {
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
		file_relay_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleMsgToUserPlatform); i {
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
			RawDescriptor: file_relay_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_relay_proto_goTypes,
		DependencyIndexes: file_relay_relay_proto_depIdxs,
		MessageInfos:      file_relay_relay_proto_msgTypes,
	}.Build()
	File_relay_relay_proto = out.File
	file_relay_relay_proto_rawDesc = nil
	file_relay_relay_proto_goTypes = nil
	file_relay_relay_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RelayClient is the client API for Relay service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelayClient interface {
	OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error)
}

type relayClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayClient(cc grpc.ClientConnInterface) RelayClient {
	return &relayClient{cc}
}

func (c *relayClient) OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error) {
	out := new(OnlinePushMsgResp)
	err := c.cc.Invoke(ctx, "/relay.relay/OnlinePushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayServer is the server API for Relay service.
type RelayServer interface {
	OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error)
}

// UnimplementedRelayServer can be embedded to have forward compatible implementations.
type UnimplementedRelayServer struct {
}

func (*UnimplementedRelayServer) OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlinePushMsg not implemented")
}

func RegisterRelayServer(s *grpc.Server, srv RelayServer) {
	s.RegisterService(&_Relay_serviceDesc, srv)
}

func _Relay_OnlinePushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlinePushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayServer).OnlinePushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relay.relay/OnlinePushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayServer).OnlinePushMsg(ctx, req.(*OnlinePushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Relay_serviceDesc = grpc.ServiceDesc{
	ServiceName: "relay.relay",
	HandlerType: (*RelayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlinePushMsg",
			Handler:    _Relay_OnlinePushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relay/relay.proto",
}
