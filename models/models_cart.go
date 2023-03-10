package models

type Cart struct {
	Id_user        int          `json:"id_user" validate:"required"`
	Id_product     int          `json:"id_product" validate:"required"`
	Qty_cart       int          `json:"qty_cart" validate:"required"`
	Total_cart     int          `json:"total_cart"`
	Delete_product []DeleteCart `json:"id_removed"`
}
type CartArray struct {
	Cart []DeleteCart `json:"cart_details"`
}

type DeleteCart struct {
	Id_product int `json:"id_product"`
}

type ViewCart struct {
	Id_product   int    `json:"id_product"`
	Nama_product string `json:"nama_product"`
	Qty_cart     int    `json:"qty_cart"`
	Total_cart   int    `json:"total_cart"`
	Harga_satuan int    `json:"price"`
}

type ViewCartArray struct {
	ViewCartArray []ViewCart
}

type TotalCartPrice struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
	CartPrice    int         `json:"total_cart"`
}

type PaymentMethod struct {
	Id_payment     int    `json:"id_payment"`
	Payment_method string `json:"payment_method"`
}
