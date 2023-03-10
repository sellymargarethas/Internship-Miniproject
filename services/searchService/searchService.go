package searchService

import (
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type searchService struct {
	Service services.UsecaseService
}

func NewSearchService(service services.UsecaseService) searchService {
	return searchService{
		Service: service,
	}
}

// Get All Transaction By Price ASC
func (svc searchService) SortProductAsc(ctx echo.Context) error {
	var result models.Response

	product, err := svc.Service.SearchRepo.ProductListAsc()

	if err != nil {
		log.Println("\nError SortProductByPriceASC- SortProductByPriceASC : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Sort Product ASC", "Failed Sort Product ASC")
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse SortProductByPriceASC ", "Success Get Product By Price ASC")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Sukses Mendapatkan Product", product)
		return ctx.JSON(http.StatusOK, result)
	}
}

// Get All Transaction By Date
func (svc searchService) SearchTxHistoryDate(ctx echo.Context) (err error) {
	var result models.Response
	var p models.SearchHistoryDate

	err = ctx.Bind(&p)

	product, err := svc.Service.SearchRepo.SearchTxDate(p.Date, p.Id_user)

	if err != nil {
		log.Println("\nError SortProductByDate- SortProductByPriceASC")
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Product By Date", "Failed Get Product By Date")
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		if product != nil {
			log.Println("\nReponse SortProductByDate ", "Success Get Product By Date")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Sukses Mendapatkan Product", product)
			return ctx.JSON(http.StatusOK, result)
		} else {
			log.Println("\nError SortProductByDate SortProductByDate")
			result = ResponseJSON(constants.FALSE_VALUE, constants.EMPTY_VALUE, "Tidak Ada Transaksi Di Tanggal Ini", p.Date)
			fmt.Println(result)
			return ctx.JSON(http.StatusBadRequest, result)
		}
	}
}

// Top Three Best Product
func (svc searchService) Top3(ctx echo.Context) (err error) {
	var result models.Response
	product, err := svc.Service.SearchRepo.Topthree()

	if err != nil {
		log.Println("\nError Get Top Three- Get Top Three")
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Get Top Three", "Failed Get Top Three")
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse Get Top Three ", "Success Get Top Three")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Sukses Mendapatkan Top Three", product)
		return ctx.JSON(http.StatusOK, result)
	}
}

// Search Product By Name
func (svc searchService) SearchProduct(ctx echo.Context) error {
	var result models.Response
	var search models.CategoryProduct

	if err := BindValidateStruct(ctx, &search); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	products, err := svc.Service.SearchRepo.SearchProduct(search.Nama_product)
	if err != nil {
		log.Println("\nError GetListProduct- GetProductList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProduct ", "Success Get Product")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, products)
	return ctx.JSON(http.StatusOK, result)

}

// Search Product By Category
func (svc searchService) SearchCategoryProduct(ctx echo.Context) error {
	var result models.Response
	var search models.CategoryProduct

	if err := BindValidateStruct(ctx, &search); err != nil {
		log.Println("Error Validate Data", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, err.Error(), nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	products, err := svc.Service.SearchRepo.GetCategoryProduct(search.Nama_kategori)
	if err != nil {
		log.Println("\nError GetListProduct- GetProductList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProduct ", "Success Get Product")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, products)
	return ctx.JSON(http.StatusOK, result)
}
