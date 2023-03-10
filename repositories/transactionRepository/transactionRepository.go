package transactionRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	. "Internship-Miniproject/utils"
	"database/sql"
	"fmt"
	"time"
)

type transactionRepository struct {
	RepoDB repositories.Repository
}

func NewTransactionRepository(repoDB repositories.Repository) transactionRepository {
	return transactionRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `id_transaction`

// Get All Transaction
func (ctx transactionRepository) GetHistoryDetails(id int) (status []models.TxHistoryDetails, err error) {
	var Details []models.TxHistoryDetails
	query := `SELECT transaction_details.id_product, product.nama_product, transaction_details.qty_transaction_details, transaction_details.total_transaction_details
	FROM transaction_details
	INNER JOIN product on transaction_details.id_product=product.id_product
	WHERE transaction_details.id_transaction=$1`

	rows, err := ctx.RepoDB.DB.Query(query, id)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listProducts models.TxHistoryDetails
		rows.Scan(&listProducts.Id_product, &listProducts.Nama_product, &listProducts.Qty_transaction, &listProducts.Total_transaction)
		Details = append(Details, listProducts)
	}

	if err != nil {
		return status, err
	}
	defer rows.Close()

	return Details, nil
}

func (ctx transactionRepository) GetHistory() (err error, data *sql.Rows) {
	query := `(SELECT transaction.id_transaction, (SELECT COUNT(transaction_details.id_product)
	GROUP BY transaction.id_transaction),
	transaction.id_user
	FROM transaction
	INNER JOIN transaction_details on transaction.id_transaction=transaction_details.id_transaction
	INNER JOIN users on transaction.id_user=users.id_user
	Group by transaction.id_transaction)`
	rows, err := ctx.RepoDB.DB.Query(query)

	return err, rows
}

func (ctx transactionRepository) GetHistoryData(id int) (nama string, tgl string, price int, err error) {
	query := `SELECT distinct
		users.nama_user,
		to_char(transaction.created_on,'DD-MM-YYYY'),
		transaction_count.total_price
		FROM transaction
		INNER JOIN users ON users.id_user=transaction.id_user
		INNER JOIN transaction_details ON transaction_details.id_transaction=transaction.id_transaction
		INNER JOIN transaction_count ON transaction_count.id_transaction=transaction.id_transaction
		where transaction.id_transaction=$1`
	rows := ctx.RepoDB.DB.QueryRow(query, id).Scan(&nama, &tgl, &price)
	fmt.Println(rows)
	return nama, tgl, price, err
}

func (ctx transactionRepository) GetHistoryData2(id int, tx *sql.Tx) (data models.HistoryTransactionId, err error) {
	query := `SELECT distinct
		transaction.id_transaction,
		users.nama_user,
		to_char(transaction.created_on,'DD-MM-YYYY'),
		transaction_count.total_price
		FROM transaction
		INNER JOIN users ON users.id_user=transaction.id_user
		INNER JOIN transaction_details ON transaction_details.id_transaction=transaction.id_transaction
		INNER JOIN transaction_count ON transaction_count.id_transaction=transaction.id_transaction
		where transaction.id_transaction=$1`
	err = ctx.RepoDB.DB.QueryRow(query, id).Scan(&data.Id_transaction, &data.Nama_user, &data.Tanggal_transaction, &data.Total_price)
	return data, err
}

func (ctx transactionRepository) PostNewTransaction(trx models.Transaction, tx *sql.Tx) (id_trx int, err error) {
	err = ctx.RepoDB.DB.QueryRow(`INSERT INTO transaction 
	(id_user, id_payment, created_on) 
	VALUES ($1, $2, $3) 
	returning id_transaction`, trx.Id_user, trx.Id_payment, time.Now()).Scan(&id_trx)
	return id_trx, err
}

// Insert Multiple Transaction
func (ctx transactionRepository) CobaInsertMultipleItemsTransaction(trx []models.TransactionDetails, id int, tx *sql.Tx) (status bool, err error) {
	sqlStatement := `INSERT INTO transaction_details (id_transaction, id_product, qty_transaction_details, total_transaction_details )VALUES`
	vals := []interface{}{}
	for _, row := range trx {
		sqlStatement += "(?,?,?,?),"

		vals = append(vals, id, row.Id_product, row.Qty_product, row.Total)
	}
	sqlStatement = sqlStatement[0 : len(sqlStatement)-1]
	sqlStatement = ReplaceSQL(sqlStatement, "?")

	stmt, _ := ctx.RepoDB.DB.Prepare(sqlStatement)

	_, err = stmt.Exec(vals...)
	fmt.Println(err)
	return true, nil

}

// Func Count Total Items To Transaction
func (ctx transactionRepository) CountTotalItemTx(id int, tx *sql.Tx) (status bool, err error) {
	query := `UPDATE transaction_details
	SET total_transaction_details=qty_transaction_details*(SELECT product.harga_product WHERE product.id_product=transaction_details.id_product)
	FROM product
	WHERE product.id_product = transaction_details.id_product AND transaction_details.id_transaction=$1`
	trx, err := ctx.RepoDB.DB.Query(query, id)
	fmt.Println(trx)
	return true, err

}

// Auto Update Inventory From Product
func (ctx transactionRepository) AutoUpdateInventory(trx int, tx *sql.Tx) (status bool, err error) {
	rows, err := ctx.RepoDB.DB.Query(`UPDATE product
	SET qty_product=qty_product - qty_transaction_details 
	FROM transaction_details
	WHERE product.id_product = transaction_details.id_product AND transaction_details.id_transaction=$1`, trx)
	fmt.Println("AutoUpdateInventory : ", rows, err)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return true, nil

}

// Func Total Price From Transaction By id_transaction
func (ctx transactionRepository) TransactionCount(trx int, tx *sql.Tx) (price int, err error) {
	err = ctx.RepoDB.DB.QueryRow(`INSERT INTO transaction_count (id_transaction, total_price) 
	VALUES ($1, (SELECT SUM(total_transaction_details) 
	FROM transaction_details WHERE id_transaction=$2 GROUP BY id_transaction)) 
	returning total_price`, trx, trx).Scan(&price)
	return price, err
}

// Get History By id_transaction (Nama_user)
func (ctx transactionRepository) GetHistoryListIdTrx2(trx int, tx *sql.Tx) (nama_user string, err error) {
	err = ctx.RepoDB.DB.QueryRow(`SELECT 
	users.nama_user
	FROM transaction
	INNER JOIN users ON users.id_user=transaction.id_user
	WHERE transaction.Id_transaction = $1`, trx).Scan(&nama_user)
	return nama_user, nil
}

// Get History By id_transaction (Tanggal_transaction)
func (ctx transactionRepository) GetHistoryListIdTrx3(trx int, tx *sql.Tx) (tanggal_transaksi string, err error) {
	err = ctx.RepoDB.DB.QueryRow(`SELECT 
	to_char(transaction.created_on,'DD-MM-YYYY')
	FROM transaction
	WHERE transaction.Id_transaction = $1`, trx).Scan(&tanggal_transaksi)
	return tanggal_transaksi, nil
}

// Get History By id_transaction (Total_price)
func (ctx transactionRepository) GetHistoryListIdTrx4(trx int, tx *sql.Tx) (price int, err error) {
	err = ctx.RepoDB.DB.QueryRow(`SELECT 
	transaction_count.total_price
	FROM transaction
	INNER JOIN transaction_count ON transaction_count.id_transaction=transaction.id_transaction WHERE transaction.Id_transaction = $1`, trx).Scan(&price)
	return price, nil
}

// Get History By id_transaction
func (ctx transactionRepository) GetHistoryListIdTrx(trx int) (status *sql.Rows, err error) {
	//var Transaction []models.HistoryTransaction
	query := `
	SELECT transaction.id_transaction, (SELECT COUNT(transaction_details.id_product)GROUP BY transaction.id_transaction),
	transaction.id_user
	FROM transaction
	INNER JOIN transaction_details on transaction.id_transaction=transaction_details.id_transaction
	INNER JOIN users on transaction.id_user=users.id_user
	where transaction.id_transaction=$1
	Group by transaction.id_transaction
	`

	rows, err := ctx.RepoDB.DB.Query(query, trx)
	fmt.Println("Di repository", rows, err)
	return rows, nil
}

// Get History By id_transaction
func (ctx transactionRepository) GetHistoryListId(trx int) (status []models.HistoryTransactionId, err error) {
	var Transaction []models.HistoryTransactionId
	query := `
		SELECT 
		transaction.id_transaction,
		users.nama_user,
		to_char(transaction.created_on,'DD-MM-YYYY'),
		transaction_count.total_price
		FROM transaction
		INNER JOIN users ON users.id_user=transaction.id_user
		INNER JOIN transaction_count ON transaction_count.id_transaction=transaction.id_transaction
		WHERE transaction.Id_user = $1`

	rows, err := ctx.RepoDB.DB.Query(query, trx)
	fmt.Println("Di repository", rows, err)
	for rows.Next() {
		var listTransaction models.HistoryTransactionId
		rows.Scan(&listTransaction.Id_transaction, &listTransaction.Nama_user, &listTransaction.Tanggal_transaction, &listTransaction.Total_price)
		Transaction = append(Transaction, listTransaction)
	}

	if err != nil {
		return status, err
	}
	defer rows.Close()

	return Transaction, nil
}
