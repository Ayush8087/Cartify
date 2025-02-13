package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: "1000"},
	{ID: "2", Name: "Smartphone", Price: "500"},
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", getProducts).Methods("GET")

	log.Println("Product service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
