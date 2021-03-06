// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/cmd.proto

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
	AgentID      int32  `json:"AgentID,omitempty"`
	UserUin      string `json:"UserUin,omitempty"`
	Gender       int32  `json:"Gender,omitempty"`
	NikeName     string `json:"NikeName,omitempty"`
	HeadImageUrl string `json:"HeadImageUrl,omitempty"`
	MachineID    string `json:"MachineID,omitempty"`
	DeviceType   int32  `json:"DeviceType,omitempty"`
}

//手机号码登陆
type Login_C_Mobile struct {
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Password    string `json:"Password,omitempty"`
	MachineID   string `json:"MachineID,omitempty"`
	DeviceType  int32  `json:"DeviceType,omitempty"`
}

//游客登陆
type Login_C_Visitor struct {
	MachineID  string `json:"MachineID,omitempty"`
	DeviceType int32  `json:"DeviceType,omitempty"`
}

//-------------------------------------------------------服务端消息----------------------------------------------------
type GameInfo struct {
	GameID         int32   `json:"GameID,omitempty"`
	KindID         int32   `json:"KindID,omitempty"`
	ServerAddr     string  `json:"ServerAddr,omitempty"`
	WsAddr         string  `json:"WsAddr,omitempty"`
	GameName       string  `json:"GameName,omitempty"`
	SortID         int32   `json:"SortID,omitempty"`
	TableCount     int32   `json:"TableCount,omitempty"`
	ChairCount     int32   `json:"ChairCount,omitempty"`
	CellScore      float32 `json:"CellScore,omitempty"`
	RevenueRatio   float32 `json:"RevenueRatio,omitempty"`
	MinEnterScore  float32 `json:"MinEnterScore,omitempty"`
	DeductionsType int32   `json:"DeductionsType,omitempty"`
}

//登陆成功
type Login_S_Success struct {
	UserID       int32       `json:"UserID,omitempty"`
	NikeName     string      `json:"NikeName,omitempty"`
	UserGold     float32     `json:"UserGold,omitempty"`
	UserDiamonds float32     `json:"UserDiamonds,omitempty"`
	MemberOrder  int32       `json:"MemberOrder,omitempty"`
	PhoneNumber  string      `json:"PhoneNumber,omitempty"`
	BinderCardNo string      `json:"BinderCardNo,omitempty"`
	FaceID       int32       `json:"FaceID,omitempty"`
	RoleID       int32       `json:"RoleID,omitempty"`
	SuitID       int32       `json:"SuitID,omitempty"`
	PhotoFrameID int32       `json:"PhotoFrameID,omitempty"`
	Gender       int32       `json:"Gender,omitempty"`
	KindID       int32       `json:"KindID,omitempty"`
	GameID       int32       `json:"GameID,omitempty"`
	Token        string      `json:"Token,omitempty"`
	GameInfoList []*GameInfo `json:"GameInfoList,omitempty"`
}

//登陆失败
type Login_S_Fail struct {
	ErrorCode int32  `json:"ErrorCode,omitempty"`
	ErrorMsg  string `json:"ErrorMsg,omitempty"`
}

func init() {
	proto.RegisterType((*Login_C_Wechat)(nil), "msg.Login_C_Wechat")
	proto.RegisterType((*Login_C_Mobile)(nil), "msg.Login_C_Mobile")
	proto.RegisterType((*Login_C_Visitor)(nil), "msg.Login_C_Visitor")
	proto.RegisterType((*GameInfo)(nil), "msg.GameInfo")
	proto.RegisterType((*Login_S_Success)(nil), "msg.Login_S_Success")
	proto.RegisterType((*Login_S_Fail)(nil), "msg.Login_S_Fail")
}

func init() { proto.RegisterFile("msg/cmd.proto", fileDescriptor_a090769f8e2c12a2) }

var fileDescriptor_a090769f8e2c12a2 = []byte{
	// 646 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0x16, 0xa1, 0x50, 0x30, 0xd0, 0x4e, 0xd1, 0x34, 0x45, 0xd5, 0x34, 0x21, 0x34, 0x4d, 0x5c,
	0x75, 0xda, 0xf6, 0x04, 0x1d, 0xf4, 0x07, 0xad, 0xd0, 0x2a, 0xb4, 0xeb, 0x65, 0x65, 0x92, 0xb3,
	0x60, 0x35, 0xb1, 0x2b, 0x3b, 0x74, 0xda, 0x3b, 0xec, 0x09, 0xf7, 0x08, 0x7b, 0x80, 0x5d, 0x4f,
	0xe7, 0xc4, 0x21, 0x71, 0x91, 0x7a, 0xe7, 0xef, 0x3b, 0xf6, 0xf1, 0x39, 0xfe, 0xbe, 0x63, 0x36,
	0xc8, 0x4c, 0xf2, 0x31, 0xca, 0xe2, 0xe3, 0x47, 0xad, 0x72, 0xe5, 0x37, 0x33, 0x93, 0x8c, 0xfe,
	0x34, 0xd8, 0xc1, 0xa5, 0x4a, 0x84, 0xbc, 0x9f, 0xdc, 0xdf, 0x41, 0xb4, 0xe6, 0xb9, 0x1f, 0xb0,
	0xfd, 0x93, 0x04, 0x64, 0x3e, 0x9b, 0x06, 0x8d, 0x61, 0x63, 0xdc, 0x0a, 0x4b, 0x88, 0x91, 0x5b,
	0x03, 0xfa, 0x56, 0xc8, 0xc0, 0x1b, 0x36, 0xc6, 0xdd, 0xb0, 0x84, 0xfe, 0x1b, 0xd6, 0x3e, 0x07,
	0x19, 0x83, 0x0e, 0x9a, 0x74, 0xc4, 0x22, 0xff, 0x88, 0x75, 0x16, 0xe2, 0x01, 0x16, 0x3c, 0x83,
	0x60, 0x8f, 0x8e, 0x6c, 0xb1, 0x3f, 0x62, 0xfd, 0x0b, 0xe0, 0xf1, 0x2c, 0xe3, 0x09, 0xdc, 0xea,
	0x34, 0x68, 0x51, 0xdc, 0xe1, 0xfc, 0xb7, 0xac, 0x3b, 0xe7, 0xd1, 0x5a, 0x48, 0x98, 0x4d, 0x83,
	0x36, 0x6d, 0xa8, 0x08, 0xff, 0x1d, 0x63, 0x53, 0x78, 0x12, 0x11, 0xdc, 0xfc, 0x7a, 0x84, 0x60,
	0x9f, 0x6e, 0xae, 0x31, 0xa3, 0xdf, 0xb5, 0xe6, 0xe6, 0x6a, 0x25, 0x52, 0xf0, 0x87, 0xac, 0x77,
	0xbd, 0x56, 0x12, 0x16, 0x9b, 0x6c, 0x05, 0x9a, 0x1a, 0xec, 0x86, 0x75, 0x0a, 0x4b, 0xbe, 0xe6,
	0xc6, 0xfc, 0x54, 0x3a, 0xb6, 0x5d, 0x6e, 0xb1, 0x5b, 0x4e, 0xf3, 0xe5, 0x72, 0xf6, 0x76, 0xca,
	0xb9, 0x62, 0x87, 0x65, 0x35, 0xdf, 0x85, 0x11, 0xb9, 0xd2, 0x6e, 0xc2, 0xc6, 0xcb, 0x09, 0xbd,
	0x9d, 0x84, 0xff, 0x3c, 0xd6, 0x39, 0xe7, 0x19, 0xcc, 0xe4, 0x0f, 0x45, 0x12, 0xe0, 0xba, 0x54,
	0xcd, 0x22, 0xe4, 0xbf, 0x09, 0x19, 0xcf, 0xa6, 0x36, 0x81, 0x45, 0x98, 0x7c, 0x09, 0xfa, 0x09,
	0xf4, 0x49, 0x1c, 0x6b, 0xdb, 0x4c, 0x8d, 0xc1, 0x73, 0x77, 0x86, 0x62, 0x85, 0x70, 0x16, 0xe1,
	0xfb, 0x60, 0x66, 0x92, 0xb4, 0x90, 0x6c, 0x8b, 0xf1, 0xcc, 0x52, 0xe9, 0xdc, 0x6a, 0xd5, 0x0a,
	0x2d, 0xc2, 0xbb, 0x6e, 0xf8, 0x2a, 0x85, 0x89, 0xda, 0xc8, 0xbc, 0x14, 0xaa, 0x62, 0x30, 0x3e,
	0x59, 0x73, 0xa1, 0x8b, 0x78, 0xa7, 0x88, 0x57, 0x0c, 0x3e, 0xd3, 0x04, 0xd2, 0x74, 0x19, 0x29,
	0x0d, 0x41, 0x77, 0xd8, 0x18, 0x7b, 0x61, 0x45, 0xa0, 0x91, 0x42, 0x78, 0x02, 0xb9, 0x81, 0x90,
	0xe7, 0x42, 0x05, 0x8c, 0x36, 0x38, 0x9c, 0xff, 0x9e, 0x0d, 0xe6, 0x42, 0x9e, 0xca, 0x1c, 0x74,
	0x91, 0xa5, 0x47, 0x9b, 0x5c, 0xd2, 0xff, 0xc0, 0x0e, 0xa6, 0x10, 0x6f, 0xa2, 0x5c, 0x28, 0x69,
	0xe8, 0xd1, 0xfb, 0x54, 0xcb, 0x33, 0x76, 0xf4, 0xb7, 0x59, 0x4a, 0xb9, 0xbc, 0x5f, 0x6e, 0xa2,
	0x08, 0x8c, 0xc1, 0xde, 0x71, 0x1a, 0xaa, 0xf7, 0x2f, 0x90, 0x33, 0x02, 0xde, 0xb3, 0x11, 0x38,
	0x62, 0x1d, 0xdc, 0x75, 0xae, 0xd2, 0x98, 0x14, 0xf0, 0xc2, 0x2d, 0xc6, 0xae, 0x70, 0x3d, 0x15,
	0x3c, 0x53, 0x32, 0x36, 0xa4, 0x82, 0x17, 0x3a, 0x1c, 0xba, 0x79, 0x0e, 0xe8, 0xda, 0x2b, 0x8d,
	0xb3, 0xd7, 0xa2, 0x8b, 0xeb, 0xd4, 0x73, 0xbf, 0xb7, 0x77, 0xfd, 0x3e, 0x62, 0xfd, 0xaf, 0x02,
	0x87, 0x75, 0xc2, 0x75, 0xbc, 0x50, 0xa4, 0x4e, 0x37, 0x74, 0x38, 0xec, 0xed, 0x8c, 0x47, 0xe8,
	0xad, 0x42, 0x1b, 0x8b, 0x90, 0x0f, 0x55, 0x8a, 0x7c, 0xb7, 0xe0, 0x0b, 0x44, 0x3e, 0xd8, 0x08,
	0xf4, 0x01, 0xb3, 0x3e, 0x20, 0x84, 0x77, 0x5d, 0xaf, 0x55, 0xae, 0xce, 0x74, 0xe1, 0xd4, 0x1e,
	0x45, 0x1d, 0xae, 0xf6, 0x95, 0xf4, 0x9d, 0xaf, 0xa4, 0xf2, 0xf1, 0xc0, 0xf1, 0x71, 0xe5, 0xfb,
	0x03, 0xc7, 0xf7, 0xaf, 0x59, 0xeb, 0x46, 0x3d, 0x80, 0x0c, 0x0e, 0xa9, 0xa1, 0x02, 0xf8, 0x9f,
	0x58, 0xbf, 0x9c, 0x98, 0x4b, 0x61, 0xf2, 0xe0, 0xd5, 0xb0, 0x39, 0xee, 0x7d, 0x1e, 0x1c, 0x67,
	0x26, 0x39, 0x2e, 0x03, 0xa1, 0xb3, 0x65, 0x74, 0xc1, 0xfa, 0xa5, 0xd6, 0x67, 0x5c, 0xd0, 0x9f,
	0x74, 0xaa, 0xb5, 0xd2, 0x13, 0x15, 0x83, 0xd5, 0xba, 0x22, 0x50, 0x52, 0x02, 0x73, 0x93, 0x94,
	0x72, 0x97, 0x78, 0xd5, 0xa6, 0x8f, 0xf7, 0xcb, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0xf9,
	0xf6, 0x25, 0x89, 0x05, 0x00, 0x00,
}
