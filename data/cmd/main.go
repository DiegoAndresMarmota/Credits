package main

import (
	"credits/data/internal/controller/data"
	"credits/data/internal/handler"
	"credits/data/internal/repository/memory"
	"log"
	"net/http"
)

func main() {

	log.Println("Starting the credits data service")
	repo := memory.NewRepo()
	cdt := data.NewController(repo)
	h := handler.NewHandler(cdt)
	http.Handle("/credits", http.HandlerFunc(h.GetData))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
