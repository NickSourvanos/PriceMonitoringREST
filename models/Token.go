package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	AccessToken        string `json:"token"`
}

type Claims struct {
	UserId string `json:"userId"`
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type TokenUser struct {
	UserId string
	Username string
	Role string
	Exp int64
}