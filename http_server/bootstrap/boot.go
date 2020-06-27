package bootstrap

import (
	"crud_multi_transport/usecase"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Bootstrap struct {
	E               *echo.Echo
	UseCaseContract usecase.UcContract
	Validator       *validator.Validate
	Translator      ut.Translator
}
