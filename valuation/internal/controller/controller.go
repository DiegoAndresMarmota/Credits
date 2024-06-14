package controller

import (
	"context"

	"credits/valuation/internal/repository"
	"credits/valuation/pkg/model"
	"errors"
)

var ErrNotFound = errors.New("valuation not found for a value")

// valuationRepository crea una interface
type valuationRepository interface {
	GetValuation(ctx context.Context, valueID model.IdentifyID, valueType model.IdentifyType) ([]model.Valuation, error)
	PutValuation(ctx context.Context, valueID model.IdentifyID, valueType model.IdentifyType, valuation *model.Valuation) error
}

// Controller define un controlador para el servicio de evaluación.
type Controller struct {
	repo valuationRepository
}

// NewValueService crea un controlador para el servicio de evaluación.
func NewValueService(repo valuationRepository) *Controller {
	return &Controller{repo}
}

func (c *Controller) GetAddingValuation(ctx context.Context, valueID model.IdentifyID, valueType model.IdentifyType) (float64, error) {
	valuations, err := c.repo.GetValuation(ctx, valueID, valueType)
	if err != nil && err == repository.ErrNotFound {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}

	calification := float64(0)
	for _, value := range valuations {
		calification += float64(value.Value)
	}
	return calification / float64(len(valuations)), nil
}

func (c *Controller) PutValuation(ctx context.Context, valueID model.IdentifyID, valueType model.IdentifyType, valuation *model.Valuation) error {
	return c.repo.PutValuation(ctx, valueID, valueType, valuation)
}
