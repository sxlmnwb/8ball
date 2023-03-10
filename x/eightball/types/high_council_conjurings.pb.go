// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/high_council_conjurings.proto

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

type HighCouncilConjurings struct {
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MagicKeyId   uint64 `protobuf:"varint,2,opt,name=magicKeyId,proto3" json:"magicKeyId,omitempty"`
	BlockStarted uint64 `protobuf:"varint,3,opt,name=blockStarted,proto3" json:"blockStarted,omitempty"`
}

func (m *HighCouncilConjurings) Reset()         { *m = HighCouncilConjurings{} }
func (m *HighCouncilConjurings) String() string { return proto.CompactTextString(m) }
func (*HighCouncilConjurings) ProtoMessage()    {}
func (*HighCouncilConjurings) Descriptor() ([]byte, []int) {
	return fileDescriptor_55422d923f860e45, []int{0}
}
func (m *HighCouncilConjurings) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HighCouncilConjurings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HighCouncilConjurings.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HighCouncilConjurings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HighCouncilConjurings.Merge(m, src)
}
func (m *HighCouncilConjurings) XXX_Size() int {
	return m.Size()
}
func (m *HighCouncilConjurings) XXX_DiscardUnknown() {
	xxx_messageInfo_HighCouncilConjurings.DiscardUnknown(m)
}

var xxx_messageInfo_HighCouncilConjurings proto.InternalMessageInfo

func (m *HighCouncilConjurings) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HighCouncilConjurings) GetMagicKeyId() uint64 {
	if m != nil {
		return m.MagicKeyId
	}
	return 0
}

func (m *HighCouncilConjurings) GetBlockStarted() uint64 {
	if m != nil {
		return m.BlockStarted
	}
	return 0
}

func init() {
	proto.RegisterType((*HighCouncilConjurings)(nil), "eightball.eightball.HighCouncilConjurings")
}

func init() {
	proto.RegisterFile("eightball/eightball/high_council_conjurings.proto", fileDescriptor_55422d923f860e45)
}

var fileDescriptor_55422d923f860e45 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x32, 0x32, 0xd3, 0x33, 0xe2, 0x93, 0xf3, 0x4b,
	0xf3, 0x92, 0x33, 0x73, 0xe2, 0x93, 0xf3, 0xf3, 0xb2, 0x4a, 0x8b, 0x32, 0xf3, 0xd2, 0x8b, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x0a, 0xf5, 0xe0, 0x2c, 0xa5, 0x6c, 0x2e, 0x51,
	0x8f, 0xcc, 0xf4, 0x0c, 0x67, 0x88, 0x26, 0x67, 0xb8, 0x1e, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x39, 0x2e, 0xae, 0xdc, 0xc4,
	0xf4, 0xcc, 0x64, 0xef, 0xd4, 0x4a, 0xcf, 0x14, 0x09, 0x26, 0xb0, 0x38, 0x92, 0x88, 0x90, 0x12,
	0x17, 0x4f, 0x52, 0x4e, 0x7e, 0x72, 0x76, 0x70, 0x49, 0x62, 0x51, 0x49, 0x6a, 0x8a, 0x04, 0x33,
	0x58, 0x05, 0x8a, 0x98, 0x93, 0xe9, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x49, 0x23, 0x3c, 0x51, 0x81, 0xe4, 0xa1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xfb,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xee, 0x0a, 0x11, 0xf4, 0x00, 0x00, 0x00,
}

func (m *HighCouncilConjurings) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HighCouncilConjurings) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HighCouncilConjurings) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockStarted != 0 {
		i = encodeVarintHighCouncilConjurings(dAtA, i, uint64(m.BlockStarted))
		i--
		dAtA[i] = 0x18
	}
	if m.MagicKeyId != 0 {
		i = encodeVarintHighCouncilConjurings(dAtA, i, uint64(m.MagicKeyId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintHighCouncilConjurings(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintHighCouncilConjurings(dAtA []byte, offset int, v uint64) int {
	offset -= sovHighCouncilConjurings(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HighCouncilConjurings) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHighCouncilConjurings(uint64(m.Id))
	}
	if m.MagicKeyId != 0 {
		n += 1 + sovHighCouncilConjurings(uint64(m.MagicKeyId))
	}
	if m.BlockStarted != 0 {
		n += 1 + sovHighCouncilConjurings(uint64(m.BlockStarted))
	}
	return n
}

func sovHighCouncilConjurings(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHighCouncilConjurings(x uint64) (n int) {
	return sovHighCouncilConjurings(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HighCouncilConjurings) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHighCouncilConjurings
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
			return fmt.Errorf("proto: HighCouncilConjurings: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HighCouncilConjurings: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHighCouncilConjurings
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MagicKeyId", wireType)
			}
			m.MagicKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHighCouncilConjurings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MagicKeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockStarted", wireType)
			}
			m.BlockStarted = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHighCouncilConjurings
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockStarted |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHighCouncilConjurings(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHighCouncilConjurings
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
func skipHighCouncilConjurings(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHighCouncilConjurings
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
					return 0, ErrIntOverflowHighCouncilConjurings
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
					return 0, ErrIntOverflowHighCouncilConjurings
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
				return 0, ErrInvalidLengthHighCouncilConjurings
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHighCouncilConjurings
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHighCouncilConjurings
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHighCouncilConjurings        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHighCouncilConjurings          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHighCouncilConjurings = fmt.Errorf("proto: unexpected end of group")
)
