// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking.proto

package systemSmartContracts

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

type StakedData struct {
	RegisterNonce uint64        `protobuf:"varint,1,opt,name=RegisterNonce,proto3" json:"RegisterNonce"`
	Staked        bool          `protobuf:"varint,2,opt,name=Staked,proto3" json:"Staked"`
	UnStakedNonce uint64        `protobuf:"varint,3,opt,name=UnStakedNonce,proto3" json:"UnStakedNonce"`
	UnStakedEpoch uint32        `protobuf:"varint,6,opt,name=UnStakedEpoch,proto3" json:"UnStakedEpoch"`
	RewardAddress []byte        `protobuf:"bytes,4,opt,name=RewardAddress,proto3" json:"RewardAddress"`
	StakeValue    *math_big.Int `protobuf:"bytes,5,opt,name=StakeValue,proto3,casttypewith=math/big.Int;github.com/ElrondNetwork/elrond-go/data.BigIntCaster" json:"StakeValue"`
	JailedRound   uint64        `protobuf:"varint,7,opt,name=JailedRound,proto3" json:"JailedRound"`
}

func (m *StakedData) Reset()      { *m = StakedData{} }
func (*StakedData) ProtoMessage() {}
func (*StakedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{0}
}
func (m *StakedData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakedData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakedData.Merge(m, src)
}
func (m *StakedData) XXX_Size() int {
	return m.Size()
}
func (m *StakedData) XXX_DiscardUnknown() {
	xxx_messageInfo_StakedData.DiscardUnknown(m)
}

var xxx_messageInfo_StakedData proto.InternalMessageInfo

func (m *StakedData) GetRegisterNonce() uint64 {
	if m != nil {
		return m.RegisterNonce
	}
	return 0
}

func (m *StakedData) GetStaked() bool {
	if m != nil {
		return m.Staked
	}
	return false
}

func (m *StakedData) GetUnStakedNonce() uint64 {
	if m != nil {
		return m.UnStakedNonce
	}
	return 0
}

func (m *StakedData) GetUnStakedEpoch() uint32 {
	if m != nil {
		return m.UnStakedEpoch
	}
	return 0
}

func (m *StakedData) GetRewardAddress() []byte {
	if m != nil {
		return m.RewardAddress
	}
	return nil
}

func (m *StakedData) GetStakeValue() *math_big.Int {
	if m != nil {
		return m.StakeValue
	}
	return nil
}

func (m *StakedData) GetJailedRound() uint64 {
	if m != nil {
		return m.JailedRound
	}
	return 0
}

func init() {
	proto.RegisterType((*StakedData)(nil), "proto.StakedData")
}

func init() { proto.RegisterFile("staking.proto", fileDescriptor_289e7c8aea278311) }

var fileDescriptor_289e7c8aea278311 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x3d, 0xee, 0xd3, 0x30,
	0x00, 0xc5, 0x63, 0xfa, 0x01, 0x0a, 0x54, 0x88, 0x88, 0x21, 0x62, 0x70, 0xa2, 0x4e, 0x59, 0x9a,
	0x08, 0x31, 0x30, 0x30, 0x35, 0xa5, 0x43, 0x3b, 0x74, 0x70, 0x05, 0x03, 0x9b, 0x13, 0x1b, 0x37,
	0x6a, 0x63, 0x57, 0xb6, 0xa3, 0x8a, 0x8d, 0x23, 0x70, 0x0c, 0xc4, 0x19, 0x38, 0x00, 0x63, 0xc7,
	0x4e, 0x81, 0xa6, 0x0b, 0xca, 0xd4, 0x23, 0xa0, 0x3a, 0x45, 0xff, 0x64, 0xb2, 0x7f, 0xef, 0xe9,
	0x3d, 0xc9, 0x4f, 0xb6, 0x47, 0x4a, 0xe3, 0x6d, 0xc6, 0x59, 0xb8, 0x97, 0x42, 0x0b, 0x67, 0x60,
	0x8e, 0x57, 0x13, 0x96, 0xe9, 0x4d, 0x91, 0x84, 0xa9, 0xc8, 0x23, 0x26, 0x98, 0x88, 0x8c, 0x9c,
	0x14, 0x9f, 0x0d, 0x19, 0x30, 0xb7, 0x26, 0x35, 0xfe, 0xd9, 0xb3, 0xed, 0xb5, 0xc6, 0x5b, 0x4a,
	0xde, 0x63, 0x8d, 0x9d, 0xb7, 0xf6, 0x08, 0x51, 0x96, 0x29, 0x4d, 0xe5, 0x4a, 0xf0, 0x94, 0xba,
	0xc0, 0x07, 0x41, 0x3f, 0x7e, 0x51, 0x97, 0x5e, 0xd7, 0x40, 0x5d, 0x74, 0xc6, 0xf6, 0xb0, 0xa9,
	0x71, 0x1f, 0xf9, 0x20, 0x78, 0x12, 0xdb, 0x75, 0xe9, 0xdd, 0x15, 0x74, 0x3f, 0x6f, 0xe5, 0x1f,
	0x78, 0x73, 0x6f, 0xca, 0x7b, 0x0f, 0xe5, 0x1d, 0x03, 0x75, 0xb1, 0x1d, 0x9c, 0xef, 0x45, 0xba,
	0x71, 0x87, 0x3e, 0x08, 0x46, 0xdd, 0xa0, 0x31, 0x50, 0x17, 0x9b, 0xe7, 0x1c, 0xb0, 0x24, 0x53,
	0x42, 0x24, 0x55, 0xca, 0xed, 0xfb, 0x20, 0x78, 0xf6, 0xff, 0x39, 0x2d, 0x03, 0x75, 0xd1, 0x51,
	0xf7, 0x55, 0x3e, 0xe2, 0x5d, 0x41, 0xdd, 0x81, 0x49, 0xad, 0xeb, 0xd2, 0x6b, 0xa9, 0x3f, 0x7e,
	0x7b, 0xd3, 0x1c, 0xeb, 0x4d, 0x94, 0x64, 0x2c, 0x5c, 0x70, 0xfd, 0xae, 0x35, 0xfc, 0x7c, 0x27,
	0x05, 0x27, 0x2b, 0xaa, 0x0f, 0x42, 0x6e, 0x23, 0x6a, 0x68, 0xc2, 0x44, 0x44, 0xb0, 0xc6, 0x61,
	0x9c, 0xb1, 0x05, 0xd7, 0x33, 0x7c, 0x1b, 0x0f, 0xb5, 0x0a, 0x9d, 0xd7, 0xf6, 0xd3, 0x25, 0xce,
	0x76, 0x94, 0x20, 0x51, 0x70, 0xe2, 0x3e, 0x36, 0xeb, 0x3c, 0xaf, 0x4b, 0xaf, 0x2d, 0xa3, 0x36,
	0xc4, 0xcb, 0xe3, 0x19, 0x5a, 0xa7, 0x33, 0xb4, 0xae, 0x67, 0x08, 0xbe, 0x56, 0x10, 0x7c, 0xaf,
	0x20, 0xf8, 0x55, 0x41, 0x70, 0xac, 0x20, 0xf8, 0x53, 0x41, 0xf0, 0xb7, 0x82, 0xd6, 0xb5, 0x82,
	0xe0, 0xdb, 0x05, 0x5a, 0xc7, 0x0b, 0xb4, 0x4e, 0x17, 0x68, 0x7d, 0x7a, 0xa9, 0xbe, 0x28, 0x4d,
	0xf3, 0x75, 0x8e, 0xa5, 0x9e, 0x09, 0xae, 0x25, 0x4e, 0xb5, 0x4a, 0x86, 0xe6, 0x47, 0xbc, 0xf9,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x27, 0xb4, 0x1a, 0x58, 0x02, 0x00, 0x00,
}

func (this *StakedData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StakedData)
	if !ok {
		that2, ok := that.(StakedData)
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
	if this.RegisterNonce != that1.RegisterNonce {
		return false
	}
	if this.Staked != that1.Staked {
		return false
	}
	if this.UnStakedNonce != that1.UnStakedNonce {
		return false
	}
	if this.UnStakedEpoch != that1.UnStakedEpoch {
		return false
	}
	if !bytes.Equal(this.RewardAddress, that1.RewardAddress) {
		return false
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		if !__caster.Equal(this.StakeValue, that1.StakeValue) {
			return false
		}
	}
	if this.JailedRound != that1.JailedRound {
		return false
	}
	return true
}
func (this *StakedData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&systemSmartContracts.StakedData{")
	s = append(s, "RegisterNonce: "+fmt.Sprintf("%#v", this.RegisterNonce)+",\n")
	s = append(s, "Staked: "+fmt.Sprintf("%#v", this.Staked)+",\n")
	s = append(s, "UnStakedNonce: "+fmt.Sprintf("%#v", this.UnStakedNonce)+",\n")
	s = append(s, "UnStakedEpoch: "+fmt.Sprintf("%#v", this.UnStakedEpoch)+",\n")
	s = append(s, "RewardAddress: "+fmt.Sprintf("%#v", this.RewardAddress)+",\n")
	s = append(s, "StakeValue: "+fmt.Sprintf("%#v", this.StakeValue)+",\n")
	s = append(s, "JailedRound: "+fmt.Sprintf("%#v", this.JailedRound)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStaking(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *StakedData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.JailedRound != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.JailedRound))
		i--
		dAtA[i] = 0x38
	}
	if m.UnStakedEpoch != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.UnStakedEpoch))
		i--
		dAtA[i] = 0x30
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		size := __caster.Size(m.StakeValue)
		i -= size
		if _, err := __caster.MarshalTo(m.StakeValue, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.RewardAddress) > 0 {
		i -= len(m.RewardAddress)
		copy(dAtA[i:], m.RewardAddress)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.RewardAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.UnStakedNonce != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.UnStakedNonce))
		i--
		dAtA[i] = 0x18
	}
	if m.Staked {
		i--
		if m.Staked {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.RegisterNonce != 0 {
		i = encodeVarintStaking(dAtA, i, uint64(m.RegisterNonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegisterNonce != 0 {
		n += 1 + sovStaking(uint64(m.RegisterNonce))
	}
	if m.Staked {
		n += 2
	}
	if m.UnStakedNonce != 0 {
		n += 1 + sovStaking(uint64(m.UnStakedNonce))
	}
	l = len(m.RewardAddress)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	{
		__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
		l = __caster.Size(m.StakeValue)
		n += 1 + l + sovStaking(uint64(l))
	}
	if m.UnStakedEpoch != 0 {
		n += 1 + sovStaking(uint64(m.UnStakedEpoch))
	}
	if m.JailedRound != 0 {
		n += 1 + sovStaking(uint64(m.JailedRound))
	}
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StakedData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StakedData{`,
		`RegisterNonce:` + fmt.Sprintf("%v", this.RegisterNonce) + `,`,
		`Staked:` + fmt.Sprintf("%v", this.Staked) + `,`,
		`UnStakedNonce:` + fmt.Sprintf("%v", this.UnStakedNonce) + `,`,
		`RewardAddress:` + fmt.Sprintf("%v", this.RewardAddress) + `,`,
		`StakeValue:` + fmt.Sprintf("%v", this.StakeValue) + `,`,
		`UnStakedEpoch:` + fmt.Sprintf("%v", this.UnStakedEpoch) + `,`,
		`JailedRound:` + fmt.Sprintf("%v", this.JailedRound) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStaking(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StakedData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
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
			return fmt.Errorf("proto: StakedData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakedData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisterNonce", wireType)
			}
			m.RegisterNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisterNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Staked = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnStakedNonce", wireType)
			}
			m.UnStakedNonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnStakedNonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAddress = append(m.RewardAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.RewardAddress == nil {
				m.RewardAddress = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeValue", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
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
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ElrondNetwork_elrond_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.StakeValue = tmp
				}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnStakedEpoch", wireType)
			}
			m.UnStakedEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnStakedEpoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedRound", wireType)
			}
			m.JailedRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailedRound |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStaking
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
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
					return 0, ErrIntOverflowStaking
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
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
