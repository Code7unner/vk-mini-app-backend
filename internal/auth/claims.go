package auth

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}