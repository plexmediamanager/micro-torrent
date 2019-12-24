// Code generated by protoc-gen-go. DO NOT EDIT.
// source: torrent.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TorrentEmpty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentEmpty) Reset()         { *m = TorrentEmpty{} }
func (m *TorrentEmpty) String() string { return proto.CompactTextString(m) }
func (*TorrentEmpty) ProtoMessage()    {}
func (*TorrentEmpty) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{0}
}

func (m *TorrentEmpty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentEmpty.Unmarshal(m, b)
}
func (m *TorrentEmpty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentEmpty.Marshal(b, m, deterministic)
}
func (m *TorrentEmpty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentEmpty.Merge(m, src)
}
func (m *TorrentEmpty) XXX_Size() int {
	return xxx_messageInfo_TorrentEmpty.Size(m)
}
func (m *TorrentEmpty) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentEmpty.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentEmpty proto.InternalMessageInfo

type TorrentResponse struct {
	Result               []byte   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentResponse) Reset()         { *m = TorrentResponse{} }
func (m *TorrentResponse) String() string { return proto.CompactTextString(m) }
func (*TorrentResponse) ProtoMessage()    {}
func (*TorrentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{1}
}

func (m *TorrentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentResponse.Unmarshal(m, b)
}
func (m *TorrentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentResponse.Marshal(b, m, deterministic)
}
func (m *TorrentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentResponse.Merge(m, src)
}
func (m *TorrentResponse) XXX_Size() int {
	return xxx_messageInfo_TorrentResponse.Size(m)
}
func (m *TorrentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentResponse proto.InternalMessageInfo

func (m *TorrentResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type TorrentOptions struct {
	Options              map[string]string `protobuf:"bytes,1,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TorrentOptions) Reset()         { *m = TorrentOptions{} }
func (m *TorrentOptions) String() string { return proto.CompactTextString(m) }
func (*TorrentOptions) ProtoMessage()    {}
func (*TorrentOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{2}
}

func (m *TorrentOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentOptions.Unmarshal(m, b)
}
func (m *TorrentOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentOptions.Marshal(b, m, deterministic)
}
func (m *TorrentOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentOptions.Merge(m, src)
}
func (m *TorrentOptions) XXX_Size() int {
	return xxx_messageInfo_TorrentOptions.Size(m)
}
func (m *TorrentOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentOptions.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentOptions proto.InternalMessageInfo

func (m *TorrentOptions) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

type TorrentString struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentString) Reset()         { *m = TorrentString{} }
func (m *TorrentString) String() string { return proto.CompactTextString(m) }
func (*TorrentString) ProtoMessage()    {}
func (*TorrentString) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{3}
}

func (m *TorrentString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentString.Unmarshal(m, b)
}
func (m *TorrentString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentString.Marshal(b, m, deterministic)
}
func (m *TorrentString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentString.Merge(m, src)
}
func (m *TorrentString) XXX_Size() int {
	return xxx_messageInfo_TorrentString.Size(m)
}
func (m *TorrentString) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentString.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentString proto.InternalMessageInfo

func (m *TorrentString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TorrentStrings struct {
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentStrings) Reset()         { *m = TorrentStrings{} }
func (m *TorrentStrings) String() string { return proto.CompactTextString(m) }
func (*TorrentStrings) ProtoMessage()    {}
func (*TorrentStrings) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{4}
}

func (m *TorrentStrings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentStrings.Unmarshal(m, b)
}
func (m *TorrentStrings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentStrings.Marshal(b, m, deterministic)
}
func (m *TorrentStrings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentStrings.Merge(m, src)
}
func (m *TorrentStrings) XXX_Size() int {
	return xxx_messageInfo_TorrentStrings.Size(m)
}
func (m *TorrentStrings) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentStrings.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentStrings proto.InternalMessageInfo

func (m *TorrentStrings) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type TorrentFilesPriority struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Priority             uint64   `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Files                []string `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TorrentFilesPriority) Reset()         { *m = TorrentFilesPriority{} }
func (m *TorrentFilesPriority) String() string { return proto.CompactTextString(m) }
func (*TorrentFilesPriority) ProtoMessage()    {}
func (*TorrentFilesPriority) Descriptor() ([]byte, []int) {
	return fileDescriptor_5390c0f93d3a953f, []int{5}
}

func (m *TorrentFilesPriority) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TorrentFilesPriority.Unmarshal(m, b)
}
func (m *TorrentFilesPriority) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TorrentFilesPriority.Marshal(b, m, deterministic)
}
func (m *TorrentFilesPriority) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TorrentFilesPriority.Merge(m, src)
}
func (m *TorrentFilesPriority) XXX_Size() int {
	return xxx_messageInfo_TorrentFilesPriority.Size(m)
}
func (m *TorrentFilesPriority) XXX_DiscardUnknown() {
	xxx_messageInfo_TorrentFilesPriority.DiscardUnknown(m)
}

var xxx_messageInfo_TorrentFilesPriority proto.InternalMessageInfo

func (m *TorrentFilesPriority) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *TorrentFilesPriority) GetPriority() uint64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *TorrentFilesPriority) GetFiles() []string {
	if m != nil {
		return m.Files
	}
	return nil
}

func init() {
	proto.RegisterType((*TorrentEmpty)(nil), "proto.TorrentEmpty")
	proto.RegisterType((*TorrentResponse)(nil), "proto.TorrentResponse")
	proto.RegisterType((*TorrentOptions)(nil), "proto.TorrentOptions")
	proto.RegisterMapType((map[string]string)(nil), "proto.TorrentOptions.OptionsEntry")
	proto.RegisterType((*TorrentString)(nil), "proto.TorrentString")
	proto.RegisterType((*TorrentStrings)(nil), "proto.TorrentStrings")
	proto.RegisterType((*TorrentFilesPriority)(nil), "proto.TorrentFilesPriority")
}

func init() { proto.RegisterFile("torrent.proto", fileDescriptor_5390c0f93d3a953f) }

var fileDescriptor_5390c0f93d3a953f = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdf, 0x6f, 0xd3, 0x3e,
	0x10, 0xc0, 0xbf, 0x59, 0xb7, 0x7d, 0xd9, 0xad, 0x1d, 0xe0, 0x75, 0xa5, 0x74, 0x2f, 0x53, 0x24,
	0xa4, 0xf2, 0xd2, 0x87, 0xf1, 0x82, 0xa6, 0x82, 0x54, 0xfa, 0x63, 0xad, 0x54, 0x89, 0x2a, 0x9d,
	0xe0, 0x85, 0x97, 0x10, 0xae, 0x8d, 0xd5, 0xd4, 0x8e, 0x6c, 0x67, 0xa8, 0x7f, 0x00, 0x2f, 0xfc,
	0xd5, 0x28, 0x8e, 0xeb, 0x36, 0x68, 0x4c, 0x4a, 0xf6, 0x14, 0xdf, 0xf9, 0xee, 0x73, 0x67, 0xdf,
	0xe5, 0x0c, 0x35, 0xc5, 0x85, 0x40, 0xa6, 0x3a, 0xb1, 0xe0, 0x8a, 0x93, 0x23, 0xfd, 0x71, 0xcf,
	0xa0, 0x7a, 0x97, 0xe9, 0x87, 0xeb, 0x58, 0x6d, 0xdc, 0xb7, 0xf0, 0xdc, 0xc8, 0x1e, 0xca, 0x98,
	0x33, 0x89, 0xa4, 0x01, 0xc7, 0x02, 0x65, 0x12, 0xa9, 0xa6, 0x73, 0xe5, 0xb4, 0xab, 0x9e, 0x91,
	0xdc, 0xdf, 0x0e, 0x9c, 0x19, 0xdb, 0xcf, 0xb1, 0xa2, 0x9c, 0x49, 0xd2, 0x85, 0xff, 0x79, 0xb6,
	0x6c, 0x3a, 0x57, 0x95, 0xf6, 0xe9, 0xb5, 0x9b, 0x45, 0xeb, 0xe4, 0xed, 0x3a, 0xe6, 0x3b, 0x64,
	0x4a, 0x6c, 0xbc, 0xad, 0x4b, 0xeb, 0x06, 0xaa, 0xfb, 0x1b, 0xe4, 0x05, 0x54, 0x56, 0xb8, 0xd1,
	0x51, 0x4f, 0xbc, 0x74, 0x49, 0xea, 0x70, 0x74, 0xef, 0x47, 0x09, 0x36, 0x0f, 0xb4, 0x2e, 0x13,
	0x6e, 0x0e, 0xde, 0x3b, 0xee, 0x1b, 0xa8, 0x99, 0x18, 0x73, 0x25, 0x28, 0x5b, 0xee, 0x4c, 0x9d,
	0x3d, 0x53, 0xb7, 0x6d, 0x53, 0xce, 0xcc, 0x64, 0x7a, 0x3a, 0xbd, 0x95, 0x65, 0x7c, 0xe2, 0x19,
	0xc9, 0xfd, 0x06, 0x75, 0x63, 0x39, 0xa2, 0x11, 0xca, 0x99, 0xa0, 0x5c, 0x50, 0xb5, 0x21, 0x04,
	0x0e, 0x43, 0x5f, 0x86, 0x06, 0xab, 0xd7, 0xa4, 0x05, 0xcf, 0x62, 0xb3, 0xaf, 0x33, 0x3b, 0xf4,
	0xac, 0x9c, 0xe6, 0xb1, 0x48, 0x01, 0xcd, 0x8a, 0xc6, 0x67, 0xc2, 0xf5, 0xaf, 0xea, 0x2e, 0x11,
	0x14, 0xf7, 0x34, 0x40, 0xd2, 0x07, 0xd2, 0x8b, 0xe3, 0x88, 0x06, 0x7e, 0x7a, 0x05, 0x5f, 0x50,
	0x48, 0xca, 0x19, 0x39, 0xcf, 0x5f, 0xa0, 0x2e, 0x52, 0xab, 0x91, 0x57, 0x6e, 0x2b, 0xe5, 0xfe,
	0x47, 0x46, 0x70, 0xb1, 0x07, 0xe9, 0xcd, 0x26, 0x25, 0x39, 0x53, 0xb8, 0xdc, 0xe3, 0x7c, 0x4a,
	0x68, 0xf4, 0x63, 0xc2, 0x16, 0x5c, 0xac, 0xb5, 0x5c, 0x94, 0x76, 0x0b, 0x8d, 0x3d, 0xda, 0x4c,
	0xe0, 0x02, 0x05, 0xb2, 0x00, 0x65, 0x51, 0xd0, 0x07, 0x80, 0x29, 0x5f, 0xa6, 0xdd, 0x41, 0x51,
	0x92, 0x8b, 0x07, 0x9b, 0xeb, 0x11, 0xf7, 0x2e, 0x9c, 0x1a, 0xe5, 0x94, 0x4a, 0x55, 0x34, 0x78,
	0x1f, 0x5e, 0x1a, 0xe5, 0x4c, 0xf0, 0x18, 0x85, 0x4a, 0x73, 0xa8, 0xe7, 0xcd, 0xb3, 0xae, 0x7a,
	0x04, 0xd2, 0xb3, 0xff, 0xd7, 0x9d, 0xf0, 0x83, 0x15, 0x8a, 0xa7, 0x20, 0xfa, 0x9c, 0x29, 0x64,
	0xaa, 0x38, 0x62, 0x00, 0x64, 0x7b, 0x14, 0x8a, 0x01, 0xca, 0xb9, 0xf2, 0x15, 0x16, 0xa6, 0x0c,
	0xe1, 0x3c, 0x47, 0x19, 0xfb, 0x32, 0x2c, 0x71, 0x25, 0x5d, 0x3b, 0x82, 0x66, 0x7e, 0x22, 0xf1,
	0xef, 0xb2, 0x9a, 0x1f, 0xb5, 0xf5, 0x50, 0xb5, 0x74, 0x4b, 0xd4, 0x76, 0xc8, 0x64, 0x5d, 0xde,
	0x7d, 0x80, 0x11, 0xaa, 0xa2, 0xee, 0x23, 0x68, 0xe4, 0xdc, 0xbf, 0x52, 0x15, 0xea, 0x79, 0x51,
	0x90, 0xf3, 0xd1, 0x8e, 0x03, 0x0f, 0xfb, 0x21, 0x06, 0xab, 0x82, 0xfe, 0x3d, 0xdb, 0x9b, 0x1e,
	0xf6, 0x18, 0xe3, 0x09, 0x0b, 0x8a, 0x1e, 0xe5, 0x16, 0x5e, 0x19, 0xcd, 0x84, 0x05, 0x02, 0x7d,
	0x89, 0x76, 0xe6, 0x95, 0x05, 0x0d, 0xf0, 0x49, 0xa0, 0x09, 0xbc, 0xb6, 0x33, 0x52, 0x8d, 0xe9,
	0x32, 0x44, 0xa9, 0x4a, 0xa2, 0xc6, 0xd0, 0xdc, 0xa1, 0xa6, 0xfc, 0x67, 0x79, 0xd2, 0xd4, 0x56,
	0x7c, 0x8e, 0xfa, 0x69, 0xb0, 0x9c, 0xcb, 0xbc, 0x43, 0xee, 0xd9, 0xf8, 0x07, 0xed, 0xfb, 0xb1,
	0xd6, 0xbe, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xf9, 0xa5, 0x50, 0x9d, 0x07, 0x00, 0x00,
}