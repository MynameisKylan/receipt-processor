package points

import "net/http"

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	// todo: unmarshal receipt json, generate id, calculate points, save to memory
	// return id
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Receipt processed successfully"))
}

func PointsHandler(w http.ResponseWriter, r *http.Request) {
	// todo: return points for the given receipt id
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Points retrieved successfully"))
}
