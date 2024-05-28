package memory

import (
	"context"
	"credits/data/internal/repository"
	model "credits/data/pkg/model"
	"sync"
)

// Repository define un repositorio de datos de creditos en memoria.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Data
}

// NewRepo crea un nuevo repositorio de datos.
func NewRepo() *Repository {
	return &Repository{
		data: make(map[string]*model.Data),
	}
}

// GetData retorna los datos de los creditos de un cliente, según su identificador.
func (r *Repository) GetData(_ context.Context, id string) (*model.Data, error) {
	r.RLock()
	defer r.RUnlock()

	md, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return md, nil
}
