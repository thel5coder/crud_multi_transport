package viewmodel

import "crud_multi_transport/usecase/viewmodel"

type ResponseVm struct {
	Message    interface{}            `json:"message"`
	Data       []viewmodel.UserVm     `json:"data"`
	Pagination viewmodel.PaginationVm `json:"pagination"`
}
