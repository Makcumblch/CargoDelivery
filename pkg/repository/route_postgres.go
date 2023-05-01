package repository

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type RoutePostgres struct {
	db *sqlx.DB
}

func NewRoutePostgres(db *sqlx.DB) *RoutePostgres {
	return &RoutePostgres{db: db}
}

func (r *RoutePostgres) CreateRoute(projectId int, routeSolution cargodelivery.RouteSolution) (int, error) {
	return 0, nil
}
