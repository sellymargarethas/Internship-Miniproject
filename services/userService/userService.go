package userService

import (
	"database/sql"

	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

type userService struct {
	Service services.UsecaseService
}

func NewUserService(service services.UsecaseService) userService {
	return userService{
		Service: service,
	}
}

func (svc userService) GetUser(ctx echo.Context) error {
	var result models.Response
	user := models.LoginDataResponse{}
	resUser, err := svc.Service.UserRepo.GetUser(user)
	if err != nil {
		log.Println("\nError GetListUser- GetUser : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("\nReponse GetListCategory ", "Success Get Data")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resUser)
	return ctx.JSON(http.StatusOK, result)
}

func (svc userService) InsertUser(ctx echo.Context) error {
	var result models.Response
	var datauser models.User
	//var detailuser models.DataUser

	if err := BindValidateStruct(ctx, &datauser); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(datauser.Password_user), bcrypt.MinCost)
	datauser.Password_user = string(hash)
	fmt.Println(datauser.Password_user)
	err = DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {

		insertdata, err := svc.Service.UserRepo.InsertUser(datauser, tx)
		fmt.Println("cek di insertdata", datauser, insertdata)
		if err != nil {
			log.Println("Error Validate Data", insertdata)
			return err
		}

		insertdatauser, err := svc.Service.UserRepo.InsertDataUser(insertdata, datauser.DataUser, tx)
		if err != nil {
			log.Println("Error Validate Data User", insertdatauser)
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Error Validate Data", "Insert User", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Berhasil insert data user ")
	//data:= svc.Service.UserRepo.GetOneUser(insertdata)
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Berhasil Sign Up", nil)
	return ctx.JSON(http.StatusOK, result)

}

func (svc userService) EditUser(ctx echo.Context) error {
	var result models.Response
	var datauser models.User

	if err := BindValidateStruct(ctx, &datauser); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	err := DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		hash, err := bcrypt.GenerateFromPassword([]byte(datauser.Password_user), bcrypt.MinCost)
		datauser.Password_user = string(hash)
		updatedata, err := svc.Service.UserRepo.UpdateUser(datauser, tx)
		if err != nil {
			log.Println("Error Validate Data", updatedata)
			return err
		}
		updatedatauser, err := svc.Service.UserRepo.UpdateDataUser(datauser.DataUser, datauser.Id_user, tx)
		if err != nil {
			log.Println("Error Validate Data", updatedatauser)
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Error Validate Data", "Update User", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	updatedUser := svc.Service.UserRepo.GetOneUser(datauser.Id_user)
	log.Println("Berhasil update data user ")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Berhasil Update Data User", updatedUser)
	return ctx.JSON(http.StatusOK, result)
}

func (svc userService) DeleteUser(ctx echo.Context) error {
	var result models.Response
	var datauser models.LoginDataResponse

	if err := BindValidateStruct(ctx, &datauser); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	err := DBTransaction(svc.Service.RepoDB, func(tx *sql.Tx) error {
		deletedata, err := svc.Service.UserRepo.DeleteUserData(datauser.Id_user, tx)
		if err != nil {
			log.Println("Error Validate Data", deletedata)
			return err
		}
		deletedatauser, err := svc.Service.UserRepo.DeleteUser(datauser.Id_user, tx)
		if err != nil {
			log.Println("Error Validate Data", deletedatauser)
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Error Validate Data", "Delete User", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("Berhasil update data user ")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Berhasil Hapus Data User", nil)
	return ctx.JSON(http.StatusOK, result)
}
