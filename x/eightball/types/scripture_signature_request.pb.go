// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/scripture_signature_request.proto

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

type ScriptureSignatureRequest struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ScriptureIndex string `protobuf:"bytes,2,opt,name=scriptureIndex,proto3" json:"scriptureIndex,omitempty"`
	BlockHeight    uint64 `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *ScriptureSignatureRequest) Reset()         { *m = ScriptureSignatureRequest{} }
func (m *ScriptureSignatureRequest) String() string { return proto.CompactTextString(m) }
func (*ScriptureSignatureRequest) ProtoMessage()    {}
func (*ScriptureSignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f48fc557442a3567, []int{0}
}
func (m *ScriptureSignatureRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScriptureSignatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScriptureSignatureRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScriptureSignatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScriptureSignatureRequest.Merge(m, src)
}
func (m *ScriptureSignatureRequest) XXX_Size() int {
	return m.Size()
}
func (m *ScriptureSignatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScriptureSignatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScriptureSignatureRequest proto.InternalMessageInfo

func (m *ScriptureSignatureRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ScriptureSignatureRequest) GetScriptureIndex() string {
	if m != nil {
		return m.ScriptureIndex
	}
	return ""
}

func (m *ScriptureSignatureRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*ScriptureSignatureRequest)(nil), "eightball.eightball.ScriptureSignatureRequest")
}

func init() {
	proto.RegisterFile("eightball/eightball/scripture_signature_request.proto", fileDescriptor_f48fc557442a3567)
}

var fileDescriptor_f48fc557442a3567 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x8a, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x4a, 0x8b,
	0x52, 0xe3, 0x8b, 0x33, 0xd3, 0xf3, 0x12, 0xc1, 0xac, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x61, 0xb8, 0x62, 0x3d, 0x38, 0x4b, 0xa9, 0x94, 0x4b,
	0x32, 0x18, 0xa6, 0x33, 0x18, 0xa6, 0x31, 0x08, 0xa2, 0x4f, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x88, 0x29, 0x33, 0x45, 0x48, 0x8d, 0x8b, 0x0f, 0x6e, 0x8d,
	0x67, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9a, 0xa8, 0x90, 0x02,
	0x17, 0x77, 0x52, 0x4e, 0x7e, 0x72, 0xb6, 0x07, 0xd8, 0x1e, 0x09, 0x66, 0xb0, 0x01, 0xc8, 0x42,
	0x4e, 0xa6, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84,
	0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x8d, 0xf0, 0x52,
	0x05, 0x92, 0xf7, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x3e, 0x31, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x87, 0x0c, 0x6e, 0x02, 0x01, 0x00, 0x00,
}

func (m *ScriptureSignatureRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScriptureSignatureRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScriptureSignatureRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintScriptureSignatureRequest(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ScriptureIndex) > 0 {
		i -= len(m.ScriptureIndex)
		copy(dAtA[i:], m.ScriptureIndex)
		i = encodeVarintScriptureSignatureRequest(dAtA, i, uint64(len(m.ScriptureIndex)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintScriptureSignatureRequest(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintScriptureSignatureRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovScriptureSignatureRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScriptureSignatureRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovScriptureSignatureRequest(uint64(m.Id))
	}
	l = len(m.ScriptureIndex)
	if l > 0 {
		n += 1 + l + sovScriptureSignatureRequest(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovScriptureSignatureRequest(uint64(m.BlockHeight))
	}
	return n
}

func sovScriptureSignatureRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScriptureSignatureRequest(x uint64) (n int) {
	return sovScriptureSignatureRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScriptureSignatureRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScriptureSignatureRequest
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
			return fmt.Errorf("proto: ScriptureSignatureRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScriptureSignatureRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScriptureSignatureRequest
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
					return ErrIntOverflowScriptureSignatureRequest
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
				return ErrInvalidLengthScriptureSignatureRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScriptureSignatureRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScriptureIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScriptureSignatureRequest
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
		default:
			iNdEx = preIndex
			skippy, err := skipScriptureSignatureRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScriptureSignatureRequest
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
func skipScriptureSignatureRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScriptureSignatureRequest
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
					return 0, ErrIntOverflowScriptureSignatureRequest
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
					return 0, ErrIntOverflowScriptureSignatureRequest
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
				return 0, ErrInvalidLengthScriptureSignatureRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScriptureSignatureRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScriptureSignatureRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScriptureSignatureRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScriptureSignatureRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScriptureSignatureRequest = fmt.Errorf("proto: unexpected end of group")
)
