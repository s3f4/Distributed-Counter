package handler

import (
	"bytes"
	"coordinator/model"
	"coordinator/processor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var p *processor.Processor

//SetProcessor usign processor once.
func SetProcessor(processor *processor.Processor) {
	p = processor
}

//Count returns merged count datas
func Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	serverAddress := fmt.Sprintf("http://127.0.0.1:%v/items/x/count", p.Nodes[p.NodeIndex].Port)
	resp, err := http.Get(serverAddress)

	var items interface{}
	fmt.Println(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&items)
	fmt.Println(items)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(items)
}

//Insert gets item and sends data to appropriate node
func Insert(w http.ResponseWriter, r *http.Request) {

	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	serverAddress := fmt.Sprintf("http://127.0.0.1:%v/items", p.Nodes[p.NodeIndex].Port)
	itemContent, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", serverAddress, bytes.NewBuffer(itemContent))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var x interface{}
	err = json.Unmarshal(body, &x)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(x)
}