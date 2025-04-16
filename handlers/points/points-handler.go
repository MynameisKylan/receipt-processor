package points

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// in-memory store for points.
// would replace with a database in a production app
var pointsStore = map[string]int{
	"abc": 100,
}

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type PointsResponse struct {
	Points int `json:"points"`
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	// todo: unmarshal receipt json, generate id, calculate points, save to memory
	// return id
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Receipt processed successfully"))
}

func PointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("ID extracted:", id)

	if id != "" {
		points, exists := pointsStore[id]
		if !exists {
			http.Error(w, "Receipt ID not found", http.StatusNotFound)
			return
		}

		response := PointsResponse{Points: points}
		marshaledData, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Error marshaling response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(marshaledData)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Invalid receipt ID"))
}
