package transport

import (
	"context"
	"crud_multi_transport/grpc_server/endpoint"
	"crud_multi_transport/grpc_server/protobuf"
	"crud_multi_transport/grpc_server/transformator"
	"github.com/go-kit/kit/transport/grpc"
)

type GRPCTransport struct {
	browse grpc.Handler
	read   grpc.Handler
	edit   grpc.Handler
	add    grpc.Handler
	delete grpc.Handler
}

func (s *GRPCTransport) Browse(ctx context.Context, req *protobuf.BrowseRequest) (*protobuf.Response, error) {
	_, resp, err := s.browse.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*protobuf.Response), nil
}

func (s *GRPCTransport) Read(ctx context.Context, req *protobuf.ReadRequest) (*protobuf.Response, error) {
	_, resp, err := s.read.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*protobuf.Response), err
}

func (s *GRPCTransport) Edit(ctx context.Context, req *protobuf.EditRequest) (*protobuf.Response, error) {
	_, resp, err := s.edit.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*protobuf.Response), err
}

func (s *GRPCTransport) Add(ctx context.Context, req *protobuf.UserRequest) (*protobuf.Response, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*protobuf.Response), err
}

func (s *GRPCTransport) Delete(ctx context.Context, req *protobuf.DeleteRequest) (*protobuf.Response, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*protobuf.Response), err
}

func NewGRPCServer(_ context.Context, endpoint endpoint.Endpoints) protobuf.UserServer {
	return &GRPCTransport{
		browse: grpc.NewServer(endpoint.BrowseEndpoint, transformator.DecodeBrowseRequest, transformator.EncodeResponse),
		read: grpc.NewServer(endpoint.ReadEndPoint, transformator.DecodeReadRequest, transformator.EncodeResponse),
		edit: grpc.NewServer(endpoint.EditEndPoint, transformator.DecodeEditRequest, transformator.EncodeResponse),
		add: grpc.NewServer(endpoint.AddEndPoint, transformator.DecodeAddRequest, transformator.EncodeResponse),
		delete: grpc.NewServer(endpoint.DeleteEndPoint, transformator.DecodeDeleteRequest, transformator.EncodeResponse),
	}
}
