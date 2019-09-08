package handler

import (
	"bytes"
	"coordinator/model"
	"coordinator/node"
	"coordinator/processor"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var p *processor.Processor

//SetProcessor usign processor once.
func SetProcessor(processor *processor.Processor) {
	p = processor
}

//UpNodes runs nodes from front-end
func UpNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if p != nil && len(p.Nodes) > 0 {
		node.KillNodes(0, p.Nodes)
	}

	params := mux.Vars(r)
	NodeCount := params["NodeCount"]
	nc, _ := strconv.Atoi(NodeCount)
	nodes, _ := node.InitNodes(nc)
	processor := processor.Processor{
		Nodes:     nodes,
		NodeCount: nc,
	}

	SetProcessor(&processor)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"nodes":   p.Nodes,
	})
}

//GetNodes returns nodes to show front end
func GetNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nodes []*model.Node
	if p != nil && len(p.Nodes) > 0 {
		nodes = p.Nodes
	} else {
		nodes = nil
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"nodes":   nodes,
	})
}

//Count returns merged count datas
func Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tenantID := params["TenantID"]
	serverAddress := fmt.Sprintf("http://127.0.0.1:%s/items/%s/count", p.Nodes[p.NodeIndex].Port, tenantID)
	resp, err := http.Get(serverAddress)

	if err != nil {
		fmt.Println(err)
	}

	var items interface{}
	err = json.NewDecoder(resp.Body).Decode(&items)

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
	p.Move()
	itemContent, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", serverAddress, bytes.NewBuffer(itemContent))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(result)
}
