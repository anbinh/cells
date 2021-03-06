// Code generated by protoc-gen-go.
// source: github.com/pydio/config-srv/proto/config/config.proto
// DO NOT EDIT!

/*
Package go_micro_srv_config_config is a generated protocol buffer package.

It is generated from these files:
	github.com/pydio/config-srv/proto/config/config.proto

It has these top-level messages:
	Change
	ChangeLog
	CreateRequest
	CreateResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	SearchRequest
	SearchResponse
	WatchRequest
	WatchResponse
	AuditLogRequest
	AuditLogResponse
*/
package go_micro_srv_config_config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_os_config "github.com/pydio/go-os/config/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Change struct {
	Id        string                        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path      string                        `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Author    string                        `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Comment   string                        `protobuf:"bytes,4,opt,name=comment" json:"comment,omitempty"`
	Timestamp int64                         `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	ChangeSet *go_micro_os_config.ChangeSet `protobuf:"bytes,6,opt,name=change_set,json=changeSet" json:"change_set,omitempty"`
}

func (m *Change) Reset()                    { *m = Change{} }
func (m *Change) String() string            { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()               {}
func (*Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Change) GetChangeSet() *go_micro_os_config.ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type ChangeLog struct {
	Action string  `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Change *Change `protobuf:"bytes,2,opt,name=change" json:"change,omitempty"`
}

func (m *ChangeLog) Reset()                    { *m = ChangeLog{} }
func (m *ChangeLog) String() string            { return proto.CompactTextString(m) }
func (*ChangeLog) ProtoMessage()               {}
func (*ChangeLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChangeLog) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type UpdateRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteRequest struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteRequest) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ReadRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ReadResponse struct {
	Change *Change `protobuf:"bytes,1,opt,name=change" json:"change,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadResponse) GetChange() *Change {
	if m != nil {
		return m.Change
	}
	return nil
}

type SearchRequest struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author" json:"author,omitempty"`
	Limit  int64  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type SearchResponse struct {
	Configs []*Change `protobuf:"bytes,1,rep,name=configs" json:"configs,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *SearchResponse) GetConfigs() []*Change {
	if m != nil {
		return m.Configs
	}
	return nil
}

type WatchRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *WatchRequest) Reset()                    { *m = WatchRequest{} }
func (m *WatchRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()               {}
func (*WatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type WatchResponse struct {
	Id        string                        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ChangeSet *go_micro_os_config.ChangeSet `protobuf:"bytes,2,opt,name=change_set,json=changeSet" json:"change_set,omitempty"`
}

func (m *WatchResponse) Reset()                    { *m = WatchResponse{} }
func (m *WatchResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()               {}
func (*WatchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *WatchResponse) GetChangeSet() *go_micro_os_config.ChangeSet {
	if m != nil {
		return m.ChangeSet
	}
	return nil
}

type AuditLogRequest struct {
	// from unix timestamp
	From int64 `protobuf:"varint,1,opt,name=from" json:"from,omitempty"`
	// to unix timestamp
	To int64 `protobuf:"varint,2,opt,name=to" json:"to,omitempty"`
	// limit number items
	Limit int64 `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	// the offset
	Offset int64 `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
	// reverse the order
	Reverse bool `protobuf:"varint,5,opt,name=reverse" json:"reverse,omitempty"`
}

func (m *AuditLogRequest) Reset()                    { *m = AuditLogRequest{} }
func (m *AuditLogRequest) String() string            { return proto.CompactTextString(m) }
func (*AuditLogRequest) ProtoMessage()               {}
func (*AuditLogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type AuditLogResponse struct {
	Changes []*ChangeLog `protobuf:"bytes,1,rep,name=changes" json:"changes,omitempty"`
}

func (m *AuditLogResponse) Reset()                    { *m = AuditLogResponse{} }
func (m *AuditLogResponse) String() string            { return proto.CompactTextString(m) }
func (*AuditLogResponse) ProtoMessage()               {}
func (*AuditLogResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *AuditLogResponse) GetChanges() []*ChangeLog {
	if m != nil {
		return m.Changes
	}
	return nil
}

func init() {
	proto.RegisterType((*Change)(nil), "go.micro.srv.config.config.Change")
	proto.RegisterType((*ChangeLog)(nil), "go.micro.srv.config.config.ChangeLog")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.config.config.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.config.config.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.config.config.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.config.config.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.config.config.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.config.config.DeleteResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.config.config.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.config.config.ReadResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.config.config.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.config.config.SearchResponse")
	proto.RegisterType((*WatchRequest)(nil), "go.micro.srv.config.config.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "go.micro.srv.config.config.WatchResponse")
	proto.RegisterType((*AuditLogRequest)(nil), "go.micro.srv.config.config.AuditLogRequest")
	proto.RegisterType((*AuditLogResponse)(nil), "go.micro.srv.config.config.AuditLogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Config service

type ConfigClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	AuditLog(ctx context.Context, in *AuditLogRequest, opts ...client.CallOption) (*AuditLogResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Config_WatchClient, error)
}

type configClient struct {
	c           client.Client
	serviceName string
}

func NewConfigClient(serviceName string, c client.Client) ConfigClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.config.config"
	}
	return &configClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *configClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) AuditLog(ctx context.Context, in *AuditLogRequest, opts ...client.CallOption) (*AuditLogResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Config.AuditLog", in)
	out := new(AuditLogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configClient) Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Config_WatchClient, error) {
	req := c.c.NewRequest(c.serviceName, "Config.Watch", &WatchRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &configWatchClient{stream}, nil
}

type Config_WatchClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*WatchResponse, error)
}

type configWatchClient struct {
	stream client.Streamer
}

func (x *configWatchClient) Close() error {
	return x.stream.Close()
}

func (x *configWatchClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *configWatchClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *configWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Config service

type ConfigHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	AuditLog(context.Context, *AuditLogRequest, *AuditLogResponse) error
	Watch(context.Context, *WatchRequest, Config_WatchStream) error
}

func RegisterConfigHandler(s server.Server, hdlr ConfigHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Config{hdlr}, opts...))
}

type Config struct {
	ConfigHandler
}

func (h *Config) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ConfigHandler.Create(ctx, in, out)
}

func (h *Config) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ConfigHandler.Update(ctx, in, out)
}

func (h *Config) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ConfigHandler.Delete(ctx, in, out)
}

func (h *Config) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.ConfigHandler.Search(ctx, in, out)
}

func (h *Config) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ConfigHandler.Read(ctx, in, out)
}

func (h *Config) AuditLog(ctx context.Context, in *AuditLogRequest, out *AuditLogResponse) error {
	return h.ConfigHandler.AuditLog(ctx, in, out)
}

func (h *Config) Watch(ctx context.Context, stream server.Streamer) error {
	m := new(WatchRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ConfigHandler.Watch(ctx, m, &configWatchStream{stream})
}

type Config_WatchStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WatchResponse) error
}

type configWatchStream struct {
	stream server.Streamer
}

func (x *configWatchStream) Close() error {
	return x.stream.Close()
}

func (x *configWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *configWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *configWatchStream) Send(m *WatchResponse) error {
	return x.stream.Send(m)
}

func init() {
	proto.RegisterFile("github.com/pydio/config-srv/proto/config/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xa6, 0x4d, 0x5b, 0xd6, 0xeb, 0x5a, 0x26, 0x0b, 0xa1, 0x28, 0x02, 0x84, 0x22, 0x21, 0x36,
	0x60, 0xe9, 0x28, 0xe2, 0x05, 0x4d, 0x42, 0x68, 0x3c, 0x01, 0xe2, 0xc1, 0x15, 0xe2, 0x01, 0xa1,
	0x91, 0xa5, 0x6e, 0x1a, 0x69, 0xa9, 0x4b, 0xe2, 0xee, 0x8d, 0x3f, 0x86, 0xf8, 0x71, 0x38, 0x67,
	0xbb, 0x4d, 0x8a, 0xea, 0xa5, 0xda, 0x9e, 0x62, 0x9f, 0xef, 0xbe, 0xef, 0x3b, 0xf7, 0x3e, 0x17,
	0xde, 0xc4, 0x89, 0x98, 0x2d, 0x2f, 0x82, 0x88, 0xa7, 0xc3, 0x34, 0x89, 0x32, 0x3e, 0x8c, 0xf8,
	0x7c, 0x9a, 0xc4, 0xc7, 0x79, 0x76, 0x35, 0x5c, 0x64, 0x5c, 0x98, 0x80, 0xfe, 0x04, 0x18, 0x23,
	0x5e, 0xcc, 0x03, 0x4c, 0x0f, 0x64, 0x5e, 0xa0, 0x8f, 0xd4, 0xc7, 0x3b, 0xf9, 0x0f, 0x32, 0xe6,
	0xc7, 0x3c, 0x37, 0x38, 0x65, 0x50, 0x85, 0xe6, 0xff, 0x69, 0x40, 0xe7, 0x6c, 0x16, 0xce, 0x63,
	0x46, 0x06, 0xd0, 0x4c, 0x26, 0x6e, 0xe3, 0x49, 0xe3, 0xb0, 0x4b, 0xe5, 0x8a, 0x10, 0x68, 0x2d,
	0x42, 0x31, 0x73, 0x9b, 0x18, 0xc1, 0x35, 0x79, 0x00, 0x9d, 0x70, 0x29, 0x66, 0x3c, 0x73, 0x1d,
	0x8c, 0xea, 0x1d, 0x71, 0xe1, 0xae, 0xe4, 0x4c, 0xd9, 0x5c, 0xb8, 0x2d, 0x3c, 0x30, 0x5b, 0xf2,
	0x10, 0xba, 0x22, 0x49, 0x59, 0x2e, 0xc2, 0x74, 0xe1, 0xb6, 0xe5, 0x99, 0x43, 0xd7, 0x01, 0x72,
	0x0a, 0x10, 0x21, 0xfb, 0x79, 0xce, 0x84, 0xdb, 0x91, 0xc7, 0xbd, 0xd1, 0xa3, 0x60, 0xd5, 0x21,
	0xcf, 0x4d, 0x83, 0x4a, 0xe3, 0x98, 0x09, 0xda, 0x8d, 0xcc, 0xd2, 0x3f, 0x87, 0xae, 0x8a, 0x7f,
	0xe6, 0x31, 0x4a, 0x8b, 0x44, 0xc2, 0xe7, 0xba, 0x05, 0xbd, 0x23, 0x6f, 0xa1, 0xa3, 0x2a, 0xb0,
	0x91, 0xde, 0xc8, 0x0f, 0xb6, 0x5f, 0xa0, 0xa6, 0xa1, 0xba, 0xc2, 0xff, 0x04, 0xfd, 0xb3, 0x8c,
	0x85, 0x82, 0x51, 0xf6, 0x6b, 0x29, 0x25, 0x97, 0xc0, 0x1a, 0x3b, 0x83, 0x1d, 0xc0, 0xc0, 0x80,
	0xe5, 0x0b, 0x3e, 0xcf, 0x11, 0xfe, 0xeb, 0x62, 0x72, 0x7b, 0xf0, 0x06, 0x6c, 0x0d, 0xff, 0x81,
	0x5d, 0xb2, 0x5b, 0x83, 0x37, 0x60, 0x1a, 0xfe, 0x15, 0xf4, 0x28, 0x0b, 0x27, 0x06, 0xbc, 0xc6,
	0xf8, 0xf8, 0x1f, 0x61, 0x5f, 0x95, 0x28, 0x88, 0x1b, 0x09, 0x62, 0xd0, 0x1f, 0xb3, 0x30, 0x8b,
	0x66, 0xdb, 0x04, 0xac, 0x67, 0xb5, 0x59, 0x99, 0xd5, 0xfb, 0xd0, 0xbe, 0x4c, 0xd2, 0x44, 0xe0,
	0x08, 0x3b, 0x54, 0x6d, 0x8a, 0x6c, 0x3e, 0x9d, 0x16, 0x53, 0xd8, 0xc2, 0xb0, 0xde, 0xf9, 0x5f,
	0x60, 0x60, 0x68, 0xb4, 0xe8, 0xd3, 0x62, 0xd6, 0x0b, 0x45, 0xb9, 0x24, 0x73, 0x6a, 0xaa, 0x36,
	0x25, 0xfe, 0x63, 0xd8, 0xff, 0x16, 0x8a, 0xad, 0xaa, 0xfd, 0x1f, 0xd0, 0xd7, 0xe7, 0x9a, 0x6e,
	0xb3, 0xad, 0xaa, 0x65, 0x9a, 0x3b, 0x5a, 0xe6, 0x37, 0xdc, 0x7b, 0xbf, 0x9c, 0x24, 0x42, 0x3a,
	0xc6, 0x28, 0x90, 0x3f, 0xd4, 0x34, 0xe3, 0x29, 0x52, 0x38, 0x14, 0xd7, 0x05, 0xa9, 0xe0, 0x08,
	0xee, 0x50, 0xb9, 0xda, 0xed, 0xce, 0x8a, 0xd7, 0x20, 0x63, 0x57, 0x2c, 0xcb, 0x19, 0x3a, 0x7e,
	0x8f, 0x9a, 0xad, 0x3f, 0x86, 0x83, 0x35, 0xbd, 0x6e, 0xf0, 0x9d, 0xbc, 0x4f, 0xd4, 0x67, 0xee,
	0xf3, 0xe9, 0xf5, 0xf7, 0x59, 0xd4, 0x9b, 0xaa, 0xd1, 0xdf, 0xb6, 0x7c, 0xc3, 0xf0, 0x94, 0x84,
	0x72, 0x85, 0x1e, 0x23, 0x47, 0x56, 0x90, 0xb2, 0xa9, 0xbd, 0xe7, 0x75, 0x52, 0xf5, 0xd0, 0xdf,
	0x29, 0x28, 0x94, 0xcf, 0xec, 0x14, 0x15, 0x63, 0xdb, 0x29, 0x36, 0x6c, 0x8b, 0x14, 0xca, 0x6b,
	0x76, 0x8a, 0x8a, 0xb9, 0xed, 0x14, 0x1b, 0xd6, 0x45, 0x0a, 0x35, 0xd6, 0x76, 0x8a, 0x8a, 0xc3,
	0xec, 0x14, 0x55, 0x97, 0x48, 0x8a, 0xef, 0xd0, 0x2a, 0xcc, 0x4e, 0x9e, 0xd9, 0xaa, 0x4a, 0x2f,
	0x88, 0x77, 0x78, 0x7d, 0xe2, 0x0a, 0x3c, 0x86, 0x3d, 0x33, 0x48, 0xe4, 0x85, 0xad, 0x6e, 0x63,
	0xda, 0xbd, 0x97, 0xf5, 0x92, 0x57, 0x44, 0x3f, 0xa1, 0x8d, 0x7e, 0x24, 0x56, 0x75, 0x65, 0x4b,
	0x7b, 0x47, 0x35, 0x32, 0x0d, 0xfe, 0x49, 0xe3, 0xa2, 0x83, 0xff, 0xc4, 0xaf, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x54, 0xb2, 0x60, 0x13, 0x10, 0x08, 0x00, 0x00,
}
