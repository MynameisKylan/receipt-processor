package points

import (
	"encoding/json"
	"net/http"

	"github.com/fetch/receipt-processor/services"
	"github.com/gorilla/mux"
)

type ProcessResponse struct {
	ID string `json:"id"`
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	var receipt services.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt data", http.StatusBadRequest)
		return
	}

	id, err := services.ProcessReceipt(receipt)
	if err != nil {
		http.Error(w, "Error processing receipt", http.StatusInternalServerError)
		return
	}

	responseData := ProcessResponse{ID: id}
	marshaledData, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, "Error marshaling response data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(marshaledData))
}

func PointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id != "" {
		pointsData, err := services.GetPointsData(id)
		if err == nil {
			marshaledData, err := json.Marshal(pointsData)
			if err != nil {
				http.Error(w, "Error marshaling points data", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			w.Write(marshaledData)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("No receipt found for that ID."))
}
