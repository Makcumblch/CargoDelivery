package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type ClientService struct {
	repo repository.IClient
}

func NewClientService(repo repository.IClient) *ClientService {
	return &ClientService{repo: repo}
}

func (c *ClientService) CreateClient(projectId int, client cargodelivery.Client) (int, error) {
	return c.repo.CreateClient(projectId, client)
}

func (c *ClientService) GetAllClients(projectId int) ([]cargodelivery.Client, error) {
	return c.repo.GetAllClients(projectId)
}

func (c *ClientService) GetClientById(projectId, clientId int) (cargodelivery.Client, error) {
	return c.repo.GetClientById(projectId, clientId)
}

func (c *ClientService) DeleteClient(projectId, clientId int) error {
	return c.repo.DeleteClient(projectId, clientId)
}

func (c *ClientService) UpdateClient(projectId, clientId int, input cargodelivery.UpdateClient) error {
	if err := input.ValidateUpdateClient(); err != nil {
		return err
	}
	return c.repo.UpdateClient(projectId, clientId, input)
}
