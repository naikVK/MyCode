package model

import jwt "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserName  string `json:"username"`
	IsAdmin   bool   `json:"isAdmin"`
	SessionId string `json:"sessionId"`
	ClientId  string `json:"clientId"`
	jwt.StandardClaims
}
