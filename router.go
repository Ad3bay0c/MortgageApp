package main

import (
	"github.com/Ad3bay0c/mortgage_app/bank"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter()
	subRouter.HandleFunc("/create", bank.Create).Methods("POST")
	subRouter.HandleFunc("/", bank.BankList).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalf("Error In Connection: %v", err.Error())
	}
	log.Printf("Server Started at localhost:3000")
}
