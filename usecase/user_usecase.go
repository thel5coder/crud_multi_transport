package usecase

import (
	"crud_multi_transport/db/repositories/actions"
	"crud_multi_transport/helpers/hashing"
	"crud_multi_transport/helpers/messages"
	"crud_multi_transport/http_server/request"
	"crud_multi_transport/usecase/viewmodel"
	"errors"
	"time"
)

type UserUseCase struct {
	*UcContract
}

func (uc UserUseCase) Browse(search, order, sort string, page, limit int) (res []viewmodel.UserVm, pagination viewmodel.PaginationVm, err error) {
	repository := actions.NewUserRepository(uc.DB)
	offset, limit, page, order, sort := uc.setPaginationParameter(page, limit, order, sort)

	users, count, err := repository.Browse(search, order, sort, limit, offset)
	if err != nil {
		return res, pagination, err
	}

	for _, user := range users {
		res = append(res, viewmodel.UserVm{
			ID:          user.ID,
			FullName:    user.FullName,
			Email:       user.Email,
			MobilePhone: user.MobilePhone,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt.String,
			DeletedAt:   user.DeletedAt.String,
		})
	}
	pagination = uc.setPaginationResponse(page, limit, count)

	return res, pagination, err
}

func (uc UserUseCase) ReadByPk(ID string) (res viewmodel.UserVm, err error) {
	repository := actions.NewUserRepository(uc.DB)

	user, err := repository.ReadBy("id",ID)
	if err != nil {
		return res, err
	}

	res = viewmodel.UserVm{
		ID:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		MobilePhone: user.MobilePhone,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt.String,
		DeletedAt:   user.DeletedAt.String,
	}

	return res, err
}

func (uc UserUseCase) Edit(input *request.UserRequest,ID string) (err error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()
	password := ""

	isExist,err := uc.IsExist(ID,input.Email, input.MobilePhone)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	if input.Password != ""{
		password,err = hashing.HashAndSalt(input.Password)
		if err != nil {
			return err
		}
	}

	body := viewmodel.UserVm{
		ID:          ID,
		FullName:    input.FullName,
		Email:       input.Email,
		MobilePhone: input.MobilePhone,
		CreatedAt:   "",
		UpdatedAt:   now.Format(time.RFC3339),
		DeletedAt:   "",
	}
	_,err = repository.Edit(body,password)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) Add(input *request.UserRequest) (err error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	isExist, err := uc.IsExist("", input.Email, input.MobilePhone)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New(messages.DataAlreadyExist)
	}

	password, err := hashing.HashAndSalt(input.Password)
	if err != nil {
		return err
	}

	body := viewmodel.UserVm{
		ID:          "",
		FullName:    input.FullName,
		Email:       input.Email,
		MobilePhone: input.MobilePhone,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}
	_,err = repository.Add(body,password)
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) Delete(ID string) (err error){
	repository := actions.NewUserRepository(uc.DB)
	now := time.Now().UTC()

	_, err = repository.Delete(ID,now.Format(time.RFC3339),now.Format(time.RFC3339))
	if err != nil {
		return err
	}

	return nil
}

func (uc UserUseCase) IsExist(ID, email, mobilePhone string) (res bool, err error){
	repository := actions.NewUserRepository(uc.DB)
	res,err = repository.IsExist(ID,email,mobilePhone)
	if err != nil {
		return res,err
	}

	return res,err
}

