package routes

import (
	"Internship-Miniproject/config"
	"Internship-Miniproject/constants"
	"Internship-Miniproject/services"
	cartService "Internship-Miniproject/services/cartService"
	categoryService "Internship-Miniproject/services/categoryService"
	loginService "Internship-Miniproject/services/loginService"
	productService "Internship-Miniproject/services/productService"
	searchService "Internship-Miniproject/services/searchService"
	transactionService "Internship-Miniproject/services/transactionService"
	userService "Internship-Miniproject/services/userService"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("supersecretkey"),
})

// RoutesApi
func RoutesApi(e echo.Echo, usecaseSvc services.UsecaseService) {

	public := e.Group("")

	private := e.Group("")
	private.Use(middleware.JWT([]byte(config.GetEnv("JWT_KEY"))))
	private.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: constants.TRUE_VALUE,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Login
	loginGroup := public.Group("/login")
	loginSvc := loginService.NewLoginService(usecaseSvc)
	loginGroup.POST("/login", loginSvc.LoginUser)

	//signup
	userGroup := public.Group("/user")
	userSvc := userService.NewUserService(usecaseSvc)
	userGroup.POST("/signup", userSvc.InsertUser)
	userGroup.GET("/getlogin", userSvc.GetUser)

	// Category
	categoryGroup := public.Group("/category")
	categorySvc := categoryService.NewCategoryService(usecaseSvc)
	categoryGroup.GET("/getcategory", categorySvc.GetListCategory, IsLoggedIn)
	categoryGroup.POST("/insertcategory", categorySvc.InsertCategory, IsLoggedIn)
	categoryGroup.POST("/deletecategory", categorySvc.DeleteCategory, IsLoggedIn)
	categoryGroup.POST("/updatecategory", categorySvc.UpdateCategory, IsLoggedIn)

	// Product
	productGroup := public.Group("/product")
	productSvc := productService.NewProductService(usecaseSvc)
	productGroup.GET("/getproduct", productSvc.GetListProduct, IsLoggedIn)
	productGroup.GET("/getproductdesc", productSvc.GetListProductDESC, IsLoggedIn)
	productGroup.POST("/insertproduct", productSvc.InsertProduct, IsLoggedIn)
	productGroup.POST("/updateproduct", productSvc.UpdateProduct, IsLoggedIn)
	productGroup.POST("/deleteproduct", productSvc.DeleteProduct, IsLoggedIn)
	productGroup.GET("/archivedproduct", productSvc.ViewArchivedProductList, IsLoggedIn)
	productGroup.POST("/restorearchivedproduct", productSvc.RestoreArchive, IsLoggedIn)
	productGroup.GET("/getproductpricedesc", productSvc.GetProductPriceDESC, IsLoggedIn)
	productGroup.POST("/getoneproduct", productSvc.GetOneProduct, IsLoggedIn)

	// Data User
	dataGroup := private.Group("/private")
	dataSvc := userService.NewUserService(usecaseSvc)
	dataGroup.POST("/updateuser", dataSvc.EditUser)
	dataGroup.POST("/deleteuser", dataSvc.DeleteUser)

	// Transaction
	transactionGroup := public.Group("/transaction")
	transactionSvc := transactionService.NewTransactionService(usecaseSvc)
	transactionGroup.GET("/gettransaction", transactionSvc.GetListTransaction, IsLoggedIn)
	transactionGroup.POST("/inserttransaction", transactionSvc.InsertTransaction, IsLoggedIn)
	transactionGroup.POST("/gettx", transactionSvc.GetTransactionByTrxId, IsLoggedIn)
	transactionGroup.POST("/gethistoryid", transactionSvc.GetHistoryListId, IsLoggedIn)

	// Cart
	cartGroup := private.Group("/cart")
	cartSvc := cartService.NewCartService(usecaseSvc)
	cartGroup.GET("/allcart", cartSvc.GetAllCart)
	cartGroup.POST("/individualcart", cartSvc.GetIndividualCart)
	cartGroup.POST("/insertcart", cartSvc.InsertCart)
	cartGroup.POST("/updatecart", cartSvc.UpdateCart)
	cartGroup.POST("/deletecart", cartSvc.DeleteCart)

	// Search
	searchGroup := public.Group("/search")
	searchSvc := searchService.NewSearchService(usecaseSvc)
	searchGroup.GET("/getproductasc", searchSvc.SortProductAsc, IsLoggedIn)
	searchGroup.POST("/historydate", searchSvc.SearchTxHistoryDate, IsLoggedIn)
	searchGroup.GET("/top3", searchSvc.Top3, IsLoggedIn)
	searchGroup.POST("/searchcategory", searchSvc.SearchCategoryProduct)
	searchGroup.POST("/searchproduct", searchSvc.SearchProduct)
}
