package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

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
	defer fmt.Println("testsegfsfgsd")
	go Down()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items/{TenantID}/count", count).Methods("GET")
	myRouter.HandleFunc("/items", insert).Methods("POST")

	port := flag.String("port", "3000", " default port is 3000")
	flag.Parse()

	http.ListenAndServe(":"+*port, myRouter)
}

//Down downs service when kill SIGINT came.
func Down() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	fmt.Println("i am dead")
	os.Exit(0)
}
