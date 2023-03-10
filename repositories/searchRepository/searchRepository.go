package searchRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	"fmt"
)

type searchRepository struct {
	RepoDB repositories.Repository
}

func NewSearchRepository(repoDB repositories.Repository) searchRepository {
	return searchRepository{
		RepoDB: repoDB,
	}
}

// Get From Name Product Order By harga_product DESC
func (ctx searchRepository) ProductListAsc() (status []models.CategoryProduct, err error) {
	ProductList := models.CategoryProductArray{}
	query := `SELECT product.id_product, kategori.nama_kategori, product.nama_product, product.deskripsi_product, product.qty_product, product.harga_product 
	FROM (product INNER JOIN kategori on product.id_kategori=kategori.id_kategori) 
	WHERE product.qty_product >0
	ORDER BY harga_product ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		ProductList.CategoryProductArray = append(ProductList.CategoryProductArray, listProduct)
	}
	return ProductList.CategoryProductArray, err
}

// Search From Name Product
func (ctx searchRepository) SearchProduct(a string) (product []models.CategoryProduct, err error) {

	ProductList := models.CategoryProductArray{}
	query := `SELECT product.id_product, kategori.nama_kategori, product.nama_product, product.deskripsi_product, product.qty_product, product.harga_product 
	FROM (product INNER JOIN kategori on product.id_kategori=kategori.id_kategori) 
	WHERE product.nama_product ILIKE $1 AND product.qty_product > 0
	ORDER BY product.nama_product ASC`
	rows, err := ctx.RepoDB.DB.Query(query, "%"+a+"%")
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		ProductList.CategoryProductArray = append(ProductList.CategoryProductArray, listProduct)
	}

	if err != nil {
		return product, err
	}
	defer rows.Close()

	return ProductList.CategoryProductArray, nil
}

// Search From Transaction Date
func (ctx searchRepository) SearchTxDate(a string, id int) (status []models.SearchHistoryDate, err error) {
	date1 := a + " 00:00:00"
	date2 := a + " 23:59:59"
	ProductList := models.SearchHistoryDateArray{}
	query := `SELECT 
	(SELECT to_char(created_on,'DD-MM-YYYY') FROM transaction WHERE transaction.created_on > $1 AND transaction.created_on <$2 AND transaction.id_user=$3),
	transaction.id_transaction,
	users.id_user,
	users.nama_user,
	product.id_product,
	product.nama_product,
	payment.type_payment,
	transaction_details.qty_transaction_details,
	transaction_details.total_transaction_details,
	transaction_count.total_price
	FROM transaction
	INNER JOIN users ON users.id_user=transaction.id_user
	INNER JOIN payment ON payment.id_payment=transaction.id_payment
	INNER JOIN transaction_details ON transaction_details.id_transaction=transaction.id_transaction
	INNER JOIN product ON product.id_product=transaction_details.id_product
	INNER JOIN transaction_count ON transaction_count.id_transaction=transaction.id_transaction 
	WHERE transaction.created_on > $4 AND transaction.created_on <$5 AND transaction.id_user=$6`

	rows, err := ctx.RepoDB.DB.Query(query, date1, date2, id, date1, date2, id)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listTransaction models.SearchHistoryDate
		rows.Scan(&listTransaction.Date, &listTransaction.Id_transaction, &listTransaction.Id_user, &listTransaction.Nama_user, &listTransaction.Id_product, &listTransaction.Nama_product, &listTransaction.Type_payment, &listTransaction.Qty_transaction, &listTransaction.Total_transaction, &listTransaction.Total_price)
		fmt.Println("listTransaction : ", listTransaction)
		ProductList.SearchHistoryDateArray = append(ProductList.SearchHistoryDateArray, listTransaction)
	}

	if err != nil {
		return status, err
	}
	defer rows.Close()

	return ProductList.SearchHistoryDateArray, nil
}

// Get Top Three
func (ctx searchRepository) Topthree() (status []models.Topthree, err error) {
	var ProductList []models.Topthree
	query := `
		SELECT 
		e.id_product,
		e.nama_product,
		(SELECT qty_product FROM product p WHERE e.id_product = p.id_product) AS qty_product,
		(SELECT deskripsi_product FROM product p WHERE e.id_product = p.id_product) AS deskripsi_product,
		(SELECT harga_product FROM product m WHERE e.id_product = m.id_product) AS harga_product,
		(SELECT nama_kategori FROM kategori l WHERE e.id_kategori = l.id_kategori) AS nama_kategori,
		SUM(transaction_details.qty_transaction_details) AS qty_product_terjual
		FROM product AS e
			INNER JOIN transaction_details ON transaction_details.id_product = e.id_product
		WHERE e.qty_product > 0
		GROUP BY e.id_product
		ORDER BY SUM(transaction_details.qty_transaction_details) DESC
		LIMIT 3;`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var listCategoryProduct models.Topthree
		rows.Scan(&listCategoryProduct.Id_product, &listCategoryProduct.Nama_product, &listCategoryProduct.Qty_product, &listCategoryProduct.Deskripsi_product, &listCategoryProduct.Harga_product, &listCategoryProduct.Nama_kategori, &listCategoryProduct.Qty_product_terjual)
		ProductList = append(ProductList, listCategoryProduct)
	}

	if err != nil {
		return status, err
	}
	defer rows.Close()

	return ProductList, err
}

// Search From Category
func (ctx searchRepository) GetCategoryProduct(kategori string) (product []models.CategoryProduct, err error) {

	var CategoryProductList []models.CategoryProduct

	rows, err := ctx.RepoDB.DB.Query(`SELECT product.id_product, product.nama_product, kategori.nama_kategori, product.deskripsi_product, product.qty_product, product.harga_product FROM product INNER JOIN kategori ON kategori.id_kategori=product.id_kategori WHERE nama_kategori ILIKE $1 AND product.qty_product>0`, "%"+kategori+"%")
	fmt.Println(err)

	for rows.Next() {
		var listCategoryProduct models.CategoryProduct
		rows.Scan(&listCategoryProduct.Id_product, &listCategoryProduct.Nama_product, &listCategoryProduct.Nama_kategori, &listCategoryProduct.Deskripsi_product, &listCategoryProduct.Qty_product, &listCategoryProduct.Harga_product)
		CategoryProductList = append(CategoryProductList, listCategoryProduct)
	}
	return CategoryProductList, nil
}
