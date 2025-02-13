package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Notification struct {
	ID      string `json:"id"`
	UserID  string `json:"userId"`
	Message string `json:"message"`
}

var notifications = []Notification{
	{ID: "1", UserID: "123", Message: "Your order has been shipped!"},
}

func getNotifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifications)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notifications", getNotifications).Methods("GET") // Ensure this is defined

	log.Println("Notification service running on port 8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
