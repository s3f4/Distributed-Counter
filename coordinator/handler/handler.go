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

func returnCode500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error!"))
}

func returnCode400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}

//Post makes post request
func Post(item model.Item, w http.ResponseWriter, r *http.Request) (int, int, error) {

	w.Header().Set("Content-Type", "application/json")
	itemContent, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	req, err := http.NewRequest("POST", p.NodePostAddress(), bytes.NewBuffer(itemContent))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	var res map[string]interface{}
	err = json.Unmarshal(result, &res)

	if err != nil {
		fmt.Println(err)
		return -1, -1, err
	}

	return int(res["lastIndexId"].(float64)), int(res["lastTenantId"].(float64)), nil
}

//Get makes get request
func Get(port int, startIndex int, endIndex int, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	resp, err := http.Get(p.NodeCountAddress(port, startIndex, endIndex))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var countObj interface{}
	err = json.NewDecoder(resp.Body).Decode(&countObj)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return countObj, nil
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

	for i, n := range p.Nodes {
		if n != nil && n.ProcessID == ProcessID {
			copy(p.Nodes[i:], p.Nodes[i+1:])
			p.Nodes[len(p.Nodes)-1] = nil
			p.Nodes = p.Nodes[:len(p.Nodes)-1]
			p.NodeCount--
			p.Move()
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}

//Count returns merged count datas
func Count(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tenantID := params["TenantID"]

	res, err := p.Count(tenantID, Get, w, r)

	if err != nil {
		returnCode500(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"count":   res,
	})
}

//Insert gets item and sends data to appropriate node
func Insert(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		fmt.Println(err)
		returnCode400(w, r)
	}

	err = p.Insert(Post, item, w, r)
	if err != nil {
		fmt.Println(err)
		returnCode500(w, r)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
	})
}
