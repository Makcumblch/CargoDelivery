package service

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/repository"
)

type ProjectService struct {
	repo repository.IProject
}

func NewProjectService(repo repository.IProject) *ProjectService {
	return &ProjectService{repo: repo}
}

func (p *ProjectService) GetUserProject(userId, projectId int) (cargodelivery.Project, error) {
	return p.repo.GetUserProject(userId, projectId)
}

func (p *ProjectService) CreateProject(userId int, project cargodelivery.Project) (int, error) {
	return p.repo.CreateProject(userId, project, cargodelivery.OWNER)
}

func (p *ProjectService) GetAllProjects(userId int) ([]cargodelivery.Project, error) {
	return p.repo.GetAllProjects(userId)
}

func (p *ProjectService) GetProjectById(userId, projectId int) (cargodelivery.Project, error) {
	return p.repo.GetProjectById(userId, projectId)
}

func (p *ProjectService) DeleteProject(userId, projectId int) error {
	return p.repo.DeleteProject(userId, projectId)
}

func (p *ProjectService) UpdateProject(userId, projectId int, input cargodelivery.UpdateProject) error {
	if err := input.ValidateUpdateProject(); err != nil {
		return err
	}
	return p.repo.UpdateProject(userId, projectId, input)
}
