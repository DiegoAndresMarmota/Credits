package http

import (
	"credits/credits/internal/controller/credits"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Handler struct {
	ctrl *credits.Controller
}

func NewHandler(ctrl *credits.Controller) *Handler {
	return &Handler{ctrl}
}

// GetCreditsDetails obtiene las peticiones GET/ de los creditos
func (h *Handler) GetCreditsDetails(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	details, err := h.ctrl.GetServiceCredits(req.Context(), id)
	if err != nil && errors.Is(err, credits.ErrorNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Error getting credits details: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(details); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
