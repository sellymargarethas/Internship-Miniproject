package cartRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	"database/sql"
	"fmt"
)

type cartRepository struct {
	RepoDB repositories.Repository
}

func NewCartRepository(repoDB repositories.Repository) cartRepository {
	return cartRepository{
		RepoDB: repoDB,
	}
}

func (ctx cartRepository) GetAllCart() (status []models.Cart, err error) {
	var carts []models.Cart

	rows, err := ctx.RepoDB.DB.Query(`SELECT cart.id_user, cart.id_product, cart.qty_cart, cart.total_cart FROM cart INNER JOIN product on product.id_product=cart.id_product
	WHERE product.qty_product>0`)

	for rows.Next() {
		var cart models.Cart
		rows.Scan(&cart.Id_user, &cart.Id_product, &cart.Qty_cart, &cart.Total_cart)
		carts = append(carts, cart)
	}
	fmt.Println(rows)
	if err != nil {
		return status, err
	}

	return carts, nil
}

func (ctx cartRepository) GetIndividualCart(id_user int) (cart []models.ViewCart, err error) {
	var c models.ViewCart
	query := `SELECT product.id_product, product.nama_product, cart.qty_cart, product.harga_product, cart.total_cart 
	FROM cart INNER JOIN product on product.id_product=cart.id_product
	WHERE product.qty_product>0 AND cart.id_user=$1`
	rows, err := ctx.RepoDB.DB.Query(query, id_user)

	for rows.Next() {
		rows.Scan(&c.Id_product, &c.Nama_product, &c.Qty_cart, &c.Harga_satuan, &c.Total_cart)
		cart = append(cart, c)
	}
	return cart, err

}

func (ctx cartRepository) CheckProductQty(transaction int) (status int, err error) {

	query := `SELECT qty_product FROM product WHERE id_product=$1`
	rows := ctx.RepoDB.DB.QueryRow(query, transaction).Scan(&status)
	fmt.Println(rows)
	return status, nil
}

func (ctx cartRepository) InsertCart(cart models.Cart) (status bool, err error) {

	scan := ctx.RepoDB.DB.QueryRow(`SELECT id_user, id_product FROM cart WHERE id_user=$1 and id_product=$2`, cart.Id_user, cart.Id_product).Scan(&cart.Id_user, &cart.Id_product)
	if scan != nil {
		if scan == sql.ErrNoRows {
			fmt.Println("Berhasil Insert Data Cart")
			items, err := ctx.RepoDB.DB.Query("INSERT INTO cart (id_user, id_product, qty_cart, total_cart) VALUES ($1, $2, $3, (SELECT product.harga_product FROM product LEFT JOIN cart ON product.id_product=cart.id_product WHERE $4=product.id_product)*$5)", cart.Id_user, cart.Id_product, cart.Qty_cart, cart.Id_product, cart.Qty_cart)
			fmt.Println(items, err)
			return true, err
		}
	} else {
		fmt.Println("Berhasil Update Data Cart")
		items, err := ctx.RepoDB.DB.Query("UPDATE cart SET qty_cart=qty_cart + $1, total_cart=((SELECT product.harga_product FROM product LEFT JOIN cart ON product.id_product=cart.id_product WHERE $2=product.id_product)*(qty_cart+$3)) WHERE id_user=$4 AND id_product=$5", cart.Qty_cart, cart.Id_product, cart.Qty_cart, cart.Id_user, cart.Id_product)
		fmt.Println(items, err)
		return true, err
	}

	return true, err
}

// Update cart
func (ctx cartRepository) UpdateCart(cart models.Cart) (status bool, err error) {

	rows, err := ctx.RepoDB.DB.Query(`UPDATE cart SET 
	id_product=$1, qty_cart=$2, 
	total_cart=((SELECT product.harga_product FROM product LEFT JOIN cart ON product.id_product=cart.id_product WHERE $3=product.id_product)*$4) 
	WHERE id_user=$5 AND id_product=$6
	`, cart.Id_product, cart.Qty_cart, cart.Id_product, cart.Qty_cart, cart.Id_user, cart.Id_product)
	fmt.Println(rows)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Delete All cart by id_product
func (ctx cartRepository) DeleteAllCart(id_product int) (status bool, err error) {

	rows, err := ctx.RepoDB.DB.Query("DELETE FROM cart WHERE id_product=$1", id_product)
	fmt.Println(rows, err)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Delete cart
func (ctx cartRepository) DeleteCart(id_user int, id_product int) (status bool, err error) {

	rows, err := ctx.RepoDB.DB.Query("DELETE FROM cart WHERE id_user=$1 AND id_product=$2", id_user, id_product)
	fmt.Println(rows, err)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ctx cartRepository) DeleteMultipleCart(id_user, id_transaction int) (status bool, err error) {

	rows, err := ctx.RepoDB.DB.Query(`DELETE FROM cart where id_user=$1 AND id_product in
	(SELECT transaction_details.id_product FROM transaction_details INNER JOIN transaction on transaction_details.id_transaction=transaction.id_transaction WHERE transaction_details.id_transaction=$2)`, id_user, id_transaction)
	fmt.Println(rows)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Func total quantity cart
func (ctx cartRepository) TotalQtyCart(cart models.Cart) (status int, err error) {
	var value int
	rows := ctx.RepoDB.DB.QueryRow(`SELECT (qty_cart +$1) FROM cart WHERE id_user=$2 AND id_product=$3`, cart.Qty_cart, cart.Id_user, cart.Id_product).Scan(&value)
	fmt.Println(rows)
	return value, err
}
