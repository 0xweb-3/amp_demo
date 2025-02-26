// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.15.7
// source: sku.proto

package protos

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

type Sku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Num   int32  `protobuf:"varint,4,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *Sku) Reset() {
	*x = Sku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sku) ProtoMessage() {}

func (x *Sku) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sku.ProtoReflect.Descriptor instead.
func (*Sku) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{0}
}

func (x *Sku) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sku) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sku) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Sku) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_sku_proto protoreflect.FileDescriptor

var file_sku_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x03, 0x53,
	0x6b, 0x75, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0x29,
	0x0a, 0x0a, 0x53, 0x6b, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x0d,
	0x64, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x04, 0x2e,
	0x53, 0x6b, 0x75, 0x1a, 0x04, 0x2e, 0x53, 0x6b, 0x75, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sku_proto_rawDescOnce sync.Once
	file_sku_proto_rawDescData = file_sku_proto_rawDesc
)

func file_sku_proto_rawDescGZIP() []byte {
	file_sku_proto_rawDescOnce.Do(func() {
		file_sku_proto_rawDescData = protoimpl.X.CompressGZIP(file_sku_proto_rawDescData)
	})
	return file_sku_proto_rawDescData
}

var file_sku_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sku_proto_goTypes = []interface{}{
	(*Sku)(nil), // 0: Sku
}
var file_sku_proto_depIdxs = []int32{
	0, // 0: SkuService.decreaseStock:input_type -> Sku
	0, // 1: SkuService.decreaseStock:output_type -> Sku
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sku_proto_init() }
func file_sku_proto_init() {
	if File_sku_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sku); i {
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
			RawDescriptor: file_sku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sku_proto_goTypes,
		DependencyIndexes: file_sku_proto_depIdxs,
		MessageInfos:      file_sku_proto_msgTypes,
	}.Build()
	File_sku_proto = out.File
	file_sku_proto_rawDesc = nil
	file_sku_proto_goTypes = nil
	file_sku_proto_depIdxs = nil
}
