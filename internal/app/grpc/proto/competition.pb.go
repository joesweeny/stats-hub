// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/app/grpc/proto/competition.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Competition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsCup *wrappers.BoolValue `protobuf:"bytes,3,opt,name=is_cup,json=isCup,proto3" json:"is_cup,omitempty"`
}

func (x *Competition) Reset() {
	*x = Competition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_competition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Competition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Competition) ProtoMessage() {}

func (x *Competition) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_competition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Competition.ProtoReflect.Descriptor instead.
func (*Competition) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_competition_proto_rawDescGZIP(), []int{0}
}

func (x *Competition) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Competition) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Competition) GetIsCup() *wrappers.BoolValue {
	if x != nil {
		return x.IsCup
	}
	return nil
}

var File_internal_app_grpc_proto_competition_proto protoreflect.FileDescriptor

var file_internal_app_grpc_proto_competition_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x63, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x69, 0x73, 0x43, 0x75, 0x70,
	0x32, 0x5b, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x19, 0x5a,
	0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_grpc_proto_competition_proto_rawDescOnce sync.Once
	file_internal_app_grpc_proto_competition_proto_rawDescData = file_internal_app_grpc_proto_competition_proto_rawDesc
)

func file_internal_app_grpc_proto_competition_proto_rawDescGZIP() []byte {
	file_internal_app_grpc_proto_competition_proto_rawDescOnce.Do(func() {
		file_internal_app_grpc_proto_competition_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_grpc_proto_competition_proto_rawDescData)
	})
	return file_internal_app_grpc_proto_competition_proto_rawDescData
}

var file_internal_app_grpc_proto_competition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_app_grpc_proto_competition_proto_goTypes = []interface{}{
	(*Competition)(nil),        // 0: proto.Competition
	(*wrappers.BoolValue)(nil), // 1: google.protobuf.BoolValue
	(*CompetitionRequest)(nil), // 2: proto.CompetitionRequest
}
var file_internal_app_grpc_proto_competition_proto_depIdxs = []int32{
	1, // 0: proto.Competition.is_cup:type_name -> google.protobuf.BoolValue
	2, // 1: proto.CompetitionService.ListCompetitions:input_type -> proto.CompetitionRequest
	0, // 2: proto.CompetitionService.ListCompetitions:output_type -> proto.Competition
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_app_grpc_proto_competition_proto_init() }
func file_internal_app_grpc_proto_competition_proto_init() {
	if File_internal_app_grpc_proto_competition_proto != nil {
		return
	}
	file_internal_app_grpc_proto_requests_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_app_grpc_proto_competition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Competition); i {
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
			RawDescriptor: file_internal_app_grpc_proto_competition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_grpc_proto_competition_proto_goTypes,
		DependencyIndexes: file_internal_app_grpc_proto_competition_proto_depIdxs,
		MessageInfos:      file_internal_app_grpc_proto_competition_proto_msgTypes,
	}.Build()
	File_internal_app_grpc_proto_competition_proto = out.File
	file_internal_app_grpc_proto_competition_proto_rawDesc = nil
	file_internal_app_grpc_proto_competition_proto_goTypes = nil
	file_internal_app_grpc_proto_competition_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompetitionServiceClient is the client API for CompetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompetitionServiceClient interface {
	ListCompetitions(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (CompetitionService_ListCompetitionsClient, error)
}

type competitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionServiceClient(cc grpc.ClientConnInterface) CompetitionServiceClient {
	return &competitionServiceClient{cc}
}

func (c *competitionServiceClient) ListCompetitions(ctx context.Context, in *CompetitionRequest, opts ...grpc.CallOption) (CompetitionService_ListCompetitionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompetitionService_serviceDesc.Streams[0], "/proto.CompetitionService/ListCompetitions", opts...)
	if err != nil {
		return nil, err
	}
	x := &competitionServiceListCompetitionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompetitionService_ListCompetitionsClient interface {
	Recv() (*Competition, error)
	grpc.ClientStream
}

type competitionServiceListCompetitionsClient struct {
	grpc.ClientStream
}

func (x *competitionServiceListCompetitionsClient) Recv() (*Competition, error) {
	m := new(Competition)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CompetitionServiceServer is the server API for CompetitionService service.
type CompetitionServiceServer interface {
	ListCompetitions(*CompetitionRequest, CompetitionService_ListCompetitionsServer) error
}

// UnimplementedCompetitionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCompetitionServiceServer struct {
}

func (*UnimplementedCompetitionServiceServer) ListCompetitions(*CompetitionRequest, CompetitionService_ListCompetitionsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCompetitions not implemented")
}

func RegisterCompetitionServiceServer(s *grpc.Server, srv CompetitionServiceServer) {
	s.RegisterService(&_CompetitionService_serviceDesc, srv)
}

func _CompetitionService_ListCompetitions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompetitionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompetitionServiceServer).ListCompetitions(m, &competitionServiceListCompetitionsServer{stream})
}

type CompetitionService_ListCompetitionsServer interface {
	Send(*Competition) error
	grpc.ServerStream
}

type competitionServiceListCompetitionsServer struct {
	grpc.ServerStream
}

func (x *competitionServiceListCompetitionsServer) Send(m *Competition) error {
	return x.ServerStream.SendMsg(m)
}

var _CompetitionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CompetitionService",
	HandlerType: (*CompetitionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCompetitions",
			Handler:       _CompetitionService_ListCompetitions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "internal/app/grpc/proto/competition.proto",
}
