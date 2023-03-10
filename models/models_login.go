package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type LoginUser struct {
	Email    string `json:"email_user" validate:"required"`
	Password string `json:"password_user" validate:"required"`
	Name     string `json:"nama_user"`
}

type LoginResponse struct {
	ResponseCode int         `json:"statuscode"`
	Token        string      `json:"token"`
	Error        ErrorMsg    `json:"error"`
	Response     interface{} `json:"response"`
	Nama         string      `json:"nama_user"`
	Id_user      int         `json:"id_user"`
	Role         string      `json:"role"`
}
type LoginDataResponse struct {
	Nama    string `json:"nama_user"`
	Id_user int    `json:"id_user"`
}

type JWTClaim struct {
	Email    string `json:"email_user"`
	Password string `json:"password_user"`
	jwt.StandardClaims
}
