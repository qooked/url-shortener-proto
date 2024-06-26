// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: shortener_api/shortener_api.proto

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

type ShortenURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *ShortenURLRequest) Reset() {
	*x = ShortenURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_api_shortener_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLRequest) ProtoMessage() {}

func (x *ShortenURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_api_shortener_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLRequest.ProtoReflect.Descriptor instead.
func (*ShortenURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_api_shortener_api_proto_rawDescGZIP(), []int{0}
}

func (x *ShortenURLRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type ShortenURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*ShortenURLResponse_ShortenedURL
	//	*ShortenURLResponse_Error
	Result isShortenURLResponse_Result `protobuf_oneof:"result"`
}

func (x *ShortenURLResponse) Reset() {
	*x = ShortenURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_api_shortener_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenURLResponse) ProtoMessage() {}

func (x *ShortenURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_api_shortener_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenURLResponse.ProtoReflect.Descriptor instead.
func (*ShortenURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_api_shortener_api_proto_rawDescGZIP(), []int{1}
}

func (m *ShortenURLResponse) GetResult() isShortenURLResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ShortenURLResponse) GetShortenedURL() string {
	if x, ok := x.GetResult().(*ShortenURLResponse_ShortenedURL); ok {
		return x.ShortenedURL
	}
	return ""
}

func (x *ShortenURLResponse) GetError() string {
	if x, ok := x.GetResult().(*ShortenURLResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isShortenURLResponse_Result interface {
	isShortenURLResponse_Result()
}

type ShortenURLResponse_ShortenedURL struct {
	ShortenedURL string `protobuf:"bytes,1,opt,name=shortenedURL,proto3,oneof"`
}

type ShortenURLResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*ShortenURLResponse_ShortenedURL) isShortenURLResponse_Result() {}

func (*ShortenURLResponse_Error) isShortenURLResponse_Result() {}

type GetURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	URL string `protobuf:"bytes,1,opt,name=URL,proto3" json:"URL,omitempty"`
}

func (x *GetURLRequest) Reset() {
	*x = GetURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_api_shortener_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLRequest) ProtoMessage() {}

func (x *GetURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_api_shortener_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLRequest.ProtoReflect.Descriptor instead.
func (*GetURLRequest) Descriptor() ([]byte, []int) {
	return file_shortener_api_shortener_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetURLRequest) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

type GetURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*GetURLResponse_OriginalURL
	//	*GetURLResponse_Error
	Result isGetURLResponse_Result `protobuf_oneof:"result"`
}

func (x *GetURLResponse) Reset() {
	*x = GetURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortener_api_shortener_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURLResponse) ProtoMessage() {}

func (x *GetURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shortener_api_shortener_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURLResponse.ProtoReflect.Descriptor instead.
func (*GetURLResponse) Descriptor() ([]byte, []int) {
	return file_shortener_api_shortener_api_proto_rawDescGZIP(), []int{3}
}

func (m *GetURLResponse) GetResult() isGetURLResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *GetURLResponse) GetOriginalURL() string {
	if x, ok := x.GetResult().(*GetURLResponse_OriginalURL); ok {
		return x.OriginalURL
	}
	return ""
}

func (x *GetURLResponse) GetError() string {
	if x, ok := x.GetResult().(*GetURLResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isGetURLResponse_Result interface {
	isGetURLResponse_Result()
}

type GetURLResponse_OriginalURL struct {
	OriginalURL string `protobuf:"bytes,1,opt,name=originalURL,proto3,oneof"`
}

type GetURLResponse_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*GetURLResponse_OriginalURL) isGetURLResponse_Result() {}

func (*GetURLResponse_Error) isGetURLResponse_Result() {}

var File_shortener_api_shortener_api_proto protoreflect.FileDescriptor

var file_shortener_api_shortener_api_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x22,
	0x5c, 0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x64, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x21, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c,
	0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x52,
	0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x7d, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x55, 0x52, 0x4c, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x12,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortener_api_shortener_api_proto_rawDescOnce sync.Once
	file_shortener_api_shortener_api_proto_rawDescData = file_shortener_api_shortener_api_proto_rawDesc
)

func file_shortener_api_shortener_api_proto_rawDescGZIP() []byte {
	file_shortener_api_shortener_api_proto_rawDescOnce.Do(func() {
		file_shortener_api_shortener_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortener_api_shortener_api_proto_rawDescData)
	})
	return file_shortener_api_shortener_api_proto_rawDescData
}

var file_shortener_api_shortener_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shortener_api_shortener_api_proto_goTypes = []interface{}{
	(*ShortenURLRequest)(nil),  // 0: api.ShortenURLRequest
	(*ShortenURLResponse)(nil), // 1: api.ShortenURLResponse
	(*GetURLRequest)(nil),      // 2: api.GetURLRequest
	(*GetURLResponse)(nil),     // 3: api.GetURLResponse
}
var file_shortener_api_shortener_api_proto_depIdxs = []int32{
	0, // 0: api.Shortener.ShortenURL:input_type -> api.ShortenURLRequest
	2, // 1: api.Shortener.GetURL:input_type -> api.GetURLRequest
	1, // 2: api.Shortener.ShortenURL:output_type -> api.ShortenURLResponse
	3, // 3: api.Shortener.GetURL:output_type -> api.GetURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortener_api_shortener_api_proto_init() }
func file_shortener_api_shortener_api_proto_init() {
	if File_shortener_api_shortener_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortener_api_shortener_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLRequest); i {
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
		file_shortener_api_shortener_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenURLResponse); i {
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
		file_shortener_api_shortener_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLRequest); i {
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
		file_shortener_api_shortener_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURLResponse); i {
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
	file_shortener_api_shortener_api_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ShortenURLResponse_ShortenedURL)(nil),
		(*ShortenURLResponse_Error)(nil),
	}
	file_shortener_api_shortener_api_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetURLResponse_OriginalURL)(nil),
		(*GetURLResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortener_api_shortener_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shortener_api_shortener_api_proto_goTypes,
		DependencyIndexes: file_shortener_api_shortener_api_proto_depIdxs,
		MessageInfos:      file_shortener_api_shortener_api_proto_msgTypes,
	}.Build()
	File_shortener_api_shortener_api_proto = out.File
	file_shortener_api_shortener_api_proto_rawDesc = nil
	file_shortener_api_shortener_api_proto_goTypes = nil
	file_shortener_api_shortener_api_proto_depIdxs = nil
}
