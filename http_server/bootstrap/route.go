package bootstrap

import (
	"crud_multi_transport/http_server/bootstrap/routes"
	"crud_multi_transport/http_server/handlers"
)

func (boot *Bootstrap) RegisterRouters(){
	 handler := handlers.Handler{
		E:               boot.E,
		UseCaseContract: &boot.UseCaseContract,
		Validate:        boot.Validator,
		Translator:      boot.Translator,
	}

	api := boot.E.Group("/api/v1")

	//auth route
	authGroup := api.Group("/auth")
	authRoute := routes.AuthRoute{
		RouteGroup: authGroup,
		Handler:    handler,
	}
	authRoute.RegisterRoute()

	//user route
	userGroup := api.Group("/user")
	userRoute := routes.UserRoute{
		RouteGroup: userGroup,
		Handler:    handler,
	}
	userRoute.RegisterRoute()
}
