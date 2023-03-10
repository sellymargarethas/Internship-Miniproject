package app

import (
	"database/sql"

	"Internship-Miniproject/repositories"
	"Internship-Miniproject/repositories/cartRepository"
	"Internship-Miniproject/repositories/categoryRepository"
	"Internship-Miniproject/repositories/loginRepository"
	"Internship-Miniproject/repositories/productRepository"
	"Internship-Miniproject/repositories/searchRepository"
	"Internship-Miniproject/repositories/transactionRepository"
	"Internship-Miniproject/repositories/userRepository"

	"Internship-Miniproject/services"
)

func SetupApp(DB *sql.DB, repo repositories.Repository) services.UsecaseService {
	categoryRepo := categoryRepository.NewCategoryRepository(repo)

	userRepo := userRepository.NewUserRepository(repo)
	loginRepo := loginRepository.NewLoginRepository(repo)

	cartRepo := cartRepository.NewCartRepository(repo)
	productRepo := productRepository.NewProductRepository(repo)
	transactionRepo := transactionRepository.NewTransactionRepository(repo)
	searchRepo := searchRepository.NewSearchRepository(repo)

	usecaseSvc := services.NewUsecaseService(
		DB, categoryRepo, cartRepo, userRepo, loginRepo, productRepo, transactionRepo, searchRepo,
	)

	return usecaseSvc
}
