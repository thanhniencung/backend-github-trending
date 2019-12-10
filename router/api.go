package router

import (
	"backend-github-trending/handler"
	"backend-github-trending/middleware"
	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", 		api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", 		api.UserHandler.HandleSignUp)

	// profile
	user := api.Echo.Group("/user", middleware.JWTMiddleware())
	user.GET("/profile", 			api.UserHandler.Profile)
	user.PUT("/profile/update", 	api.UserHandler.UpdateProfile)
}
