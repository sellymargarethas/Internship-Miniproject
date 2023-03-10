package loginRepository

import (
	"fmt"

	"Internship-Miniproject/models"
	"Internship-Miniproject/repositories"
)

type loginRepository struct {
	RepoDB repositories.Repository
}

func NewLoginRepository(repoDB repositories.Repository) loginRepository {
	return loginRepository{
		RepoDB: repoDB,
	}
}

const defineColumn = `nama_user, id_user`

// Func Cheking Hash
func (ctx loginRepository) CheckHash(user models.LoginUser) (status string, err error) {
	err = ctx.RepoDB.DB.QueryRow(`SELECT password_user FROM users WHERE email_user=$1`, user.Email).Scan(&status)
	return status, err
}

// Func Return Login
func (ctx loginRepository) ReturnLogin(user models.LoginUser) (status models.LoginDataResponse, err error) {
	rows := ctx.RepoDB.DB.QueryRow(`SELECT `+defineColumn+` FROM users WHERE email_user=$1`, user.Email).Scan(&status.Nama, &status.Id_user)

	fmt.Println(rows, err)
	return status, nil
}
