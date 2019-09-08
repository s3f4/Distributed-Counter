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

//Initialize nodes and assign to processor
func initializeNodes(nodeCount string) {
	nc, _ := strconv.Atoi(nodeCount)
	nodes, _ := node.InitNodes(nc)

	p.Nodes = nodes
	p.NodeCount = nc
}

//Post makes post request
func Post(item model.Item, w http.ResponseWriter, r *http.Request) (int, int) {

	w.Header().Set("Content-Type", "application/json")
	itemContent, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	req, err := http.NewRequest("POST", p.NodeAddress(), bytes.NewBuffer(itemContent))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	var res map[string]interface{}
	err = json.Unmarshal(result, &res)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	return int(res["lastIndexId"].(float64)), int(res["lastTenantId"].(float64))
}

//UpNodes runs nodes from front-end
func UpNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if p != nil && len(p.Nodes) > 0 {
		node.KillNodes(0, p.Nodes)
		p.NodeCount = 0
	}

	params := mux.Vars(r)
	NodeCount := params["NodeCount"]
	initializeNodes(NodeCount)

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

//Shutdown shutdowns node by given processID
func Shutdown(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	ProcessID, _ := strconv.Atoi(params["ProcessID"])
	node.KillNodes(ProcessID, nil)

	for i, node := range p.Nodes {
		if node.ProcessID == ProcessID {
			copy(p.Nodes[i:], p.Nodes[i+1:])
			p.Nodes[len(p.Nodes)-1] = nil
			p.Nodes = p.Nodes[:len(p.Nodes)-1]
			p.NodeCount--
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

//Count returns merged count datas
func Count(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	tenantID := params["TenantID"]
	serverAddress := fmt.Sprintf("http://127.0.0.1:%v/items/%s/count", p.Nodes[p.NodeIndex].Port, tenantID)
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

	p.Insert(Post, item, w, r)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
