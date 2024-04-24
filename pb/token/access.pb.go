// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: access.proto

package token

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GenReq) Reset() {
	*x = GenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenReq) ProtoMessage() {}

func (x *GenReq) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenReq.ProtoReflect.Descriptor instead.
func (*GenReq) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{0}
}

func (x *GenReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type TokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenReq) Reset() {
	*x = TokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq) ProtoMessage() {}

func (x *TokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq.ProtoReflect.Descriptor instead.
func (*TokenReq) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{2}
}

func (x *TokenReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TokenResp) Reset() {
	*x = TokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResp) ProtoMessage() {}

func (x *TokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_access_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResp.ProtoReflect.Descriptor instead.
func (*TokenResp) Descriptor() ([]byte, []int) {
	return file_access_proto_rawDescGZIP(), []int{3}
}

func (x *TokenResp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TokenResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TokenResp) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_access_proto protoreflect.FileDescriptor

var file_access_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x20, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x38, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4d, 0x0a, 0x09, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xd2, 0x01, 0x0a, 0x06, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0b, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x0b, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_access_proto_rawDescOnce sync.Once
	file_access_proto_rawDescData = file_access_proto_rawDesc
)

func file_access_proto_rawDescGZIP() []byte {
	file_access_proto_rawDescOnce.Do(func() {
		file_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_access_proto_rawDescData)
	})
	return file_access_proto_rawDescData
}

var file_access_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_access_proto_goTypes = []interface{}{
	(*GenReq)(nil),    // 0: access.GenReq
	(*Result)(nil),    // 1: access.Result
	(*TokenReq)(nil),  // 2: access.TokenReq
	(*TokenResp)(nil), // 3: access.TokenResp
}
var file_access_proto_depIdxs = []int32{
	0, // 0: access.Access.GenToken:input_type -> access.GenReq
	2, // 1: access.Access.FreezeToken:input_type -> access.TokenReq
	2, // 2: access.Access.ResetToken:input_type -> access.TokenReq
	2, // 3: access.Access.VerifyToken:input_type -> access.TokenReq
	3, // 4: access.Access.GenToken:output_type -> access.TokenResp
	3, // 5: access.Access.FreezeToken:output_type -> access.TokenResp
	3, // 6: access.Access.ResetToken:output_type -> access.TokenResp
	3, // 7: access.Access.VerifyToken:output_type -> access.TokenResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_access_proto_init() }
func file_access_proto_init() {
	if File_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenReq); i {
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
		file_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq); i {
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
		file_access_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenResp); i {
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
			RawDescriptor: file_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_access_proto_goTypes,
		DependencyIndexes: file_access_proto_depIdxs,
		MessageInfos:      file_access_proto_msgTypes,
	}.Build()
	File_access_proto = out.File
	file_access_proto_rawDesc = nil
	file_access_proto_goTypes = nil
	file_access_proto_depIdxs = nil
}
