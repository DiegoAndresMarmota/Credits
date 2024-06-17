package model

import "credits/data/pkg/model"

//CreditsDetails entrega la "data" de las "evaluaciones" agregadas
type CreditsDetails struct {
	Valuation *float64   `json:"valuation,omitempty"`
	Data      model.Data `json:"data"`
}
