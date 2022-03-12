package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	//	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	//	router.HandleFunc("/customers", addCustomer).Methods("POST")
	//	router.HandleFunc("/customers", updateCustomer).Methods("PUT")
	//	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	http.ListenAndServe(conf.server.port, router)
}
