package middleware

import (
	"GitHub-Trending/model"
	"GitHub-Trending/model/req"
	"github.com/labstack/echo/v4"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// handle logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    err.Error(),
					Data:       nil,
				})
			}
			if req.Email != "lkquan1609" {
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message:    "Error",
					Data:       nil,
				})
			}
			return next(c)
		}
	}
}
