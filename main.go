package main

import (
	"fmt"
	"net/http"

	"github.com/fetch/receipt-processor/handlers/points"
)

func main() {
	fmt.Println("Initializing receipt processor service...")
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/receipts/process", points.ProcessHandler)
	mux.HandleFunc("/receipts/{id}/points", points.PointsHandler)

	fmt.Println("Server starting on :8181")
	err := http.ListenAndServe(":8181", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
