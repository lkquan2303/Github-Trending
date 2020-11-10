package security

import (
	"GitHub-Trending/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const SECRECT_KEY = "screct_key"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRECT_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
