package cmd

import (
	"context"
	generate "credits/balancer/pkg/discovery"
	search "credits/balancer/pkg/discovery/search"
	"credits/data/internal/controller/data"
	"credits/data/internal/handler"
	"credits/data/internal/repository/memory"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

const serviceName = "data"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "Port to listen on the handler")
	flag.Parse()
	log.Printf("Starting Data Service on port %d\n", port)
	discovery, err := search.NewRegistry("localhost:8500")
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	instanceID := generate.GenerateInstanceID(serviceName)

	if err := discovery.Instance(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := discovery.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer discovery.NotInstance(ctx, instanceID, serviceName)

	repo := memory.NewRepo()
	cdt := data.NewController(repo)
	h := handler.NewHandler(cdt)
	http.Handle("/data", http.HandlerFunc(h.GetData))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
