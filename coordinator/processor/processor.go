package processor

import (
	"coordinator/model"
	"fmt"
	"net/http"
)

type Processor struct {
	NodeIndex    int
	NodeCount    int
	Nodes        []*model.Node
	Partitions   map[string]map[int][]*model.Partition // tenantIds partitions
	lastTenantId int
}

type sendToServerFn func(model.Item, http.ResponseWriter, *http.Request) (int, int)

func (p *Processor) createPartition(item model.Item) {
	if p.Partitions[item.TenantID] == nil {
		p.Partitions[item.TenantID] = map[int][]*model.Partition{
			p.NodeIndex: []*model.Partition{&model.Partition{
				PartitionIndexes: make([]int, 0),
			}},
		}
	}
}

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
	p.Move()
	p.createPartition(item)
	lastIndexId, lastTenantId := sendToServer(item, w, r)
	p.Move().Move()
	p.createPartition(item)
	fmt.Println(p.Partitions)
	lastIndexId, lastTenantId = sendToServer(item, w, r)
	fmt.Printf("lastItemId:%v , lastTenantId:%v\n", lastIndexId, lastTenantId)
	return p
}

func (p *Processor) Merge() *Processor {
	return p
}
