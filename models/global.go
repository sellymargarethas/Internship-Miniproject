package models

import "time"

type Response struct {
	StatusCode       string      `json:"statusCode"`
	Success          bool        `json:"success"`
	ResponseDatetime time.Time   `json:"responseDatetime"`
	Result           interface{} `json:"result"`
	Message          string      `json:"message"`
}

type ResponseInsertTrx struct {
	StatusCode       string               `json:"statusCode"`
	Success          bool                 `json:"success"`
	ResponseDatetime time.Time            `json:"responseDatetime"`
	Result           HistoryTransactionId `json:"result"`
	Message          string               `json:"message"`
}

type ResponseTransaction struct {
	StatusCode       string      `json:"statusCode"`
	Success          bool        `json:"success"`
	ResponseDatetime time.Time   `json:"responseDatetime"`
	Nama             interface{} `json:"nama"`
	Tanggal          interface{} `json:"tanggal"`
	Price            int         `json:"price"`
	Trx_id           int         `json:"transaction_id"`
	Message          string      `json:"message"`
}

type ResponseTransactionByTrxId struct {
	StatusCode       string               `json:"statusCode"`
	Success          bool                 `json:"success"`
	ResponseDatetime time.Time            `json:"responseDatetime"`
	Nama             interface{}          `json:"nama"`
	Tanggal          interface{}          `json:"tanggal"`
	Price            int                  `json:"price"`
	Trx_id           int                  `json:"transaction_id"`
	Respons          []HistoryTransaction `json:"respons"`
	Message          string               `json:"message"`
}

type ErrorMsg struct {
	Status    string `json:"status"`
	ErrorDesc string `json:"errordesc"`
}
