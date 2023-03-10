package models

type Category struct {
	Id_kategori   int    `json:"id_kategori"`
	Nama_kategori string `json:"nama_kategori" validate:"required"`
}
type CategoryArray struct {
	CategoryArray []Category `json:"category_array" validate:"required, dive"`
}

type CategoryResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
}
