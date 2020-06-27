package service

import (
	input "crud_multi_transport/http_server/request"
	"crud_multi_transport/usecase"
	"crud_multi_transport/usecase/viewmodel"
)

type Service interface {
	Browse(search,sort,order string, page,limit int) ([]viewmodel.UserVm,viewmodel.PaginationVm,error)
	Read(ID string)(viewmodel.UserVm,error)
	Edit(fullName,email,password,mobilePhone,ID string) error
	Add(fullName,email,password,mobilePhone string) error
	Delete(ID string) error
}

type UserService struct {
	*usecase.UcContract
}

func (s *UserService) Browse(search,sort,order string, page, limit int)([]viewmodel.UserVm,viewmodel.PaginationVm,error){
	userUc := usecase.UserUseCase{UcContract:s.UcContract}
	return userUc.Browse(search,order,sort,page,limit)
}

func (s *UserService) Read(ID string)(viewmodel.UserVm,error){
	userUc := usecase.UserUseCase{UcContract:s.UcContract}
	return userUc.ReadByPk(ID)
}

func (s *UserService) Edit(fullName,email,password,mobilePhone,ID string) error{
	userUc := usecase.UserUseCase{UcContract:s.UcContract}
	req := input.UserRequest{
		FullName: fullName,
		Email: email,
		Password: password,
		MobilePhone: mobilePhone,
	}
	return userUc.Edit(&req,ID)
}

func (s *UserService) Add(fullName,email,password,mobilePhone string) error{
	userUc := usecase.UserUseCase{UcContract:s.UcContract}
	req := input.UserRequest{
		FullName: fullName,
		Email: email,
		Password: password,
		MobilePhone: mobilePhone,
	}
	return userUc.Add(&req)
}

func (s *UserService) Delete(ID string) error{
	userUc := usecase.UserUseCase{UcContract:s.UcContract}
	return userUc.Delete(ID)
}
