// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/comprehension/file.proto

package comprehension

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for FileService service

func NewFileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FileService service

type FileService interface {
	DownloadProductImage(ctx context.Context, opts ...client.CallOption) (FileService_DownloadProductImageService, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) DownloadProductImage(ctx context.Context, opts ...client.CallOption) (FileService_DownloadProductImageService, error) {
	req := c.c.NewRequest(c.name, "FileService.DownloadProductImage", &DataChunk{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &fileServiceDownloadProductImage{stream}, nil
}

type FileService_DownloadProductImageService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*DataChunk) error
}

type fileServiceDownloadProductImage struct {
	stream client.Stream
}

func (x *fileServiceDownloadProductImage) Close() error {
	return x.stream.Close()
}

func (x *fileServiceDownloadProductImage) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceDownloadProductImage) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceDownloadProductImage) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceDownloadProductImage) Send(m *DataChunk) error {
	return x.stream.Send(m)
}

// Server API for FileService service

type FileServiceHandler interface {
	DownloadProductImage(context.Context, FileService_DownloadProductImageStream) error
}

func RegisterFileServiceHandler(s server.Server, hdlr FileServiceHandler, opts ...server.HandlerOption) error {
	type fileService interface {
		DownloadProductImage(ctx context.Context, stream server.Stream) error
	}
	type FileService struct {
		fileService
	}
	h := &fileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FileService{h}, opts...))
}

type fileServiceHandler struct {
	FileServiceHandler
}

func (h *fileServiceHandler) DownloadProductImage(ctx context.Context, stream server.Stream) error {
	return h.FileServiceHandler.DownloadProductImage(ctx, &fileServiceDownloadProductImageStream{stream})
}

type FileService_DownloadProductImageStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*DataChunk, error)
}

type fileServiceDownloadProductImageStream struct {
	stream server.Stream
}

func (x *fileServiceDownloadProductImageStream) Close() error {
	return x.stream.Close()
}

func (x *fileServiceDownloadProductImageStream) Context() context.Context {
	return x.stream.Context()
}

func (x *fileServiceDownloadProductImageStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *fileServiceDownloadProductImageStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *fileServiceDownloadProductImageStream) Recv() (*DataChunk, error) {
	m := new(DataChunk)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Api Endpoints for FileService service

func NewFileServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for FileService service

type FileService interface {
	SubmitFile(ctx context.Context, in *Submission, opts ...client.CallOption) (*Response, error)
}

type fileService struct {
	c    client.Client
	name string
}

func NewFileService(name string, c client.Client) FileService {
	return &fileService{
		c:    c,
		name: name,
	}
}

func (c *fileService) SubmitFile(ctx context.Context, in *Submission, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "FileService.SubmitFile", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FileService service

type FileServiceHandler interface {
	SubmitFile(context.Context, *Submission, *Response) error
}

func RegisterFileServiceHandler(s server.Server, hdlr FileServiceHandler, opts ...server.HandlerOption) error {
	type fileService interface {
		SubmitFile(ctx context.Context, in *Submission, out *Response) error
	}
	type FileService struct {
		fileService
	}
	h := &fileServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&FileService{h}, opts...))
}

type fileServiceHandler struct {
	FileServiceHandler
}

func (h *fileServiceHandler) SubmitFile(ctx context.Context, in *Submission, out *Response) error {
	return h.FileServiceHandler.SubmitFile(ctx, in, out)
}
