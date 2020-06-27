package routes

import (
	"crud_multi_transport/http_server/handlers"
	"github.com/labstack/echo"
)

type AuthRoute struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route AuthRoute) RegisterRoute(){
	authHandler := handlers.AuthHandler{Handler: route.Handler}
	userHandler := handlers.UserHandler{Handler: route.Handler}

	route.RouteGroup.POST("/login", authHandler.Login)
	route.RouteGroup.POST("/register",userHandler.Add)

}
