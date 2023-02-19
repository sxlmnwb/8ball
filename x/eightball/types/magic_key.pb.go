// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/magic_key.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type MagicKey struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EcPointX       string `protobuf:"bytes,2,opt,name=ecPointX,proto3" json:"ecPointX,omitempty"`
	EcPointY       string `protobuf:"bytes,3,opt,name=ecPointY,proto3" json:"ecPointY,omitempty"`
	BlockHeight    uint64 `protobuf:"varint,6,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BitcoinAlias   string `protobuf:"bytes,7,opt,name=bitcoinAlias,proto3" json:"bitcoinAlias,omitempty"`
	EthereumAlias  string `protobuf:"bytes,8,opt,name=ethereumAlias,proto3" json:"ethereumAlias,omitempty"`
	TronAlias      string `protobuf:"bytes,9,opt,name=tronAlias,proto3" json:"tronAlias,omitempty"`
	ZclassicAlias  string `protobuf:"bytes,10,opt,name=zclassicAlias,proto3" json:"zclassicAlias,omitempty"`
	EightballAlias string `protobuf:"bytes,11,opt,name=eightballAlias,proto3" json:"eightballAlias,omitempty"`
}

func (m *MagicKey) Reset()         { *m = MagicKey{} }
func (m *MagicKey) String() string { return proto.CompactTextString(m) }
func (*MagicKey) ProtoMessage()    {}
func (*MagicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_0884930af9c7a043, []int{0}
}
func (m *MagicKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MagicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MagicKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MagicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MagicKey.Merge(m, src)
}
func (m *MagicKey) XXX_Size() int {
	return m.Size()
}
func (m *MagicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MagicKey.DiscardUnknown(m)
}

var xxx_messageInfo_MagicKey proto.InternalMessageInfo

func (m *MagicKey) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MagicKey) GetEcPointX() string {
	if m != nil {
		return m.EcPointX
	}
	return ""
}

func (m *MagicKey) GetEcPointY() string {
	if m != nil {
		return m.EcPointY
	}
	return ""
}

func (m *MagicKey) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MagicKey) GetBitcoinAlias() string {
	if m != nil {
		return m.BitcoinAlias
	}
	return ""
}

func (m *MagicKey) GetEthereumAlias() string {
	if m != nil {
		return m.EthereumAlias
	}
	return ""
}

func (m *MagicKey) GetTronAlias() string {
	if m != nil {
		return m.TronAlias
	}
	return ""
}

func (m *MagicKey) GetZclassicAlias() string {
	if m != nil {
		return m.ZclassicAlias
	}
	return ""
}

func (m *MagicKey) GetEightballAlias() string {
	if m != nil {
		return m.EightballAlias
	}
	return ""
}

func init() {
	proto.RegisterType((*MagicKey)(nil), "eightball.eightball.MagicKey")
}

func init() {
	proto.RegisterFile("eightball/eightball/magic_key.proto", fileDescriptor_0884930af9c7a043)
}

var fileDescriptor_0884930af9c7a043 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x72, 0x13, 0xd3, 0x33, 0x93, 0xe3, 0xb3, 0x53,
	0x2b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x52, 0x7a, 0x70, 0x96, 0xd2, 0x2a,
	0x26, 0x2e, 0x0e, 0x5f, 0x90, 0x42, 0xef, 0xd4, 0x4a, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x29, 0x2e, 0x8e, 0xd4, 0xe4, 0x80,
	0xfc, 0xcc, 0xbc, 0x92, 0x08, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x38, 0x1f, 0x49, 0x2e,
	0x52, 0x82, 0x19, 0x45, 0x2e, 0x52, 0x48, 0x81, 0x8b, 0x3b, 0x29, 0x27, 0x3f, 0x39, 0xdb, 0x03,
	0x6c, 0x8f, 0x04, 0x1b, 0xd8, 0x40, 0x64, 0x21, 0x21, 0x25, 0x2e, 0x9e, 0xa4, 0xcc, 0x92, 0xe4,
	0xfc, 0xcc, 0x3c, 0xc7, 0x9c, 0xcc, 0xc4, 0x62, 0x09, 0x76, 0xb0, 0x09, 0x28, 0x62, 0x42, 0x2a,
	0x5c, 0xbc, 0xa9, 0x25, 0x19, 0xa9, 0x45, 0xa9, 0xa5, 0xb9, 0x10, 0x45, 0x1c, 0x60, 0x45, 0xa8,
	0x82, 0x42, 0x32, 0x5c, 0x9c, 0x25, 0x45, 0xf9, 0x50, 0x63, 0x38, 0xc1, 0x2a, 0x10, 0x02, 0x20,
	0x33, 0xaa, 0x92, 0x73, 0x12, 0x8b, 0x8b, 0x33, 0x93, 0x21, 0x2a, 0xb8, 0x20, 0x66, 0xa0, 0x08,
	0x0a, 0xa9, 0x71, 0xf1, 0xc1, 0x43, 0x04, 0xa2, 0x8c, 0x1b, 0xac, 0x0c, 0x4d, 0xd4, 0xc9, 0xf4,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xa4, 0x11, 0xc1, 0x5e, 0x81, 0x14,
	0x05, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xf0, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xe7, 0xd3, 0x5d, 0xbf, 0xa6, 0x01, 0x00, 0x00,
}

func (m *MagicKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MagicKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MagicKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EightballAlias) > 0 {
		i -= len(m.EightballAlias)
		copy(dAtA[i:], m.EightballAlias)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.EightballAlias)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ZclassicAlias) > 0 {
		i -= len(m.ZclassicAlias)
		copy(dAtA[i:], m.ZclassicAlias)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.ZclassicAlias)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.TronAlias) > 0 {
		i -= len(m.TronAlias)
		copy(dAtA[i:], m.TronAlias)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.TronAlias)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.EthereumAlias) > 0 {
		i -= len(m.EthereumAlias)
		copy(dAtA[i:], m.EthereumAlias)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.EthereumAlias)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BitcoinAlias) > 0 {
		i -= len(m.BitcoinAlias)
		copy(dAtA[i:], m.BitcoinAlias)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.BitcoinAlias)))
		i--
		dAtA[i] = 0x3a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintMagicKey(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x30
	}
	if len(m.EcPointY) > 0 {
		i -= len(m.EcPointY)
		copy(dAtA[i:], m.EcPointY)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.EcPointY)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EcPointX) > 0 {
		i -= len(m.EcPointX)
		copy(dAtA[i:], m.EcPointX)
		i = encodeVarintMagicKey(dAtA, i, uint64(len(m.EcPointX)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMagicKey(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMagicKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovMagicKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MagicKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMagicKey(uint64(m.Id))
	}
	l = len(m.EcPointX)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	l = len(m.EcPointY)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMagicKey(uint64(m.BlockHeight))
	}
	l = len(m.BitcoinAlias)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	l = len(m.EthereumAlias)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	l = len(m.TronAlias)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	l = len(m.ZclassicAlias)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	l = len(m.EightballAlias)
	if l > 0 {
		n += 1 + l + sovMagicKey(uint64(l))
	}
	return n
}

func sovMagicKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMagicKey(x uint64) (n int) {
	return sovMagicKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MagicKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMagicKey
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
			return fmt.Errorf("proto: MagicKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MagicKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EcPointX", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EcPointX = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EcPointY", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EcPointY = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcoinAlias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BitcoinAlias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthereumAlias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthereumAlias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TronAlias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TronAlias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZclassicAlias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZclassicAlias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EightballAlias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMagicKey
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
				return ErrInvalidLengthMagicKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMagicKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EightballAlias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMagicKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMagicKey
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
func skipMagicKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMagicKey
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
					return 0, ErrIntOverflowMagicKey
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
					return 0, ErrIntOverflowMagicKey
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
				return 0, ErrInvalidLengthMagicKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMagicKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMagicKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMagicKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMagicKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMagicKey = fmt.Errorf("proto: unexpected end of group")
)
