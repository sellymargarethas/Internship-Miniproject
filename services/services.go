package services

import (
	"Internship-Miniproject/repositories"
	"database/sql"
)

type UsecaseService struct {
	RepoDB          *sql.DB
	CategoryRepo    repositories.CategoryRepository
	UserRepo        repositories.UserRepository
	LoginRepo       repositories.LoginRepository
	ProductRepo     repositories.ProductRepository
	TransactionRepo repositories.TransactionRepository
	CartRepo        repositories.CartRepository
	SearchRepo      repositories.SearchRepository
}

func NewUsecaseService(
	repoDB *sql.DB,
	CategoryRepo repositories.CategoryRepository,
	CartRepo repositories.CartRepository,
	UserRepo repositories.UserRepository,
	LoginRepo repositories.LoginRepository,
	ProductRepo repositories.ProductRepository,
	TransactionRepo repositories.TransactionRepository,
	SearchRepo repositories.SearchRepository,
) UsecaseService {
	return UsecaseService{
		RepoDB:          repoDB,
		CategoryRepo:    CategoryRepo,
		UserRepo:        UserRepo,
		LoginRepo:       LoginRepo,
		ProductRepo:     ProductRepo,
		TransactionRepo: TransactionRepo,
		CartRepo:        CartRepo,
		SearchRepo:      SearchRepo,
	}
}
