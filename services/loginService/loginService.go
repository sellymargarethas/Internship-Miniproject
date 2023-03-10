package loginService

import (
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"golang.org/x/crypto/bcrypt"
)

type loginService struct {
	Login services.UsecaseService
}

func NewLoginService(service services.UsecaseService) loginService {
	return loginService{

		Login: service,
	}
}

// Validate Struct
func validateStruct(x interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(x)
	if err != nil {
		return err
	}
	return nil
}

func (svc loginService) LoginUser(ctx echo.Context) error {
	var result models.Response
	//user := models.LoginDataResponse{}
	var credentials models.LoginUser
	//credentials := new(models.LoginUser)
	if err := BindValidateStruct(ctx, &credentials); err != nil {
		log.Println("Error Validate Login", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	resUser, err := svc.Login.LoginRepo.CheckHash(credentials)
	//user2, err := repositories.CheckHash(credentials)
	if err != nil {
		log.Println("Error Validate Login", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	err = bcrypt.CompareHashAndPassword([]byte(resUser), []byte(credentials.Password))
	if err != nil {
		msg := "email atau password salah"
		log.Println("Email atau Password Salah", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.DATA_NOT_FOUND_CODE, msg, nil)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		token, err := GenerateToken(credentials)
		log.Println(err)
		data, err := svc.Login.LoginRepo.ReturnLogin(credentials)
		log.Println("\nReponse Login ", "Success Login", err)
		//result.Message = token
		fmt.Println(data)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, token, data)
		return ctx.JSON(http.StatusOK, result)

	}
}
