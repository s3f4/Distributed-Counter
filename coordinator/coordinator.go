package main

import (
	"coordinator/handler"
	"coordinator/model"
	"coordinator/processor"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	handler.SetProcessor(&processor.Processor{
		Nodes:     make([]*model.Node, 0),
		NodeCount: 0,
	})

	fmt.Println("servers are ready....")
	//Latency  waiting to up nodes
	time.Sleep(time.Second * 2)

	myRouter := mux.NewRouter().StrictSlash(true)
	/*
		For Frontend requests
	*/
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	myRouter.Use(cors)

	myRouter.HandleFunc("/items/{TenantID}/count", handler.Count).Methods("GET")
	myRouter.HandleFunc("/items", handler.Insert).Methods("POST")
	myRouter.HandleFunc("/nodes", handler.GetNodes).Methods("GET")
	myRouter.HandleFunc("/upNodes/{NodeCount}", handler.UpNodes).Methods("GET")
	myRouter.HandleFunc("/shutdown/{ProcessID}", handler.Shutdown).Methods("GET")
	err := http.ListenAndServe(":3001", myRouter)
	if err != nil {
		fmt.Print(err)
	}
}
