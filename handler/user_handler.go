package handler

import (
	"GitHub-Trending/banana"
	"GitHub-Trending/model"
	req "GitHub-Trending/model/req"
	"GitHub-Trending/repository"
	"GitHub-Trending/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	req := req.ReqSignIn{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//checkpass
	isPasswordTrue := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isPasswordTrue {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token
	user.Password = ""
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}
func (u *UserHandler) HandleSignUp(c echo.Context) error {
	req := req.ReqSignUp{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()
	userID, err := uuid.NewUUID()

	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user := model.User{
		UserId:   userID.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	token, err := security.GenToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error",
			Data:       nil,
		})
	}
	user.Token = token
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}

func (u *UserHandler) Profile(c echo.Context) error  {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user,err := u.UserRepo.SelectUserById(c.Request().Context(), claims.UserId)
	if err != nil{
		if err == banana.UserNotFound{
			return c.JSON(http.StatusNotFound, model.Response{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
				Data:       nil,
			})

		}

		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       user,
	})
}