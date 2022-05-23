package main

import (
	"log"
	"net/http"
	"time"
	"tracker/handlers"
	"tracker/keepers"

	"github.com/gorilla/mux"
)

var keeper *keepers.TransactionKeeper = keepers.NewTransactionKeeper()
var handler, _ = handlers.NewTransactionHandler(keeper)

func main() {
	r := mux.NewRouter()
	// Add your routes as needed

	r.HandleFunc("/transaction", handler.HandleFunc).Methods("GET")
	r.HandleFunc("/transaction", handler.HandleFunc).Methods("POST")

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Println("Corriendo en: 0.0.0.0:8080")
	panic(srv.ListenAndServe())
}
