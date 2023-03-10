package productService

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

type productService struct {
	Service services.UsecaseService
}

func NewProductService(service services.UsecaseService) productService {
	return productService{
		Service: service,
	}
}

// Validate Struct
func validateStructProduct(x interface{}) error {
	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(x)
	if err != nil {
		return err
	}
	return nil
}

// Get All Product List
func (svc productService) GetListProduct(ctx echo.Context) error {
	var result models.Response
	product := models.CategoryProduct{}
	fmt.Println(product)
	resProduct, err := svc.Service.ProductRepo.GetProductList()
	if err != nil {
		log.Println("\nError GetListProduct- GetProductList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProduct ", "Success Get Product")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resProduct)
	return ctx.JSON(http.StatusOK, result)
}

// Insert Product
func (svc productService) InsertProduct(ctx echo.Context) error {
	var p models.Product
	var result models.Response

	err := ctx.Bind(&p)
	err = validateStructProduct(p)
	if err != nil {
		fmt.Println(err)
		log.Println("\nError InsertProduct- InsertProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	check, err := svc.Service.ProductRepo.CheckNameProduct(p.Nama_product)

	if check == p.Nama_product {
		log.Println("\nError InsertProduct- InsertProduct")
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Nama Product Sudah Digunakan", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		id, err := svc.Service.ProductRepo.InsertProduct(p)
		id2, err := svc.Service.ProductRepo.GetProduct(p.Nama_product)
		fmt.Println(id, err)
		if err != nil {
			fmt.Println(err.Error(), id)
		}

		if err != nil {
			log.Println("\nError InsertProduct- InsertProduct : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Insert Product", nil)
			fmt.Println(result)
			return ctx.JSON(http.StatusBadRequest, result)
		} else {
			log.Println("\nReponse InsertProduct- InsertProduct")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Insert Product", id2)
			fmt.Println(result)
			return ctx.JSON(http.StatusOK, result)
		}
	}
}

// Update Product
func (svc productService) UpdateProduct(ctx echo.Context) error {
	var result models.Response
	var p models.Product

	err := ctx.Bind(&p)
	err = validateStructProduct(p)

	if err != nil {
		log.Println("\nError UpdateProduct- UpdateProduct")
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	product, err := svc.Service.ProductRepo.UpdateProduct(p)
	if err != nil {
		fmt.Println(err.Error(), product)
	} else {
		if err != nil {
			log.Println("\nError UpdateProduct- UpdateProduct : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Update Product", nil)
			fmt.Println(result)
			return ctx.JSON(http.StatusBadRequest, result)
		} else {
			log.Println("\nReponse UpdateProduct- UpdateProduct")
			result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Update Product", p)
			fmt.Println(result)
			return ctx.JSON(http.StatusOK, result)
		}
	}
	return ctx.JSON(http.StatusCreated, result)
}

// Delete Product
func (svc productService) DeleteProduct(ctx echo.Context) error {
	var result models.Response
	var p models.Product

	err := ctx.Bind(&p)
	product, err := svc.Service.ProductRepo.DeleteProduct(p)
	fmt.Println(product)
	if err != nil {
		log.Println("\nError DeleteProduct- DeleteProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Delete Product", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse DeleteProduct- DeleteProduct")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Delete Product", p)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}
}

// Mengembalikan produk dari tabel archive ke product (auto delete yg di archive)
func (svc productService) RestoreArchive(ctx echo.Context) error {
	var result models.Response
	var p models.Product
	err := ctx.Bind(&p)
	if err != nil {
		fmt.Println(err)
		log.Println("\nError RestoreProduct- RestoreProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	prod, err := svc.Service.ProductRepo.RestoreProduct(p)
	prod2, err := svc.Service.ProductRepo.GetRestoreProduct(p.Id_product)
	fmt.Println(prod)
	fmt.Println(prod2)
	if err != nil {
		log.Println("\nError RestoreProduct- RestoreProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.FAILED_CODE, "Failed Restore Product", nil)
		fmt.Println(result)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		log.Println("\nReponse RestoreProduct- RestoreProduct")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, "Success Restore Product", prod2)
		fmt.Println(result)
		return ctx.JSON(http.StatusOK, result)
	}
}

// Get All archived product list
func (svc productService) ViewArchivedProductList(ctx echo.Context) error {
	var result models.Response
	product := models.CategoryProduct{}
	fmt.Println(product)
	resProduct, err := svc.Service.ProductRepo.GetArchivedProduct()
	if err != nil {
		log.Println("\nError GetListProduct- GetProductList : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProduct ", "Success Get Product")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resProduct)
	return ctx.JSON(http.StatusOK, result)
}

// Get Product List DESC
func (svc productService) GetListProductDESC(ctx echo.Context) error {
	var result models.Response
	product := models.CategoryProduct{}
	fmt.Println(product)
	resProduct, err := svc.Service.ProductRepo.GetProductListDESC()
	if err != nil {
		log.Println("\nError GetListProductDESC- GetProductListDESC : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProductDESC", "Success Get Product DESC")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resProduct)
	return ctx.JSON(http.StatusOK, result)
}

// Get Product List By Price DESC
func (svc productService) GetProductPriceDESC(ctx echo.Context) error {
	var result models.Response
	product := models.CategoryProduct{}
	fmt.Println(product)
	resProduct, err := svc.Service.ProductRepo.GetProductPriceDESC()
	if err != nil {
		log.Println("\nError GetListProductPriceDESC- GetProductListPriceDESC : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}

	log.Println("\nReponse GetListProductPriceDESC ", "Success Get Product By Price DESC")
	result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resProduct)
	return ctx.JSON(http.StatusOK, result)
}

// Get One Product By id_product
func (svc productService) GetOneProduct(ctx echo.Context) (err error) {
	var result models.Response
	var p models.Product
	err = ctx.Bind(&p)
	if err != nil {
		log.Println("\nError RestoreProduct- RestoreProduct : ", err.Error())
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Validate Data", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	}
	check, err := svc.Service.ProductRepo.CheckProduct(p.Id_product)

	if check != p.Id_product {
		log.Println("\nError GetOneProduct- GetOneProduct")
		result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "id_product tidak ditemukan", nil)
		return ctx.JSON(http.StatusBadRequest, result)
	} else {
		resProduct, err := svc.Service.ProductRepo.GetOneProduct(p.Id_product)
		if err != nil {
			log.Println("\nError GetOneProduct- GetOneProduct : ", err.Error())
			result = ResponseJSON(constants.FALSE_VALUE, constants.VALIDATE_ERROR_CODE, "Failed Get Product", nil)
			return ctx.JSON(http.StatusBadRequest, result)
		}

		log.Println("\nReponse GetOneProduct ", "Success Get One Product")
		result = ResponseJSON(constants.TRUE_VALUE, constants.SUCCESS_CODE, constants.EMPTY_VALUE, resProduct)
		return ctx.JSON(http.StatusOK, result)
	}
}
