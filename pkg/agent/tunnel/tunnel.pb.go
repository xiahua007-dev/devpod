// protoc -I . tunnel.proto  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.3
// source: tunnel.proto

package tunnel

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

type LogLevel int32

const (
	LogLevel_DEBUG   LogLevel = 0
	LogLevel_INFO    LogLevel = 1
	LogLevel_DONE    LogLevel = 2
	LogLevel_WARNING LogLevel = 3
	LogLevel_ERROR   LogLevel = 4
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "DONE",
		3: "WARNING",
		4: "ERROR",
	}
	LogLevel_value = map[string]int32{
		"DEBUG":   0,
		"INFO":    1,
		"DONE":    2,
		"WARNING": 3,
		"ERROR":   4,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_tunnel_proto_enumTypes[0].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_tunnel_proto_enumTypes[0]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{1}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type LogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogLevel LogLevel `protobuf:"varint,1,opt,name=logLevel,proto3,enum=tunnel.LogLevel" json:"logLevel,omitempty"`
	Message  string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogMessage) Reset() {
	*x = LogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessage) ProtoMessage() {}

func (x *LogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessage.ProtoReflect.Descriptor instead.
func (*LogMessage) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{2}
}

func (x *LogMessage) GetLogLevel() LogLevel {
	if x != nil {
		return x.LogLevel
	}
	return LogLevel_DEBUG
}

func (x *LogMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tunnel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_tunnel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_tunnel_proto_rawDescGZIP(), []int{3}
}

var File_tunnel_proto protoreflect.FileDescriptor

var file_tunnel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x43,
	0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x54,
	0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10,
	0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2a, 0x41, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04,
	0x32, 0xdb, 0x02, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x2e, 0x74, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d,
	0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x0d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x2e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x0f, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x11, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x0f, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x47,
	0x69, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x0f, 0x2e,
	0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0f,
	0x2e, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x74,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0f, 0x2e, 0x74, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x66,
	0x74, 0x2d, 0x73, 0x68, 0x2f, 0x64, 0x65, 0x76, 0x70, 0x6f, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tunnel_proto_rawDescOnce sync.Once
	file_tunnel_proto_rawDescData = file_tunnel_proto_rawDesc
)

func file_tunnel_proto_rawDescGZIP() []byte {
	file_tunnel_proto_rawDescOnce.Do(func() {
		file_tunnel_proto_rawDescData = protoimpl.X.CompressGZIP(file_tunnel_proto_rawDescData)
	})
	return file_tunnel_proto_rawDescData
}

var file_tunnel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tunnel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tunnel_proto_goTypes = []interface{}{
	(LogLevel)(0),      // 0: tunnel.LogLevel
	(*Message)(nil),    // 1: tunnel.Message
	(*Chunk)(nil),      // 2: tunnel.Chunk
	(*LogMessage)(nil), // 3: tunnel.LogMessage
	(*Empty)(nil),      // 4: tunnel.Empty
}
var file_tunnel_proto_depIdxs = []int32{
	0, // 0: tunnel.LogMessage.logLevel:type_name -> tunnel.LogLevel
	4, // 1: tunnel.Tunnel.Ping:input_type -> tunnel.Empty
	3, // 2: tunnel.Tunnel.Log:input_type -> tunnel.LogMessage
	4, // 3: tunnel.Tunnel.ReadWorkspace:input_type -> tunnel.Empty
	1, // 4: tunnel.Tunnel.SendResult:input_type -> tunnel.Message
	1, // 5: tunnel.Tunnel.DockerCredentials:input_type -> tunnel.Message
	1, // 6: tunnel.Tunnel.GitCredentials:input_type -> tunnel.Message
	4, // 7: tunnel.Tunnel.GitUser:input_type -> tunnel.Empty
	4, // 8: tunnel.Tunnel.Ping:output_type -> tunnel.Empty
	4, // 9: tunnel.Tunnel.Log:output_type -> tunnel.Empty
	2, // 10: tunnel.Tunnel.ReadWorkspace:output_type -> tunnel.Chunk
	4, // 11: tunnel.Tunnel.SendResult:output_type -> tunnel.Empty
	1, // 12: tunnel.Tunnel.DockerCredentials:output_type -> tunnel.Message
	1, // 13: tunnel.Tunnel.GitCredentials:output_type -> tunnel.Message
	1, // 14: tunnel.Tunnel.GitUser:output_type -> tunnel.Message
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tunnel_proto_init() }
func file_tunnel_proto_init() {
	if File_tunnel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tunnel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_tunnel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_tunnel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMessage); i {
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
		file_tunnel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_tunnel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tunnel_proto_goTypes,
		DependencyIndexes: file_tunnel_proto_depIdxs,
		EnumInfos:         file_tunnel_proto_enumTypes,
		MessageInfos:      file_tunnel_proto_msgTypes,
	}.Build()
	File_tunnel_proto = out.File
	file_tunnel_proto_rawDesc = nil
	file_tunnel_proto_goTypes = nil
	file_tunnel_proto_depIdxs = nil
}
