// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qomapp/vesting/v1/vesting.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_x_auth_vesting_types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClawbackVestingAccount implements the VestingAccount interface. It provides
// an account that can hold contributions subject to "lockup" (like a
// PeriodicVestingAccount), or vesting which is subject to clawback
// of unvested tokens, or a combination (tokens vest, but are still locked).
type ClawbackVestingAccount struct {
	// base_vesting_account implements the VestingAccount interface. It contains
	// all the necessary fields needed for any vesting account implementation
	*types.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	// funder_address specifies the account which can perform clawback
	FunderAddress string `protobuf:"bytes,2,opt,name=funder_address,json=funderAddress,proto3" json:"funder_address,omitempty"`
	// start_time defines the time at which the vesting period begins
	StartTime time.Time `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// lockup_periods defines the unlocking schedule relative to the start_time
	LockupPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,4,rep,name=lockup_periods,json=lockupPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"lockup_periods"`
	// vesting_periods defines the vesting schedule relative to the start_time
	VestingPeriods github_com_cosmos_cosmos_sdk_x_auth_vesting_types.Periods `protobuf:"bytes,5,rep,name=vesting_periods,json=vestingPeriods,proto3,castrepeated=github.com/cosmos/cosmos-sdk/x/auth/vesting/types.Periods" json:"vesting_periods"`
}

func (m *ClawbackVestingAccount) Reset()      { *m = ClawbackVestingAccount{} }
func (*ClawbackVestingAccount) ProtoMessage() {}
func (*ClawbackVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae64364fd8921f39, []int{0}
}
func (m *ClawbackVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClawbackVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClawbackVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClawbackVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClawbackVestingAccount.Merge(m, src)
}
func (m *ClawbackVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ClawbackVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ClawbackVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ClawbackVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClawbackVestingAccount)(nil), "qomapp.vesting.v1.ClawbackVestingAccount")
}

func init() { proto.RegisterFile("qomapp/vesting/v1/vesting.proto", fileDescriptor_ae64364fd8921f39) }

var fileDescriptor_ae64364fd8921f39 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x4d, 0xaf, 0xd2, 0x40,
	0x14, 0xed, 0xf8, 0xd0, 0xbc, 0x37, 0x2f, 0x0f, 0x63, 0xf3, 0x62, 0x1a, 0x16, 0x1d, 0x62, 0x34,
	0x21, 0x46, 0x66, 0x02, 0xae, 0x74, 0x47, 0x49, 0x5c, 0x9b, 0xc6, 0xb8, 0x70, 0xd3, 0xcc, 0xb4,
	0x43, 0x69, 0xa0, 0x9d, 0xd2, 0x99, 0x22, 0xfe, 0x03, 0xc2, 0x8a, 0xa5, 0x4b, 0xd6, 0xfe, 0x12,
	0x96, 0x2c, 0x5d, 0x81, 0x81, 0x3f, 0x62, 0xda, 0x69, 0x0d, 0x68, 0xdc, 0xbe, 0x55, 0xef, 0x3d,
	0xf7, 0xe3, 0x9c, 0xd3, 0x3b, 0x10, 0xcd, 0x44, 0x4c, 0xd3, 0x94, 0xcc, 0xb9, 0x54, 0x51, 0x12,
	0x92, 0x79, 0xaf, 0x0e, 0x71, 0x9a, 0x09, 0x25, 0xcc, 0x67, 0xba, 0x01, 0xd7, 0xe8, 0xbc, 0xd7,
	0x7a, 0xe9, 0x0b, 0x19, 0x0b, 0x79, 0x36, 0xc3, 0xb8, 0xa2, 0x7f, 0x0d, 0xb6, 0xee, 0x43, 0x11,
	0x8a, 0x32, 0x24, 0x45, 0x54, 0xa1, 0x28, 0x14, 0x22, 0x9c, 0x72, 0x52, 0x66, 0x2c, 0x1f, 0x11,
	0x15, 0xc5, 0x5c, 0x2a, 0x1a, 0xa7, 0xba, 0xe1, 0xc5, 0xaa, 0x01, 0x9f, 0x0f, 0xa7, 0xf4, 0x2b,
	0xa3, 0xfe, 0xe4, 0xb3, 0x5e, 0x38, 0xf0, 0x7d, 0x91, 0x27, 0xca, 0x64, 0xf0, 0x9e, 0x51, 0xc9,
	0xbd, 0x8a, 0xc7, 0xa3, 0x1a, 0xb7, 0x40, 0x1b, 0x74, 0x6e, 0xfb, 0xaf, 0xb1, 0x96, 0x75, 0xa6,
	0xb4, 0x94, 0x85, 0x1d, 0x2a, 0xf9, 0xe5, 0x26, 0xa7, 0xb1, 0xdb, 0x23, 0xe0, 0x9a, 0xec, 0x9f,
	0x8a, 0xf9, 0x0a, 0x36, 0x47, 0x79, 0x12, 0xf0, 0xcc, 0xa3, 0x41, 0x90, 0x71, 0x29, 0xad, 0x47,
	0x6d, 0xd0, 0xb9, 0x71, 0xef, 0x34, 0x3a, 0xd0, 0xa0, 0x39, 0x84, 0x50, 0x2a, 0x9a, 0x29, 0xaf,
	0x90, 0x6f, 0x5d, 0x95, 0x02, 0x5a, 0x58, 0x7b, 0xc3, 0xb5, 0x37, 0xfc, 0xa9, 0xf6, 0xe6, 0x5c,
	0x6f, 0xf7, 0xc8, 0x58, 0x1f, 0x10, 0x70, 0x6f, 0xca, 0xb9, 0xa2, 0x62, 0x2e, 0x01, 0x6c, 0x4e,
	0x85, 0x3f, 0xc9, 0x53, 0x2f, 0xe5, 0x59, 0x24, 0x02, 0x69, 0x35, 0xda, 0x57, 0x9d, 0xdb, 0xbe,
	0xfd, 0x3f, 0x2b, 0x1f, 0xcb, 0x36, 0x67, 0x50, 0x6c, 0xfb, 0x71, 0x40, 0xef, 0xc2, 0x48, 0x8d,
	0x73, 0x86, 0x7d, 0x11, 0x93, 0xea, 0x26, 0xfa, 0xd3, 0x95, 0xc1, 0x84, 0x2c, 0x08, 0xcd, 0xd5,
	0xf8, 0xcf, 0x95, 0xd4, 0xb7, 0x94, 0xcb, 0x6a, 0x83, 0x74, 0xef, 0x34, 0x71, 0x95, 0x9a, 0x2b,
	0x00, 0x9f, 0xd6, 0xbf, 0xb5, 0xd6, 0xf2, 0xf8, 0xa1, 0xb4, 0x34, 0x2b, 0xb8, 0xca, 0xdf, 0x5f,
	0x2f, 0x37, 0xc8, 0xf8, 0xbe, 0x41, 0x86, 0xf3, 0x61, 0x7b, 0xb4, 0xc1, 0xee, 0x68, 0x83, 0x5f,
	0x47, 0x1b, 0xac, 0x4f, 0xb6, 0xb1, 0x3b, 0xd9, 0xc6, 0xcf, 0x93, 0x6d, 0x7c, 0x79, 0x73, 0x46,
	0x37, 0x13, 0x71, 0x57, 0x24, 0x9c, 0xd4, 0x4f, 0xb9, 0x47, 0x16, 0x97, 0x4c, 0xec, 0x49, 0x79,
	0x92, 0xb7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x25, 0xc4, 0x44, 0xee, 0x02, 0x00, 0x00,
}

func (m *ClawbackVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClawbackVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClawbackVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VestingPeriods) > 0 {
		for iNdEx := len(m.VestingPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LockupPeriods) > 0 {
		for iNdEx := len(m.LockupPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LockupPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintVesting(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	if len(m.FunderAddress) > 0 {
		i -= len(m.FunderAddress)
		copy(dAtA[i:], m.FunderAddress)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.FunderAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClawbackVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	l = len(m.FunderAddress)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovVesting(uint64(l))
	if len(m.LockupPeriods) > 0 {
		for _, e := range m.LockupPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	if len(m.VestingPeriods) > 0 {
		for _, e := range m.VestingPeriods {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClawbackVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ClawbackVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClawbackVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FunderAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FunderAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockupPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockupPeriods = append(m.LockupPeriods, types.Period{})
			if err := m.LockupPeriods[len(m.LockupPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingPeriods", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestingPeriods = append(m.VestingPeriods, types.Period{})
			if err := m.VestingPeriods[len(m.VestingPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
