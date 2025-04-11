package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In-memory storage for receipts
var receipts = make(map[string]Receipt)

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "The receipt is invalid. Please verify input.", http.StatusBadRequest)
		return
	}

	// Validate receipt
	if err := receipt.Validate(); err != nil {
		http.Error(w, "The receipt is invalid. Please verify input.", http.StatusBadRequest)
		return
	}

	// Generate UUID
	id := generateUUID()
	receipts[id] = receipt

	// Return response
	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)
	response := map[string]int{"points": points}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
