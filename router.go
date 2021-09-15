package main

import (
	"github.com/Ad3bay0c/mortgage_app/bank"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//subRouter := router.PathPrefix("/api").Subrouter()
	router.HandleFunc("/create", bank.Create).Methods("POST")
	router.HandleFunc("/", bank.List).Methods("GET")
	router.HandleFunc("/{id}", bank.GetBank).Methods("GET")
	router.HandleFunc("/update/{id}", bank.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", bank.Delete).Methods("DELETE")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalf("Error In Connection: %v", err.Error())
	}
	log.Printf("Server Started at localhost:3000")
}
