package transactionService

import (
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type transactionService struct {
	Service services.UsecaseService
}

func NewTransactionService(service services.UsecaseService) transactionService {
	return transactionService{
		Service: service,
	}
}

// Validate Struct
func validateStructTransaction(x interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(x)
	if err != nil {
		return err
	}
	return nil
}

// Get All Transaction
func (svc transactionService) GetListTransaction(ctx echo.Context) error {
	var result models.Response
	var transaction []models.TxHistoryResume
	err, rows := svc.Service.TransactionRepo.GetHistory()
	for rows.Next() {
		var dataTx models.TxHistoryResume

		rows.Scan(&dataTx.Id_transaction, &dataTx.Jml_item, &dataTx.Id_user)
		for i := 0; i <= dataTx.Jml_item; i++ {
			dataTx.TransactionDetails, _ = svc.Service.TransactionRepo.GetHistoryDetails(dataTx.Id_transaction)
		}
		nama, tgl, price, err := svc.Service.TransactionRepo.GetHistoryData(dataTx.Id_transaction)
		dataTx.Nama_user = nama
		dataTx.Tanggal_transaction = tgl
		dataTx.Total_price = price

		fmt.Println(rows, err)

		transaction = append(transaction, dataTx)
	}

	if err != nil {
		log.Println("\nError InsertProduct- InsertProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Product", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListTransaction ", "Success Get Transaction")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, transaction)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Transaction
func (svc transactionService) InsertTransaction(ctx echo.Context) (err error) {
	var result models.ResponseInsertTrx
	var result2 models.Response
	var p models.Transaction

	err = ctx.Bind(&p)
	err = validateStructTransaction(p)

	if err != nil {
		log.Println("\nError InsertTransaction- InsertTransaction : ", err.Error())
		result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		fmt.Println(result2)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		id, err := svc.Service.TransactionRepo.PostNewTransaction(p, tx)
		fmt.Println(id)
		if err != nil {
			log.Println("\nError PostNewTransaction- PostNewTransaction : ")
			result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed PostNewTransaction", nil)
			fmt.Println(result2)
			return ctx.JSON(http.StatusBadRequest, result)
		} else {
			id2, err := svc.Service.TransactionRepo.CobaInsertMultipleItemsTransaction(p.TransactionDetailsArray, id, tx)
			fmt.Println(id2)
			if err != nil {
				log.Println("\nError CobaInsertMultipleItemsTransaction- CobaInsertMultipleItemsTransaction : ")
				result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed ICobaInsertMultipleItemsTransaction", nil)
				fmt.Println(result2)
				return ctx.JSON(http.StatusBadRequest, result)
			} else {
				id5, err := svc.Service.TransactionRepo.CountTotalItemTx(id, tx)
				fmt.Println(id5)
				if err != nil {
					log.Println("\nError CountTotalItemTx- CountTotalItemTx : ")
					result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed CountTotalItemTx", nil)
					fmt.Println(result2)
					return ctx.JSON(http.StatusBadRequest, result)
				} else {
					id3, err := svc.Service.TransactionRepo.AutoUpdateInventory(id, tx)
					fmt.Println(id3)
					if err != nil {
						log.Println("\nError AutoUpdateInventory- AutoUpdateInventory : ")
						result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed AutoUpdateInventory", nil)
						fmt.Println(result2)
						return ctx.JSON(http.StatusBadRequest, result)
					} else {
						id4, err := svc.Service.TransactionRepo.TransactionCount(id, tx)
						fmt.Println(id4)
						if err != nil {
							log.Println("\nError TransactionCount- TransactionCount : ")
							result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed TransactionCount", nil)
							fmt.Println(result2)
							return ctx.JSON(http.StatusBadRequest, result)
						} else {
							data, err := svc.Service.TransactionRepo.GetHistoryData2(id, tx)
							fmt.Println("data : ", data)
							if err != nil {
								log.Println("\nError GetHistoryListIdTrx4- GetHistoryListIdTrx4", err.Error())
								result = ResponseInsertTrxJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed GetHistoryListIdTrx4", data)
								return ctx.JSON(http.StatusBadRequest, result)
							} else {
								log.Println("\nReponse InsertTransaction- InsertTransaction")
								result = ResponseInsertTrxJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Transaction", data)
								fmt.Println(result)
								return ctx.JSON(http.StatusOK, result)
							}
						}
					}
				}
			}
		}
	})
	return nil
}

// Get Transaction By id_transaction
func (svc transactionService) GetTransactionByTrxId(ctx echo.Context) (err error) {
	var result2 models.Response
	var data models.HistoryTransactionId
	var transaction []models.TxHistoryResume
	fmt.Println("di service id: ", data.Id_transaction)
	if err := BindValidateStruct(ctx, &data); err != nil {
		log.Println("Error Validate Data", err.Error())
		result2 = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result2)
	}

	rows, err := svc.Service.TransactionRepo.GetHistoryListIdTrx(data.Id_transaction)
	for rows.Next() {
		var dataTx models.TxHistoryResume

		rows.Scan(&dataTx.Id_transaction, &dataTx.Jml_item, &dataTx.Id_user)
		for i := 0; i <= dataTx.Jml_item; i++ {
			dataTx.TransactionDetails, _ = svc.Service.TransactionRepo.GetHistoryDetails(dataTx.Id_transaction)
		}
		nama, tgl, price, err := svc.Service.TransactionRepo.GetHistoryData(dataTx.Id_transaction)
		dataTx.Nama_user = nama
		dataTx.Tanggal_transaction = tgl
		dataTx.Total_price = price

		fmt.Println(rows, err)

		transaction = append(transaction, dataTx)
	}

	if err != nil {
		log.Println("\nError InsertProduct- InsertProduct : ", err.Error())
		result2 = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Product", nil)
		fmt.Println(result2)
		return ctx.JSON(http.StatusBadRequest, result2)
	}

	log.Println("\nReponse GetListTransaction ", "Success Get Transaction")
	result2 = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, transaction)
	return ctx.JSON(http.StatusOK, result2)
}

// Get Transaction By id_user
func (svc transactionService) GetHistoryListId(ctx echo.Context) (err error) {
	var result models.Response
	var p models.Transaction
	err = ctx.Bind(&p)
	if err != nil {
		log.Println("\nError GetListTransaction- GetTransactionList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Transaction", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	cart, err := svc.Service.TransactionRepo.GetHistoryListId(p.Id_user)
	log.Println("\nReponse GetListTransaction ", "Success Get Transaction")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, cart)
	return ctx.JSON(http.StatusOK, result)
}
