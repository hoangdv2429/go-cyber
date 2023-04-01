// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cyber/bandwidth/v1beta1/types.proto

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
	RecoveryPeriod    uint64                                 `protobuf:"varint,1,opt,name=recovery_period,json=recoveryPeriod,proto3" json:"recovery_period,omitempty"`
	AdjustPricePeriod uint64                                 `protobuf:"varint,2,opt,name=adjust_price_period,json=adjustPricePeriod,proto3" json:"adjust_price_period,omitempty"`
	BasePrice         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=base_price,json=basePrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_price"`
	BaseLoad          github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=base_load,json=baseLoad,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"base_load"`
	MaxBlockBandwidth uint64                                 `protobuf:"varint,5,opt,name=max_block_bandwidth,json=maxBlockBandwidth,proto3" json:"max_block_bandwidth,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76904de4f1717b1, []int{0}
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

func (m *Params) GetRecoveryPeriod() uint64 {
	if m != nil {
		return m.RecoveryPeriod
	}
	return 0
}

func (m *Params) GetAdjustPricePeriod() uint64 {
	if m != nil {
		return m.AdjustPricePeriod
	}
	return 0
}

func (m *Params) GetMaxBlockBandwidth() uint64 {
	if m != nil {
		return m.MaxBlockBandwidth
	}
	return 0
}

type NeuronBandwidth struct {
	Neuron           string `protobuf:"bytes,1,opt,name=neuron,proto3" json:"neuron,omitempty"`
	RemainedValue    uint64 `protobuf:"varint,2,opt,name=remained_value,json=remainedValue,proto3" json:"remained_value,omitempty"`
	LastUpdatedBlock uint64 `protobuf:"varint,3,opt,name=last_updated_block,json=lastUpdatedBlock,proto3" json:"last_updated_block,omitempty"`
	MaxValue         uint64 `protobuf:"varint,4,opt,name=max_value,json=maxValue,proto3" json:"max_value,omitempty"`
}

func (m *NeuronBandwidth) Reset()         { *m = NeuronBandwidth{} }
func (m *NeuronBandwidth) String() string { return proto.CompactTextString(m) }
func (*NeuronBandwidth) ProtoMessage()    {}
func (*NeuronBandwidth) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76904de4f1717b1, []int{1}
}
func (m *NeuronBandwidth) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NeuronBandwidth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NeuronBandwidth.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NeuronBandwidth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeuronBandwidth.Merge(m, src)
}
func (m *NeuronBandwidth) XXX_Size() int {
	return m.Size()
}
func (m *NeuronBandwidth) XXX_DiscardUnknown() {
	xxx_messageInfo_NeuronBandwidth.DiscardUnknown(m)
}

var xxx_messageInfo_NeuronBandwidth proto.InternalMessageInfo

type Price struct {
	Price github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price"`
}

func (m *Price) Reset()         { *m = Price{} }
func (m *Price) String() string { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()    {}
func (*Price) Descriptor() ([]byte, []int) {
	return fileDescriptor_b76904de4f1717b1, []int{2}
}
func (m *Price) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Price) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Price.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Price) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Price.Merge(m, src)
}
func (m *Price) XXX_Size() int {
	return m.Size()
}
func (m *Price) XXX_DiscardUnknown() {
	xxx_messageInfo_Price.DiscardUnknown(m)
}

var xxx_messageInfo_Price proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "cyber.bandwidth.v1beta1.Params")
	proto.RegisterType((*NeuronBandwidth)(nil), "cyber.bandwidth.v1beta1.NeuronBandwidth")
	proto.RegisterType((*Price)(nil), "cyber.bandwidth.v1beta1.Price")
}

func init() {
	proto.RegisterFile("cyber/bandwidth/v1beta1/types.proto", fileDescriptor_b76904de4f1717b1)
}

var fileDescriptor_b76904de4f1717b1 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x6e, 0xd4, 0x30,
	0x1c, 0xc6, 0xe3, 0x92, 0x3b, 0xdd, 0x59, 0xa2, 0x05, 0x83, 0x20, 0x02, 0x29, 0xa9, 0x0e, 0x01,
	0x1d, 0x68, 0xa2, 0xaa, 0x5b, 0xc7, 0xa8, 0x1b, 0x14, 0x9d, 0x22, 0x01, 0x12, 0x4b, 0xe4, 0xd8,
	0x56, 0x1a, 0x9a, 0xc4, 0x91, 0xed, 0x1c, 0xb9, 0x37, 0x60, 0xe4, 0x11, 0x6e, 0xe1, 0x09, 0x78,
	0x89, 0x8e, 0x1d, 0x11, 0x43, 0x85, 0xee, 0x96, 0x3e, 0x06, 0xb2, 0x9d, 0x5c, 0x3b, 0x77, 0x4a,
	0xfc, 0x7d, 0xbf, 0x7c, 0xf1, 0xff, 0xd3, 0x1f, 0xbe, 0x22, 0xcb, 0x8c, 0x89, 0x28, 0xc3, 0x35,
	0xfd, 0x5e, 0x50, 0x75, 0x1e, 0x2d, 0x8e, 0x32, 0xa6, 0xf0, 0x51, 0xa4, 0x96, 0x0d, 0x93, 0x61,
	0x23, 0xb8, 0xe2, 0xe8, 0xb9, 0x81, 0xc2, 0x2d, 0x14, 0xf6, 0xd0, 0x8b, 0xa7, 0x39, 0xcf, 0xb9,
	0x61, 0x22, 0xfd, 0x66, 0xf1, 0xd9, 0xef, 0x1d, 0x38, 0x9e, 0x63, 0x81, 0x2b, 0x89, 0xde, 0xc2,
	0x3d, 0xc1, 0x08, 0x5f, 0x30, 0xb1, 0x4c, 0x1b, 0x26, 0x0a, 0x4e, 0x3d, 0xb0, 0x0f, 0x0e, 0xdc,
	0x64, 0x77, 0x90, 0xe7, 0x46, 0x45, 0x21, 0x7c, 0x82, 0xe9, 0xb7, 0x56, 0xaa, 0xb4, 0x11, 0x05,
	0x61, 0x03, 0xbc, 0x63, 0xe0, 0xc7, 0xd6, 0x9a, 0x6b, 0xa7, 0xe7, 0xcf, 0x20, 0xcc, 0xb0, 0x64,
	0x96, 0xf6, 0x1e, 0xec, 0x83, 0x83, 0x69, 0x1c, 0x5e, 0x5e, 0x07, 0xce, 0xdf, 0xeb, 0xe0, 0x4d,
	0x5e, 0xa8, 0xf3, 0x36, 0x0b, 0x09, 0xaf, 0x22, 0xc2, 0x65, 0xc5, 0x65, 0xff, 0x38, 0x94, 0xf4,
	0xa2, 0x1f, 0xec, 0x94, 0x91, 0x64, 0xaa, 0x13, 0x4c, 0x28, 0x7a, 0x0f, 0xcd, 0x21, 0x2d, 0x39,
	0xa6, 0x9e, 0x7b, 0xaf, 0xb4, 0x89, 0x0e, 0xf8, 0xc0, 0xb1, 0x99, 0xa5, 0xc2, 0x5d, 0x9a, 0x95,
	0x9c, 0x5c, 0xa4, 0xdb, 0xd2, 0xbc, 0x91, 0x9d, 0xa5, 0xc2, 0x5d, 0xac, 0x9d, 0x78, 0x30, 0x4e,
	0xdc, 0x9b, 0x55, 0x00, 0x66, 0xbf, 0x00, 0xdc, 0xfb, 0xc8, 0x5a, 0xc1, 0xeb, 0xad, 0x83, 0x9e,
	0xc1, 0x71, 0x6d, 0x24, 0xd3, 0xda, 0x34, 0xe9, 0x4f, 0xe8, 0x35, 0xdc, 0x15, 0xac, 0xc2, 0x45,
	0xcd, 0x68, 0xba, 0xc0, 0x65, 0xcb, 0xfa, 0xa2, 0x1e, 0x0e, 0xea, 0x67, 0x2d, 0xa2, 0x77, 0x10,
	0x95, 0x58, 0xaa, 0xb4, 0x6d, 0x28, 0x56, 0x8c, 0xda, 0x1b, 0x99, 0xb2, 0xdc, 0xe4, 0x91, 0x76,
	0x3e, 0x59, 0xc3, 0xdc, 0x07, 0xbd, 0x84, 0x53, 0x7d, 0x6d, 0x9b, 0xe7, 0x1a, 0x68, 0x52, 0xe1,
	0xce, 0x44, 0x9d, 0x4c, 0x7e, 0xac, 0x02, 0xe7, 0x66, 0x15, 0x38, 0xb3, 0x2f, 0x70, 0x64, 0x3b,
	0x3b, 0x85, 0x23, 0xdb, 0x3e, 0xb8, 0x57, 0x5f, 0xf6, 0xe3, 0xdb, 0xe0, 0xf8, 0xec, 0x72, 0xed,
	0x83, 0xab, 0xb5, 0x0f, 0xfe, 0xad, 0x7d, 0xf0, 0x73, 0xe3, 0x3b, 0x57, 0x1b, 0xdf, 0xf9, 0xb3,
	0xf1, 0x9d, 0xaf, 0xc7, 0x77, 0x23, 0xf5, 0x2a, 0x12, 0x5e, 0xe7, 0x82, 0x49, 0x19, 0xe5, 0xfc,
	0xd0, 0x2e, 0x70, 0x77, 0x67, 0x85, 0xcd, 0x3f, 0xb2, 0xb1, 0x59, 0xc6, 0xe3, 0xff, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x1d, 0x38, 0x0c, 0x71, 0xe2, 0x02, 0x00, 0x00,
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
	if this.RecoveryPeriod != that1.RecoveryPeriod {
		return false
	}
	if this.AdjustPricePeriod != that1.AdjustPricePeriod {
		return false
	}
	if !this.BasePrice.Equal(that1.BasePrice) {
		return false
	}
	if !this.BaseLoad.Equal(that1.BaseLoad) {
		return false
	}
	if this.MaxBlockBandwidth != that1.MaxBlockBandwidth {
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
	if m.MaxBlockBandwidth != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxBlockBandwidth))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.BaseLoad.Size()
		i -= size
		if _, err := m.BaseLoad.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.BasePrice.Size()
		i -= size
		if _, err := m.BasePrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AdjustPricePeriod != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.AdjustPricePeriod))
		i--
		dAtA[i] = 0x10
	}
	if m.RecoveryPeriod != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RecoveryPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NeuronBandwidth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NeuronBandwidth) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NeuronBandwidth) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxValue != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MaxValue))
		i--
		dAtA[i] = 0x20
	}
	if m.LastUpdatedBlock != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.LastUpdatedBlock))
		i--
		dAtA[i] = 0x18
	}
	if m.RemainedValue != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.RemainedValue))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Neuron) > 0 {
		i -= len(m.Neuron)
		copy(dAtA[i:], m.Neuron)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Neuron)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Price) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Price) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Price) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	if m.RecoveryPeriod != 0 {
		n += 1 + sovTypes(uint64(m.RecoveryPeriod))
	}
	if m.AdjustPricePeriod != 0 {
		n += 1 + sovTypes(uint64(m.AdjustPricePeriod))
	}
	l = m.BasePrice.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.BaseLoad.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.MaxBlockBandwidth != 0 {
		n += 1 + sovTypes(uint64(m.MaxBlockBandwidth))
	}
	return n
}

func (m *NeuronBandwidth) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Neuron)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.RemainedValue != 0 {
		n += 1 + sovTypes(uint64(m.RemainedValue))
	}
	if m.LastUpdatedBlock != 0 {
		n += 1 + sovTypes(uint64(m.LastUpdatedBlock))
	}
	if m.MaxValue != 0 {
		n += 1 + sovTypes(uint64(m.MaxValue))
	}
	return n
}

func (m *Price) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovTypes(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field RecoveryPeriod", wireType)
			}
			m.RecoveryPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecoveryPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustPricePeriod", wireType)
			}
			m.AdjustPricePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AdjustPricePeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasePrice", wireType)
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
			if err := m.BasePrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseLoad", wireType)
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
			if err := m.BaseLoad.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockBandwidth", wireType)
			}
			m.MaxBlockBandwidth = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBlockBandwidth |= uint64(b&0x7F) << shift
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
func (m *NeuronBandwidth) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: NeuronBandwidth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NeuronBandwidth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Neuron", wireType)
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
			m.Neuron = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainedValue", wireType)
			}
			m.RemainedValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RemainedValue |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastUpdatedBlock", wireType)
			}
			m.LastUpdatedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastUpdatedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValue", wireType)
			}
			m.MaxValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxValue |= uint64(b&0x7F) << shift
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
func (m *Price) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Price: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Price: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
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
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
