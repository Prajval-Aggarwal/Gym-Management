package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId   string
	Username string `json:"user_id"`
	jwt.RegisteredClaims
}
