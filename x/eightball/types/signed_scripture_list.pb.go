// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/signed_scripture_list.proto

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

type SignedScriptureList struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ScriptureIndex string `protobuf:"bytes,2,opt,name=scriptureIndex,proto3" json:"scriptureIndex,omitempty"`
	Creator        string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SignedScriptureList) Reset()         { *m = SignedScriptureList{} }
func (m *SignedScriptureList) String() string { return proto.CompactTextString(m) }
func (*SignedScriptureList) ProtoMessage()    {}
func (*SignedScriptureList) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b5e7a9dcc5d8a41, []int{0}
}
func (m *SignedScriptureList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedScriptureList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedScriptureList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedScriptureList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedScriptureList.Merge(m, src)
}
func (m *SignedScriptureList) XXX_Size() int {
	return m.Size()
}
func (m *SignedScriptureList) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedScriptureList.DiscardUnknown(m)
}

var xxx_messageInfo_SignedScriptureList proto.InternalMessageInfo

func (m *SignedScriptureList) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignedScriptureList) GetScriptureIndex() string {
	if m != nil {
		return m.ScriptureIndex
	}
	return ""
}

func (m *SignedScriptureList) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SignedScriptureList)(nil), "eightball.eightball.SignedScriptureList")
}

func init() {
	proto.RegisterFile("eightball/eightball/signed_scripture_list.proto", fileDescriptor_5b5e7a9dcc5d8a41)
}

var fileDescriptor_5b5e7a9dcc5d8a41 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0x41, 0x62, 0x15, 0x67, 0xa6, 0xe7, 0xa5, 0xa6, 0xc4, 0x17, 0x27,
	0x17, 0x65, 0x16, 0x94, 0x94, 0x16, 0xa5, 0xc6, 0xe7, 0x64, 0x16, 0x97, 0xe8, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x09, 0xc3, 0x95, 0xe9, 0xc1, 0x59, 0x4a, 0xe9, 0x5c, 0xc2, 0xc1, 0x60, 0x3d,
	0xc1, 0x30, 0x2d, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x6a, 0x5c, 0x7c, 0x70, 0x33, 0x3d, 0xf3, 0x52,
	0x52, 0x2b, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xd0, 0x44, 0x85, 0x24, 0xb8, 0xd8, 0x93,
	0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x98, 0xc1, 0x0a, 0x60, 0x5c, 0x27, 0xd3, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x46, 0x38, 0xbf, 0x02, 0xc9, 0x2b, 0x25,
	0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xb7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x45, 0x35, 0x3c, 0xee, 0x00, 0x00, 0x00,
}

func (m *SignedScriptureList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedScriptureList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedScriptureList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignedScriptureList(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ScriptureIndex) > 0 {
		i -= len(m.ScriptureIndex)
		copy(dAtA[i:], m.ScriptureIndex)
		i = encodeVarintSignedScriptureList(dAtA, i, uint64(len(m.ScriptureIndex)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignedScriptureList(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignedScriptureList(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignedScriptureList(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignedScriptureList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignedScriptureList(uint64(m.Id))
	}
	l = len(m.ScriptureIndex)
	if l > 0 {
		n += 1 + l + sovSignedScriptureList(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignedScriptureList(uint64(l))
	}
	return n
}

func sovSignedScriptureList(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignedScriptureList(x uint64) (n int) {
	return sovSignedScriptureList(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignedScriptureList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignedScriptureList
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
			return fmt.Errorf("proto: SignedScriptureList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedScriptureList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedScriptureList
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
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptureIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedScriptureList
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
				return ErrInvalidLengthSignedScriptureList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedScriptureList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScriptureIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedScriptureList
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
				return ErrInvalidLengthSignedScriptureList
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedScriptureList
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSignedScriptureList(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignedScriptureList
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
func skipSignedScriptureList(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignedScriptureList
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
					return 0, ErrIntOverflowSignedScriptureList
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
					return 0, ErrIntOverflowSignedScriptureList
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
				return 0, ErrInvalidLengthSignedScriptureList
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignedScriptureList
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignedScriptureList
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignedScriptureList        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignedScriptureList          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignedScriptureList = fmt.Errorf("proto: unexpected end of group")
)