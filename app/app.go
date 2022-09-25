package app

import (
	"log"
	"net/http"

	"github.com/brpdevel83/bank/service"
	"github.com/gorilla/mux"
)

func Start() {
	//mux := http.NewServeMux()
	mux := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	mux.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)

	//mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	//mux.HandleFunc("/customers/{id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:9900", mux))
}
