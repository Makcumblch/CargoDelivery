package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type CargoService struct {
	repo repository.ICargo
}

func NewCargoService(repo repository.ICargo) *CargoService {
	return &CargoService{repo: repo}
}

func (c *CargoService) CreateCargo(projectId int, cargo cargodelivery.Cargo) (int, error) {
	return c.repo.CreateCargo(projectId, cargo)
}

func (c *CargoService) GetAllCargos(projectId int) ([]cargodelivery.Cargo, error) {
	return c.repo.GetAllCargos(projectId)
}

func (c *CargoService) GetCargoById(projectId, cargoId int) (cargodelivery.Cargo, error) {
	return c.repo.GetCargoById(projectId, cargoId)
}

func (c *CargoService) DeleteCargo(projectId, cargoId int) error {
	return c.repo.DeleteCargo(projectId, cargoId)
}

func (c *CargoService) UpdateCargo(projectId, cargoId int, input cargodelivery.UpdateCargo) error {
	if err := input.ValidateUpdateCargo(); err != nil {
		return err
	}
	return c.repo.UpdateCargo(projectId, cargoId, input)
}
