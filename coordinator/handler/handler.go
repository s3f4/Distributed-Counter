package handler

import (
	 "encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["TenantID"])
	_, source := http.Get("http://golang.org")
	fmt.Println(source)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"test": "test",
	})
}

func Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"test": "test",
	})
}
