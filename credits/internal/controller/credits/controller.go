package credits

import (
	"context"
	"credits/credits/internal/gateway"
	credit "credits/credits/pkg/model"
	data "credits/data/pkg/model"
	"credits/valuation/pkg/model"
	"errors"
)

var ErrorNotFound = errors.New("credit not found")

type valuationGateway interface {
	GetAddValuation(ctx context.Context, valuaID model.IdentifyID, valuaType model.IdentifyType) (float64, error)
	PutEditValuation(ctx context.Context, valuaID model.IdentifyID, valuaType model.IdentifyType, valuation *model.Valuation) error
}

type dataGateway interface {
	GetDataID(ctx context.Context, id string) (*data.Data, error)
}

type Controller struct {
	valuation valuationGateway
	data      dataGateway
}

// NewController crea una nueva instancia del controlador para el servicio de creditos.
func NewController(valuation valuationGateway, data dataGateway) *Controller {
	return &Controller{
		valuation: valuation,
		data:      data,
	}
}

// GetServiceCredits retorna el detalle del credito incluyendo la evaluacion y los datos del credito
func (c *Controller) GetServiceCredits(ctx context.Context, id string) (*credit.CreditDetails, error) {
	data, err := c.data.GetDataID(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrorNotFound) {
		return nil, ErrorNotFound
	} else if err != nil {
		return nil, err
	}
	details := &credit.CreditDetails{
		Data: *data,
	}
	valuation, err := c.valuation.GetAddValuation(ctx, model.IdentifyID(id), model.IdentifyTypeData)
	if err != nil && !errors.Is(err, gateway.ErrorNotFound) {
		return nil, err
	} else if err != nil {
		return nil, err
	} else {
		details.Valuation = &valuation
	}
	return details, nil
}
