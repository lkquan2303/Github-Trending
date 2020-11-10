package repository

import (
	"GitHub-Trending/model"
	"GitHub-Trending/model/req"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error)
	SelectUserById(context context.Context, userId string) (model.User, error)
}
