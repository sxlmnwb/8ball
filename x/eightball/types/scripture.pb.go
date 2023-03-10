// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/scripture.proto

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

type Scripture struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Alias    string `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias,omitempty"`
	Value    string `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Scripture) Reset()         { *m = Scripture{} }
func (m *Scripture) String() string { return proto.CompactTextString(m) }
func (*Scripture) ProtoMessage()    {}
func (*Scripture) Descriptor() ([]byte, []int) {
	return fileDescriptor_023bf207a65c9973, []int{0}
}
func (m *Scripture) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scripture) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scripture.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scripture) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scripture.Merge(m, src)
}
func (m *Scripture) XXX_Size() int {
	return m.Size()
}
func (m *Scripture) XXX_DiscardUnknown() {
	xxx_messageInfo_Scripture.DiscardUnknown(m)
}

var xxx_messageInfo_Scripture proto.InternalMessageInfo

func (m *Scripture) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Scripture) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Scripture) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Scripture) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Scripture) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Scripture) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Scripture)(nil), "eightball.eightball.Scripture")
}

func init() {
	proto.RegisterFile("eightball/eightball/scripture.proto", fileDescriptor_023bf207a65c9973)
}

var fileDescriptor_023bf207a65c9973 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x8a, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x4a, 0x8b,
	0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x52, 0x7a, 0x70, 0x96, 0xd2, 0x74,
	0x46, 0x2e, 0xce, 0x60, 0x98, 0x42, 0x21, 0x11, 0x2e, 0xd6, 0xcc, 0xbc, 0x94, 0xd4, 0x0a, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x88, 0x8b, 0x25, 0x23, 0xb1, 0x38, 0x43,
	0x82, 0x09, 0x2c, 0x08, 0x66, 0x0b, 0x49, 0x71, 0x71, 0xe4, 0xe4, 0x27, 0x27, 0x96, 0x64, 0xe6,
	0xe7, 0x49, 0x30, 0x83, 0xc5, 0xe1, 0x7c, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4,
	0xe2, 0x62, 0x09, 0x16, 0xb0, 0x14, 0x8c, 0x0b, 0x32, 0x3f, 0x31, 0x27, 0x33, 0xb1, 0x58, 0x82,
	0x15, 0x62, 0x3e, 0x98, 0x03, 0x12, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x83, 0x88, 0x82,
	0x39, 0x4e, 0xa6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x8d, 0xf0,
	0x63, 0x05, 0x92, 0x7f, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x9e, 0x35, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x44, 0xc0, 0x37, 0xd5, 0x13, 0x01, 0x00, 0x00,
}

func (m *Scripture) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scripture) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Scripture) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Alias) > 0 {
		i -= len(m.Alias)
		copy(dAtA[i:], m.Alias)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Alias)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintScripture(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScripture(dAtA []byte, offset int, v uint64) int {
	offset -= sovScripture(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Scripture) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	l = len(m.Alias)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovScripture(uint64(l))
	}
	return n
}

func sovScripture(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScripture(x uint64) (n int) {
	return sovScripture(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scripture) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScripture
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
			return fmt.Errorf("proto: Scripture: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scripture: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alias", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Alias = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScripture
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
				return ErrInvalidLengthScripture
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScripture
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScripture(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScripture
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
func skipScripture(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScripture
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
					return 0, ErrIntOverflowScripture
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
					return 0, ErrIntOverflowScripture
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
				return 0, ErrInvalidLengthScripture
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScripture
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScripture
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScripture        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScripture          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScripture = fmt.Errorf("proto: unexpected end of group")
)
