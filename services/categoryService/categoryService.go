package categoryService

import (
	"Internship-Miniproject/constants"
	"Internship-Miniproject/models"
	"Internship-Miniproject/services"
	. "Internship-Miniproject/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type categoryService struct {
	Service services.UsecaseService
}

func NewCategoryService(service services.UsecaseService) categoryService {
	return categoryService{
		Service: service,
	}
}

// Validate Struct
func validateStructCategory(x interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(x)
	if err != nil {
		return err
	}
	return nil
}

func (svc categoryService) GetListCategory(ctx echo.Context) error {
	var result models.Response
	category := models.Category{}
	resCategory, err := svc.Service.CategoryRepo.GetCategoryList(category)
	if err != nil {
		log.Println("\nError GetListCategory- GetCategoryList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Category", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListCategory ", "Success Get Category")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resCategory)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Category
func (svc categoryService) InsertCategory(ctx echo.Context) error {
	var p models.Category
	var result models.Response

	err := ctx.Bind(&p)
	err = validateStructCategory(p)
	if err != nil {
		fmt.Println(err)
		log.Println("\nError InsertCategory- InsertCategory : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	check, err := svc.Service.CategoryRepo.CheckName(p.Nama_kategori)

	if check == p.Nama_kategori {
		log.Println("\nError InsertCategory- InsertCategory")
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Nama Kategori Sudah Digunakan", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		id, err := svc.Service.CategoryRepo.InsertCategory(p)
		id2, err := svc.Service.CategoryRepo.GetCategoryByName(p.Nama_kategori)
		fmt.Println("id : ", id, err)
		fmt.Println("id2 : ", id2, err)
		if err != nil {
			fmt.Println(err.Error(), id)
		}

		if err != nil {
			log.Println("\nError InsertCategory- InsertCategory : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Category", nil)
			fmt.Println(result)
			return ctx.JSON(http.StatusBadRequest, result)
		} else {
			log.Println("\nReponse InsertCategory- InsertCategory")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Category", id2)
			fmt.Println(result)
			return ctx.JSON(http.StatusOK, result)
		}
	}
}

// Delete Category
func (svc categoryService) DeleteCategory(ctx echo.Context) error {
	var result models.Response
	var p models.Category

	err := ctx.Bind(&p)
	category, err := svc.Service.CategoryRepo.DeleteCategory(p)
	fmt.Println(category)
	if err != nil {
		log.Println("\nError DeleteCategory- DeleteCategory : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Category", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse DeleteCategory- DeleteCategory")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Category", p)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}
}

// Update Category
func (svc categoryService) UpdateCategory(ctx echo.Context) error {
	var result models.Response
	var p models.Category

	err := ctx.Bind(&p)
	err = validateStructCategory(p)
	if err != nil {
		fmt.Println(err)
		log.Println("\nError UpdateCategory- UpdateCategory : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	check, err := svc.Service.CategoryRepo.CheckName(p.Nama_kategori)

	if check == p.Nama_kategori {
		log.Println("\nError UpdateCategory- UpdateCategory")
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Nama Kategori Sudah Digunakan", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		category, err := svc.Service.CategoryRepo.UpdateCategory(p)
		if err != nil {
			fmt.Println(err.Error(), category)
		} else {
			if err != nil {
				log.Println("\nError UpdateCategory- UpdateCategory : ", err.Error())
				result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Category", nil)
				fmt.Println(result)
				return ctx.JSON(http.StatusBadRequest, result)
			} else {
				log.Println("\nReponse UpdateCategory- UpdateCategory")
				result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Category", p)
				fmt.Println(result)
				return ctx.JSON(http.StatusOK, result)
			}
		}
		return ctx.JSON(http.StatusCreated, result)
	}
}
