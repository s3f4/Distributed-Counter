package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Item is tenant's item
type Item struct {
	Id     string
	Tenant string
}

var items = []Item{
	Item{Id: "test", Tenant: "sefa sahin"},
	Item{Id: "test2", Tenant: "sefa sahin2"},
}

func count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["TenantID"])
	_, source := http.Get("https://tutorialedge.net/golang/creating-restful-api-with-golang/")
	fmt.Println(source)
	json.NewEncoder(w).Encode(items)
}

func insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(items)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items/{TenantID}/count", count).Methods("GET")
	myRouter.HandleFunc("/items", insert).Methods("POST")

	port := flag.String("port", "3000", " default port is 3000")
	flag.Parse()
	fmt.Printf("Port: %v", *port)
	// http.ListenAndServe(":"+*port, myRouter)
	for {

	}
	fmt.Println("I'm dead")
}
