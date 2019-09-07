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
	ItemID   string
	TenantID string
}

//Storage holds datas
type Storage struct {
	items []Item
}

var storage = Storage{
	items: make([]Item, 0),
}

func count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("count:")
	fmt.Println(storage.items)
	json.NewEncoder(w).Encode(storage.items)
}

func insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	storage.items = append(storage.items, item)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(item)
}

func main() {

	go Down()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items", insert).Methods("POST", "GET")
	myRouter.HandleFunc("/items/{TenantID}/count", count).Methods("GET")

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
