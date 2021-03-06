// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: price_service.proto

package pb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ask    float64 `protobuf:"fixed64,2,opt,name=ask,proto3" json:"ask,omitempty"`
	Bid    float64 `protobuf:"fixed64,3,opt,name=bid,proto3" json:"bid,omitempty"`
	Symbol string  `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_price_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_price_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_price_service_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Price) GetAsk() float64 {
	if x != nil {
		return x.Ask
	}
	return 0
}

func (x *Price) GetBid() float64 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *Price) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

var File_price_service_proto protoreflect.FileDescriptor

var file_price_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x53, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x73, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x62, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x32, 0x3f, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x06, 0x2e, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_price_service_proto_rawDescOnce sync.Once
	file_price_service_proto_rawDescData = file_price_service_proto_rawDesc
)

func file_price_service_proto_rawDescGZIP() []byte {
	file_price_service_proto_rawDescOnce.Do(func() {
		file_price_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_price_service_proto_rawDescData)
	})
	return file_price_service_proto_rawDescData
}

var file_price_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_price_service_proto_goTypes = []interface{}{
	(*Price)(nil),       // 0: Price
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
}
var file_price_service_proto_depIdxs = []int32{
	1, // 0: PriceService.GetPrices:input_type -> google.protobuf.Empty
	0, // 1: PriceService.GetPrices:output_type -> Price
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_price_service_proto_init() }
func file_price_service_proto_init() {
	if File_price_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_price_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
			RawDescriptor: file_price_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_price_service_proto_goTypes,
		DependencyIndexes: file_price_service_proto_depIdxs,
		MessageInfos:      file_price_service_proto_msgTypes,
	}.Build()
	File_price_service_proto = out.File
	file_price_service_proto_rawDesc = nil
	file_price_service_proto_goTypes = nil
	file_price_service_proto_depIdxs = nil
}
