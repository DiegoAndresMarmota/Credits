package model

//Tipos
type IdentifyID string
type IdentifyType string
type UserID string
type ValuationValue int

const IdentifyTypeData = IdentifyType("data")

//Valuation define la calificación individual.
type Valuation struct {
	IdentifyID   string         `json:"identify_id"`
	IdentifyType string         `json:"identify_type"`
	UserID       UserID         `json:"user_id"`
	Value        ValuationValue `json:"value"`
}
