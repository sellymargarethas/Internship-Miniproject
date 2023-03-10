package productRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	"fmt"
)

type productRepository struct {
	RepoDB repositories.Repository
}

func NewProductRepository(repoDB repositories.Repository) productRepository {
	return productRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `nama_product, deskripsi_product, qty_product, harga_product`

// Get All Product List
func (ctx productRepository) GetProductList() ([]models.CategoryProduct, error) {
	result := models.CategoryProductArray{}
	query := `
		SELECT id_product, kategori.nama_kategori, ` + defineColumn + `
		FROM product INNER JOIN kategori on product.id_kategori=kategori.id_kategori
		WHERE qty_product > 0
		ORDER BY product.nama_product ASC`

	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		result.CategoryProductArray = append(result.CategoryProductArray, listProduct)
	}

	if err != nil {
		return result.CategoryProductArray, err
	}
	defer rows.Close()

	return result.CategoryProductArray, nil
}

// Get All Product List DESC
func (ctx productRepository) GetProductListDESC() ([]models.CategoryProduct, error) {
	result := models.CategoryProductArray{}
	query := `
		SELECT id_product, kategori.nama_kategori, ` + defineColumn + `
		FROM product INNER JOIN kategori on product.id_kategori=kategori.id_kategori
		WHERE qty_product > 0
		ORDER BY product.nama_product DESC`

	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		result.CategoryProductArray = append(result.CategoryProductArray, listProduct)
	}

	if err != nil {
		return result.CategoryProductArray, err
	}
	defer rows.Close()

	return result.CategoryProductArray, nil
}

// Insert Product
func (ctx productRepository) InsertProduct(product models.Product) (id int, err error) {
	err = ctx.RepoDB.DB.QueryRow(`INSERT INTO product 
	(id_kategori, `+defineColumn+`) 
	VALUES ($1, $2, $3, $4, $5) 
	returning id_product`, product.Id_kategori, product.Nama_product, product.Deskripsi_product, product.Qty_product, product.Harga_product).Scan(&id)
	return id, err
}

// Update Product
func (ctx productRepository) UpdateProduct(product models.Product) (status bool, err error) {
	rows, err := ctx.RepoDB.DB.Query(`UPDATE product SET id_kategori=$1, nama_product=$2, deskripsi_product=$3, qty_product=$4, harga_product=$5 WHERE id_product=$6`, product.Id_kategori, product.Nama_product, product.Deskripsi_product, product.Qty_product, product.Harga_product, product.Id_product)
	fmt.Println(rows, err)
	return true, nil
}

// Delete Product
func (ctx productRepository) DeleteProduct(product models.Product) (status bool, err error) {
	_, err = ctx.RepoDB.DB.Query(`UPDATE product SET qty_product=0
	WHERE id_product=$1`, product.Id_product)
	return true, err
}

// dapetin list product yg di archived
func (ctx productRepository) GetArchivedProduct() (product []models.CategoryProduct, err error) {
	ProductList := models.CategoryProductArray{}
	query := `SELECT product.id_product, kategori.nama_kategori, ` + defineColumn + ` 
	FROM (product INNER JOIN kategori on product.id_kategori=kategori.id_kategori)
	WHERE qty_product=0`
	rows, err := ctx.RepoDB.DB.Query(query)
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		ProductList.CategoryProductArray = append(ProductList.CategoryProductArray, listProduct)
	}
	defer rows.Close()
	return ProductList.CategoryProductArray, err
}

// func insert produk dari archive ke product dari func sebelumnya
func (ctx productRepository) RestoreProduct(p models.Product) (status bool, err error) {
	var id int
	query := `UPDATE product SET qty_product=$1 WHERE id_product=$2 RETURNING id_product`

	err = ctx.RepoDB.DB.QueryRow(query, p.Qty_product, p.Id_product).Scan(&id)
	if err != nil {
		fmt.Println("ada error")
		return false, err
	}
	return true, err
}

// Get Restore Product
func (ctx productRepository) GetRestoreProduct(id int) ([]models.Product, error) {
	var Product []models.Product

	rows, err := ctx.RepoDB.DB.Query(`SELECT product.id_product, product.id_kategori, `+defineColumn+` FROM product WHERE product.id_product=$1`, id)
	fmt.Println(rows, err)

	for rows.Next() {
		var product models.Product
		rows.Scan(&product.Id_product, &product.Id_kategori, &product.Nama_product, &product.Deskripsi_product, &product.Qty_product, &product.Harga_product)
		Product = append(Product, product)
	}
	if err != nil {
		return Product, err
	}
	defer rows.Close()
	return Product, err
}

// Func Check Name
func (ctx productRepository) CheckNameProduct(product string) (status string, err error) {
	var value string
	rows := ctx.RepoDB.DB.QueryRow(`SELECT nama_product FROM product WHERE nama_product=$1`, product).Scan(&value)
	fmt.Println(rows, err)
	return value, nil
}

// Get Product Yang Diinsert
func (ctx productRepository) GetProduct(category string) ([]models.Product, error) {
	var result []models.Product
	query := `
		SELECT id_product, id_kategori, ` + defineColumn + `
		FROM product WHERE nama_product=$1 `

	rows, err := ctx.RepoDB.DB.Query(query, category)
	fmt.Println(rows, err)
	for rows.Next() {
		var listProduct models.Product
		rows.Scan(&listProduct.Id_product, &listProduct.Id_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		result = append(result, listProduct)
	}

	if err != nil {
		return result, err
	}
	defer rows.Close()

	return result, nil
}

// Get All Product List By Price DESC
func (ctx productRepository) GetProductPriceDESC() ([]models.CategoryProduct, error) {
	result := models.CategoryProductArray{}
	query := `
		SELECT id_product, kategori.nama_kategori, ` + defineColumn + `
		FROM product INNER JOIN kategori on product.id_kategori=kategori.id_kategori
		WHERE qty_product > 0
		ORDER BY product.harga_product DESC`

	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listProduct models.CategoryProduct
		rows.Scan(&listProduct.Id_product, &listProduct.Nama_kategori, &listProduct.Nama_product, &listProduct.Deskripsi_product, &listProduct.Qty_product, &listProduct.Harga_product)
		result.CategoryProductArray = append(result.CategoryProductArray, listProduct)
	}

	if err != nil {
		return result.CategoryProductArray, err
	}
	defer rows.Close()

	return result.CategoryProductArray, nil
}

// Get One Product
func (ctx productRepository) GetOneProduct(a int) (product models.CategoryProduct, err error) {
	query := `
		SELECT id_product, kategori.nama_kategori, ` + defineColumn + `
		FROM product
		INNER JOIN kategori on product.id_kategori=kategori.id_kategori
		WHERE product.id_product=$1`

	err = ctx.RepoDB.DB.QueryRow(query, a).Scan(&product.Id_product, &product.Nama_kategori, &product.Nama_product, &product.Deskripsi_product, &product.Qty_product, &product.Harga_product)
	return product, err
}

// Check id_product
func (ctx productRepository) CheckProduct(id int) (status int, err error) {
	var value int
	rows := ctx.RepoDB.DB.QueryRow(`SELECT id_product FROM product WHERE id_product=$1`, id).Scan(&value)
	fmt.Println(rows)
	return value, nil
}
