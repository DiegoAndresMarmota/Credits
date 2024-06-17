package http

import (
	"context"
	"credits/valuation/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// Gateway define el http de los creditos para la evaluación.
type Gateway struct {
	address string
}

// NewGateway define una nueva instancia para el constructor para el servicio de creditos en la evaluación.
func NewGateway(address string) *Gateway {
	return &Gateway{address: address}
}

// GetAddValuation obtiene la valoración de un credito.
func (g *Gateway) GetAddValuation(ctx context.Context, valuaID model.IdentifyID, valuaType model.IdentifyType) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, g.address+"/v1/valuation", nil)
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)
	params := req.URL.Query()
	params.Add("id", string(valuaID))
	params.Add("type", fmt.Sprintf("%v", valuaType))
	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result, nil
}

// PutEditValuation edita las valoraciones de los creditos.
func (g *Gateway) PutEditValuation(ctx context.Context, valuaID model.IdentifyID, valuaType model.IdentifyType, valuation *model.Valuation) error {
	req, err := http.NewRequest(http.MethodPut, g.address+"/v1/valuation", nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)
	params := req.URL.Query()
	params.Add("id", string(valuaID))
	params.Add("type", fmt.Sprintf("%v", valuaType))
	params.Add("userID", string(valuation.UserID))
	params.Add("value", fmt.Sprintf("%v", valuation.Value))

	req.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}
	return nil
}
