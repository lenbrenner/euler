// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/matilda.proto

package api

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

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_matilda_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_api_matilda_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_api_matilda_proto_rawDescGZIP(), []int{0}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Square struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"` //Todo - add transitions and make this more like Location
}

func (x *Square) Reset() {
	*x = Square{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_matilda_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Square) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Square) ProtoMessage() {}

func (x *Square) ProtoReflect() protoreflect.Message {
	mi := &file_api_matilda_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Square.ProtoReflect.Descriptor instead.
func (*Square) Descriptor() ([]byte, []int) {
	return file_api_matilda_proto_rawDescGZIP(), []int{1}
}

func (x *Square) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Square) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_api_matilda_proto protoreflect.FileDescriptor

var file_api_matilda_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x69, 0x6c, 0x64, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x44, 0x0a, 0x06, 0x53,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x32, 0x31, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x69, 0x6c, 0x64, 0x61, 0x12, 0x26, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x12, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x22, 0x00, 0x42, 0x39, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6b, 0x65,
	0x6f, 0x66, 0x66, 0x2e, 0x6d, 0x61, 0x74, 0x69, 0x6c, 0x64, 0x61, 0x42, 0x07, 0x4d, 0x61, 0x74,
	0x69, 0x6c, 0x64, 0x61, 0x50, 0x01, 0x5a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x69, 0x6c, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_matilda_proto_rawDescOnce sync.Once
	file_api_matilda_proto_rawDescData = file_api_matilda_proto_rawDesc
)

func file_api_matilda_proto_rawDescGZIP() []byte {
	file_api_matilda_proto_rawDescOnce.Do(func() {
		file_api_matilda_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_matilda_proto_rawDescData)
	})
	return file_api_matilda_proto_rawDescData
}

var file_api_matilda_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_matilda_proto_goTypes = []interface{}{
	(*Point)(nil),  // 0: api.Point
	(*Square)(nil), // 1: api.Square
}
var file_api_matilda_proto_depIdxs = []int32{
	0, // 0: api.Square.location:type_name -> api.Point
	0, // 1: api.Matilda.GetSquare:input_type -> api.Point
	1, // 2: api.Matilda.GetSquare:output_type -> api.Square
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_matilda_proto_init() }
func file_api_matilda_proto_init() {
	if File_api_matilda_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_matilda_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_api_matilda_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Square); i {
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
			RawDescriptor: file_api_matilda_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_matilda_proto_goTypes,
		DependencyIndexes: file_api_matilda_proto_depIdxs,
		MessageInfos:      file_api_matilda_proto_msgTypes,
	}.Build()
	File_api_matilda_proto = out.File
	file_api_matilda_proto_rawDesc = nil
	file_api_matilda_proto_goTypes = nil
	file_api_matilda_proto_depIdxs = nil
}