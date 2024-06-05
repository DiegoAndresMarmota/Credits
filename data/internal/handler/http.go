package handler

import (
	"credits/data/internal/controller/data"
	"credits/data/internal/repository"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Handler define el handler http de los datos de los creditos.
type Handler struct {
	cdt *data.Controller
}

// NewHandler crea un nuevo handler http de los datos de los nuevos creditos.
func NewHandler(cdt *data.Controller) *Handler {
	return &Handler{cdt}
}

// GetData maneja las solicitudes GET/ de data.
func (h *Handler) GetData(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := req.Context()
	data, err := h.cdt.GetData(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Response encoding error: %v\n", err)
	}

}
