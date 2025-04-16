package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/fetch/receipt-processor/handlers/points"
	"github.com/gorilla/mux"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8181, "Port number to run the server on")
	flag.Parse()

	fmt.Println("Initializing receipt processor service...")
	r := mux.NewRouter()

	// routes with HTTP methods specified
	r.HandleFunc("/receipts/process", points.ProcessHandler)
	r.HandleFunc("/receipts/{id}/points", points.PointsHandler)

	fmt.Printf("Server starting on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
