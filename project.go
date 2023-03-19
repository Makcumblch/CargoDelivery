package cargodelivery

import "errors"

const (
	OWNER = "owner"
	READ  = "read"
	WRITE = "write"
)

type Project struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" binding:"required" db:"name"`
	Access string `json:"access" db:"access"`
}

type UpdateProject struct {
	Name *string `json:"name"`
}

func (i *UpdateProject) ValidateUpdateProject() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
