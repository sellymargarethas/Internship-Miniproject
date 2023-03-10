package models

import (
	"time"
)

type User struct {
	Id_user       int      `json:"id_user"`
	Nama_user     string   `json:"nama_user" validate:"required,min=3"`
	Email_user    string   `json:"email_user" validate:"required,email"`
	Password_user string   `json:"password_user" validate:"required,min=8,max=50"`
	Role_user     string   `json:"role_user"`
	DataUser      DataUser `json:"data_user"`
	Created_on    time.Time
	Updated_on    time.Time
}

type DataUser struct {
	Gender_user string `json:"gender_user"`
	Foto_user   string `json:"foto_user"`
	Phone_user  string `json:"phone_user"`
	Ttl_user    string `json:"ttl_user"`
	Kota_user   string `json:"kota_user"`
}

type UserProfile struct {
	Id_user     int    `json:"id_user"`
	Nama_user   string `json:"nama_user"`
	Email_user  string `json:"email_user"`
	Gender_user string `json:"gender_user"`
	Foto_user   string `json:"foto_user"`
	Phone_user  string `json:"phone_user"`
	TTl_user    string `json:"ttl_user"`
	Kota_user   string `json:"kota_user"`
}

type UpdateUser struct {
	Nama_user     string `json:"nama_user" validate:"required"`
	Email_user    string `json:"email_user" validate:"required"`
	Password_user string `json:"password_user" validate:"required,min=8,max=50"`
}

type UserResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
}
