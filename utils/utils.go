package utils

import (
	"Internship-Miniproject/config"
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"database/sql"
	"encoding/json"
	"errors"

	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func BindValidateStruct(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}

	if err := ctx.Validate(i); err != nil {
		return err
	}

	return nil
}

func ResponseJSON(success bool, code string, msg string, result interface{}) models.Response {
	tm := time.Now()
	response := models.Response{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return response
}

func ResponseInsertTrxJSON(success bool, code string, msg string, result models.HistoryTransactionId) models.ResponseInsertTrx {
	tm := time.Now()
	response := models.ResponseInsertTrx{
		Success:          success,
		StatusCode:       code,
		Result:           result,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return response
}

func ResponseTransactionJSON(success bool, code string, msg string, nama interface{}, tanggal interface{}, price int, trx_id int) models.ResponseTransaction {
	tm := time.Now()
	responseTransaction := models.ResponseTransaction{
		Success:          success,
		StatusCode:       code,
		Nama:             nama,
		Tanggal:          tanggal,
		Price:            price,
		Trx_id:           trx_id,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return responseTransaction
}

func ResponseTransactionByTrxIdJSON(success bool, code string, msg string, nama interface{}, tanggal interface{}, price int, trx_id int, respons []models.HistoryTransaction) models.ResponseTransactionByTrxId {
	tm := time.Now()
	responseTransaction := models.ResponseTransactionByTrxId{
		Success:          success,
		StatusCode:       code,
		Nama:             nama,
		Tanggal:          tanggal,
		Price:            price,
		Trx_id:           trx_id,
		Respons:          respons,
		Message:          msg,
		ResponseDatetime: tm,
	}

	return responseTransaction
}

func TimeStampNow() string {
	return time.Now().Format(constants.LAYOUT_TIMESTAMP)
}

func ReplaceSQL(old, searchPattern string) string {
	tmpCount := strings.Count(old, searchPattern)
	for m := 1; m <= tmpCount; m++ {
		old = strings.Replace(old, searchPattern, "$"+strconv.Itoa(m), 1)
	}
	return old
}

func DBTransaction(db *sql.DB, txFunc func(*sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // Rollback Panic
		} else if err != nil {
			tx.Rollback() // err is not nill
		} else {
			err = tx.Commit() // err is nil
		}
	}()
	err = txFunc(tx)
	return err
}

func Stringify(input interface{}) string {
	bytes, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}
	strings := string(bytes)
	bytes, err = json.Marshal(strings)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func JSONPrettyfy(data interface{}) string {
	bytesData, _ := json.MarshalIndent(data, "", "  ")
	return string(bytesData)
}

func ToString(i interface{}) string {
	log, _ := json.Marshal(i)
	logString := string(log)

	return logString
}

func GenerateToken(request models.LoginUser) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = request.Email
	claims["password"] = request.Password
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()

	// Generate encoded token and send it as response.
	restoken, err := token.SignedString([]byte(config.GetEnv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return restoken, nil
}

// Validate Token
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("JWT_KEY")), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
