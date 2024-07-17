package main

import (
	"context"
	"credits/balancer/pkg/discovery"
	search "credits/balancer/pkg/discovery/search"
	"credits/credits/internal/controller/credits"
	data "credits/credits/internal/gateway/data/http"
	valuation "credits/credits/internal/gateway/valuation/http"
	handler "credits/credits/internal/handler/http"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const serviceName = "credits"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API Handler Port")
	flag.Parse()
	log.Printf("Starting the credits service on port %d", port)
	registry, err := search.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Instance(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.NotInstance(ctx, instanceID, serviceName)

	data := data.NewGateway(registry)
	valuation := valuation.NewGateway(registry)
	ctrl := credits.NewController(valuation, data)
	h := handler.NewHandler(ctrl)
	http.Handle("/credits", http.HandlerFunc(h.GetCreditsDetails))
	if err := http.ListenAndServe(":8099", nil); err != nil {
		panic(err)
	}
}
