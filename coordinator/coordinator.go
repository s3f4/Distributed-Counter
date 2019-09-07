package main

import (
	"coordinator/handler"
	"coordinator/model"
	"time"

	"github.com/gorilla/mux"
)

//Item is tenant's item

var items = []model.Item{
	model.Item{ID: "test", TenantID: "sefa sahin"},
	model.Item{ID: "test2", TenantID: "sefa sahin2"},
}

func main() {

	servers, _ := model.InitServers(3)
	time.Sleep(5 * time.Second)
	model.KillServers(0, servers)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items/{TenantID}/count", handler.Count).Methods("GET")
	myRouter.HandleFunc("/items", handler.Insert).Methods("POST")
	// err := http.ListenAndServe(":3001", myRouter)
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
