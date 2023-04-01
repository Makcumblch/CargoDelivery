package repository

import (
	"fmt"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user cargodelivery.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, salt) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := a.db.QueryRow(query, user.Username, user.Password, user.Salt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) GetUser(username string) (cargodelivery.User, error) {
	var user cargodelivery.User
	query := fmt.Sprintf("SELECT id, password_hash, salt FROM %s WHERE username=$1", usersTable)
	err := a.db.Get(&user, query, username)
	return user, err
}

func (a *AuthPostgres) GetUserById(userId int) (cargodelivery.User, error) {
	var user cargodelivery.User
	query := fmt.Sprintf("SELECT id, password_hash, salt FROM %s WHERE id=$1", usersTable)
	err := a.db.Get(&user, query, userId)
	return user, err
}
