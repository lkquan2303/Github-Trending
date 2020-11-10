package middleware

import (
	"GitHub-Trending/model"
	"GitHub-Trending/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: security.SECRECT_KEY,
	}
	return middleware.JWTWithConfig(config)
}
