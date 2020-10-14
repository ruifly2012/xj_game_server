// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

//微信登陆
type Login_C_Wechat struct {
	AgentID              int32    `protobuf:"varint,1,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	UserUin              string   `protobuf:"bytes,2,opt,name=UserUin,proto3" json:"UserUin,omitempty"`
	Gender               bool     `protobuf:"varint,3,opt,name=Gender,proto3" json:"Gender,omitempty"`
	NikeName             string   `protobuf:"bytes,4,opt,name=NikeName,proto3" json:"NikeName,omitempty"`
	HeadImageUrl         string   `protobuf:"bytes,5,opt,name=HeadImageUrl,proto3" json:"HeadImageUrl,omitempty"`
	MachineID            string   `protobuf:"bytes,6,opt,name=MachineID,proto3" json:"MachineID,omitempty"`
	DeviceType           int32    `protobuf:"varint,7,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login_C_Wechat) Reset()         { *m = Login_C_Wechat{} }
func (m *Login_C_Wechat) String() string { return proto.CompactTextString(m) }
func (*Login_C_Wechat) ProtoMessage()    {}
func (*Login_C_Wechat) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{0}
}

func (m *Login_C_Wechat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login_C_Wechat.Unmarshal(m, b)
}
func (m *Login_C_Wechat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login_C_Wechat.Marshal(b, m, deterministic)
}
func (m *Login_C_Wechat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login_C_Wechat.Merge(m, src)
}
func (m *Login_C_Wechat) XXX_Size() int {
	return xxx_messageInfo_Login_C_Wechat.Size(m)
}
func (m *Login_C_Wechat) XXX_DiscardUnknown() {
	xxx_messageInfo_Login_C_Wechat.DiscardUnknown(m)
}

var xxx_messageInfo_Login_C_Wechat proto.InternalMessageInfo

func (m *Login_C_Wechat) GetAgentID() int32 {
	if m != nil {
		return m.AgentID
	}
	return 0
}

func (m *Login_C_Wechat) GetUserUin() string {
	if m != nil {
		return m.UserUin
	}
	return ""
}

func (m *Login_C_Wechat) GetGender() bool {
	if m != nil {
		return m.Gender
	}
	return false
}

func (m *Login_C_Wechat) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *Login_C_Wechat) GetHeadImageUrl() string {
	if m != nil {
		return m.HeadImageUrl
	}
	return ""
}

func (m *Login_C_Wechat) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *Login_C_Wechat) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

//手机号码登陆
type Login_C_Mobile struct {
	PhoneNumber          string   `protobuf:"bytes,1,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	MachineID            string   `protobuf:"bytes,3,opt,name=MachineID,proto3" json:"MachineID,omitempty"`
	DeviceType           int32    `protobuf:"varint,4,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login_C_Mobile) Reset()         { *m = Login_C_Mobile{} }
func (m *Login_C_Mobile) String() string { return proto.CompactTextString(m) }
func (*Login_C_Mobile) ProtoMessage()    {}
func (*Login_C_Mobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{1}
}

func (m *Login_C_Mobile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login_C_Mobile.Unmarshal(m, b)
}
func (m *Login_C_Mobile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login_C_Mobile.Marshal(b, m, deterministic)
}
func (m *Login_C_Mobile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login_C_Mobile.Merge(m, src)
}
func (m *Login_C_Mobile) XXX_Size() int {
	return xxx_messageInfo_Login_C_Mobile.Size(m)
}
func (m *Login_C_Mobile) XXX_DiscardUnknown() {
	xxx_messageInfo_Login_C_Mobile.DiscardUnknown(m)
}

var xxx_messageInfo_Login_C_Mobile proto.InternalMessageInfo

func (m *Login_C_Mobile) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Login_C_Mobile) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Login_C_Mobile) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *Login_C_Mobile) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

//游客登陆
type Login_C_Visitor struct {
	MachineID            string   `protobuf:"bytes,1,opt,name=MachineID,proto3" json:"MachineID,omitempty"`
	DeviceType           int32    `protobuf:"varint,2,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login_C_Visitor) Reset()         { *m = Login_C_Visitor{} }
func (m *Login_C_Visitor) String() string { return proto.CompactTextString(m) }
func (*Login_C_Visitor) ProtoMessage()    {}
func (*Login_C_Visitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{2}
}

func (m *Login_C_Visitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login_C_Visitor.Unmarshal(m, b)
}
func (m *Login_C_Visitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login_C_Visitor.Marshal(b, m, deterministic)
}
func (m *Login_C_Visitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login_C_Visitor.Merge(m, src)
}
func (m *Login_C_Visitor) XXX_Size() int {
	return xxx_messageInfo_Login_C_Visitor.Size(m)
}
func (m *Login_C_Visitor) XXX_DiscardUnknown() {
	xxx_messageInfo_Login_C_Visitor.DiscardUnknown(m)
}

var xxx_messageInfo_Login_C_Visitor proto.InternalMessageInfo

func (m *Login_C_Visitor) GetMachineID() string {
	if m != nil {
		return m.MachineID
	}
	return ""
}

func (m *Login_C_Visitor) GetDeviceType() int32 {
	if m != nil {
		return m.DeviceType
	}
	return 0
}

//-------------------------------------------------------服务端消息----------------------------------------------------
type GameInfo struct {
	GameID               int32    `protobuf:"varint,1,opt,name=GameID,proto3" json:"GameID,omitempty"`
	KindID               int32    `protobuf:"varint,2,opt,name=KindID,proto3" json:"KindID,omitempty"`
	ServerAddr           string   `protobuf:"bytes,3,opt,name=ServerAddr,proto3" json:"ServerAddr,omitempty"`
	WsAddr               string   `protobuf:"bytes,4,opt,name=WsAddr,proto3" json:"WsAddr,omitempty"`
	GameName             string   `protobuf:"bytes,5,opt,name=GameName,proto3" json:"GameName,omitempty"`
	SortID               int32    `protobuf:"varint,6,opt,name=SortID,proto3" json:"SortID,omitempty"`
	TableCount           int32    `protobuf:"varint,7,opt,name=TableCount,proto3" json:"TableCount,omitempty"`
	ChairCount           int32    `protobuf:"varint,8,opt,name=ChairCount,proto3" json:"ChairCount,omitempty"`
	CellScore            float32  `protobuf:"fixed32,9,opt,name=CellScore,proto3" json:"CellScore,omitempty"`
	RevenueRatio         float32  `protobuf:"fixed32,10,opt,name=RevenueRatio,proto3" json:"RevenueRatio,omitempty"`
	MinEnterScore        float32  `protobuf:"fixed32,11,opt,name=MinEnterScore,proto3" json:"MinEnterScore,omitempty"`
	DeductionsType       int32    `protobuf:"varint,12,opt,name=DeductionsType,proto3" json:"DeductionsType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameInfo) Reset()         { *m = GameInfo{} }
func (m *GameInfo) String() string { return proto.CompactTextString(m) }
func (*GameInfo) ProtoMessage()    {}
func (*GameInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{3}
}

func (m *GameInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameInfo.Unmarshal(m, b)
}
func (m *GameInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameInfo.Marshal(b, m, deterministic)
}
func (m *GameInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameInfo.Merge(m, src)
}
func (m *GameInfo) XXX_Size() int {
	return xxx_messageInfo_GameInfo.Size(m)
}
func (m *GameInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GameInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GameInfo proto.InternalMessageInfo

func (m *GameInfo) GetGameID() int32 {
	if m != nil {
		return m.GameID
	}
	return 0
}

func (m *GameInfo) GetKindID() int32 {
	if m != nil {
		return m.KindID
	}
	return 0
}

func (m *GameInfo) GetServerAddr() string {
	if m != nil {
		return m.ServerAddr
	}
	return ""
}

func (m *GameInfo) GetWsAddr() string {
	if m != nil {
		return m.WsAddr
	}
	return ""
}

func (m *GameInfo) GetGameName() string {
	if m != nil {
		return m.GameName
	}
	return ""
}

func (m *GameInfo) GetSortID() int32 {
	if m != nil {
		return m.SortID
	}
	return 0
}

func (m *GameInfo) GetTableCount() int32 {
	if m != nil {
		return m.TableCount
	}
	return 0
}

func (m *GameInfo) GetChairCount() int32 {
	if m != nil {
		return m.ChairCount
	}
	return 0
}

func (m *GameInfo) GetCellScore() float32 {
	if m != nil {
		return m.CellScore
	}
	return 0
}

func (m *GameInfo) GetRevenueRatio() float32 {
	if m != nil {
		return m.RevenueRatio
	}
	return 0
}

func (m *GameInfo) GetMinEnterScore() float32 {
	if m != nil {
		return m.MinEnterScore
	}
	return 0
}

func (m *GameInfo) GetDeductionsType() int32 {
	if m != nil {
		return m.DeductionsType
	}
	return 0
}

//登陆成功
type Login_S_Success struct {
	UserID               int32       `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NikeName             string      `protobuf:"bytes,2,opt,name=NikeName,proto3" json:"NikeName,omitempty"`
	UserGold             float32     `protobuf:"fixed32,3,opt,name=UserGold,proto3" json:"UserGold,omitempty"`
	UserDiamonds         float32     `protobuf:"fixed32,4,opt,name=UserDiamonds,proto3" json:"UserDiamonds,omitempty"`
	MemberOrder          int32       `protobuf:"varint,5,opt,name=MemberOrder,proto3" json:"MemberOrder,omitempty"`
	PhoneNumber          string      `protobuf:"bytes,6,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	BinderCardNo         string      `protobuf:"bytes,7,opt,name=BinderCardNo,proto3" json:"BinderCardNo,omitempty"`
	FaceID               int32       `protobuf:"varint,8,opt,name=FaceID,proto3" json:"FaceID,omitempty"`
	RoleID               int32       `protobuf:"varint,9,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	SuitID               int32       `protobuf:"varint,10,opt,name=SuitID,proto3" json:"SuitID,omitempty"`
	PhotoFrameID         int32       `protobuf:"varint,11,opt,name=PhotoFrameID,proto3" json:"PhotoFrameID,omitempty"`
	Gender               int32       `protobuf:"varint,12,opt,name=Gender,proto3" json:"Gender,omitempty"`
	KindID               int32       `protobuf:"varint,13,opt,name=KindID,proto3" json:"KindID,omitempty"`
	GameID               int32       `protobuf:"varint,14,opt,name=GameID,proto3" json:"GameID,omitempty"`
	Token                string      `protobuf:"bytes,15,opt,name=Token,proto3" json:"Token,omitempty"`
	GameInfoList         []*GameInfo `protobuf:"bytes,16,rep,name=GameInfoList,proto3" json:"GameInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Login_S_Success) Reset()         { *m = Login_S_Success{} }
func (m *Login_S_Success) String() string { return proto.CompactTextString(m) }
func (*Login_S_Success) ProtoMessage()    {}
func (*Login_S_Success) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{4}
}

func (m *Login_S_Success) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login_S_Success.Unmarshal(m, b)
}
func (m *Login_S_Success) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login_S_Success.Marshal(b, m, deterministic)
}
func (m *Login_S_Success) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login_S_Success.Merge(m, src)
}
func (m *Login_S_Success) XXX_Size() int {
	return xxx_messageInfo_Login_S_Success.Size(m)
}
func (m *Login_S_Success) XXX_DiscardUnknown() {
	xxx_messageInfo_Login_S_Success.DiscardUnknown(m)
}

var xxx_messageInfo_Login_S_Success proto.InternalMessageInfo

func (m *Login_S_Success) GetUserID() int32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Login_S_Success) GetNikeName() string {
	if m != nil {
		return m.NikeName
	}
	return ""
}

func (m *Login_S_Success) GetUserGold() float32 {
	if m != nil {
		return m.UserGold
	}
	return 0
}

func (m *Login_S_Success) GetUserDiamonds() float32 {
	if m != nil {
		return m.UserDiamonds
	}
	return 0
}

func (m *Login_S_Success) GetMemberOrder() int32 {
	if m != nil {
		return m.MemberOrder
	}
	return 0
}

func (m *Login_S_Success) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Login_S_Success) GetBinderCardNo() string {
	if m != nil {
		return m.BinderCardNo
	}
	return ""
}

func (m *Login_S_Success) GetFaceID() int32 {
	if m != nil {
		return m.FaceID
	}
	return 0
}

func (m *Login_S_Success) GetRoleID() int32 {
	if m != nil {
		return m.RoleID
	}
	return 0
}

func (m *Login_S_Success) GetSuitID() int32 {
	if m != nil {
		return m.SuitID
	}
	return 0
}

func (m *Login_S_Success) GetPhotoFrameID() int32 {
	if m != nil {
		return m.PhotoFrameID
	}
	return 0
}

func (m *Login_S_Success) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *Login_S_Success) GetKindID() int32 {
	if m != nil {
		return m.KindID
	}
	return 0
}

func (m *Login_S_Success) GetGameID() int32 {
	if m != nil {
		return m.GameID
	}
	return 0
}

func (m *Login_S_Success) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Login_S_Success) GetGameInfoList() []*GameInfo {
	if m != nil {
		return m.GameInfoList
	}
	return nil
}

//登陆失败
type Login_S_Fail struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=ErrorMsg,proto3" json:"ErrorMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login_S_Fail) Reset()         { *m = Login_S_Fail{} }
func (m *Login_S_Fail) String() string { return proto.CompactTextString(m) }
func (*Login_S_Fail) ProtoMessage()    {}
func (*Login_S_Fail) Descriptor() ([]byte, []int) {
	return fileDescriptor_7520252fb01eaf30, []int{5}
}

func (m *Login_S_Fail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login_S_Fail.Unmarshal(m, b)
}
func (m *Login_S_Fail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login_S_Fail.Marshal(b, m, deterministic)
}
func (m *Login_S_Fail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login_S_Fail.Merge(m, src)
}
func (m *Login_S_Fail) XXX_Size() int {
	return xxx_messageInfo_Login_S_Fail.Size(m)
}
func (m *Login_S_Fail) XXX_DiscardUnknown() {
	xxx_messageInfo_Login_S_Fail.DiscardUnknown(m)
}

var xxx_messageInfo_Login_S_Fail proto.InternalMessageInfo

func (m *Login_S_Fail) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Login_S_Fail) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Login_C_Wechat)(nil), "msg.Login_C_Wechat")
	proto.RegisterType((*Login_C_Mobile)(nil), "msg.Login_C_Mobile")
	proto.RegisterType((*Login_C_Visitor)(nil), "msg.Login_C_Visitor")
	proto.RegisterType((*GameInfo)(nil), "msg.GameInfo")
	proto.RegisterType((*Login_S_Success)(nil), "msg.Login_S_Success")
	proto.RegisterType((*Login_S_Fail)(nil), "msg.Login_S_Fail")
}

func init() { proto.RegisterFile("cmd.proto", fileDescriptor_7520252fb01eaf30) }

var fileDescriptor_7520252fb01eaf30 = []byte{
	// 647 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdf, 0x6a, 0xdb, 0x3e,
	0x14, 0xc6, 0x4e, 0x93, 0xc6, 0x4a, 0xda, 0xfe, 0x30, 0x3f, 0x86, 0x29, 0x63, 0x84, 0x30, 0x46,
	0xae, 0x0a, 0xdb, 0x9e, 0xa0, 0x4b, 0xfa, 0x27, 0xac, 0x49, 0x8b, 0xd2, 0xae, 0x97, 0x45, 0xb1,
	0xcf, 0x12, 0x51, 0x5b, 0x2a, 0x92, 0xd3, 0xb1, 0x77, 0xd8, 0x13, 0xee, 0x11, 0xf6, 0x00, 0xbb,
	0x1e, 0x47, 0x92, 0x63, 0xab, 0x81, 0xde, 0xf9, 0xfb, 0xce, 0xf1, 0xd1, 0x39, 0xfa, 0xbe, 0x23,
	0x12, 0xa5, 0x45, 0x76, 0xf2, 0xa4, 0x64, 0x29, 0xe3, 0x56, 0xa1, 0x57, 0xc3, 0xdf, 0x01, 0x39,
	0xbc, 0x92, 0x2b, 0x2e, 0x1e, 0xc6, 0x0f, 0xf7, 0x90, 0xae, 0x59, 0x19, 0x27, 0x64, 0xff, 0x74,
	0x05, 0xa2, 0x9c, 0x4e, 0x92, 0x60, 0x10, 0x8c, 0xda, 0xb4, 0x82, 0x18, 0xb9, 0xd3, 0xa0, 0xee,
	0xb8, 0x48, 0xc2, 0x41, 0x30, 0x8a, 0x68, 0x05, 0xe3, 0x37, 0xa4, 0x73, 0x01, 0x22, 0x03, 0x95,
	0xb4, 0x06, 0xc1, 0xa8, 0x4b, 0x1d, 0x8a, 0x8f, 0x49, 0x77, 0xce, 0x1f, 0x61, 0xce, 0x0a, 0x48,
	0xf6, 0xcc, 0x2f, 0x5b, 0x1c, 0x0f, 0x49, 0xff, 0x12, 0x58, 0x36, 0x2d, 0xd8, 0x0a, 0xee, 0x54,
	0x9e, 0xb4, 0x4d, 0xdc, 0xe3, 0xe2, 0xb7, 0x24, 0x9a, 0xb1, 0x74, 0xcd, 0x05, 0x4c, 0x27, 0x49,
	0xc7, 0x24, 0xd4, 0x44, 0xfc, 0x8e, 0x90, 0x09, 0x3c, 0xf3, 0x14, 0x6e, 0x7f, 0x3e, 0x41, 0xb2,
	0x6f, 0x9a, 0x6d, 0x30, 0xc3, 0x5f, 0x8d, 0xe1, 0x66, 0x72, 0xc9, 0x73, 0x88, 0x07, 0xa4, 0x77,
	0xb3, 0x96, 0x02, 0xe6, 0x9b, 0x62, 0x09, 0xca, 0x0c, 0x18, 0xd1, 0x26, 0x85, 0x2d, 0xdf, 0x30,
	0xad, 0x7f, 0x48, 0x95, 0xb9, 0x29, 0xb7, 0xd8, 0x6f, 0xa7, 0xf5, 0x7a, 0x3b, 0x7b, 0x3b, 0xed,
	0x5c, 0x93, 0xa3, 0xaa, 0x9b, 0x6f, 0x5c, 0xf3, 0x52, 0x2a, 0xbf, 0x60, 0xf0, 0x7a, 0xc1, 0x70,
	0xa7, 0xe0, 0xdf, 0x90, 0x74, 0x2f, 0x58, 0x01, 0x53, 0xf1, 0x5d, 0x1a, 0x09, 0xf0, 0xbb, 0x52,
	0xcd, 0x21, 0xe4, 0xbf, 0x72, 0x91, 0x4d, 0x27, 0xae, 0x80, 0x43, 0x58, 0x7c, 0x01, 0xea, 0x19,
	0xd4, 0x69, 0x96, 0x29, 0x37, 0x4c, 0x83, 0xc1, 0xff, 0xee, 0xb5, 0x89, 0x59, 0xe1, 0x1c, 0xc2,
	0xfb, 0xc1, 0xca, 0x46, 0x52, 0x2b, 0xd9, 0x16, 0xe3, 0x3f, 0x0b, 0xa9, 0x4a, 0xa7, 0x55, 0x9b,
	0x3a, 0x84, 0x67, 0xdd, 0xb2, 0x65, 0x0e, 0x63, 0xb9, 0x11, 0x65, 0x25, 0x54, 0xcd, 0x60, 0x7c,
	0xbc, 0x66, 0x5c, 0xd9, 0x78, 0xd7, 0xc6, 0x6b, 0x06, 0xaf, 0x69, 0x0c, 0x79, 0xbe, 0x48, 0xa5,
	0x82, 0x24, 0x1a, 0x04, 0xa3, 0x90, 0xd6, 0x04, 0x1a, 0x89, 0xc2, 0x33, 0x88, 0x0d, 0x50, 0x56,
	0x72, 0x99, 0x10, 0x93, 0xe0, 0x71, 0xf1, 0x7b, 0x72, 0x30, 0xe3, 0xe2, 0x4c, 0x94, 0xa0, 0x6c,
	0x95, 0x9e, 0x49, 0xf2, 0xc9, 0xf8, 0x03, 0x39, 0x9c, 0x40, 0xb6, 0x49, 0x4b, 0x2e, 0x85, 0x36,
	0x97, 0xde, 0x37, 0xbd, 0xbc, 0x60, 0x87, 0x7f, 0x5a, 0x95, 0x94, 0x8b, 0x87, 0xc5, 0x26, 0x4d,
	0x41, 0x6b, 0x9c, 0x1d, 0xb7, 0xa1, 0xbe, 0x7f, 0x8b, 0xbc, 0x15, 0x08, 0x5f, 0xac, 0xc0, 0x31,
	0xe9, 0x62, 0xd6, 0x85, 0xcc, 0x33, 0xa3, 0x40, 0x48, 0xb7, 0x18, 0xa7, 0xc2, 0xef, 0x09, 0x67,
	0x85, 0x14, 0x99, 0x36, 0x2a, 0x84, 0xd4, 0xe3, 0xd0, 0xcd, 0x33, 0x40, 0xd7, 0x5e, 0x2b, 0xdc,
	0xbd, 0xb6, 0x39, 0xb8, 0x49, 0xbd, 0xf4, 0x7b, 0x67, 0xd7, 0xef, 0x43, 0xd2, 0xff, 0xc2, 0x71,
	0x59, 0xc7, 0x4c, 0x65, 0x73, 0x69, 0xd4, 0x89, 0xa8, 0xc7, 0xe1, 0x6c, 0xe7, 0x2c, 0x45, 0x6f,
	0x59, 0x6d, 0x1c, 0x42, 0x9e, 0xca, 0x1c, 0xf9, 0xc8, 0xf2, 0x16, 0x19, 0x1f, 0x6c, 0x38, 0xfa,
	0x80, 0x38, 0x1f, 0x18, 0x84, 0x67, 0xdd, 0xac, 0x65, 0x29, 0xcf, 0x95, 0x75, 0x6a, 0xcf, 0x44,
	0x3d, 0xae, 0xf1, 0x94, 0xf4, 0x9d, 0x8f, 0xed, 0x53, 0x52, 0xfb, 0xf8, 0xc0, 0xf3, 0x71, 0xed,
	0xfb, 0x43, 0xcf, 0xf7, 0xff, 0x93, 0xf6, 0xad, 0x7c, 0x04, 0x91, 0x1c, 0x99, 0x81, 0x2c, 0x88,
	0x3f, 0x92, 0x7e, 0xb5, 0x31, 0x57, 0x5c, 0x97, 0xc9, 0x7f, 0x83, 0xd6, 0xa8, 0xf7, 0xe9, 0xe0,
	0xa4, 0xd0, 0xab, 0x93, 0x2a, 0x40, 0xbd, 0x94, 0xe1, 0x25, 0xe9, 0x57, 0x5a, 0x9f, 0x33, 0x6e,
	0xde, 0xa4, 0x33, 0xa5, 0xa4, 0x1a, 0xcb, 0x0c, 0x9c, 0xd6, 0x35, 0x81, 0x92, 0x1a, 0x30, 0xd3,
	0xab, 0x4a, 0xee, 0x0a, 0x2f, 0x3b, 0xe6, 0xe1, 0xfd, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0xb7, 0x8f, 0x82, 0x85, 0x05, 0x00, 0x00,
}