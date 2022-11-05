// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: app/proto/top_up_service/top_up_service.proto

package top_up_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TopUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance  string `protobuf:"bytes,1,opt,name=Balance,proto3" json:"Balance,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	NoKartu  string `protobuf:"bytes,4,opt,name=NoKartu,proto3" json:"NoKartu,omitempty"`
}

func (x *TopUpRequest) Reset() {
	*x = TopUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_top_up_service_top_up_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopUpRequest) ProtoMessage() {}

func (x *TopUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_top_up_service_top_up_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopUpRequest.ProtoReflect.Descriptor instead.
func (*TopUpRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_top_up_service_top_up_service_proto_rawDescGZIP(), []int{0}
}

func (x *TopUpRequest) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *TopUpRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TopUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TopUpRequest) GetNoKartu() string {
	if x != nil {
		return x.NoKartu
	}
	return ""
}

type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BillerId string `protobuf:"bytes,1,opt,name=BillerId,proto3" json:"BillerId,omitempty"`
	UserId   string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	NoKartu  string `protobuf:"bytes,4,opt,name=NoKartu,proto3" json:"NoKartu,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_top_up_service_top_up_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_top_up_service_top_up_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_top_up_service_top_up_service_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentRequest) GetBillerId() string {
	if x != nil {
		return x.BillerId
	}
	return ""
}

func (x *PaymentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PaymentRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PaymentRequest) GetNoKartu() string {
	if x != nil {
		return x.NoKartu
	}
	return ""
}

var File_app_proto_top_up_service_top_up_service_proto protoreflect.FileDescriptor

var file_app_proto_top_up_service_top_up_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x5f,
	0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x6f, 0x70, 0x5f, 0x75,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x74, 0x6f, 0x70, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x76, 0x0a, 0x0c,
	0x54, 0x6f, 0x70, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x6f,
	0x4b, 0x61, 0x72, 0x74, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x6f, 0x4b,
	0x61, 0x72, 0x74, 0x75, 0x22, 0x7a, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4e, 0x6f, 0x4b, 0x61, 0x72, 0x74,
	0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4e, 0x6f, 0x4b, 0x61, 0x72, 0x74, 0x75,
	0x32, 0x91, 0x01, 0x0a, 0x0d, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x3d, 0x0a, 0x05, 0x54, 0x6f, 0x70, 0x55, 0x70, 0x12, 0x1c, 0x2e, 0x74, 0x6f,
	0x70, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x54, 0x6f, 0x70,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x41, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x74,
	0x6f, 0x70, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x73, 0x37, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x6f, 0x70, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x74,
	0x6f, 0x70, 0x5f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proto_top_up_service_top_up_service_proto_rawDescOnce sync.Once
	file_app_proto_top_up_service_top_up_service_proto_rawDescData = file_app_proto_top_up_service_top_up_service_proto_rawDesc
)

func file_app_proto_top_up_service_top_up_service_proto_rawDescGZIP() []byte {
	file_app_proto_top_up_service_top_up_service_proto_rawDescOnce.Do(func() {
		file_app_proto_top_up_service_top_up_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_top_up_service_top_up_service_proto_rawDescData)
	})
	return file_app_proto_top_up_service_top_up_service_proto_rawDescData
}

var file_app_proto_top_up_service_top_up_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_proto_top_up_service_top_up_service_proto_goTypes = []interface{}{
	(*TopUpRequest)(nil),   // 0: top_up_service.TopUpRequest
	(*PaymentRequest)(nil), // 1: top_up_service.PaymentRequest
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_app_proto_top_up_service_top_up_service_proto_depIdxs = []int32{
	0, // 0: top_up_service.TopUpServices.TopUp:input_type -> top_up_service.TopUpRequest
	1, // 1: top_up_service.TopUpServices.Payment:input_type -> top_up_service.PaymentRequest
	2, // 2: top_up_service.TopUpServices.TopUp:output_type -> google.protobuf.Empty
	2, // 3: top_up_service.TopUpServices.Payment:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_proto_top_up_service_top_up_service_proto_init() }
func file_app_proto_top_up_service_top_up_service_proto_init() {
	if File_app_proto_top_up_service_top_up_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_top_up_service_top_up_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopUpRequest); i {
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
		file_app_proto_top_up_service_top_up_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
			RawDescriptor: file_app_proto_top_up_service_top_up_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proto_top_up_service_top_up_service_proto_goTypes,
		DependencyIndexes: file_app_proto_top_up_service_top_up_service_proto_depIdxs,
		MessageInfos:      file_app_proto_top_up_service_top_up_service_proto_msgTypes,
	}.Build()
	File_app_proto_top_up_service_top_up_service_proto = out.File
	file_app_proto_top_up_service_top_up_service_proto_rawDesc = nil
	file_app_proto_top_up_service_top_up_service_proto_goTypes = nil
	file_app_proto_top_up_service_top_up_service_proto_depIdxs = nil
}
