package contracts

import (
	"crud_multi_transport/db/models"
	"crud_multi_transport/usecase/viewmodel"
)

type IUserRepository interface {
	Browse(search, order, sort string, limit, offset int) (res []models.Users, count int, err error)

	ReadBy(column,value string) (res models.Users, err error)

	Edit(body viewmodel.UserVm, password string) (res string, err error)

	Add(body viewmodel.UserVm, password string) (res string, err error)

	Delete(ID string, updatedAt, deletedAt string) (res string, err error)

	IsExist(ID, email,mobilePhone string) (res bool, err error)
}
