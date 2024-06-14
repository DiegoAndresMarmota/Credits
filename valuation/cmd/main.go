package cmd

import (
	"credits/valuation/internal/controller"
	"credits/valuation/internal/handler"
	"credits/valuation/internal/repository/memory"
	"log"
	"net/http"
)

func main() {
	log.Println("Inicializando el servicio de evaluación")
	repo := memory.NewRepoValuation()
	ctrl := controller.NewValueService(repo)
	h := handler.NewHandValuation(ctrl)
	http.Handle("/valuation", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
