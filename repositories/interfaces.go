package repositories

import (
	"Internship-Miniproject/models"
	"database/sql"
)

type CategoryRepository interface {
	GetCategoryList(category models.Category) ([]models.Category, error)
	InsertCategory(category models.Category) (id int, err error)
	GetCategoryByName(category string) ([]models.Category, error)
	DeleteCategory(category models.Category) (status bool, err error)
	UpdateCategory(category models.Category) (status bool, err error)
	CheckName(category string) (status string, err error)
}

type LoginRepository interface {
	CheckHash(user models.LoginUser) (status string, err error)
	ReturnLogin(user models.LoginUser) (status models.LoginDataResponse, err error)
}

type ProductRepository interface {
	GetProductList() ([]models.CategoryProduct, error)
	GetProductListDESC() ([]models.CategoryProduct, error)
	InsertProduct(product models.Product) (id int, err error)
	UpdateProduct(product models.Product) (status bool, err error)
	DeleteProduct(product models.Product) (status bool, err error)
	GetArchivedProduct() (product []models.CategoryProduct, err error)
	RestoreProduct(p models.Product) (status bool, err error)
	GetRestoreProduct(id int) (product []models.Product, err error)
	CheckNameProduct(product string) (status string, err error)
	GetProduct(category string) ([]models.Product, error)
	GetProductPriceDESC() ([]models.CategoryProduct, error)
	GetOneProduct(a int) (product models.CategoryProduct, err error)
	CheckProduct(id int) (status int, err error)
}

type UserRepository interface {
	InsertUser(user models.User, tx *sql.Tx) (id int, err error)
	InsertDataUser(id int, user models.DataUser, tx *sql.Tx) (id_user int, err error)
	GetUser(user models.LoginDataResponse) ([]models.LoginDataResponse, error)
	UpdateUser(user models.User, tx *sql.Tx) (status bool, err error)
	UpdateDataUser(user models.DataUser, id_user int, tx *sql.Tx) (id int, err error)
	GetOneUser(i int) (user models.UserProfile)
	DeleteUser(id_user int, tx *sql.Tx) (status bool, err error)
	DeleteUserData(id_user int, tx *sql.Tx) (status bool, err error)
}

type TransactionRepository interface {
	GetHistoryDetails(id int) (status []models.TxHistoryDetails, err error)
	GetHistory() (err error, data *sql.Rows)
	GetHistoryData(id int) (nama string, tgl string, price int, err error)
	PostNewTransaction(trx models.Transaction, tx *sql.Tx) (id_trx int, err error)
	CobaInsertMultipleItemsTransaction(trx []models.TransactionDetails, id int, tx *sql.Tx) (status bool, err error)
	CountTotalItemTx(id int, tx *sql.Tx) (status bool, err error)
	AutoUpdateInventory(trx int, tx *sql.Tx) (status bool, err error)
	TransactionCount(trx int, tx *sql.Tx) (price int, err error)
	GetHistoryListIdTrx2(trx int, tx *sql.Tx) (nama_user string, err error)
	GetHistoryListIdTrx3(trx int, tx *sql.Tx) (tanggal_transaksi string, err error)
	GetHistoryListIdTrx4(trx int, tx *sql.Tx) (price int, err error)
	GetHistoryListIdTrx(trx int) (status *sql.Rows, err error)
	GetHistoryListId(trx int) (status []models.HistoryTransactionId, err error)
	GetHistoryData2(id int, tx *sql.Tx) (data models.HistoryTransactionId, err error)
}

type CartRepository interface {
	GetAllCart() (status []models.Cart, err error)
	GetIndividualCart(id_user int) (cart []models.ViewCart, err error)
	CheckProductQty(transaction int) (status int, err error)
	InsertCart(cart models.Cart) (status bool, err error)
	UpdateCart(cart models.Cart) (status bool, err error)
	DeleteAllCart(id_product int) (status bool, err error)
	DeleteCart(id_user int, id_product int) (status bool, err error)
	DeleteMultipleCart(id_user, id_transaction int) (status bool, err error)
	TotalQtyCart(cart models.Cart) (status int, err error)
}

type SearchRepository interface {
	ProductListAsc() (status []models.CategoryProduct, err error)
	SearchTxDate(a string, id int) (status []models.SearchHistoryDate, err error)
	Topthree() (status []models.Topthree, err error)
	SearchProduct(a string) (product []models.CategoryProduct, err error)
	GetCategoryProduct(kategori string) (product []models.CategoryProduct, err error)
}
