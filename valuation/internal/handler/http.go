package handler

import (
	"credits/valuation/internal/controller"
	"credits/valuation/internal/repository"
	"credits/valuation/pkg/model"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

// Handler define el controlador del servicio de evaluación.
type Handler struct {
	ctrl *controller.Controller
}

// NewHandValuation crea un nuevo manejador del servicio de evaluación.
func NewHandValuation(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	identifyID := model.IdentifyID(req.FormValue("id"))
	if identifyID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	identifyType := model.IdentifyType(req.FormValue("type"))
	if identifyType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		k, err := h.ctrl.GetAddingValuation(req.Context(), identifyID, identifyType)
		if err != nil && errors.Is(err, repository.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(k); err != nil {
			log.Printf("Error al codificar la respuesta JSON: %\n", err)
		}
	case http.MethodPut:
		valueID := model.IdentifyID(req.FormValue("valueID"))
		k, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutValuation(
			req.Context(),
			identifyID,
			identifyType,
			&model.Valuation{
				IdentifyID: string(valueID),
				Value:      model.ValuationValue(k)}); err != nil {
			log.Printf("Error al procesar la petición PUT: %\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
