package main

import (
	"credits/credits/internal/controller/credits"
	data "credits/credits/internal/gateway/data/http"
	valuation "credits/credits/internal/gateway/valuation/http"
	handler "credits/credits/internal/handler/http"
	"log"
	"net/http"
)

func main() {
	log.Println("Inicializando el servicio de creditos")
	data := data.NewGateway("localhost:8090")
	valuation := valuation.NewGateway("localhost:8091")
	ctrl := credits.NewController(valuation, data)
	h := handler.NewHandler(ctrl)
	http.Handle("/credits", http.HandlerFunc(h.GetCreditsDetails))
	if err := http.ListenAndServe(":8099", nil); err != nil {
		panic(err)
	}
}
