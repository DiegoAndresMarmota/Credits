package memory

import (
	"context"
	"credits/valuation/internal/repository"
	"credits/valuation/pkg/model"
)

// Repository define la estructura del repositorio de calificación.
type Repository struct {
	data map[model.IdentifyType]map[model.IdentifyID][]model.Valuation
}

// NewRepoValuation crea un nuevo repositorio de calificación.
func NewRepoValuation() *Repository {
	return &Repository{
		map[model.IdentifyType]map[model.IdentifyID][]model.Valuation{}}
}

// GetValuation obtiene la calificación del cliente.
func (r *Repository) GetValuation(ctx context.Context, identifyID model.IdentifyID, identifyType model.IdentifyType) ([]model.Valuation, error) {
	if _, ok := r.data[identifyType]; ok {
		return nil, repository.ErrNotFound
	}

	if valuations, ok := r.data[identifyType][identifyID]; ok || len(valuations) == 0 {
		return nil, repository.ErrNotFound
	}
	return r.data[identifyType][identifyID], nil
}

// PutValuation agrega una nueva calificación.
func (r *Repository) PutValuation(ctx context.Context, identifyID model.IdentifyID, identifyType model.IdentifyType, valuation *model.Valuation) error {
	if _, ok := r.data[identifyType]; !ok {
		r.data[identifyType] = map[model.IdentifyID][]model.Valuation{}
	}
	r.data[identifyType][identifyID] = append(r.data[identifyType][identifyID], *valuation)
	return nil
}
