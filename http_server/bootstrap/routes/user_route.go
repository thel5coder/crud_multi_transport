package routes

import (
	"crud_multi_transport/http_server/handlers"
	middleware "crud_multi_transport/http_server/middlewares"
	"github.com/labstack/echo"
)

type UserRoute struct {
	RouteGroup *echo.Group
	Handler    handlers.Handler
}

func (route UserRoute) RegisterRoute(){
	handler := handlers.UserHandler{Handler: route.Handler}
	jwtMiddleware := middleware.JwtVerify{UcContract:route.Handler.UseCaseContract}

	route.RouteGroup.Use(jwtMiddleware.JWTWithConfig)
	route.RouteGroup.GET("", handler.Browse)
	route.RouteGroup.GET("/:id", handler.Read)
	route.RouteGroup.PUT("/:id", handler.Edit)
	route.RouteGroup.POST("", handler.Add)
	route.RouteGroup.DELETE("/:id", handler.Delete)
}
