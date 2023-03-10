package userRepository

import (
	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
	"database/sql"
	"fmt"
	"time"
)

type userRepository struct {
	RepoDB repositories.Repository
}

func NewUserRepository(repoDB repositories.Repository) userRepository {
	return userRepository{
		RepoDB: repoDB,
	}
}

// Insert Admin & User
func (ctx userRepository) InsertUser(user models.User, tx *sql.Tx) (id int, err error) {
	//created_on := time.Now()

	//db transaction nya belom berhasil

	query := `INSERT INTO users 
	(nama_user, email_user, password_user, role_user, created_on) 
	VALUES ($1, $2, $3, $4, $5) 
	RETURNING id_user`

	if tx != nil {
		err = tx.QueryRow(query, user.Nama_user, user.Email_user, user.Password_user, "user", time.Now()).Scan(&id)
	} else {
		err = ctx.RepoDB.DB.QueryRow(query, user.Nama_user, user.Email_user, user.Password_user, "user", time.Now()).Scan(&id)
	}

	return id, err
}

func (ctx userRepository) InsertDataUser(id int, user models.DataUser, tx *sql.Tx) (id_user int, err error) {
	layout := "2006-01-02"
	date, _ := time.Parse(layout, user.Ttl_user)
	fmt.Println(date)
	query := `INSERT INTO data_users 
	(id_user, gender_user, foto_user, phone_user, ttl_user, kota_user) 
	VALUES ( $1, $2, $3, $4, $5, $6) 
	RETURNING id_user`
	//(SELECT MAX(id_user) FROM users),
	if tx != nil {
		err = tx.QueryRow(query, id, user.Gender_user, user.Foto_user, user.Phone_user, user.Ttl_user, user.Kota_user).Scan(&id_user)
	} else {
		err = ctx.RepoDB.DB.QueryRow(query, user.Gender_user, user.Foto_user, user.Phone_user, user.Ttl_user, user.Kota_user).Scan(&id_user)
	}

	return id_user, err
}

func (ctx userRepository) UpdateDataUser(user models.DataUser, id_user int, tx *sql.Tx) (id int, err error) {

	query := `UPDATE data_users SET gender_user=$1, foto_user=$2, phone_user=$3, ttl_user=$4, kota_user=$5 WHERE id_user=$6 RETURNING id_user`
	if tx != nil {
		err = tx.QueryRow(query, user.Gender_user, user.Foto_user, user.Phone_user, user.Ttl_user, user.Kota_user, id_user).Scan(&id)
	} else {
		err = ctx.RepoDB.DB.QueryRow(query, user.Gender_user, user.Foto_user, user.Phone_user, user.Ttl_user, user.Kota_user, id_user).Scan(&id)
	}
	return id, err
}

func (ctx userRepository) UpdateUser(user models.User, tx *sql.Tx) (status bool, err error) {
	query := `UPDATE users SET nama_user=$1, email_user=$2, password_user=$3, updated_on=$4 WHERE id_user=$5`

	if tx != nil {
		_, err = tx.Query(query, user.Nama_user, user.Email_user, user.Password_user, time.Now(), user.Id_user)
	} else {
		_, err = ctx.RepoDB.DB.Query(query, user.Nama_user, user.Email_user, user.Password_user, time.Now(), user.Id_user)
	}
	return true, err
}

// Delete User & Admin
func (ctx userRepository) DeleteUser(id_user int, tx *sql.Tx) (status bool, err error) {

	query := "DELETE FROM users WHERE id_user = $1"
	if tx != nil {
		_, err = tx.Query(query, id_user)

	} else {
		_, err = ctx.RepoDB.DB.Query(query, id_user)
	}

	return true, err
}

// Delete User & Admin
func (ctx userRepository) DeleteUserData(id_user int, tx *sql.Tx) (status bool, err error) {

	query := "DELETE FROM data_users WHERE id_user = $1"
	if tx != nil {
		_, err = tx.Query(query, id_user)
	} else {
		_, err = ctx.RepoDB.DB.Query(query, id_user)
	}

	return true, err
}

//const defineColumn = `nama_user, id_user`

// Get User & Admin
// func (ctx userRepository) GetUser() ([]models.LoginDataResponse, error) {
// 	var people []models.LoginDataResponse

// 	query := `SELECT ` + defineColumn + ` FROM users ORDER BY id_user`
// 	rows, err := ctx.RepoDB.DB.Query(query)
// 	fmt.Println(rows, err)

// 	for rows.Next() {
// 		var user models.LoginDataResponse
// 		rows.Scan(&user.Nama, &user.Id_user)
// 		people = append(people, user)
// 	}

// 	if err != nil {
// 		return people, err
// 	}

// 	defer rows.Close()
// 	return people, nil

// }
func (ctx userRepository) GetUser(user models.LoginDataResponse) ([]models.LoginDataResponse, error) {
	var result []models.LoginDataResponse
	query := `
	SELECT id_user, nama_user 
	FROM users 
	ORDER BY id_user
	`
	rows, err := ctx.RepoDB.DB.Query(query)
	fmt.Println(rows)
	for rows.Next() {
		var listUsers models.LoginDataResponse
		rows.Scan(&listUsers.Id_user, &listUsers.Nama)
		result = append(result, listUsers)

	}
	if err != nil {
		return result, err
	}
	defer rows.Close()

	return result, nil

}

func (ctx userRepository) GetOneUser(i int) (user models.UserProfile) {
	//var people []models.UserProfile
	//var user models.UserProfile
	query := `SELECT users.id_user, users.nama_user, users.email_user, data_users.gender_user, data_users.foto_user, data_users.phone_user, to_char(data_users.ttl_user,'DD-MM-YYYY'), data_users.kota_user
	FROM  data_users INNER JOIN users on users.id_user=data_users.id_user
	WHERE users.id_user=$1`

	rows := ctx.RepoDB.DB.QueryRow(query, i).Scan(&user.Id_user, &user.Nama_user, &user.Email_user, &user.Gender_user, &user.Foto_user, &user.Phone_user, &user.TTl_user, &user.Kota_user)
	fmt.Println(rows)

	return user
}
