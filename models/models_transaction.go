package models

import "time"

type Transaction struct {
	Id_transaction          int `json:"id_transaction"`
	Id_user                 int `json:"id_user"`
	Id_payment              int `json:"id_payment"`
	Created_on              time.Time
	TransactionDetailsArray []TransactionDetails `json:"transactiondetailsarray"`
}

type TxHistoryResume struct {
	Id_transaction      int                `json:"id_transaction"`
	Id_user             int                `json:"id_user"`
	Nama_user           string             `json:"nama_user"`
	Tanggal_transaction string             `json:"tanggal_transaction"`
	Total_price         int                `json:"total_price"`
	Jml_item            int                `json:"jml_item"`
	TransactionDetails  []TxHistoryDetails `json:"details"`
}

type TransactionId struct {
	Id_transaction          int                   `json:"id_transaction"`
	Id_user                 int                   `json:"id_user"`
	Nama_user               string                `json:"nama_user"`
	Id_payment              int                   `json:"id_payment"`
	Created_on              string                `json:"created_on"`
	Total                   int                   `json:"total"`
	TransactionDetailsArray []TransactionDetails2 `json:"transactiondetailsarray"`
}

type TransactionDetails struct {
	Id_transaction int `json:"id_transaction"`
	Id_product     int `json:"id_product"`
	Qty_product    int `json:"qty_product"`
	Total          int `json:"total_transaction_details"`
}

type TransactionDetails2 struct {
	Id_transaction int `json:"id_transaction"`
	Id_product     int `json:"id_product"`
	Qty_product    int `json:"qty_product"`
}

type TransactionCount struct {
	Id_transaction int `json:"id_transaction" validate:"required"`
	Total_price    int `json:"total_price"`
}

type HistoryTransaction struct {
	Id_transaction    int    `json:"id_transaction"`
	Nama_user         string `json:"nama_user"`
	Id_product        int    `json:"id_product"`
	Nama_product      string `json:"nama_product"`
	Type_payment      string `json:"type_payment"`
	Qty_transaction   int    `json:"qty_transaction"`
	Total_transaction int    `json:"total_transaction"`
	Total_price       int    `json:"total_price"`
}

type HistoryTransactionId struct {
	Id_transaction      int    `json:"id_transaction"`
	Nama_user           string `json:"nama_user"`
	Tanggal_transaction string `json:"tanggal_transaction"`
	Total_price         int    `json:"total_price"`
}

type TxHistoryDetails struct {
	Id_product        int    `json:"id_product"`
	Nama_product      string `json:"nama_product"`
	Qty_transaction   int    `json:"qty_transaction"`
	Total_transaction int    `json:"total_transaction"`
}

type TransactionResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
}

type TransactionHistoryResponse struct {
	ResponseCode int         `json:"responseCode"`
	Message      string      `json:"message"`
	Response     interface{} `json:"response"`
	Nama         interface{} `json:"nama"`
	Tanggal      interface{} `json:"tanggal"`
	Price        int         `json:"price"`
	Trx_id       int         `json:"trassssnsaction_id"`
}

type TransactionArray struct {
	TransactionArray []Transaction
}

type SearchHistoryDate struct {
	Date              string `json:"date"`
	Id_transaction    int    `json:"id_transaction"`
	Id_user           int    `json:"id_user"`
	Nama_user         string `json:"nama_user"`
	Id_product        int    `json:"id_product"`
	Nama_product      string `json:"nama_product"`
	Type_payment      string `json:"type_payment"`
	Qty_transaction   int    `json:"qty_transaction"`
	Total_transaction int    `json:"total_transaction"`
	Total_price       int    `json:"total_price"`
}

type SearchHistoryDateArray struct {
	SearchHistoryDateArray []SearchHistoryDate
}
type IniSementara struct {
	Id_transaction int `json:"id_transaction"`
}
