package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string
	City string
	zip  string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello and doing greeting.")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Babak", "Istanbul", "34340"},
		{"Nuğdiş", "Istanbul", "34340"},
	}

	if r.Header.Get("Content-Type") == "application-xml" {
		w.Header().Set("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["id"])
}
