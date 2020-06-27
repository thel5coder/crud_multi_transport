package endpoint

import (
	"context"
	"crud_multi_transport/grpc_server/service"
	"crud_multi_transport/grpc_server/transformator"
	"crud_multi_transport/grpc_server/viewmodel"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	BrowseEndpoint endpoint.Endpoint
	ReadEndPoint   endpoint.Endpoint
	EditEndPoint   endpoint.Endpoint
	AddEndPoint    endpoint.Endpoint
	DeleteEndPoint endpoint.Endpoint
}

func MakeBrowseEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		grpcReq := req.(transformator.BrowseRequest)
		data, pagination, err := service.Browse(grpcReq.Search, grpcReq.Sort, grpcReq.Order, int(grpcReq.Page), int(grpcReq.Limit))
		if err != nil {
			return viewmodel.ResponseVm{
				Message:    err.Error(),
				Data:       nil,
				Pagination: pagination,
			}, err
		}

		return viewmodel.ResponseVm{
			Message:    nil,
			Data:       data,
			Pagination: pagination,
		}, nil
	}
}

func MakeReadEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		grpcReq := req.(transformator.ReadRequest)
		data, err := service.Read(grpcReq.ID)
		if err != nil {
			return viewmodel.ResponseVm{
				Message:    err.Error(),
			}, err
		}

		res := viewmodel.ResponseVm{}
		res.Data = append(res.Data, data)
		return res, nil
	}
}

func MakeEditEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		grpcReq := req.(transformator.EditRequest)
		err := service.Edit(grpcReq.User.FullName,grpcReq.User.Email,grpcReq.User.Password,grpcReq.User.MobilePhone, grpcReq.ID)
		if err != nil {
			return viewmodel.ResponseVm{
				Message:    err.Error(),
			}, err
		}

		return viewmodel.ResponseVm{
			Message:    err.Error(),
		}, err
	}
}

func MakeAddEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		grpcReq := req.(transformator.UserRequest)
		err := service.Add(grpcReq.FullName,grpcReq.Email,grpcReq.Password,grpcReq.MobilePhone)
		if err != nil {
			return viewmodel.ResponseVm{
				Message:    err.Error(),
			}, err
		}

		return viewmodel.ResponseVm{}, err
	}
}

func MakeDeleteEndPoint(service service.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		grpcReq := req.(transformator.DeleteRequest)
		err := service.Delete(grpcReq.ID)
		if err != nil {
			return viewmodel.ResponseVm{}, err
		}

		return viewmodel.ResponseVm{}, err
	}
}
