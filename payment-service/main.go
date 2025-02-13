package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Payment struct {
	ID     string  `json:"id"`
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

var payments = []Payment{}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var newPayment Payment
	if err := json.NewDecoder(r.Body).Decode(&newPayment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payments = append(payments, newPayment)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPayment)
}

func listPayments(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payments)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/payments", createPayment).Methods("POST") // Ensure POST is defined
	r.HandleFunc("/payments", listPayments).Methods("GET")   // Ensure GET is defined

	log.Println("Payment service running on port 8084")
	log.Fatal(http.ListenAndServe(":8084", r))
}
