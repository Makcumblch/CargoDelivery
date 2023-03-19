package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type CarService struct {
	repo repository.ICar
}

func NewCarService(repo repository.ICar) *CarService {
	return &CarService{repo: repo}
}

func (c *CarService) CreateCar(projectId int, car cargodelivery.Car) (int, error) {
	return c.repo.CreateCar(projectId, car)
}

func (c *CarService) GetAllCars(projectId int) ([]cargodelivery.Car, error) {
	return c.repo.GetAllCars(projectId)
}

func (c *CarService) GetCarById(projectId, carId int) (cargodelivery.Car, error) {
	return c.repo.GetCarById(projectId, carId)
}

func (c *CarService) DeleteCar(projectId, carId int) error {
	return c.repo.DeleteCar(projectId, carId)
}

func (c *CarService) UpdateCar(projectId, carId int, input cargodelivery.UpdateCar) error {
	if err := input.ValidateUpdateCar(); err != nil {
		return err
	}
	return c.repo.UpdateCar(projectId, carId, input)
}
