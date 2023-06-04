package repository

import (
	"encoding/json"
	"fmt"
	"time"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type RoutePostgres struct {
	db *sqlx.DB
}

func NewRoutePostgres(db *sqlx.DB) *RoutePostgres {
	return &RoutePostgres{db: db}
}

func (r *RoutePostgres) CreateRoute(projectId int, solutionToDb cargodelivery.SolutionToDb, date time.Time) (int, error) {
	var id int
	strSolution, err := json.Marshal(solutionToDb)
	if err != nil {
		return 0, err
	}
	createCarQuery := fmt.Sprintf("INSERT INTO %s (project_id, dt, solution) VALUES ($1, $2, $3) RETURNING id", routeTable)
	row := r.db.QueryRow(createCarQuery, projectId, time.Now(), strSolution)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *RoutePostgres) GetAllRoutes(projectId int) ([]cargodelivery.Routedb, error) {
	var routes []cargodelivery.Routedb

	query := fmt.Sprintf("SELECT * FROM %s WHERE project_id = $1 ORDER BY dt DESC", routeTable)
	err := r.db.Select(&routes, query, projectId)

	return routes, err
}

func (r *RoutePostgres) DeleteRoute(projectId, routeId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE project_id = $1 AND id = $2", routeTable)
	_, err := r.db.Exec(query, projectId, routeId)

	return err
}
