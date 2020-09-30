// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: esdt.proto

package builtInFunctions

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ElrondNetwork_elrond_go_data "github.com/ElrondNetwork/elrond-go/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
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

// ESDigitalToken holds the data for a elrond standard digital token transaction
type ESDigitalToken struct {
	Value      *math_big.Int `protobuf:"bytes,1,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"value"`
	Properties []byte        `protobuf:"bytes,3,opt,name=Properties,proto3" json:"properties"`
}

func (m *ESDigitalToken) Reset()      { *m = ESDigitalToken{} }
func (*ESDigitalToken) ProtoMessage() {}
func (*ESDigitalToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_e413e402abc6a34c, []int{0}
}
func (m *ESDigitalToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ESDigitalToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ESDigitalToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ESDigitalToken.Merge(m, src)
}
func (m *ESDigitalToken) XXX_Size() int {
	return m.Size()
}
func (m *ESDigitalToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ESDigitalToken.DiscardUnknown(m)
}

var xxx_messageInfo_ESDigitalToken proto.InternalMessageInfo

func (m *ESDigitalToken) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ESDigitalToken) GetProperties() []byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*ESDigitalToken)(nil), "protoBuiltInFunctions.ESDigitalToken")
}

func init() { proto.RegisterFile("esdt.proto", fileDescriptor_e413e402abc6a34c) }

var fileDescriptor_e413e402abc6a34c = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4a, 0x33, 0x41,
	0x14, 0x85, 0x67, 0xf8, 0xc9, 0x5f, 0x0c, 0x12, 0x24, 0x20, 0x04, 0x8b, 0x1b, 0xb1, 0xb2, 0xc9,
	0x6e, 0x61, 0x69, 0xe5, 0x6a, 0x84, 0xb5, 0x08, 0x12, 0xc5, 0xc2, 0x6e, 0x36, 0x3b, 0x4e, 0x86,
	0x6c, 0x66, 0x96, 0xd9, 0xbb, 0xda, 0xfa, 0x08, 0x3e, 0x85, 0x88, 0x4f, 0x62, 0xb9, 0xe5, 0x56,
	0xd1, 0x9d, 0x6d, 0x24, 0x55, 0x1e, 0x41, 0x9c, 0x80, 0x04, 0xab, 0x7b, 0xbf, 0xc3, 0xe1, 0x1c,
	0xee, 0x65, 0x4c, 0x14, 0x29, 0x06, 0xb9, 0x35, 0x68, 0x7a, 0x7b, 0x7e, 0x44, 0xa5, 0xca, 0x30,
	0xd6, 0x17, 0xa5, 0x9e, 0xa2, 0x32, 0xba, 0xd8, 0x1f, 0x4a, 0x85, 0xb3, 0x32, 0x09, 0xa6, 0x66,
	0x11, 0x4a, 0x23, 0x4d, 0xe8, 0x6d, 0x49, 0x79, 0xef, 0xc9, 0x83, 0xdf, 0x36, 0x29, 0x87, 0x2f,
	0x94, 0x75, 0x47, 0xd7, 0xe7, 0x4a, 0x2a, 0xe4, 0xd9, 0x8d, 0x99, 0x0b, 0xdd, 0x4b, 0x59, 0xe7,
	0x96, 0x67, 0xa5, 0xe8, 0xd3, 0x03, 0x7a, 0xb4, 0x13, 0x8d, 0x57, 0xcb, 0x41, 0xe7, 0xe1, 0x47,
	0x78, 0xfb, 0x18, 0x9c, 0x2e, 0x38, 0xce, 0xc2, 0x44, 0xc9, 0x20, 0xd6, 0x78, 0xb2, 0x55, 0x35,
	0xca, 0xac, 0xd1, 0xe9, 0x58, 0xe0, 0xa3, 0xb1, 0xf3, 0x50, 0x78, 0x1a, 0x4a, 0x13, 0xa6, 0x1c,
	0x79, 0x10, 0x29, 0x19, 0x6b, 0x3c, 0xe3, 0x05, 0x0a, 0x3b, 0xd9, 0x84, 0xf7, 0x02, 0xc6, 0xae,
	0xac, 0xc9, 0x85, 0x45, 0x25, 0x8a, 0xfe, 0x3f, 0x5f, 0xd5, 0x5d, 0x2d, 0x07, 0x2c, 0xff, 0x55,
	0x27, 0x5b, 0x8e, 0xe8, 0xb2, 0x6a, 0x80, 0xd4, 0x0d, 0x90, 0x75, 0x03, 0xf4, 0xc9, 0x01, 0x7d,
	0x75, 0x40, 0xdf, 0x1d, 0xd0, 0xca, 0x01, 0xad, 0x1d, 0xd0, 0x4f, 0x07, 0xf4, 0xcb, 0x01, 0x59,
	0x3b, 0xa0, 0xcf, 0x2d, 0x90, 0xaa, 0x05, 0x52, 0xb7, 0x40, 0xee, 0x76, 0x93, 0x3f, 0x3f, 0x4a,
	0xfe, 0xfb, 0xdb, 0x8f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x3c, 0x43, 0xb7, 0x4f, 0x01,
	0x00, 0x00,
}

func (this *ESDigitalToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ESDigitalToken)
	if !ok {
		that2, ok := that.(ESDigitalToken)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.Properties, that1.Properties) {
		return false
	}
	return true
}
func (this *ESDigitalToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&builtInFunctions.ESDigitalToken{")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Properties: "+fmt.Sprintf("%#v", this.Properties)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringEsdt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ESDigitalToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ESDigitalToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ESDigitalToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintEsdt(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEsdt(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEsdt(dAtA []byte, offset int, v uint64) int {
	offset -= sovEsdt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ESDigitalToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovEsdt(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovEsdt(uint64(l))
	}
	return n
}

func sovEsdt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEsdt(x uint64) (n int) {
	return sovEsdt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ESDigitalToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ESDigitalToken{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Properties:` + fmt.Sprintf("%v", this.Properties) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEsdt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ESDigitalToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEsdt
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
			return fmt.Errorf("proto: ESDigitalToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ESDigitalToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEsdt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEsdt
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEsdt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties[:0], dAtA[iNdEx:postIndex]...)
			if m.Properties == nil {
				m.Properties = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEsdt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEsdt
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthEsdt
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
func skipEsdt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
					return 0, ErrIntOverflowEsdt
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
				return 0, ErrInvalidLengthEsdt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEsdt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEsdt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEsdt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEsdt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEsdt = fmt.Errorf("proto: unexpected end of group")
)
