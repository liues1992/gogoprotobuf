// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumcustomname.proto

package enumcustomname

/*
	Package enumcustomname tests the behavior of enum_customname and
	enumvalue_customname extensions.
*/

import proto "github.com/liues1992/gogoprotobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/liues1992/gogoprotobuf/gogoproto"
import test "github.com/liues1992/gogoprotobuf/test"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MyCustomEnum int32

const (
	// The following field will take on the custom name and the prefix, joined
	// by an underscore.
	MyCustomEnum_MyBetterNameA MyCustomEnum = 0
	MyCustomEnum_B             MyCustomEnum = 1
)

var MyCustomEnum_name = map[int32]string{
	0: "A",
	1: "B",
}
var MyCustomEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyCustomEnum) Enum() *MyCustomEnum {
	p := new(MyCustomEnum)
	*p = x
	return p
}
func (x MyCustomEnum) String() string {
	return proto.EnumName(MyCustomEnum_name, int32(x))
}
func (x *MyCustomEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomEnum_value, data, "MyCustomEnum")
	if err != nil {
		return err
	}
	*x = MyCustomEnum(value)
	return nil
}
func (MyCustomEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enumcustomname_d428393be9e73607, []int{0}
}

type MyCustomUnprefixedEnum int32

const (
	MyBetterNameUnprefixedA MyCustomUnprefixedEnum = 0
	UNPREFIXED_B            MyCustomUnprefixedEnum = 1
)

var MyCustomUnprefixedEnum_name = map[int32]string{
	0: "UNPREFIXED_A",
	1: "UNPREFIXED_B",
}
var MyCustomUnprefixedEnum_value = map[string]int32{
	"UNPREFIXED_A": 0,
	"UNPREFIXED_B": 1,
}

func (x MyCustomUnprefixedEnum) Enum() *MyCustomUnprefixedEnum {
	p := new(MyCustomUnprefixedEnum)
	*p = x
	return p
}
func (x MyCustomUnprefixedEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyCustomUnprefixedEnum_name, int32(x))
}
func (x *MyCustomUnprefixedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomUnprefixedEnum_value, data, "MyCustomUnprefixedEnum")
	if err != nil {
		return err
	}
	*x = MyCustomUnprefixedEnum(value)
	return nil
}
func (MyCustomUnprefixedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enumcustomname_d428393be9e73607, []int{1}
}

type MyEnumWithEnumStringer int32

const (
	MyEnumWithEnumStringer_EnumValueStringerA MyEnumWithEnumStringer = 0
	MyEnumWithEnumStringer_STRINGER_B         MyEnumWithEnumStringer = 1
)

var MyEnumWithEnumStringer_name = map[int32]string{
	0: "STRINGER_A",
	1: "STRINGER_B",
}
var MyEnumWithEnumStringer_value = map[string]int32{
	"STRINGER_A": 0,
	"STRINGER_B": 1,
}

func (x MyEnumWithEnumStringer) Enum() *MyEnumWithEnumStringer {
	p := new(MyEnumWithEnumStringer)
	*p = x
	return p
}
func (x MyEnumWithEnumStringer) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyEnumWithEnumStringer_name, int32(x))
}
func (x *MyEnumWithEnumStringer) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyEnumWithEnumStringer_value, data, "MyEnumWithEnumStringer")
	if err != nil {
		return err
	}
	*x = MyEnumWithEnumStringer(value)
	return nil
}
func (MyEnumWithEnumStringer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enumcustomname_d428393be9e73607, []int{2}
}

type OnlyEnums struct {
	MyEnum                         *MyCustomEnum               `protobuf:"varint,1,opt,name=my_enum,json=myEnum,enum=enumcustomname.MyCustomEnum" json:"my_enum,omitempty"`
	MyEnumDefaultA                 *MyCustomEnum               `protobuf:"varint,2,opt,name=my_enum_default_a,json=myEnumDefaultA,enum=enumcustomname.MyCustomEnum,def=0" json:"my_enum_default_a,omitempty"`
	MyEnumDefaultB                 *MyCustomEnum               `protobuf:"varint,3,opt,name=my_enum_default_b,json=myEnumDefaultB,enum=enumcustomname.MyCustomEnum,def=1" json:"my_enum_default_b,omitempty"`
	MyUnprefixedEnum               *MyCustomUnprefixedEnum     `protobuf:"varint,4,opt,name=my_unprefixed_enum,json=myUnprefixedEnum,enum=enumcustomname.MyCustomUnprefixedEnum" json:"my_unprefixed_enum,omitempty"`
	MyUnprefixedEnumDefaultA       *MyCustomUnprefixedEnum     `protobuf:"varint,5,opt,name=my_unprefixed_enum_default_a,json=myUnprefixedEnumDefaultA,enum=enumcustomname.MyCustomUnprefixedEnum,def=0" json:"my_unprefixed_enum_default_a,omitempty"`
	MyUnprefixedEnumDefaultB       *MyCustomUnprefixedEnum     `protobuf:"varint,6,opt,name=my_unprefixed_enum_default_b,json=myUnprefixedEnumDefaultB,enum=enumcustomname.MyCustomUnprefixedEnum,def=1" json:"my_unprefixed_enum_default_b,omitempty"`
	YetAnotherTestEnum             *test.YetAnotherTestEnum    `protobuf:"varint,7,opt,name=yet_another_test_enum,json=yetAnotherTestEnum,enum=test.YetAnotherTestEnum" json:"yet_another_test_enum,omitempty"`
	YetAnotherTestEnumDefaultAa    *test.YetAnotherTestEnum    `protobuf:"varint,8,opt,name=yet_another_test_enum_default_aa,json=yetAnotherTestEnumDefaultAa,enum=test.YetAnotherTestEnum,def=0" json:"yet_another_test_enum_default_aa,omitempty"`
	YetAnotherTestEnumDefaultBb    *test.YetAnotherTestEnum    `protobuf:"varint,9,opt,name=yet_another_test_enum_default_bb,json=yetAnotherTestEnumDefaultBb,enum=test.YetAnotherTestEnum,def=1" json:"yet_another_test_enum_default_bb,omitempty"`
	YetYetAnotherTestEnum          *test.YetYetAnotherTestEnum `protobuf:"varint,10,opt,name=yet_yet_another_test_enum,json=yetYetAnotherTestEnum,enum=test.YetYetAnotherTestEnum" json:"yet_yet_another_test_enum,omitempty"`
	YetYetAnotherTestEnumDefaultCc *test.YetYetAnotherTestEnum `protobuf:"varint,11,opt,name=yet_yet_another_test_enum_default_cc,json=yetYetAnotherTestEnumDefaultCc,enum=test.YetYetAnotherTestEnum,def=0" json:"yet_yet_another_test_enum_default_cc,omitempty"`
	YetYetAnotherTestEnumDefaultDd *test.YetYetAnotherTestEnum `protobuf:"varint,12,opt,name=yet_yet_another_test_enum_default_dd,json=yetYetAnotherTestEnumDefaultDd,enum=test.YetYetAnotherTestEnum,def=1" json:"yet_yet_another_test_enum_default_dd,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                    `json:"-"`
	XXX_unrecognized               []byte                      `json:"-"`
	XXX_sizecache                  int32                       `json:"-"`
}

func (m *OnlyEnums) Reset()         { *m = OnlyEnums{} }
func (m *OnlyEnums) String() string { return proto.CompactTextString(m) }
func (*OnlyEnums) ProtoMessage()    {}
func (*OnlyEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_enumcustomname_d428393be9e73607, []int{0}
}
func (m *OnlyEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlyEnums.Unmarshal(m, b)
}
func (m *OnlyEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlyEnums.Marshal(b, m, deterministic)
}
func (dst *OnlyEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlyEnums.Merge(dst, src)
}
func (m *OnlyEnums) XXX_Size() int {
	return xxx_messageInfo_OnlyEnums.Size(m)
}
func (m *OnlyEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlyEnums.DiscardUnknown(m)
}

var xxx_messageInfo_OnlyEnums proto.InternalMessageInfo

const Default_OnlyEnums_MyEnumDefaultA MyCustomEnum = MyCustomEnum_MyBetterNameA
const Default_OnlyEnums_MyEnumDefaultB MyCustomEnum = MyCustomEnum_B
const Default_OnlyEnums_MyUnprefixedEnumDefaultA MyCustomUnprefixedEnum = MyBetterNameUnprefixedA
const Default_OnlyEnums_MyUnprefixedEnumDefaultB MyCustomUnprefixedEnum = UNPREFIXED_B
const Default_OnlyEnums_YetAnotherTestEnumDefaultAa test.YetAnotherTestEnum = test.AA
const Default_OnlyEnums_YetAnotherTestEnumDefaultBb test.YetAnotherTestEnum = test.BetterYetBB
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_CC
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_BetterYetDD

func (m *OnlyEnums) GetMyEnum() MyCustomEnum {
	if m != nil && m.MyEnum != nil {
		return *m.MyEnum
	}
	return MyCustomEnum_MyBetterNameA
}

func (m *OnlyEnums) GetMyEnumDefaultA() MyCustomEnum {
	if m != nil && m.MyEnumDefaultA != nil {
		return *m.MyEnumDefaultA
	}
	return Default_OnlyEnums_MyEnumDefaultA
}

func (m *OnlyEnums) GetMyEnumDefaultB() MyCustomEnum {
	if m != nil && m.MyEnumDefaultB != nil {
		return *m.MyEnumDefaultB
	}
	return Default_OnlyEnums_MyEnumDefaultB
}

func (m *OnlyEnums) GetMyUnprefixedEnum() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnum != nil {
		return *m.MyUnprefixedEnum
	}
	return MyBetterNameUnprefixedA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultA() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultA != nil {
		return *m.MyUnprefixedEnumDefaultA
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultB() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultB != nil {
		return *m.MyUnprefixedEnumDefaultB
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultB
}

func (m *OnlyEnums) GetYetAnotherTestEnum() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnum != nil {
		return *m.YetAnotherTestEnum
	}
	return test.AA
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultAa() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultAa != nil {
		return *m.YetAnotherTestEnumDefaultAa
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultAa
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultBb() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultBb != nil {
		return *m.YetAnotherTestEnumDefaultBb
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultBb
}

func (m *OnlyEnums) GetYetYetAnotherTestEnum() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnum != nil {
		return *m.YetYetAnotherTestEnum
	}
	return test.YetYetAnotherTestEnum_CC
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultCc() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultCc != nil {
		return *m.YetYetAnotherTestEnumDefaultCc
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultDd() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultDd != nil {
		return *m.YetYetAnotherTestEnumDefaultDd
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd
}

func init() {
	proto.RegisterType((*OnlyEnums)(nil), "enumcustomname.OnlyEnums")
	proto.RegisterEnum("enumcustomname.MyCustomEnum", MyCustomEnum_name, MyCustomEnum_value)
	proto.RegisterEnum("enumcustomname.MyCustomUnprefixedEnum", MyCustomUnprefixedEnum_name, MyCustomUnprefixedEnum_value)
	proto.RegisterEnum("enumcustomname.MyEnumWithEnumStringer", MyEnumWithEnumStringer_name, MyEnumWithEnumStringer_value)
}
func (x MyEnumWithEnumStringer) String() string {
	s, ok := MyEnumWithEnumStringer_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() {
	proto.RegisterFile("enumcustomname.proto", fileDescriptor_enumcustomname_d428393be9e73607)
}

var fileDescriptor_enumcustomname_d428393be9e73607 = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0xc6, 0x29, 0xba, 0x2c, 0x3b, 0x22, 0xe9, 0x4e, 0x14, 0x47, 0x30, 0x4d, 0xb3, 0x31, 0xc6,
	0x60, 0x16, 0x12, 0x8f, 0x78, 0x9a, 0x52, 0x34, 0x1b, 0x03, 0x9a, 0xee, 0xe2, 0xbf, 0x4b, 0xd3,
	0x96, 0xe1, 0x4f, 0xc2, 0xb4, 0x9b, 0x32, 0x8d, 0xf6, 0x1b, 0x18, 0xbe, 0x03, 0x27, 0x39, 0x78,
	0xf4, 0xbc, 0x67, 0x3f, 0x98, 0x99, 0xe9, 0xc2, 0x42, 0x5b, 0x0a, 0xf1, 0x34, 0xed, 0x9b, 0xe7,
	0x7d, 0x7e, 0xf3, 0x3e, 0x79, 0x07, 0x3c, 0x22, 0x6e, 0x40, 0x9d, 0x60, 0xc6, 0x3c, 0xea, 0x5a,
	0x94, 0x34, 0xae, 0x7d, 0x8f, 0x79, 0xb0, 0xbc, 0x5d, 0xad, 0x9e, 0x8f, 0x26, 0x6c, 0x1c, 0xd8,
	0x0d, 0xc7, 0xa3, 0xcd, 0x91, 0x37, 0xf2, 0x9a, 0x42, 0x66, 0x07, 0x43, 0xf1, 0x27, 0x7e, 0xc4,
	0x57, 0xd4, 0x5e, 0x7d, 0xb5, 0x53, 0xce, 0xc8, 0x8c, 0x35, 0xd9, 0x98, 0xf0, 0x33, 0x12, 0x9f,
	0xfd, 0x2d, 0x82, 0x93, 0x0f, 0xee, 0x34, 0xec, 0xb8, 0x01, 0x9d, 0xc1, 0x26, 0x38, 0xa6, 0xa1,
	0xc9, 0xf1, 0x48, 0x52, 0xa5, 0x97, 0xe5, 0xd7, 0x95, 0x46, 0xec, 0x86, 0x5d, 0xa1, 0x34, 0x0a,
	0x54, 0x9c, 0x50, 0x07, 0xa7, 0xb7, 0x0d, 0xe6, 0x80, 0x0c, 0xad, 0x60, 0xca, 0x4c, 0x0b, 0xe5,
	0xb3, 0x5a, 0x5b, 0x12, 0x36, 0xca, 0x51, 0xb7, 0x1e, 0x75, 0xe0, 0x34, 0x17, 0x1b, 0xdd, 0xcb,
	0x76, 0xd1, 0x62, 0x2e, 0x1a, 0xec, 0x01, 0x48, 0x43, 0x33, 0x70, 0xaf, 0x7d, 0x32, 0x9c, 0xfc,
	0x20, 0x83, 0x68, 0x8e, 0xfb, 0xc2, 0x46, 0x4d, 0xda, 0xf4, 0xd7, 0x42, 0x31, 0x91, 0x4c, 0x63,
	0x15, 0xe8, 0x82, 0x67, 0x49, 0xbf, 0x8d, 0x31, 0x8f, 0x0e, 0x73, 0x6e, 0x95, 0xfa, 0xbd, 0x8f,
	0x46, 0xe7, 0xed, 0xc5, 0x97, 0x8e, 0x6e, 0x62, 0x03, 0xc5, 0x39, 0xeb, 0x14, 0xb2, 0x79, 0x36,
	0x2a, 0xfc, 0x07, 0x4f, 0xdb, 0xc9, 0xd3, 0xe0, 0x7b, 0xf0, 0x38, 0x24, 0xcc, 0xb4, 0x5c, 0x8f,
	0x8d, 0x89, 0x6f, 0xf2, 0xa5, 0x88, 0x22, 0x3b, 0x16, 0x20, 0xd4, 0x10, 0x6b, 0xf2, 0x95, 0x30,
	0x1c, 0x29, 0xae, 0xc8, 0x8c, 0x89, 0xa8, 0x60, 0x98, 0xa8, 0x41, 0x07, 0xa8, 0xa9, 0x66, 0x77,
	0x79, 0x59, 0xa8, 0x98, 0xed, 0xdb, 0xca, 0x63, 0x6c, 0xd4, 0x92, 0xde, 0xab, 0x80, 0xac, 0xfd,
	0x10, 0xdb, 0x46, 0x27, 0xfb, 0x20, 0x9a, 0x96, 0x01, 0xd1, 0x6c, 0xd8, 0x07, 0x4f, 0x39, 0x24,
	0x3d, 0x1a, 0x20, 0xdc, 0x6b, 0x6b, 0xf7, 0x94, 0x74, 0x78, 0xa8, 0xc9, 0x32, 0xa4, 0xe0, 0xf9,
	0x4e, 0xdb, 0xf5, 0xfd, 0x1d, 0x07, 0x3d, 0xd8, 0x4b, 0x68, 0xe5, 0xdb, 0x6d, 0x43, 0x49, 0xa5,
	0xdc, 0x4e, 0xd1, 0x76, 0x0e, 0xc3, 0x0d, 0x06, 0xa8, 0x74, 0x00, 0x4e, 0xd7, 0xb3, 0x71, 0xfa,
	0xa0, 0xfe, 0x06, 0x14, 0xa2, 0x87, 0x09, 0x11, 0x90, 0xb0, 0x9c, 0xab, 0x9e, 0xce, 0x17, 0xea,
	0xc3, 0x6e, 0xa8, 0x11, 0xc6, 0x88, 0xdf, 0xb3, 0x28, 0xc1, 0xf0, 0x08, 0x48, 0x9a, 0x2c, 0x55,
	0xe5, 0x9b, 0xa5, 0x52, 0xea, 0x86, 0x6d, 0xb1, 0xc1, 0xbc, 0xa5, 0xfe, 0x1d, 0xc8, 0xf1, 0x25,
	0x86, 0xe7, 0x60, 0xeb, 0xd9, 0xc8, 0xb9, 0x6a, 0x6d, 0xbe, 0x50, 0x9f, 0x6c, 0x3a, 0xde, 0x75,
	0x60, 0x28, 0x6f, 0xc9, 0x39, 0xe6, 0xec, 0xe7, 0x2f, 0x25, 0xf7, 0x7b, 0xa9, 0xe4, 0x6e, 0x96,
	0x4a, 0x65, 0x85, 0xdb, 0x86, 0xd4, 0xbf, 0x81, 0x4a, 0x74, 0xeb, 0xcf, 0x13, 0x36, 0xe6, 0xe7,
	0x25, 0xf3, 0x27, 0xee, 0x88, 0xf8, 0xf0, 0x05, 0x00, 0x97, 0x57, 0xc6, 0x45, 0xef, 0x5d, 0xc7,
	0x10, 0xf0, 0xca, 0x7c, 0xa1, 0x42, 0xae, 0xf8, 0x64, 0x4d, 0x03, 0xb2, 0x92, 0x61, 0x58, 0xde,
	0xd0, 0x71, 0x6a, 0x91, 0x13, 0xff, 0x2c, 0x15, 0xe9, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0x65, 0x55, 0xe7, 0xdb, 0x05, 0x00, 0x00,
}
