package main

import (
	"context"
	generate "credits/balancer/pkg/discovery"
	search "credits/balancer/pkg/discovery/search"
	"credits/valuation/internal/controller"
	"credits/valuation/internal/handler"
	"credits/valuation/internal/repository/memory"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const serviceName = "valuation"

func main() {
	var port int
	flag.IntVar(&port, "port", 8082, "Port to listen on tne handler")
	flag.Parse()
	log.Printf("Starting Valuation Service on port %d\n", port)

	registry, err := search.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	instanceID := generate.GenerateInstanceID(serviceName)

	if err := registry.Instance(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Printf("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	defer registry.NotInstance(ctx, instanceID, serviceName)

	repo := memory.NewRepoValuation()
	ctrl := controller.NewValueService(repo)
	h := handler.NewHandValuation(ctrl)
	http.Handle("/valuation", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8092", nil); err != nil {
		log.Fatal(err)
	}
}
