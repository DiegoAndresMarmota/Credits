package data

import (
	"context"
	"credits/data/internal/repository"
	"credits/data/pkg/model"
	"errors"
)

var ErrorNotFound = errors.New("not found")

type dataRepository interface{
	GetData( ctx context.Context, id string ) (*model.Data, error)
}

//Controller define un controlador para el servicio de data.
type Controller struct {
	repo dataRepository
}

//NewController crea un nuevo controlador para el servicio de data.
func NewController(repo dataRepository) *Controller {
	return &Controller{repo: repo}
}

//GetData obtiene los datos según id.
func (c *Controller) GetData(ctx context.Context, id string) (*model.Data, error) {
	res, err := c.repo.GetData(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound){
		return nil, ErrorNotFound
	}
	return res, err
}