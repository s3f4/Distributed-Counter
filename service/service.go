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
	items       []Item
	lastIndexId int
}

var storage = Storage{
	items:       make([]Item, 0),
	lastIndexId: 0,
}

//GetStorage returns db to show front end
func GetStorage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"items":   storage.items,
	})
}

//Count returns items by tenantId
func Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tenantID := params["TenantID"]
	count := 0

	for _, item := range storage.items {
		if item.TenantID == tenantID {
			count++
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   count,
	})
}

//Insert inserts a item
func Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	storage.items = append(storage.items, item)
	storage.lastIndexId++
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"item":        item,
		"lastIndexId": storage.lastIndexId,
	})
}

func main() {

	go Down()

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items", Insert).Methods("POST", "GET")
	myRouter.HandleFunc("/items/{TenantID}/count", Count).Methods("GET")
	myRouter.HandleFunc("/storage", GetStorage).Methods("GET")
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
