package categoryRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	"fmt"
)

type categoryRepository struct {
	RepoDB repositories.Repository
}

func NewCategoryRepository(repoDB repositories.Repository) categoryRepository {
	return categoryRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `nama_kategori`

func (ctx categoryRepository) GetCategoryList(category models.Category) ([]models.Category, error) {
	var result []models.Category
	query := `
		SELECT id_kategori, ` + defineColumn + `
		FROM kategori ORDER BY id_kategori `

	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println(rows, err)
	for rows.Next() {
		var listCategory models.Category
		rows.Scan(&listCategory.Id_kategori, &listCategory.Nama_kategori)
		result = append(result, listCategory)
	}

	if err != nil {
		return result, err
	}
	defer rows.Close()

	return result, nil
}

// Insert Category
func (ctx categoryRepository) InsertCategory(category models.Category) (id int, err error) {
	err = ctx.RepoDB.DB.QueryRow(`INSERT INTO kategori (`+defineColumn+`) VALUES ($1) returning id_kategori`, category.Nama_kategori).Scan(&id)
	return id, nil
}

// Get Category Yang Diinsert
func (ctx categoryRepository) GetCategoryByName(category string) ([]models.Category, error) {
	var result []models.Category
	query := `
		SELECT id_kategori, ` + defineColumn + `
		FROM kategori WHERE nama_kategori=$1 `

	rows, err := ctx.RepoDB.DB.Query(query, category)
	fmt.Println(rows, err)
	for rows.Next() {
		var listCategory models.Category
		rows.Scan(&listCategory.Id_kategori, &listCategory.Nama_kategori)
		result = append(result, listCategory)
	}

	if err != nil {
		return result, err
	}
	defer rows.Close()

	return result, nil
}

// Func Check Name
func (ctx categoryRepository) CheckName(category string) (status string, err error) {
	var value string
	rows := ctx.RepoDB.DB.QueryRow(`SELECT `+defineColumn+` FROM kategori WHERE nama_kategori=$1`, category).Scan(&value)
	fmt.Println(rows, err)
	return value, nil
}

// Delete Category
func (ctx categoryRepository) DeleteCategory(category models.Category) (status bool, err error) {
	_, err = ctx.RepoDB.DB.Query("DELETE FROM kategori WHERE id_kategori=$1", category.Id_kategori)
	return true, err
}

// Update Category
func (ctx categoryRepository) UpdateCategory(category models.Category) (status bool, err error) {
	rows, err := ctx.RepoDB.DB.Query(`UPDATE kategori SET `+defineColumn+`=$1 WHERE id_kategori=$2`, category.Nama_kategori, category.Id_kategori)
	fmt.Println(rows, err)
	return true, nil
}
