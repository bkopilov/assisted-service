// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/openshift/api/cloudnetwork/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func (m *CloudPrivateIPConfig) Reset()      { *m = CloudPrivateIPConfig{} }
func (*CloudPrivateIPConfig) ProtoMessage() {}
func (*CloudPrivateIPConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_454253a7ab01c6d0, []int{0}
}
func (m *CloudPrivateIPConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloudPrivateIPConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CloudPrivateIPConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudPrivateIPConfig.Merge(m, src)
}
func (m *CloudPrivateIPConfig) XXX_Size() int {
	return m.Size()
}
func (m *CloudPrivateIPConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudPrivateIPConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CloudPrivateIPConfig proto.InternalMessageInfo

func (m *CloudPrivateIPConfigList) Reset()      { *m = CloudPrivateIPConfigList{} }
func (*CloudPrivateIPConfigList) ProtoMessage() {}
func (*CloudPrivateIPConfigList) Descriptor() ([]byte, []int) {
	return fileDescriptor_454253a7ab01c6d0, []int{1}
}
func (m *CloudPrivateIPConfigList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloudPrivateIPConfigList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CloudPrivateIPConfigList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudPrivateIPConfigList.Merge(m, src)
}
func (m *CloudPrivateIPConfigList) XXX_Size() int {
	return m.Size()
}
func (m *CloudPrivateIPConfigList) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudPrivateIPConfigList.DiscardUnknown(m)
}

var xxx_messageInfo_CloudPrivateIPConfigList proto.InternalMessageInfo

func (m *CloudPrivateIPConfigSpec) Reset()      { *m = CloudPrivateIPConfigSpec{} }
func (*CloudPrivateIPConfigSpec) ProtoMessage() {}
func (*CloudPrivateIPConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_454253a7ab01c6d0, []int{2}
}
func (m *CloudPrivateIPConfigSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloudPrivateIPConfigSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CloudPrivateIPConfigSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudPrivateIPConfigSpec.Merge(m, src)
}
func (m *CloudPrivateIPConfigSpec) XXX_Size() int {
	return m.Size()
}
func (m *CloudPrivateIPConfigSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudPrivateIPConfigSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CloudPrivateIPConfigSpec proto.InternalMessageInfo

func (m *CloudPrivateIPConfigStatus) Reset()      { *m = CloudPrivateIPConfigStatus{} }
func (*CloudPrivateIPConfigStatus) ProtoMessage() {}
func (*CloudPrivateIPConfigStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_454253a7ab01c6d0, []int{3}
}
func (m *CloudPrivateIPConfigStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CloudPrivateIPConfigStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *CloudPrivateIPConfigStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloudPrivateIPConfigStatus.Merge(m, src)
}
func (m *CloudPrivateIPConfigStatus) XXX_Size() int {
	return m.Size()
}
func (m *CloudPrivateIPConfigStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CloudPrivateIPConfigStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CloudPrivateIPConfigStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CloudPrivateIPConfig)(nil), "github.com.openshift.api.cloudnetwork.v1.CloudPrivateIPConfig")
	proto.RegisterType((*CloudPrivateIPConfigList)(nil), "github.com.openshift.api.cloudnetwork.v1.CloudPrivateIPConfigList")
	proto.RegisterType((*CloudPrivateIPConfigSpec)(nil), "github.com.openshift.api.cloudnetwork.v1.CloudPrivateIPConfigSpec")
	proto.RegisterType((*CloudPrivateIPConfigStatus)(nil), "github.com.openshift.api.cloudnetwork.v1.CloudPrivateIPConfigStatus")
}

func init() {
	proto.RegisterFile("github.com/openshift/api/cloudnetwork/v1/generated.proto", fileDescriptor_454253a7ab01c6d0)
}

var fileDescriptor_454253a7ab01c6d0 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0xae, 0x9b, 0x86, 0x07, 0x08, 0x45, 0x1c, 0xa2, 0x1e, 0xbc, 0xaa, 0xa7, 0x5e,
	0xb0, 0xe9, 0x84, 0xd0, 0x0e, 0x88, 0x43, 0xca, 0x65, 0x12, 0x8c, 0x29, 0xdc, 0x10, 0x07, 0x5c,
	0xc7, 0x4d, 0x4d, 0x17, 0x3b, 0x8a, 0x9d, 0x22, 0x6e, 0x3c, 0x02, 0xef, 0xc0, 0xcb, 0xf4, 0xc0,
	0x61, 0xc7, 0x5d, 0x98, 0x68, 0x78, 0x11, 0x64, 0x37, 0x6d, 0x23, 0xd6, 0x69, 0x91, 0x7a, 0xcb,
	0xf7, 0x25, 0xff, 0xff, 0xef, 0xfb, 0xfe, 0x8e, 0x0c, 0x4f, 0x13, 0x61, 0x26, 0xc5, 0x08, 0x33,
	0x95, 0x12, 0x95, 0x71, 0xa9, 0x27, 0x62, 0x6c, 0x08, 0xcd, 0x04, 0x61, 0x97, 0xaa, 0x88, 0x25,
	0x37, 0x5f, 0x55, 0x3e, 0x25, 0xb3, 0x01, 0x49, 0xb8, 0xe4, 0x39, 0x35, 0x3c, 0xc6, 0x59, 0xae,
	0x8c, 0xf2, 0xfb, 0x1b, 0x25, 0x5e, 0x2b, 0x31, 0xcd, 0x04, 0xae, 0x2b, 0xf1, 0x6c, 0xd0, 0x79,
	0x56, 0x63, 0x24, 0x2a, 0x51, 0xc4, 0x19, 0x8c, 0x8a, 0xb1, 0xab, 0x5c, 0xe1, 0x9e, 0x96, 0xc6,
	0x9d, 0x17, 0xd3, 0x53, 0x8d, 0x85, 0xb2, 0x43, 0xa4, 0x94, 0x4d, 0x84, 0xe4, 0xf9, 0x37, 0x92,
	0x4d, 0x13, 0xdb, 0xd0, 0x24, 0xe5, 0x86, 0x6e, 0x19, 0xa7, 0x43, 0xee, 0x52, 0xe5, 0x85, 0x34,
	0x22, 0xe5, 0xb7, 0x04, 0x2f, 0xef, 0x13, 0x68, 0x36, 0xe1, 0x29, 0xfd, 0x5f, 0xd7, 0xfb, 0xd5,
	0x82, 0x4f, 0x87, 0x76, 0xc3, 0x8b, 0x5c, 0xcc, 0xa8, 0xe1, 0x67, 0x17, 0x43, 0x25, 0xc7, 0x22,
	0xf1, 0x3f, 0xc3, 0x43, 0x3b, 0x5c, 0x4c, 0x0d, 0x0d, 0x40, 0x17, 0xf4, 0x8f, 0x4e, 0x9e, 0xe3,
	0x25, 0x03, 0xd7, 0x19, 0x38, 0x9b, 0x26, 0xb6, 0xa1, 0xb1, 0xfd, 0x1a, 0xcf, 0x06, 0xf8, 0xfd,
	0xe8, 0x0b, 0x67, 0xe6, 0x1d, 0x37, 0x34, 0xf4, 0xe7, 0x37, 0xc7, 0x5e, 0x79, 0x73, 0x0c, 0x37,
	0xbd, 0x68, 0xed, 0xea, 0xc7, 0xb0, 0xad, 0x33, 0xce, 0x82, 0x96, 0x73, 0x0f, 0x71, 0xd3, 0x13,
	0xc0, 0xdb, 0xe6, 0xfd, 0x90, 0x71, 0x16, 0x3e, 0xac, 0x78, 0x6d, 0x5b, 0x45, 0xce, 0xdd, 0xbf,
	0x84, 0x07, 0xda, 0x50, 0x53, 0xe8, 0x60, 0xcf, 0x71, 0xde, 0xec, 0xc8, 0x71, 0x5e, 0xe1, 0xe3,
	0x8a, 0x74, 0xb0, 0xac, 0xa3, 0x8a, 0xd1, 0xfb, 0x0d, 0x60, 0xb0, 0x4d, 0xf6, 0x56, 0x68, 0xe3,
	0x7f, 0xba, 0x15, 0x29, 0x6e, 0x16, 0xa9, 0x55, 0xbb, 0x40, 0x9f, 0x54, 0xd8, 0xc3, 0x55, 0xa7,
	0x16, 0x27, 0x83, 0xfb, 0xc2, 0xf0, 0x54, 0x07, 0xad, 0xee, 0x5e, 0xff, 0xe8, 0xe4, 0xf5, 0x6e,
	0x7b, 0x86, 0x8f, 0x2a, 0xd4, 0xfe, 0x99, 0x35, 0x8d, 0x96, 0xde, 0xbd, 0x57, 0xdb, 0xd7, 0xb3,
	0x79, 0xfb, 0x5d, 0xd8, 0x96, 0x2a, 0xe6, 0x6e, 0xb5, 0x07, 0x9b, 0xb3, 0x38, 0x57, 0x31, 0x8f,
	0xdc, 0x9b, 0xde, 0x4f, 0x00, 0x3b, 0x77, 0x87, 0x7a, 0xbf, 0x81, 0xcf, 0x20, 0x64, 0x4a, 0xc6,
	0xc2, 0x08, 0x25, 0x57, 0x8b, 0x92, 0x66, 0x19, 0x0e, 0x57, 0xba, 0xcd, 0x5f, 0xb9, 0x6e, 0xe9,
	0xa8, 0x66, 0x1b, 0x9e, 0xcf, 0x17, 0xc8, 0xbb, 0x5a, 0x20, 0xef, 0x7a, 0x81, 0xbc, 0xef, 0x25,
	0x02, 0xf3, 0x12, 0x81, 0xab, 0x12, 0x81, 0xeb, 0x12, 0x81, 0x3f, 0x25, 0x02, 0x3f, 0xfe, 0x22,
	0xef, 0x63, 0xbf, 0xe9, 0x55, 0xf3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xf0, 0xc5, 0x6e, 0x95,
	0x04, 0x00, 0x00,
}

func (m *CloudPrivateIPConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudPrivateIPConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudPrivateIPConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CloudPrivateIPConfigList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudPrivateIPConfigList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudPrivateIPConfigList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CloudPrivateIPConfigSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudPrivateIPConfigSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudPrivateIPConfigSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.Node)
	copy(dAtA[i:], m.Node)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Node)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CloudPrivateIPConfigStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CloudPrivateIPConfigStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CloudPrivateIPConfigStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Conditions) > 0 {
		for iNdEx := len(m.Conditions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Conditions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	i -= len(m.Node)
	copy(dAtA[i:], m.Node)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Node)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CloudPrivateIPConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CloudPrivateIPConfigList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *CloudPrivateIPConfigSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Node)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *CloudPrivateIPConfigStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Node)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Conditions) > 0 {
		for _, e := range m.Conditions {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CloudPrivateIPConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CloudPrivateIPConfig{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "CloudPrivateIPConfigSpec", "CloudPrivateIPConfigSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "CloudPrivateIPConfigStatus", "CloudPrivateIPConfigStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CloudPrivateIPConfigList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]CloudPrivateIPConfig{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "CloudPrivateIPConfig", "CloudPrivateIPConfig", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&CloudPrivateIPConfigList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *CloudPrivateIPConfigSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CloudPrivateIPConfigSpec{`,
		`Node:` + fmt.Sprintf("%v", this.Node) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CloudPrivateIPConfigStatus) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForConditions := "[]Condition{"
	for _, f := range this.Conditions {
		repeatedStringForConditions += fmt.Sprintf("%v", f) + ","
	}
	repeatedStringForConditions += "}"
	s := strings.Join([]string{`&CloudPrivateIPConfigStatus{`,
		`Node:` + fmt.Sprintf("%v", this.Node) + `,`,
		`Conditions:` + repeatedStringForConditions + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CloudPrivateIPConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloudPrivateIPConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudPrivateIPConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloudPrivateIPConfigList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloudPrivateIPConfigList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudPrivateIPConfigList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, CloudPrivateIPConfig{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloudPrivateIPConfigSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloudPrivateIPConfigSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudPrivateIPConfigSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CloudPrivateIPConfigStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CloudPrivateIPConfigStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CloudPrivateIPConfigStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Node", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Node = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Conditions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Conditions = append(m.Conditions, v1.Condition{})
			if err := m.Conditions[len(m.Conditions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)