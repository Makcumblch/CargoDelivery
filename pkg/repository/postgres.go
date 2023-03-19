package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable       = "users"
	projectsTable    = "projects"
	projectUserTable = "projects_users"
	carsTable        = "cars"
	cargosTable      = "cargos"
	ordersTable      = "orders"
	clientsTable     = "clients"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
