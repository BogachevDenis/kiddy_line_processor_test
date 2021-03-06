// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: sport.proto

package tutorialpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X []string `protobuf:"bytes,1,rep,name=x,proto3" json:"x,omitempty"`
	Y int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_sport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_sport_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetX() []string {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *Request) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xresult []string           `protobuf:"bytes,1,rep,name=xresult,proto3" json:"xresult,omitempty"`
	Result  map[string]float32 `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_sport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_sport_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetXresult() []string {
	if x != nil {
		return x.Xresult
	}
	return nil
}

func (x *Response) GetResult() map[string]float32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_sport_proto protoreflect.FileDescriptor

var file_sport_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x25, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x97,
	0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x78,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x78, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x4e, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x45, 0x0a, 0x16, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4f, 0x6e,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sport_proto_rawDescOnce sync.Once
	file_sport_proto_rawDescData = file_sport_proto_rawDesc
)

func file_sport_proto_rawDescGZIP() []byte {
	file_sport_proto_rawDescOnce.Do(func() {
		file_sport_proto_rawDescData = protoimpl.X.CompressGZIP(file_sport_proto_rawDescData)
	})
	return file_sport_proto_rawDescData
}

var file_sport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sport_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: protobuf.Request
	(*Response)(nil), // 1: protobuf.Response
	nil,              // 2: protobuf.Response.ResultEntry
}
var file_sport_proto_depIdxs = []int32{
	2, // 0: protobuf.Response.result:type_name -> protobuf.Response.ResultEntry
	0, // 1: protobuf.Sport.SubscribeOnSportsLines:input_type -> protobuf.Request
	1, // 2: protobuf.Sport.SubscribeOnSportsLines:output_type -> protobuf.Response
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sport_proto_init() }
func file_sport_proto_init() {
	if File_sport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_sport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_sport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sport_proto_goTypes,
		DependencyIndexes: file_sport_proto_depIdxs,
		MessageInfos:      file_sport_proto_msgTypes,
	}.Build()
	File_sport_proto = out.File
	file_sport_proto_rawDesc = nil
	file_sport_proto_goTypes = nil
	file_sport_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SportClient is the client API for Sport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SportClient interface {
	SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (Sport_SubscribeOnSportsLinesClient, error)
}

type sportClient struct {
	cc grpc.ClientConnInterface
}

func NewSportClient(cc grpc.ClientConnInterface) SportClient {
	return &sportClient{cc}
}

func (c *sportClient) SubscribeOnSportsLines(ctx context.Context, opts ...grpc.CallOption) (Sport_SubscribeOnSportsLinesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Sport_serviceDesc.Streams[0], "/protobuf.Sport/SubscribeOnSportsLines", opts...)
	if err != nil {
		return nil, err
	}
	x := &sportSubscribeOnSportsLinesClient{stream}
	return x, nil
}

type Sport_SubscribeOnSportsLinesClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type sportSubscribeOnSportsLinesClient struct {
	grpc.ClientStream
}

func (x *sportSubscribeOnSportsLinesClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sportSubscribeOnSportsLinesClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SportServer is the server API for Sport service.
type SportServer interface {
	SubscribeOnSportsLines(Sport_SubscribeOnSportsLinesServer) error
}

// UnimplementedSportServer can be embedded to have forward compatible implementations.
type UnimplementedSportServer struct {
}

func (*UnimplementedSportServer) SubscribeOnSportsLines(Sport_SubscribeOnSportsLinesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeOnSportsLines not implemented")
}

func RegisterSportServer(s *grpc.Server, srv SportServer) {
	s.RegisterService(&_Sport_serviceDesc, srv)
}

func _Sport_SubscribeOnSportsLines_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SportServer).SubscribeOnSportsLines(&sportSubscribeOnSportsLinesServer{stream})
}

type Sport_SubscribeOnSportsLinesServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type sportSubscribeOnSportsLinesServer struct {
	grpc.ServerStream
}

func (x *sportSubscribeOnSportsLinesServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sportSubscribeOnSportsLinesServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Sport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Sport",
	HandlerType: (*SportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeOnSportsLines",
			Handler:       _Sport_SubscribeOnSportsLines_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sport.proto",
}
