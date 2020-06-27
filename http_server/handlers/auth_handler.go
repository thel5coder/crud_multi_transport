package handlers

import (
	"crud_multi_transport/http_server/request"
	"crud_multi_transport/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"net/http"
)

type AuthHandler struct {
	Handler
}

func (handler AuthHandler) Login(ctx echo.Context) error {
	input := new(request.LoginRequest)

	if err := ctx.Bind(input); err != nil {
		return handler.SendResponseBadRequest(ctx, http.StatusBadRequest, err.Error())
	}
	if err := handler.Validate.Struct(input); err != nil {
		return handler.SendResponseErrorValidation(ctx, err.(validator.ValidationErrors))
	}

	uc := usecase.AuthUseCase{UcContract: handler.UseCaseContract}
	res, err := uc.Login(input)
	if err != nil {
		return handler.SendResponseUnauthorized(ctx, err)
	}

	return handler.SendResponse(ctx, res, nil, err)
}
