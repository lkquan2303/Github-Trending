package router

import (
	"GitHub-Trending/handler"
	"GitHub-Trending/middleware"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	//	g := api.Echo.Group("/user", middleware.AddTrailingSlash())
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
