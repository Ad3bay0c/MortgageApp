package main

import (
	"github.com/Ad3bay0c/mortgage_app/bank"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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

	c := cors.New(cors.Options{
		AllowCredentials: true,
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	handler := c.Handler(router)
	log.Printf("Server Started at localhost:5000")
	err := http.ListenAndServe(":5000", handler)
	if err != nil {
		log.Fatalf("Error In Connection: %v", err.Error())
	}

}
