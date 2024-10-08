// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: ip/v1/ip.proto

package ip

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

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Addr:
	//
	//	*IP_V4
	//	*IP_V6
	Addr isIP_Addr `protobuf_oneof:"addr"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ip_v1_ip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_ip_v1_ip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_ip_v1_ip_proto_rawDescGZIP(), []int{0}
}

func (m *IP) GetAddr() isIP_Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (x *IP) GetV4() *IPv4 {
	if x, ok := x.GetAddr().(*IP_V4); ok {
		return x.V4
	}
	return nil
}

func (x *IP) GetV6() *IPv6 {
	if x, ok := x.GetAddr().(*IP_V6); ok {
		return x.V6
	}
	return nil
}

type isIP_Addr interface {
	isIP_Addr()
}

type IP_V4 struct {
	V4 *IPv4 `protobuf:"bytes,1,opt,name=v4,proto3,oneof"`
}

type IP_V6 struct {
	V6 *IPv6 `protobuf:"bytes,2,opt,name=v6,proto3,oneof"`
}

func (*IP_V4) isIP_Addr() {}

func (*IP_V6) isIP_Addr() {}

type AddrPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr *IP    `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AddrPort) Reset() {
	*x = AddrPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ip_v1_ip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddrPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddrPort) ProtoMessage() {}

func (x *AddrPort) ProtoReflect() protoreflect.Message {
	mi := &file_ip_v1_ip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddrPort.ProtoReflect.Descriptor instead.
func (*AddrPort) Descriptor() ([]byte, []int) {
	return file_ip_v1_ip_proto_rawDescGZIP(), []int{1}
}

func (x *AddrPort) GetAddr() *IP {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *AddrPort) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Prefix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr *IP   `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Bits int32 `protobuf:"varint,2,opt,name=bits,proto3" json:"bits,omitempty"`
}

func (x *Prefix) Reset() {
	*x = Prefix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ip_v1_ip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prefix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prefix) ProtoMessage() {}

func (x *Prefix) ProtoReflect() protoreflect.Message {
	mi := &file_ip_v1_ip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prefix.ProtoReflect.Descriptor instead.
func (*Prefix) Descriptor() ([]byte, []int) {
	return file_ip_v1_ip_proto_rawDescGZIP(), []int{2}
}

func (x *Prefix) GetAddr() *IP {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *Prefix) GetBits() int32 {
	if x != nil {
		return x.Bits
	}
	return 0
}

type IPv4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip uint32 `protobuf:"fixed32,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IPv4) Reset() {
	*x = IPv4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ip_v1_ip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4) ProtoMessage() {}

func (x *IPv4) ProtoReflect() protoreflect.Message {
	mi := &file_ip_v1_ip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4.ProtoReflect.Descriptor instead.
func (*IPv4) Descriptor() ([]byte, []int) {
	return file_ip_v1_ip_proto_rawDescGZIP(), []int{3}
}

func (x *IPv4) GetIp() uint32 {
	if x != nil {
		return x.Ip
	}
	return 0
}

type IPv6 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hi uint64 `protobuf:"fixed64,1,opt,name=hi,proto3" json:"hi,omitempty"`
	Lo uint64 `protobuf:"fixed64,2,opt,name=lo,proto3" json:"lo,omitempty"`
}

func (x *IPv6) Reset() {
	*x = IPv6{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ip_v1_ip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv6) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv6) ProtoMessage() {}

func (x *IPv6) ProtoReflect() protoreflect.Message {
	mi := &file_ip_v1_ip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv6.ProtoReflect.Descriptor instead.
func (*IPv6) Descriptor() ([]byte, []int) {
	return file_ip_v1_ip_proto_rawDescGZIP(), []int{4}
}

func (x *IPv6) GetHi() uint64 {
	if x != nil {
		return x.Hi
	}
	return 0
}

func (x *IPv6) GetLo() uint64 {
	if x != nil {
		return x.Lo
	}
	return 0
}

var File_ip_v1_ip_proto protoreflect.FileDescriptor

var file_ip_v1_ip_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x22, 0x4a, 0x0a, 0x02, 0x49, 0x50, 0x12, 0x1d, 0x0a,
	0x02, 0x76, 0x34, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x48, 0x00, 0x52, 0x02, 0x76, 0x34, 0x12, 0x1d, 0x0a, 0x02,
	0x76, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x50, 0x76, 0x36, 0x48, 0x00, 0x52, 0x02, 0x76, 0x36, 0x42, 0x06, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x22, 0x3d, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x1d, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x69, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x3b, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x69, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x22,
	0x16, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x07, 0x52, 0x02, 0x69, 0x70, 0x22, 0x26, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x36, 0x12,
	0x0e, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x02, 0x68, 0x69, 0x12,
	0x0e, 0x0a, 0x02, 0x6c, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x02, 0x6c, 0x6f, 0x42,
	0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2d, 0x70, 0x62, 0x2f, 0x64, 0x74, 0x2f, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x69,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ip_v1_ip_proto_rawDescOnce sync.Once
	file_ip_v1_ip_proto_rawDescData = file_ip_v1_ip_proto_rawDesc
)

func file_ip_v1_ip_proto_rawDescGZIP() []byte {
	file_ip_v1_ip_proto_rawDescOnce.Do(func() {
		file_ip_v1_ip_proto_rawDescData = protoimpl.X.CompressGZIP(file_ip_v1_ip_proto_rawDescData)
	})
	return file_ip_v1_ip_proto_rawDescData
}

var file_ip_v1_ip_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ip_v1_ip_proto_goTypes = []any{
	(*IP)(nil),       // 0: ip.v1.IP
	(*AddrPort)(nil), // 1: ip.v1.AddrPort
	(*Prefix)(nil),   // 2: ip.v1.Prefix
	(*IPv4)(nil),     // 3: ip.v1.IPv4
	(*IPv6)(nil),     // 4: ip.v1.IPv6
}
var file_ip_v1_ip_proto_depIdxs = []int32{
	3, // 0: ip.v1.IP.v4:type_name -> ip.v1.IPv4
	4, // 1: ip.v1.IP.v6:type_name -> ip.v1.IPv6
	0, // 2: ip.v1.AddrPort.addr:type_name -> ip.v1.IP
	0, // 3: ip.v1.Prefix.addr:type_name -> ip.v1.IP
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ip_v1_ip_proto_init() }
func file_ip_v1_ip_proto_init() {
	if File_ip_v1_ip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ip_v1_ip_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IP); i {
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
		file_ip_v1_ip_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddrPort); i {
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
		file_ip_v1_ip_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Prefix); i {
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
		file_ip_v1_ip_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IPv4); i {
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
		file_ip_v1_ip_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*IPv6); i {
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
	file_ip_v1_ip_proto_msgTypes[0].OneofWrappers = []any{
		(*IP_V4)(nil),
		(*IP_V6)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ip_v1_ip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ip_v1_ip_proto_goTypes,
		DependencyIndexes: file_ip_v1_ip_proto_depIdxs,
		MessageInfos:      file_ip_v1_ip_proto_msgTypes,
	}.Build()
	File_ip_v1_ip_proto = out.File
	file_ip_v1_ip_proto_rawDesc = nil
	file_ip_v1_ip_proto_goTypes = nil
	file_ip_v1_ip_proto_depIdxs = nil
}
