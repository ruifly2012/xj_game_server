// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChangeReq struct {
	UserID               int32    `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Score                float32  `protobuf:"fixed32,2,opt,name=Score,proto3" json:"Score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeReq) Reset()         { *m = ChangeReq{} }
func (m *ChangeReq) String() string { return proto.CompactTextString(m) }
func (*ChangeReq) ProtoMessage()    {}
func (*ChangeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *ChangeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeReq.Unmarshal(m, b)
}
func (m *ChangeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeReq.Marshal(b, m, deterministic)
}
func (m *ChangeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeReq.Merge(m, src)
}
func (m *ChangeReq) XXX_Size() int {
	return xxx_messageInfo_ChangeReq.Size(m)
}
func (m *ChangeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeReq proto.InternalMessageInfo

func (m *ChangeReq) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *ChangeReq) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

type ChangeReply struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=ErrorMsg,proto3" json:"ErrorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeReply) Reset()         { *m = ChangeReply{} }
func (m *ChangeReply) String() string { return proto.CompactTextString(m) }
func (*ChangeReply) ProtoMessage()    {}
func (*ChangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *ChangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeReply.Unmarshal(m, b)
}
func (m *ChangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeReply.Marshal(b, m, deterministic)
}
func (m *ChangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeReply.Merge(m, src)
}
func (m *ChangeReply) XXX_Size() int {
	return xxx_messageInfo_ChangeReply.Size(m)
}
func (m *ChangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeReply proto.InternalMessageInfo

func (m *ChangeReply) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ChangeReply) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

// 游戏配置文件
type ModifyStockReq struct {
	GameID               int32    `protobuf:"varint,1,opt,name=GameID,proto3" json:"GameID,omitempty"`
	KindID               int32    `protobuf:"varint,2,opt,name=KindID,proto3" json:"KindID,omitempty"`
	KindName             string   `protobuf:"bytes,3,opt,name=KindName,proto3" json:"KindName,omitempty"`
	GameName             string   `protobuf:"bytes,4,opt,name=GameName,proto3" json:"GameName,omitempty"`
	SortID               int32    `protobuf:"varint,5,opt,name=SortID,proto3" json:"SortID,omitempty"`
	TableCount           int32    `protobuf:"varint,6,opt,name=TableCount,proto3" json:"TableCount,omitempty"`
	ChairCount           int32    `protobuf:"varint,7,opt,name=ChairCount,proto3" json:"ChairCount,omitempty"`
	CellScore            float32  `protobuf:"fixed32,8,opt,name=CellScore,proto3" json:"CellScore,omitempty"`
	RevenueRatio         float32  `protobuf:"fixed32,9,opt,name=RevenueRatio,proto3" json:"RevenueRatio,omitempty"`
	MinEnterScore        float32  `protobuf:"fixed32,10,opt,name=MinEnterScore,proto3" json:"MinEnterScore,omitempty"`
	DeductionsType       int32    `protobuf:"varint,11,opt,name=DeductionsType,proto3" json:"DeductionsType,omitempty"`
	StoresDecay          float32  `protobuf:"fixed32,12,opt,name=StoresDecay,proto3" json:"StoresDecay,omitempty"`
	StartStores          float32  `protobuf:"fixed32,13,opt,name=StartStores,proto3" json:"StartStores,omitempty"`
	StartWinRate         float32  `protobuf:"fixed32,14,opt,name=StartWinRate,proto3" json:"StartWinRate,omitempty"`
	Threshold1           float32  `protobuf:"fixed32,15,opt,name=Threshold1,proto3" json:"Threshold1,omitempty"`
	WinRate1             float32  `protobuf:"fixed32,16,opt,name=WinRate1,proto3" json:"WinRate1,omitempty"`
	Threshold2           float32  `protobuf:"fixed32,17,opt,name=Threshold2,proto3" json:"Threshold2,omitempty"`
	WinRate2             float32  `protobuf:"fixed32,18,opt,name=WinRate2,proto3" json:"WinRate2,omitempty"`
	Threshold3           float32  `protobuf:"fixed32,19,opt,name=Threshold3,proto3" json:"Threshold3,omitempty"`
	WinRate3             float32  `protobuf:"fixed32,20,opt,name=WinRate3,proto3" json:"WinRate3,omitempty"`
	Threshold4           float32  `protobuf:"fixed32,21,opt,name=Threshold4,proto3" json:"Threshold4,omitempty"`
	WinRate4             float32  `protobuf:"fixed32,22,opt,name=WinRate4,proto3" json:"WinRate4,omitempty"`
	Threshold5           float32  `protobuf:"fixed32,23,opt,name=Threshold5,proto3" json:"Threshold5,omitempty"`
	WinRate5             float32  `protobuf:"fixed32,24,opt,name=WinRate5,proto3" json:"WinRate5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStockReq) Reset()         { *m = ModifyStockReq{} }
func (m *ModifyStockReq) String() string { return proto.CompactTextString(m) }
func (*ModifyStockReq) ProtoMessage()    {}
func (*ModifyStockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *ModifyStockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStockReq.Unmarshal(m, b)
}
func (m *ModifyStockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStockReq.Marshal(b, m, deterministic)
}
func (m *ModifyStockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStockReq.Merge(m, src)
}
func (m *ModifyStockReq) XXX_Size() int {
	return xxx_messageInfo_ModifyStockReq.Size(m)
}
func (m *ModifyStockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStockReq.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStockReq proto.InternalMessageInfo

func (m *ModifyStockReq) GetGameID() int32 {
	if m != nil {
		return m.GameID
	}
	return 0
}

func (m *ModifyStockReq) GetKindID() int32 {
	if m != nil {
		return m.KindID
	}
	return 0
}

func (m *ModifyStockReq) GetKindName() string {
	if m != nil {
		return m.KindName
	}
	return ""
}

func (m *ModifyStockReq) GetGameName() string {
	if m != nil {
		return m.GameName
	}
	return ""
}

func (m *ModifyStockReq) GetSortID() int32 {
	if m != nil {
		return m.SortID
	}
	return 0
}

func (m *ModifyStockReq) GetTableCount() int32 {
	if m != nil {
		return m.TableCount
	}
	return 0
}

func (m *ModifyStockReq) GetChairCount() int32 {
	if m != nil {
		return m.ChairCount
	}
	return 0
}

func (m *ModifyStockReq) GetCellScore() float32 {
	if m != nil {
		return m.CellScore
	}
	return 0
}

func (m *ModifyStockReq) GetRevenueRatio() float32 {
	if m != nil {
		return m.RevenueRatio
	}
	return 0
}

func (m *ModifyStockReq) GetMinEnterScore() float32 {
	if m != nil {
		return m.MinEnterScore
	}
	return 0
}

func (m *ModifyStockReq) GetDeductionsType() int32 {
	if m != nil {
		return m.DeductionsType
	}
	return 0
}

func (m *ModifyStockReq) GetStoresDecay() float32 {
	if m != nil {
		return m.StoresDecay
	}
	return 0
}

func (m *ModifyStockReq) GetStartStores() float32 {
	if m != nil {
		return m.StartStores
	}
	return 0
}

func (m *ModifyStockReq) GetStartWinRate() float32 {
	if m != nil {
		return m.StartWinRate
	}
	return 0
}

func (m *ModifyStockReq) GetThreshold1() float32 {
	if m != nil {
		return m.Threshold1
	}
	return 0
}

func (m *ModifyStockReq) GetWinRate1() float32 {
	if m != nil {
		return m.WinRate1
	}
	return 0
}

func (m *ModifyStockReq) GetThreshold2() float32 {
	if m != nil {
		return m.Threshold2
	}
	return 0
}

func (m *ModifyStockReq) GetWinRate2() float32 {
	if m != nil {
		return m.WinRate2
	}
	return 0
}

func (m *ModifyStockReq) GetThreshold3() float32 {
	if m != nil {
		return m.Threshold3
	}
	return 0
}

func (m *ModifyStockReq) GetWinRate3() float32 {
	if m != nil {
		return m.WinRate3
	}
	return 0
}

func (m *ModifyStockReq) GetThreshold4() float32 {
	if m != nil {
		return m.Threshold4
	}
	return 0
}

func (m *ModifyStockReq) GetWinRate4() float32 {
	if m != nil {
		return m.WinRate4
	}
	return 0
}

func (m *ModifyStockReq) GetThreshold5() float32 {
	if m != nil {
		return m.Threshold5
	}
	return 0
}

func (m *ModifyStockReq) GetWinRate5() float32 {
	if m != nil {
		return m.WinRate5
	}
	return 0
}

type ModifyStockReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyStockReply) Reset()         { *m = ModifyStockReply{} }
func (m *ModifyStockReply) String() string { return proto.CompactTextString(m) }
func (*ModifyStockReply) ProtoMessage()    {}
func (*ModifyStockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}

func (m *ModifyStockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyStockReply.Unmarshal(m, b)
}
func (m *ModifyStockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyStockReply.Marshal(b, m, deterministic)
}
func (m *ModifyStockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyStockReply.Merge(m, src)
}
func (m *ModifyStockReply) XXX_Size() int {
	return xxx_messageInfo_ModifyStockReply.Size(m)
}
func (m *ModifyStockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyStockReply.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyStockReply proto.InternalMessageInfo

type QueryStockReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStockReq) Reset()         { *m = QueryStockReq{} }
func (m *QueryStockReq) String() string { return proto.CompactTextString(m) }
func (*QueryStockReq) ProtoMessage()    {}
func (*QueryStockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4}
}

func (m *QueryStockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStockReq.Unmarshal(m, b)
}
func (m *QueryStockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStockReq.Marshal(b, m, deterministic)
}
func (m *QueryStockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStockReq.Merge(m, src)
}
func (m *QueryStockReq) XXX_Size() int {
	return xxx_messageInfo_QueryStockReq.Size(m)
}
func (m *QueryStockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStockReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStockReq proto.InternalMessageInfo

type QueryStockReply struct {
	NowStores            float32  `protobuf:"fixed32,1,opt,name=NowStores,proto3" json:"NowStores,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryStockReply) Reset()         { *m = QueryStockReply{} }
func (m *QueryStockReply) String() string { return proto.CompactTextString(m) }
func (*QueryStockReply) ProtoMessage()    {}
func (*QueryStockReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5}
}

func (m *QueryStockReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStockReply.Unmarshal(m, b)
}
func (m *QueryStockReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStockReply.Marshal(b, m, deterministic)
}
func (m *QueryStockReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStockReply.Merge(m, src)
}
func (m *QueryStockReply) XXX_Size() int {
	return xxx_messageInfo_QueryStockReply.Size(m)
}
func (m *QueryStockReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStockReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStockReply proto.InternalMessageInfo

func (m *QueryStockReply) GetNowStores() float32 {
	if m != nil {
		return m.NowStores
	}
	return 0
}

func init() {
	proto.RegisterType((*ChangeReq)(nil), "grpc.ChangeReq")
	proto.RegisterType((*ChangeReply)(nil), "grpc.ChangeReply")
	proto.RegisterType((*ModifyStockReq)(nil), "grpc.ModifyStockReq")
	proto.RegisterType((*ModifyStockReply)(nil), "grpc.ModifyStockReply")
	proto.RegisterType((*QueryStockReq)(nil), "grpc.QueryStockReq")
	proto.RegisterType((*QueryStockReply)(nil), "grpc.QueryStockReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xda, 0x4c,
	0x10, 0x86, 0x3f, 0xf2, 0x05, 0x1a, 0x86, 0x00, 0xc9, 0x86, 0xd0, 0x11, 0xaa, 0x2a, 0x64, 0x55,
	0x55, 0x4e, 0xa9, 0x62, 0xe0, 0xd0, 0x4a, 0x3d, 0xe1, 0x08, 0xa1, 0x8a, 0x48, 0x35, 0xa9, 0x7a,
	0x76, 0xec, 0x29, 0x58, 0x35, 0x5e, 0x77, 0x31, 0xad, 0xf8, 0xd7, 0x3d, 0xf4, 0x07, 0x54, 0xbb,
	0x6b, 0xcc, 0xae, 0xdb, 0x43, 0x6f, 0x7e, 0xdf, 0x99, 0x67, 0x34, 0x1e, 0xcf, 0x18, 0x60, 0x25,
	0xb2, 0xf0, 0x36, 0x13, 0x3c, 0xe7, 0xec, 0x54, 0x3e, 0x3b, 0x6f, 0xa1, 0x39, 0x5d, 0x07, 0xe9,
	0x8a, 0x7c, 0xfa, 0xc6, 0xfa, 0xd0, 0xf8, 0xb4, 0x25, 0x31, 0xf7, 0xb0, 0x36, 0xac, 0xdd, 0xd4,
	0xfd, 0x42, 0xb1, 0x1e, 0xd4, 0x97, 0x21, 0x17, 0x84, 0x27, 0xc3, 0xda, 0xcd, 0x89, 0xaf, 0x85,
	0x33, 0x83, 0xd6, 0x01, 0xcd, 0x92, 0x3d, 0x7b, 0x01, 0xcd, 0x7b, 0x21, 0xb8, 0x98, 0xf2, 0x88,
	0x0a, 0xfe, 0x68, 0xb0, 0x01, 0x9c, 0x29, 0xb1, 0xd8, 0xae, 0x54, 0x95, 0xa6, 0x5f, 0x6a, 0xe7,
	0x57, 0x1d, 0x3a, 0x0b, 0x1e, 0xc5, 0x5f, 0xf6, 0xcb, 0x9c, 0x87, 0x5f, 0x8b, 0x4e, 0x66, 0xc1,
	0x86, 0x8e, 0x9d, 0x68, 0x25, 0xfd, 0x0f, 0x71, 0x1a, 0xcd, 0x3d, 0x55, 0xa4, 0xee, 0x17, 0x4a,
	0x96, 0x97, 0x4f, 0x0f, 0xc1, 0x86, 0xf0, 0x7f, 0x5d, 0xfe, 0xa0, 0x65, 0x4c, 0xd2, 0x2a, 0x76,
	0xaa, 0x63, 0x07, 0x2d, 0xeb, 0x2d, 0xb9, 0xc8, 0xe7, 0x1e, 0xd6, 0x75, 0x3d, 0xad, 0xd8, 0x4b,
	0x80, 0xc7, 0xe0, 0x29, 0xa1, 0x29, 0xdf, 0xa5, 0x39, 0x36, 0x54, 0xcc, 0x70, 0x64, 0x7c, 0xba,
	0x0e, 0x62, 0xa1, 0xe3, 0xcf, 0x74, 0xfc, 0xe8, 0xc8, 0x61, 0x4c, 0x29, 0x49, 0xf4, 0xd4, 0xce,
	0xd4, 0xd4, 0x8e, 0x06, 0x73, 0xe0, 0xdc, 0xa7, 0xef, 0x94, 0xee, 0xc8, 0x0f, 0xf2, 0x98, 0x63,
	0x53, 0x25, 0x58, 0x1e, 0x7b, 0x05, 0xed, 0x45, 0x9c, 0xde, 0xa7, 0x39, 0x09, 0x5d, 0x05, 0x54,
	0x92, 0x6d, 0xb2, 0xd7, 0xd0, 0xf1, 0x28, 0xda, 0x85, 0x79, 0xcc, 0xd3, 0xed, 0xe3, 0x3e, 0x23,
	0x6c, 0xa9, 0x5e, 0x2a, 0x2e, 0x1b, 0x42, 0x6b, 0x99, 0x73, 0x41, 0x5b, 0x8f, 0xc2, 0x60, 0x8f,
	0xe7, 0xaa, 0x96, 0x69, 0xe9, 0x8c, 0x40, 0xe4, 0xda, 0xc3, 0xf6, 0x21, 0xa3, 0xb4, 0x64, 0xd7,
	0x4a, 0x7e, 0x8e, 0x53, 0x3f, 0xc8, 0x09, 0x3b, 0xba, 0x6b, 0xd3, 0x53, 0x73, 0x5b, 0x0b, 0xda,
	0xae, 0x79, 0x12, 0xdd, 0x61, 0x57, 0x65, 0x18, 0x8e, 0xfc, 0x16, 0x45, 0xea, 0x1d, 0x5e, 0xa8,
	0x68, 0xa9, 0x2d, 0xd6, 0xc5, 0xcb, 0x0a, 0xeb, 0x1a, 0xac, 0x8b, 0xcc, 0x62, 0x5d, 0x8b, 0x1d,
	0xe1, 0x55, 0x85, 0x1d, 0x19, 0xec, 0x08, 0x7b, 0x16, 0x3b, 0xb2, 0xd8, 0x31, 0x5e, 0x57, 0xd8,
	0xb1, 0xc1, 0x8e, 0xb1, 0x6f, 0xb1, 0x63, 0x8b, 0x9d, 0xe0, 0xf3, 0x0a, 0x3b, 0x31, 0xd8, 0x09,
	0xa2, 0xc5, 0x4e, 0x1c, 0x06, 0x17, 0xd6, 0xd6, 0x67, 0xc9, 0xde, 0xe9, 0x42, 0xfb, 0xe3, 0x8e,
	0x44, 0x79, 0x08, 0xce, 0x1b, 0xe8, 0x9a, 0x46, 0x71, 0x68, 0x0f, 0xfc, 0x47, 0xf1, 0x9d, 0x6a,
	0x7a, 0xb7, 0x4a, 0xc3, 0xfd, 0x59, 0xd3, 0xeb, 0x3e, 0x13, 0x59, 0xc8, 0x5c, 0xb5, 0xa6, 0xe9,
	0x8a, 0x66, 0x3c, 0x89, 0x58, 0xf7, 0x56, 0x9d, 0x7f, 0x79, 0xef, 0x83, 0x4b, 0xdb, 0x90, 0x0d,
	0xfc, 0xc7, 0x26, 0xd0, 0xd6, 0x86, 0x17, 0x07, 0x1b, 0x9e, 0xfe, 0x2b, 0xf6, 0x1e, 0x5a, 0xc6,
	0xdb, 0xb0, 0x9e, 0xce, 0xb1, 0xcf, 0x7a, 0xd0, 0xff, 0x8b, 0xab, 0xf1, 0x77, 0x00, 0xc7, 0xf7,
	0x64, 0x57, 0x3a, 0xcf, 0x1a, 0xc5, 0xe0, 0xfa, 0x4f, 0x53, 0xb1, 0x4f, 0x0d, 0xf5, 0x43, 0x1b,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x48, 0x7d, 0x96, 0xde, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameGrpcClient is the client API for GameGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameGrpcClient interface {
	//金币变动
	ChangeGold(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeReply, error)
	//余额变动
	ChangeDiamond(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeReply, error)
	//修改库存值
	ModifyStock(ctx context.Context, in *ModifyStockReq, opts ...grpc.CallOption) (*ModifyStockReply, error)
	//查询库存值
	QueryStock(ctx context.Context, in *QueryStockReq, opts ...grpc.CallOption) (*QueryStockReply, error)
}

type gameGrpcClient struct {
	cc *grpc.ClientConn
}

func NewGameGrpcClient(cc *grpc.ClientConn) GameGrpcClient {
	return &gameGrpcClient{cc}
}

func (c *gameGrpcClient) ChangeGold(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeReply, error) {
	out := new(ChangeReply)
	err := c.cc.Invoke(ctx, "/grpc.GameGrpc/ChangeGold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameGrpcClient) ChangeDiamond(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeReply, error) {
	out := new(ChangeReply)
	err := c.cc.Invoke(ctx, "/grpc.GameGrpc/ChangeDiamond", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameGrpcClient) ModifyStock(ctx context.Context, in *ModifyStockReq, opts ...grpc.CallOption) (*ModifyStockReply, error) {
	out := new(ModifyStockReply)
	err := c.cc.Invoke(ctx, "/grpc.GameGrpc/ModifyStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameGrpcClient) QueryStock(ctx context.Context, in *QueryStockReq, opts ...grpc.CallOption) (*QueryStockReply, error) {
	out := new(QueryStockReply)
	err := c.cc.Invoke(ctx, "/grpc.GameGrpc/QueryStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameGrpcServer is the server API for GameGrpc service.
type GameGrpcServer interface {
	//金币变动
	ChangeGold(context.Context, *ChangeReq) (*ChangeReply, error)
	//余额变动
	ChangeDiamond(context.Context, *ChangeReq) (*ChangeReply, error)
	//修改库存值
	ModifyStock(context.Context, *ModifyStockReq) (*ModifyStockReply, error)
	//查询库存值
	QueryStock(context.Context, *QueryStockReq) (*QueryStockReply, error)
}

// UnimplementedGameGrpcServer can be embedded to have forward compatible implementations.
type UnimplementedGameGrpcServer struct {
}

func (*UnimplementedGameGrpcServer) ChangeGold(ctx context.Context, req *ChangeReq) (*ChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeGold not implemented")
}
func (*UnimplementedGameGrpcServer) ChangeDiamond(ctx context.Context, req *ChangeReq) (*ChangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDiamond not implemented")
}
func (*UnimplementedGameGrpcServer) ModifyStock(ctx context.Context, req *ModifyStockReq) (*ModifyStockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyStock not implemented")
}
func (*UnimplementedGameGrpcServer) QueryStock(ctx context.Context, req *QueryStockReq) (*QueryStockReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStock not implemented")
}

func RegisterGameGrpcServer(s *grpc.Server, srv GameGrpcServer) {
	s.RegisterService(&_GameGrpc_serviceDesc, srv)
}

func _GameGrpc_ChangeGold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameGrpcServer).ChangeGold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GameGrpc/ChangeGold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameGrpcServer).ChangeGold(ctx, req.(*ChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameGrpc_ChangeDiamond_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameGrpcServer).ChangeDiamond(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GameGrpc/ChangeDiamond",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameGrpcServer).ChangeDiamond(ctx, req.(*ChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameGrpc_ModifyStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameGrpcServer).ModifyStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GameGrpc/ModifyStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameGrpcServer).ModifyStock(ctx, req.(*ModifyStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameGrpc_QueryStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameGrpcServer).QueryStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.GameGrpc/QueryStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameGrpcServer).QueryStock(ctx, req.(*QueryStockReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.GameGrpc",
	HandlerType: (*GameGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeGold",
			Handler:    _GameGrpc_ChangeGold_Handler,
		},
		{
			MethodName: "ChangeDiamond",
			Handler:    _GameGrpc_ChangeDiamond_Handler,
		},
		{
			MethodName: "ModifyStock",
			Handler:    _GameGrpc_ModifyStock_Handler,
		},
		{
			MethodName: "QueryStock",
			Handler:    _GameGrpc_QueryStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
