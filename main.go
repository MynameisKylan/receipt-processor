package main

import (
	"fmt"
	"net/http"

	"github.com/fetch/receipt-processor/handlers/points"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Initializing receipt processor service...")
	r := mux.NewRouter()

	// routes with HTTP methods specified
	r.HandleFunc("/receipts/process", points.ProcessHandler)
	r.HandleFunc("/receipts/{id}/points", points.PointsHandler)

	fmt.Println("Server starting on :8181")
	err := http.ListenAndServe(":8181", r)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
