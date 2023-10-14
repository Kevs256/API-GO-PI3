package models

import "github.com/golang-jwt/jwt/v4"

type ClaimsJWT struct {
	Id int `json:"id_user"`
	jwt.RegisteredClaims
}
