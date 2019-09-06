package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/gorilla/mux"
)

//Item is tenant's item
type Item struct {
	ID     string
	Tenant string
}

//ServiceLookup keeps services port and tenants IDs
type ServiceLookup struct {
	ID      string
	Port    int
	Tenants []string
}

var items = []Item{
	Item{ID: "test", Tenant: "sefa sahin"},
	Item{ID: "test2", Tenant: "sefa sahin2"},
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["TenantID"])
	_, source := http.Get("http://golang.org")
	fmt.Println(source)
	json.NewEncoder(w).Encode(items)
}

func insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	freeport, err := GetFreePort()
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("./service/service", "--port", strconv.Itoa(freeport))
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
	fmt.Printf("Service ProcessID : %v\n", cmd.Process.Pid)
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/items/{TenantID}/count", count).Methods("GET")
	// myRouter.HandleFunc("/items", insert).Methods("POST")
	// err = http.ListenAndServe(":3001", myRouter)
	// if(err != nil){
	// 	fmt.Print(err)
	// }
}
