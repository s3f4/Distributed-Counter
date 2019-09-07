package main

import (
	"coordinator/handler"
	"coordinator/node"
	"coordinator/processor"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	nodes, _ := node.InitNodes(2)
	processor := processor.Processor{
		Nodes:     nodes,
		NodeCount: 2,
	}

	//Latency  waiting to up nodes
	time.Sleep(time.Second * 2)

	handler.SetProcessor(&processor)

	// time.Sleep(5 * time.Second)
	// node.KillNodes(0, nodes)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items/{TenantID}/count", handler.Count).Methods("GET")
	myRouter.HandleFunc("/items", handler.Insert).Methods("POST")
	myRouter.HandleFunc("/nodes", handler.GetNodes).Methods("GET")
	err := http.ListenAndServe(":3001", myRouter)
	if err != nil {
		fmt.Print(err)
	}
}
