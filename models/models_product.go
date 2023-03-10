package models

type Product struct {
	Id_product        int    `json:"id_product"`
	Id_kategori       int    `json:"id_kategori" validate:"required"`
	Nama_product      string `json:"nama_product" validate:"required"`
	Deskripsi_product string `json:"deskripsi_product" validate:"required"`
	Qty_product       int    `json:"qty_product" validate:"required"`
	Harga_product     int    `json:"harga_product" validate:"required"`
}

type CategoryProduct struct {
	Id_product        int    `json:"id_product"`
	Nama_kategori     string `json:"nama_kategori"`
	Nama_product      string `json:"nama_product"`
	Deskripsi_product string `json:"deskripsi_product"`
	Qty_product       int    `json:"qty_product"`
	Harga_product     int    `json:"harga_product"`
}

type Topthree struct {
	Id_product          int    `json:"id_product"`
	Nama_kategori       string `json:"nama_kategori"`
	Nama_product        string `json:"nama_product"`
	Qty_product         int    `json:"qty_product"`
	Deskripsi_product   string `json:"deskripsi_product"`
	Qty_product_terjual int    `json:"qty_product_terjual"`
	Harga_product       int    `json:"harga_product"`
}

type CategoryProductArray struct {
	CategoryProductArray []CategoryProduct
}

type ProductArray struct {
	ProductArray []Product
}

type ProductResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
}

type OneProductResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"id_product"`
	Kategori     interface{} `json:"nama_kategori"`
	Id_kategori  int         `json:"id_kategori"`
	Product      interface{} `json:"nama_product"`
	Deskripsi    interface{} `json:"deskripsi_product"`
	Qty          interface{} `json:"qty_product"`
	Harga        interface{} `json:"harga_product"`
}
