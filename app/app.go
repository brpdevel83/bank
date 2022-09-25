package app

import (
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:9900", nil))
}
