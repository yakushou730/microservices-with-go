// Code generated by protoc-gen-go. DO NOT EDIT.
// source: section-3-4/src/proto/wta.proto

package WTA

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
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

type Player struct {
	Id                   uint32               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string               `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	IsRightHanded        bool                 `protobuf:"varint,4,opt,name=isRightHanded,proto3" json:"isRightHanded,omitempty"`
	BirthDate            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	CountryCode          string               `protobuf:"bytes,6,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_4057945fb2ade8ac, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Player) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Player) GetIsRightHanded() bool {
	if m != nil {
		return m.IsRightHanded
	}
	return false
}

func (m *Player) GetBirthDate() *timestamp.Timestamp {
	if m != nil {
		return m.BirthDate
	}
	return nil
}

func (m *Player) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

type Ranking struct {
	PlayerId             uint32               `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	RankingDate          *timestamp.Timestamp `protobuf:"bytes,2,opt,name=rankingDate,proto3" json:"rankingDate,omitempty"`
	Ranking              uint32               `protobuf:"varint,3,opt,name=ranking,proto3" json:"ranking,omitempty"`
	RankingPoints        float32              `protobuf:"fixed32,4,opt,name=rankingPoints,proto3" json:"rankingPoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Ranking) Reset()         { *m = Ranking{} }
func (m *Ranking) String() string { return proto.CompactTextString(m) }
func (*Ranking) ProtoMessage()    {}
func (*Ranking) Descriptor() ([]byte, []int) {
	return fileDescriptor_4057945fb2ade8ac, []int{1}
}

func (m *Ranking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ranking.Unmarshal(m, b)
}
func (m *Ranking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ranking.Marshal(b, m, deterministic)
}
func (m *Ranking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ranking.Merge(m, src)
}
func (m *Ranking) XXX_Size() int {
	return xxx_messageInfo_Ranking.Size(m)
}
func (m *Ranking) XXX_DiscardUnknown() {
	xxx_messageInfo_Ranking.DiscardUnknown(m)
}

var xxx_messageInfo_Ranking proto.InternalMessageInfo

func (m *Ranking) GetPlayerId() uint32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Ranking) GetRankingDate() *timestamp.Timestamp {
	if m != nil {
		return m.RankingDate
	}
	return nil
}

func (m *Ranking) GetRanking() uint32 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *Ranking) GetRankingPoints() float32 {
	if m != nil {
		return m.RankingPoints
	}
	return 0
}

type PlayerWithRanking struct {
	Player               *Player  `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Ranking              *Ranking `protobuf:"bytes,2,opt,name=ranking,proto3" json:"ranking,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerWithRanking) Reset()         { *m = PlayerWithRanking{} }
func (m *PlayerWithRanking) String() string { return proto.CompactTextString(m) }
func (*PlayerWithRanking) ProtoMessage()    {}
func (*PlayerWithRanking) Descriptor() ([]byte, []int) {
	return fileDescriptor_4057945fb2ade8ac, []int{2}
}

func (m *PlayerWithRanking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerWithRanking.Unmarshal(m, b)
}
func (m *PlayerWithRanking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerWithRanking.Marshal(b, m, deterministic)
}
func (m *PlayerWithRanking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerWithRanking.Merge(m, src)
}
func (m *PlayerWithRanking) XXX_Size() int {
	return xxx_messageInfo_PlayerWithRanking.Size(m)
}
func (m *PlayerWithRanking) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerWithRanking.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerWithRanking proto.InternalMessageInfo

func (m *PlayerWithRanking) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *PlayerWithRanking) GetRanking() *Ranking {
	if m != nil {
		return m.Ranking
	}
	return nil
}

type PlayerIdRequest struct {
	PlayerId             uint32   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerIdRequest) Reset()         { *m = PlayerIdRequest{} }
func (m *PlayerIdRequest) String() string { return proto.CompactTextString(m) }
func (*PlayerIdRequest) ProtoMessage()    {}
func (*PlayerIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4057945fb2ade8ac, []int{3}
}

func (m *PlayerIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerIdRequest.Unmarshal(m, b)
}
func (m *PlayerIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerIdRequest.Marshal(b, m, deterministic)
}
func (m *PlayerIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerIdRequest.Merge(m, src)
}
func (m *PlayerIdRequest) XXX_Size() int {
	return xxx_messageInfo_PlayerIdRequest.Size(m)
}
func (m *PlayerIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerIdRequest proto.InternalMessageInfo

func (m *PlayerIdRequest) GetPlayerId() uint32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type PlayersReply struct {
	Player               []*Player `protobuf:"bytes,1,rep,name=player,proto3" json:"player,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PlayersReply) Reset()         { *m = PlayersReply{} }
func (m *PlayersReply) String() string { return proto.CompactTextString(m) }
func (*PlayersReply) ProtoMessage()    {}
func (*PlayersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4057945fb2ade8ac, []int{4}
}

func (m *PlayersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayersReply.Unmarshal(m, b)
}
func (m *PlayersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayersReply.Marshal(b, m, deterministic)
}
func (m *PlayersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayersReply.Merge(m, src)
}
func (m *PlayersReply) XXX_Size() int {
	return xxx_messageInfo_PlayersReply.Size(m)
}
func (m *PlayersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayersReply.DiscardUnknown(m)
}

var xxx_messageInfo_PlayersReply proto.InternalMessageInfo

func (m *PlayersReply) GetPlayer() []*Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func init() {
	proto.RegisterType((*Player)(nil), "WTA.Player")
	proto.RegisterType((*Ranking)(nil), "WTA.Ranking")
	proto.RegisterType((*PlayerWithRanking)(nil), "WTA.PlayerWithRanking")
	proto.RegisterType((*PlayerIdRequest)(nil), "WTA.PlayerIdRequest")
	proto.RegisterType((*PlayersReply)(nil), "WTA.PlayersReply")
}

func init() { proto.RegisterFile("section-3-4/src/proto/wta.proto", fileDescriptor_4057945fb2ade8ac) }

var fileDescriptor_4057945fb2ade8ac = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xeb, 0x14, 0xba, 0xdb, 0xc9, 0x16, 0xb4, 0x16, 0xa0, 0xa8, 0x20, 0x6d, 0x64, 0x10,
	0xca, 0xa5, 0x29, 0x6a, 0x39, 0x80, 0x84, 0x84, 0x16, 0x90, 0x76, 0x39, 0x80, 0x2a, 0xab, 0x52,
	0xaf, 0xb8, 0x8d, 0x37, 0xb5, 0x68, 0xe3, 0x10, 0xbb, 0x42, 0x79, 0x1e, 0x5e, 0x88, 0x03, 0x0f,
	0x84, 0x6a, 0xc7, 0xf9, 0x73, 0x28, 0xdc, 0x3c, 0x33, 0xdf, 0xd8, 0xbf, 0xef, 0x4b, 0xe0, 0x4a,
	0xf1, 0x8d, 0x16, 0x32, 0x9b, 0xcc, 0x27, 0xaf, 0xa7, 0xaa, 0xd8, 0x4c, 0xf3, 0x42, 0x6a, 0x39,
	0xfd, 0xa9, 0x59, 0x6c, 0x4e, 0xb8, 0xbf, 0x5a, 0x5e, 0x8f, 0xaf, 0x52, 0x29, 0xd3, 0x1d, 0xb7,
	0xc3, 0xf5, 0xe1, 0x6e, 0xaa, 0xc5, 0x9e, 0x2b, 0xcd, 0xf6, 0xb9, 0x55, 0x91, 0x3f, 0x08, 0x06,
	0x8b, 0x1d, 0x2b, 0x79, 0x81, 0x1f, 0x80, 0x27, 0x92, 0x00, 0x85, 0x28, 0x1a, 0x51, 0x4f, 0x24,
	0xf8, 0x19, 0x0c, 0xef, 0x44, 0xa1, 0xf4, 0x57, 0xb6, 0xe7, 0x81, 0x17, 0xa2, 0x68, 0x48, 0x9b,
	0x06, 0x1e, 0xc3, 0xf9, 0x8e, 0x55, 0xc3, 0xbe, 0x19, 0xd6, 0x35, 0x7e, 0x01, 0x23, 0xa1, 0xa8,
	0x48, 0xb7, 0xfa, 0x96, 0x65, 0x09, 0x4f, 0x82, 0x7b, 0x21, 0x8a, 0xce, 0x69, 0xb7, 0x89, 0xdf,
	0xc0, 0x70, 0x2d, 0x0a, 0xbd, 0xfd, 0xc4, 0x34, 0x0f, 0xee, 0x87, 0x28, 0xf2, 0x67, 0xe3, 0xd8,
	0xf2, 0xc6, 0x8e, 0x37, 0x5e, 0x3a, 0x5e, 0xda, 0x88, 0x71, 0x08, 0xfe, 0x46, 0x1e, 0x32, 0x5d,
	0x94, 0x1f, 0x65, 0xc2, 0x83, 0x81, 0x79, 0xbe, 0xdd, 0x22, 0xbf, 0x10, 0x9c, 0x51, 0x96, 0x7d,
	0x17, 0x59, 0x7a, 0x24, 0xcd, 0x8d, 0xc3, 0xcf, 0xce, 0x5d, 0x5d, 0xe3, 0x77, 0xe0, 0x17, 0x56,
	0x66, 0x28, 0xbc, 0xff, 0x52, 0xb4, 0xe5, 0x38, 0x80, 0xb3, 0xaa, 0x34, 0x11, 0x8c, 0xa8, 0x2b,
	0x8f, 0x09, 0x54, 0xc7, 0x85, 0x14, 0x99, 0x56, 0x26, 0x01, 0x8f, 0x76, 0x9b, 0xe4, 0x1b, 0x5c,
	0xda, 0xec, 0x57, 0x42, 0x6f, 0x1d, 0xee, 0x73, 0x18, 0x58, 0x3c, 0x03, 0xeb, 0xcf, 0xfc, 0x78,
	0xb5, 0xbc, 0x8e, 0xad, 0x8e, 0x56, 0x23, 0xfc, 0xb2, 0x79, 0xd9, 0x32, 0x5f, 0x18, 0x55, 0x75,
	0x47, 0xcd, 0x41, 0x26, 0xf0, 0x70, 0x51, 0x79, 0xa5, 0xfc, 0xc7, 0x81, 0x2b, 0xfd, 0xaf, 0x38,
	0xc8, 0x1c, 0x2e, 0xac, 0x5c, 0x51, 0x9e, 0xef, 0xca, 0x0e, 0x4b, 0xff, 0x04, 0xcb, 0xec, 0x37,
	0x82, 0xe3, 0xbf, 0x86, 0xbf, 0xc0, 0xd3, 0x1b, 0xae, 0x1b, 0x43, 0xb7, 0x22, 0xdd, 0x72, 0xa5,
	0x9d, 0xaf, 0x47, 0xad, 0xdd, 0x9a, 0x66, 0xfc, 0xa4, 0xd5, 0x6d, 0xa5, 0x40, 0x7a, 0xf8, 0x3d,
	0x3c, 0xbe, 0xe1, 0x6e, 0x5b, 0x7d, 0x28, 0xdd, 0xe6, 0x89, 0x8b, 0x3a, 0x01, 0x90, 0xde, 0x2b,
	0x84, 0xdf, 0x02, 0xd4, 0x3c, 0xea, 0xc4, 0xd6, 0x65, 0xab, 0x6b, 0x3d, 0x93, 0x5e, 0x84, 0xd6,
	0x03, 0xf3, 0xe5, 0xe7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce, 0x2a, 0x40, 0xe6, 0x65, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WTAClient is the client API for WTA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WTAClient interface {
	GetPlayerWithHighestRanking(ctx context.Context, in *PlayerIdRequest, opts ...grpc.CallOption) (*PlayerWithRanking, error)
	GetRankingsByPlayerId(ctx context.Context, in *PlayerIdRequest, opts ...grpc.CallOption) (WTA_GetRankingsByPlayerIdClient, error)
	GetPlayers(ctx context.Context, opts ...grpc.CallOption) (WTA_GetPlayersClient, error)
}

type wTAClient struct {
	cc *grpc.ClientConn
}

func NewWTAClient(cc *grpc.ClientConn) WTAClient {
	return &wTAClient{cc}
}

func (c *wTAClient) GetPlayerWithHighestRanking(ctx context.Context, in *PlayerIdRequest, opts ...grpc.CallOption) (*PlayerWithRanking, error) {
	out := new(PlayerWithRanking)
	err := c.cc.Invoke(ctx, "/WTA.WTA/GetPlayerWithHighestRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wTAClient) GetRankingsByPlayerId(ctx context.Context, in *PlayerIdRequest, opts ...grpc.CallOption) (WTA_GetRankingsByPlayerIdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WTA_serviceDesc.Streams[0], "/WTA.WTA/GetRankingsByPlayerId", opts...)
	if err != nil {
		return nil, err
	}
	x := &wTAGetRankingsByPlayerIdClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WTA_GetRankingsByPlayerIdClient interface {
	Recv() (*Ranking, error)
	grpc.ClientStream
}

type wTAGetRankingsByPlayerIdClient struct {
	grpc.ClientStream
}

func (x *wTAGetRankingsByPlayerIdClient) Recv() (*Ranking, error) {
	m := new(Ranking)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wTAClient) GetPlayers(ctx context.Context, opts ...grpc.CallOption) (WTA_GetPlayersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WTA_serviceDesc.Streams[1], "/WTA.WTA/GetPlayers", opts...)
	if err != nil {
		return nil, err
	}
	x := &wTAGetPlayersClient{stream}
	return x, nil
}

type WTA_GetPlayersClient interface {
	Send(*PlayerIdRequest) error
	CloseAndRecv() (*PlayersReply, error)
	grpc.ClientStream
}

type wTAGetPlayersClient struct {
	grpc.ClientStream
}

func (x *wTAGetPlayersClient) Send(m *PlayerIdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wTAGetPlayersClient) CloseAndRecv() (*PlayersReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PlayersReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WTAServer is the server API for WTA service.
type WTAServer interface {
	GetPlayerWithHighestRanking(context.Context, *PlayerIdRequest) (*PlayerWithRanking, error)
	GetRankingsByPlayerId(*PlayerIdRequest, WTA_GetRankingsByPlayerIdServer) error
	GetPlayers(WTA_GetPlayersServer) error
}

func RegisterWTAServer(s *grpc.Server, srv WTAServer) {
	s.RegisterService(&_WTA_serviceDesc, srv)
}

func _WTA_GetPlayerWithHighestRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WTAServer).GetPlayerWithHighestRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WTA.WTA/GetPlayerWithHighestRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WTAServer).GetPlayerWithHighestRanking(ctx, req.(*PlayerIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WTA_GetRankingsByPlayerId_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PlayerIdRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WTAServer).GetRankingsByPlayerId(m, &wTAGetRankingsByPlayerIdServer{stream})
}

type WTA_GetRankingsByPlayerIdServer interface {
	Send(*Ranking) error
	grpc.ServerStream
}

type wTAGetRankingsByPlayerIdServer struct {
	grpc.ServerStream
}

func (x *wTAGetRankingsByPlayerIdServer) Send(m *Ranking) error {
	return x.ServerStream.SendMsg(m)
}

func _WTA_GetPlayers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WTAServer).GetPlayers(&wTAGetPlayersServer{stream})
}

type WTA_GetPlayersServer interface {
	SendAndClose(*PlayersReply) error
	Recv() (*PlayerIdRequest, error)
	grpc.ServerStream
}

type wTAGetPlayersServer struct {
	grpc.ServerStream
}

func (x *wTAGetPlayersServer) SendAndClose(m *PlayersReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wTAGetPlayersServer) Recv() (*PlayerIdRequest, error) {
	m := new(PlayerIdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WTA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WTA.WTA",
	HandlerType: (*WTAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerWithHighestRanking",
			Handler:    _WTA_GetPlayerWithHighestRanking_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRankingsByPlayerId",
			Handler:       _WTA_GetRankingsByPlayerId_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetPlayers",
			Handler:       _WTA_GetPlayers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "section-3-4/src/proto/wta.proto",
}
