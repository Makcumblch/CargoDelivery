package repository

import (
	"fmt"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/jmoiron/sqlx"
)

type ProjectPostgres struct {
	db *sqlx.DB
}

func NewProjectPostgres(db *sqlx.DB) *ProjectPostgres {
	return &ProjectPostgres{db: db}
}

func (p *ProjectPostgres) GetUserProject(userId, projectId int) (cargodelivery.Project, error) {
	var projectUser cargodelivery.Project
	query := fmt.Sprintf("SELECT access FROM %s WHERE project_id=$1 AND user_id=$2", projectUserTable)
	err := p.db.Get(&projectUser, query, projectId, userId)
	return projectUser, err
}

func (p *ProjectPostgres) CreateProject(userId int, project cargodelivery.Project, access string) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProjectQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", projectsTable)
	row := tx.QueryRow(createProjectQuery, project.Name)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createProjectUserQuery := fmt.Sprintf("INSERT INTO %s (project_id, user_id, access) VALUES ($1, $2, $3)", projectUserTable)
	_, err = tx.Exec(createProjectUserQuery, id, userId, access)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (p *ProjectPostgres) GetAllProjects(userId int) ([]cargodelivery.Project, error) {
	var projects []cargodelivery.Project

	query := fmt.Sprintf("SELECT pr.id, pr.name, pu.access FROM %s pr INNER JOIN %s pu ON pr.id = pu.project_id WHERE pu.user_id = $1", projectsTable, projectUserTable)
	err := p.db.Select(&projects, query, userId)

	return projects, err
}

func (p *ProjectPostgres) GetProjectById(userId, projectId int) (cargodelivery.Project, error) {
	var project cargodelivery.Project

	query := fmt.Sprintf("SELECT pr.id, pr.name, pu.access FROM %s pr INNER JOIN %s pu ON pr.id = pu.project_id WHERE pu.user_id = $1 AND pu.project_id = $2", projectsTable, projectUserTable)
	err := p.db.Get(&project, query, userId, projectId)

	return project, err
}

func (p *ProjectPostgres) DeleteProject(userId, projectId int) error {
	query := fmt.Sprintf("DELETE FROM %s pr USING %s pu WHERE pr.id = pu.project_id AND pu.user_id = $1 AND pu.project_id = $2 AND pu.access = $3", projectsTable, projectUserTable)
	_, err := p.db.Exec(query, userId, projectId, cargodelivery.OWNER)

	return err
}

func (p *ProjectPostgres) UpdateProject(userId, projectId int, input cargodelivery.UpdateProject) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s pr SET %s FROM %s pu WHERE pr.id = pu.project_id AND pu.user_id = $%d AND pu.project_id = $%d AND pu.access = $%d", projectsTable, setQuery, projectUserTable, argId, argId+1, argId+2)
	args = append(args, userId, projectId, cargodelivery.OWNER)

	_, err := p.db.Exec(query, args...)

	return err
}
