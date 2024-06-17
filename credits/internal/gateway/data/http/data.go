package http

import (
	"context"
	"credits/credits/internal/gateway"
	"credits/data/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// Gateway define el http de de datos para la evaluación.
type Gateway struct {
	address string
}

// Gateway define una nueva instancia para el constructor para el servicio de datos de la evaluación.
func NewGateway(address string) *Gateway {
	return &Gateway{address}
}

func (ga *Gateway) GetDataID(ctx context.Context, id string) (*model.Data, error) {
	req, err := http.NewRequest(http.MethodGet, ga.address+"/v1/data/"+id, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	req.URL.RawQuery = values.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrorNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx status code: %v", resp.StatusCode)
	}

	var data *model.Data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
