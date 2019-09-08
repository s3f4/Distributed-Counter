package processor

import (
	"coordinator/model"
	"fmt"
	"net/http"
)

//Partition ...
type Partition struct {
	ServerID         int
	PartitionIndexes []int // Partition holds first index and last index
}

type Processor struct {
	NodeIndex  int
	NodeCount  int
	Nodes      []*model.Node
	Partitions map[string][]*Partition // tenantIds partitions
}

type sendToServerFn func(model.Item, http.ResponseWriter, *http.Request) int

//Move moves between nodes
func (p *Processor) Move() *Processor {
	p.NodeIndex = (p.NodeIndex + 1) % p.NodeCount
	return p
}

//NodeAddress returns current node's address
func (p *Processor) NodeAddress() string {
	return fmt.Sprintf("http://127.0.0.1:%v/items", p.Nodes[p.NodeIndex].Port)
}

//Insert ...
func (p *Processor) Insert(sendToServer sendToServerFn, item model.Item, w http.ResponseWriter, r *http.Request) *Processor {
	lastIndexId := sendToServer(item, w, r)
	p.Move().Move()
	lastIndexId = sendToServer(item, w, r)
	fmt.Println(lastIndexId)
	return p
}

func (p *Processor) Merge() *Processor {
	return p
}
