package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (ur *UserPostgres) ResetPassword(username string, newPassword string) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return err
	}
	query := fmt.Sprintf(
		"UPDATE %s SET pass_hash=$1 WHERE email=$2",
		userDb, // Теперь безопасно, т.к. проверено
	)
	_, err = tx.Exec(query, newPassword, username)
	if err != nil {
		fmt.Println(err)
		if err = tx.Rollback(); err != nil {
			return err
		}
	}
	return tx.Commit()
}
