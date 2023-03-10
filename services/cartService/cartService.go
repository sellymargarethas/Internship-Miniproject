package cartservice

import (
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type cartService struct {
	Service services.UsecaseService
}

func NewCartService(service services.UsecaseService) cartService {
	return cartService{
		Service: service,
	}
}

func (svc cartService) GetAllCart(ctx echo.Context) error {
	var result models.Response

	cart, err := svc.Service.CartRepo.GetAllCart()
	if err != nil {
		log.Println("\nError Get All Cart : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	log.Println("\nReponse Get All Cart ", "Success Get Data")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, cart)
	return ctx.JSON(http.StatusOK, result)

}

func (svc cartService) GetIndividualCart(ctx echo.Context) error {
	var result models.Response
	var cart_data models.Cart

	cart, err := svc.Service.CartRepo.GetIndividualCart(cart_data.Id_user)
	if err != nil {
		log.Println("\nError Get Cart Individual : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	if cart != nil {
		log.Println("\nReponse Get Individual Cart ", "Success Get Data")
		cartQty, _ := svc.Service.CartRepo.TotalQtyCart(cart_data)
		var cart_value string
		cart_value = "Total Cart: " + strconv.Itoa(cartQty)
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, cart_value, cart)
		return ctx.JSON(http.StatusOK, result)
	}
	log.Println("\nError Get Cart Individual : ", err.Error())
	result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Empty cart", nil)
	return ctx.JSON(http.StatusBadRequest, result)

}

func (svc cartService) InsertCart(ctx echo.Context) error {
	var result models.Response
	var cart_data models.Cart

	if err := BindValidateStruct(ctx, &cart_data); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	newcart, err := svc.Service.CartRepo.InsertCart(cart_data)
	if err != nil {
		log.Println("\nError InsertCart- InsertCart : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Cart", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse InsertCart- InsertCart")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Cart", newcart)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}
}

func (svc cartService) UpdateCart(ctx echo.Context) error {
	var result models.Response
	var cart_data models.Cart

	if err := BindValidateStruct(ctx, &cart_data); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	newcart, err := svc.Service.CartRepo.UpdateCart(cart_data)
	if err != nil {
		log.Println("\nError Update Cart Failed : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Cart", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse Update Cart Success")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Cart", newcart)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}
}

func (svc cartService) DeleteCart(ctx echo.Context) error {
	var result models.Response
	var cart_data models.Cart
	if err := BindValidateStruct(ctx, &cart_data); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	newcart, err := svc.Service.CartRepo.DeleteCart(cart_data.Id_user, cart_data.Id_product)
	if err != nil {
		log.Println("\nError Delete Cart Failed : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Cart", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse Delete Cart Success")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Cart", newcart)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}

}
