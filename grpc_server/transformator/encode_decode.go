package transformator

import (
	"context"
	"crud_multi_transport/grpc_server/protobuf"
	"crud_multi_transport/grpc_server/viewmodel"
)

func EncodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(viewmodel.ResponseVm)
	resProto := &protobuf.Response{}
	for _,data := range res.Data {
		temp := &protobuf.UserResponse{
			ID:                   data.ID,
			FullName:             data.FullName,
			Email:                data.Email,
			MobilePhone:          data.MobilePhone,
			CreatedAt:            data.CreatedAt,
			UpdatedAt:            data.UpdatedAt,
			DeletedAt:            data.DeletedAt,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		}
		resProto.Data = append(resProto.Data,temp)
	}
	resProto.ErrorMessage = res.Message.(string)
	resProto.Pagination = &protobuf.PaginationResponse{
		CurrentPage:          int32(res.Pagination.CurrentPage),
		LastPage:             int32(res.Pagination.LastPage),
		Total:                int32(res.Pagination.Total),
		PerPage:              int32(res.Pagination.PerPage),
	}
	return resProto,nil
}


func DecodeBrowseRequest(_ context.Context, request interface{}) (interface{},error){
	req := request.(*protobuf.BrowseRequest)
	return BrowseRequest{
		Search: req.Search,
		Sort:   req.Sort,
		Order:  req.Order,
		Page:   req.Page,
		Limit:  req.Limit,
	},nil
}

func DecodeAddRequest(_ context.Context, request interface{})(interface{},error){
	req := request.(*protobuf.UserRequest)
	return UserRequest{
		FullName:    req.FullName,
		Email:       req.Email,
		Password:    req.Password,
		MobilePhone: req.MobilePhone,
	},nil
}

func DecodeEditRequest(_ context.Context,request interface{}) (interface{},error){
	req := request.(*protobuf.EditRequest)
	return EditRequest{
		ID:   req.ID,
		User: UserRequest{
			FullName:    req.UserRequest.FullName,
			Email:       req.UserRequest.Email,
			Password:    req.UserRequest.Password,
			MobilePhone: req.UserRequest.MobilePhone,
		},
	},nil
}

func DecodeDeleteRequest(_ context.Context,request interface{}) (interface{},error){
	req := request.(*protobuf.DeleteRequest)
	return DeleteRequest{ID: req.ID},nil
}

func DecodeReadRequest(_ context.Context,request interface{}) (interface{},error){
	req := request.(*protobuf.ReadRequest)
	return ReadRequest{ID: req.ID},nil
}