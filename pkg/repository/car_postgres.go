package repository

import (
	"fmt"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type CarPostgres struct {
	db *sqlx.DB
}

func NewCarPostgres(db *sqlx.DB) *CarPostgres {
	return &CarPostgres{db: db}
}

func (c *CarPostgres) CreateCar(projectId int, car cargodelivery.Car) (int, error) {
	var id int
	createCarQuery := fmt.Sprintf("INSERT INTO %s (project_id, load_capacity, width, height, length, fuel_consumption, name) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", carsTable)
	row := c.db.QueryRow(createCarQuery, projectId, car.LoadCapacity, car.Width, car.Height, car.Length, car.FuelConsumption, car.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (c *CarPostgres) GetAllCars(projectId int) ([]cargodelivery.Car, error) {
	var cars []cargodelivery.Car

	query := fmt.Sprintf("SELECT id, name, load_capacity, width, height, length, fuel_consumption FROM %s WHERE project_id = $1", carsTable)
	err := c.db.Select(&cars, query, projectId)

	return cars, err
}

func (c *CarPostgres) GetCarById(projectId, carId int) (cargodelivery.Car, error) {
	var car cargodelivery.Car

	query := fmt.Sprintf("SELECT id, name, load_capacity, width, height, length, fuel_consumption FROM %s WHERE project_id = $1 AND id = $2", carsTable)
	err := c.db.Get(&car, query, projectId, carId)

	return car, err
}

func (c *CarPostgres) DeleteCar(projectId, carId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE project_id = $1 AND id = $2", carsTable)
	_, err := c.db.Exec(query, projectId, carId)

	return err
}

func (c *CarPostgres) UpdateCar(projectId, carId int, input cargodelivery.UpdateCar) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.FuelConsumption != nil {
		setValues = append(setValues, fmt.Sprintf("fuel_consumption=$%d", argId))
		args = append(args, *input.FuelConsumption)
		argId++
	}

	if input.Height != nil {
		setValues = append(setValues, fmt.Sprintf("height=$%d", argId))
		args = append(args, *input.Height)
		argId++
	}

	if input.Length != nil {
		setValues = append(setValues, fmt.Sprintf("length=$%d", argId))
		args = append(args, *input.Length)
		argId++
	}

	if input.LoadCapacity != nil {
		setValues = append(setValues, fmt.Sprintf("load_capacity=$%d", argId))
		args = append(args, *input.LoadCapacity)
		argId++
	}

	if input.Width != nil {
		setValues = append(setValues, fmt.Sprintf("width=$%d", argId))
		args = append(args, *input.Width)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE project_id = $%d AND id = $%d", carsTable, setQuery, argId, argId+1)
	args = append(args, projectId, carId)

	_, err := c.db.Exec(query, args...)

	return err
}
