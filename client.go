package cargodelivery

import "errors"

type Client struct {
	Id      int    `json:"id" db:"id"`
	Address string `json:"address" binding:"required" db:"address"`
	Name    string `json:"name" binding:"required" db:"name"`
}

type UpdateClient struct {
	Address *string `json:"address"`
	Name    *string `json:"name"`
}

func (i *UpdateClient) ValidateUpdateClient() error {
	if i.Address == nil && i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
