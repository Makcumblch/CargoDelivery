package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type DepoService struct {
	repo repository.IDepo
}

func NewDepoService(repo repository.IDepo) *DepoService {
	return &DepoService{repo: repo}
}

func (c *DepoService) CreateDepo(projectId int, depo cargodelivery.Depo) (int, error) {
	return c.repo.CreateDepo(projectId, depo)
}

func (c *DepoService) GetDepo(projectId int) (cargodelivery.Client, error) {
	depo, err := c.repo.GetDepo(projectId)
	if err != nil {
		return depo, err
	}
	depo.Name = "Депо"
	return depo, nil
}

func (c *DepoService) UpdateDepo(projectId int, input cargodelivery.UpdateDepo) error {
	if err := input.ValidateUpdateDepo(); err != nil {
		return err
	}
	return c.repo.UpdateDepo(projectId, input)
}
