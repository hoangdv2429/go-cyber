// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cyber/rank/v1beta1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Params struct {
	CalculationPeriod int64                                  `protobuf:"varint,1,opt,name=calculation_period,json=calculationPeriod,proto3" json:"calculation_period,omitempty"`
	DampingFactor     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=damping_factor,json=dampingFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"damping_factor"`
	Tolerance         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=tolerance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tolerance"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e38a68eff7d025f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCalculationPeriod() int64 {
	if m != nil {
		return m.CalculationPeriod
	}
	return 0
}

type RankedParticle struct {
	Particle string `protobuf:"bytes,1,opt,name=particle,proto3" json:"particle,omitempty"`
	Rank     uint64 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (m *RankedParticle) Reset()         { *m = RankedParticle{} }
func (m *RankedParticle) String() string { return proto.CompactTextString(m) }
func (*RankedParticle) ProtoMessage()    {}
func (*RankedParticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e38a68eff7d025f, []int{1}
}
func (m *RankedParticle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RankedParticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RankedParticle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RankedParticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankedParticle.Merge(m, src)
}
func (m *RankedParticle) XXX_Size() int {
	return m.Size()
}
func (m *RankedParticle) XXX_DiscardUnknown() {
	xxx_messageInfo_RankedParticle.DiscardUnknown(m)
}

var xxx_messageInfo_RankedParticle proto.InternalMessageInfo

func (m *RankedParticle) GetParticle() string {
	if m != nil {
		return m.Particle
	}
	return ""
}

func (m *RankedParticle) GetRank() uint64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "cyber.rank.v1beta1.Params")
	proto.RegisterType((*RankedParticle)(nil), "cyber.rank.v1beta1.RankedParticle")
}

func init() { proto.RegisterFile("cyber/rank/v1beta1/types.proto", fileDescriptor_5e38a68eff7d025f) }

var fileDescriptor_5e38a68eff7d025f = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0xad, 0x3f, 0x21, 0xfc, 0x11, 0x34, 0x50, 0xd1, 0x21, 0x64, 0x50, 0x42, 0x86, 0x92,
	0x25, 0x16, 0xa1, 0x5b, 0xa7, 0x12, 0x4a, 0xe9, 0xd0, 0x21, 0x18, 0xba, 0x74, 0x09, 0xb2, 0xac,
	0xaa, 0x26, 0xb6, 0xae, 0x91, 0x94, 0xd2, 0xbc, 0x45, 0x1f, 0xa1, 0x8f, 0x93, 0x31, 0x63, 0xe9,
	0x10, 0x8a, 0xb3, 0xf4, 0x31, 0x8a, 0x65, 0xd3, 0x66, 0xee, 0xa4, 0x23, 0xbe, 0xcb, 0xe1, 0xdc,
	0x73, 0x31, 0x15, 0x9b, 0x58, 0x1a, 0x66, 0xb8, 0x5e, 0xb1, 0xe7, 0x59, 0x2c, 0x1d, 0x9f, 0x31,
	0xb7, 0x29, 0xa4, 0x0d, 0x0b, 0x03, 0x0e, 0x08, 0xf1, 0x3c, 0xac, 0x78, 0xd8, 0xf0, 0xc1, 0x99,
	0x02, 0x05, 0x1e, 0xb3, 0x4a, 0xd5, 0x93, 0xe3, 0x12, 0xe1, 0xce, 0x82, 0x1b, 0x9e, 0x5b, 0x32,
	0xc5, 0x44, 0xf0, 0x4c, 0xac, 0x33, 0xee, 0x52, 0xd0, 0xcb, 0x42, 0x9a, 0x14, 0x92, 0x3e, 0x1a,
	0xa1, 0x49, 0x2b, 0x3a, 0x3d, 0x22, 0x0b, 0x0f, 0xc8, 0x3d, 0xee, 0x25, 0x3c, 0x2f, 0x52, 0xad,
	0x96, 0x8f, 0x5c, 0x38, 0x30, 0xfd, 0x7f, 0x23, 0x34, 0xe9, 0xce, 0xc3, 0xed, 0x7e, 0x18, 0x7c,
	0xec, 0x87, 0xe7, 0x2a, 0x75, 0x4f, 0xeb, 0x38, 0x14, 0x90, 0x33, 0x01, 0x36, 0x07, 0xdb, 0x3c,
	0x53, 0x9b, 0xac, 0x9a, 0xb4, 0xd7, 0x52, 0x44, 0x27, 0x8d, 0xcb, 0x8d, 0x37, 0x21, 0x77, 0xb8,
	0xeb, 0x20, 0x93, 0x86, 0x6b, 0x21, 0xfb, 0xad, 0x3f, 0x39, 0xfe, 0x1a, 0x5c, 0xb6, 0xbf, 0xde,
	0x86, 0x68, 0x7c, 0x85, 0x7b, 0x11, 0xd7, 0x2b, 0x99, 0x2c, 0xb8, 0x71, 0xa9, 0xc8, 0x24, 0x19,
	0xe0, 0xff, 0x45, 0xa3, 0xfd, 0x86, 0xdd, 0xe8, 0xe7, 0x4f, 0x08, 0x6e, 0x57, 0xc5, 0xf9, 0x75,
	0xda, 0x91, 0xd7, 0xf3, 0xdb, 0x6d, 0x49, 0xd1, 0xae, 0xa4, 0xe8, 0xb3, 0xa4, 0xe8, 0xf5, 0x40,
	0x83, 0xdd, 0x81, 0x06, 0xef, 0x07, 0x1a, 0x3c, 0x84, 0xc7, 0xa1, 0xaa, 0xd6, 0x05, 0x68, 0x65,
	0xa4, 0xb5, 0x4c, 0xc1, 0xb4, 0x3e, 0xd3, 0x4b, 0x7d, 0x28, 0x1f, 0x30, 0xee, 0xf8, 0xde, 0x2f,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x00, 0x90, 0x8d, 0xc3, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.CalculationPeriod != that1.CalculationPeriod {
		return false
	}
	if !this.DampingFactor.Equal(that1.DampingFactor) {
		return false
	}
	if !this.Tolerance.Equal(that1.Tolerance) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Tolerance.Size()
		i -= size
		if _, err := m.Tolerance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.DampingFactor.Size()
		i -= size
		if _, err := m.DampingFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.CalculationPeriod != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CalculationPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RankedParticle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RankedParticle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RankedParticle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rank != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Particle) > 0 {
		i -= len(m.Particle)
		copy(dAtA[i:], m.Particle)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Particle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CalculationPeriod != 0 {
		n += 1 + sovTypes(uint64(m.CalculationPeriod))
	}
	l = m.DampingFactor.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Tolerance.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *RankedParticle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Particle)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovTypes(uint64(m.Rank))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CalculationPeriod", wireType)
			}
			m.CalculationPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CalculationPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DampingFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DampingFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tolerance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tolerance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RankedParticle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RankedParticle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RankedParticle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Particle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Particle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
