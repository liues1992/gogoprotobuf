// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: import_public/sub/b.proto

package sub // import "github.com/liues1992/gogoprotobuf/protoc-gen-gogog/testdata/import_public/sub"

import proto "github.com/liues1992/gogoprotobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_b_d16d7ba92a37c489, []int{0}
}
func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (dst *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(dst, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "goproto.test.import_public.sub.M2")
}

func init() { proto.RegisterFile("import_public/sub/b.proto", fileDescriptor_b_d16d7ba92a37c489) }

var fileDescriptor_b_d16d7ba92a37c489 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x89, 0x2f, 0x28, 0x4d, 0xca, 0xc9, 0x4c, 0xd6, 0x2f, 0x2e, 0x4d, 0xd2, 0x4f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4b, 0xcf, 0x07, 0x33, 0xf4, 0x4a, 0x52, 0x8b, 0x4b,
	0xf4, 0x50, 0xd4, 0xe9, 0x15, 0x97, 0x26, 0x29, 0xb1, 0x70, 0x31, 0xf9, 0x1a, 0x39, 0xb9, 0x46,
	0x39, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0x75, 0x25, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0x60, 0x09,
	0x90, 0x41, 0x29, 0x89, 0x25, 0x89, 0xfa, 0x18, 0x96, 0x26, 0xb1, 0x81, 0xd5, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0xe4, 0xbf, 0x60, 0x90, 0x00, 0x00, 0x00,
}
