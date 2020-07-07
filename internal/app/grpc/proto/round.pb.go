// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: internal/app/grpc/proto/round.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

type Round struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SeasonId uint64 `protobuf:"varint,3,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	StartDate string `protobuf:"bytes,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// RFC3339 formatted string i.e "2006-01-02T15:04:05Z07:00"
	EndDate string `protobuf:"bytes,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *Round) Reset() {
	*x = Round{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_app_grpc_proto_round_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Round) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Round) ProtoMessage() {}

func (x *Round) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_grpc_proto_round_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Round.ProtoReflect.Descriptor instead.
func (*Round) Descriptor() ([]byte, []int) {
	return file_internal_app_grpc_proto_round_proto_rawDescGZIP(), []int{0}
}

func (x *Round) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Round) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Round) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *Round) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Round) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

var File_internal_app_grpc_proto_round_proto protoreflect.FileDescriptor

var file_internal_app_grpc_proto_round_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a,
	0x05, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x42, 0x19, 0x5a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_app_grpc_proto_round_proto_rawDescOnce sync.Once
	file_internal_app_grpc_proto_round_proto_rawDescData = file_internal_app_grpc_proto_round_proto_rawDesc
)

func file_internal_app_grpc_proto_round_proto_rawDescGZIP() []byte {
	file_internal_app_grpc_proto_round_proto_rawDescOnce.Do(func() {
		file_internal_app_grpc_proto_round_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_app_grpc_proto_round_proto_rawDescData)
	})
	return file_internal_app_grpc_proto_round_proto_rawDescData
}

var file_internal_app_grpc_proto_round_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_app_grpc_proto_round_proto_goTypes = []interface{}{
	(*Round)(nil), // 0: proto.Round
}
var file_internal_app_grpc_proto_round_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_app_grpc_proto_round_proto_init() }
func file_internal_app_grpc_proto_round_proto_init() {
	if File_internal_app_grpc_proto_round_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_app_grpc_proto_round_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Round); i {
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
			RawDescriptor: file_internal_app_grpc_proto_round_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_app_grpc_proto_round_proto_goTypes,
		DependencyIndexes: file_internal_app_grpc_proto_round_proto_depIdxs,
		MessageInfos:      file_internal_app_grpc_proto_round_proto_msgTypes,
	}.Build()
	File_internal_app_grpc_proto_round_proto = out.File
	file_internal_app_grpc_proto_round_proto_rawDesc = nil
	file_internal_app_grpc_proto_round_proto_goTypes = nil
	file_internal_app_grpc_proto_round_proto_depIdxs = nil
}
