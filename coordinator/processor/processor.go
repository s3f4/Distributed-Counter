package processor

import (
	"coordinator/model"
	"fmt"
	"net/http"
	"strconv"
)

type Processor struct {
	NodeIndex     int
	NodeProcessID int
	NodeCount     int
	Nodes         []*model.Node
	Partitions    map[string]map[int][]*model.Partition // tenantIds partitions
	lastIndexId   int
	lastTenantId  string
}

type sendToServerFn func(model.Item, http.ResponseWriter, *http.Request) (int, int)

func (p *Processor) createPartition(item model.Item) *Processor {
	if p.Partitions[item.TenantID] == nil {
		p.Partitions[item.TenantID] = map[int][]*model.Partition{
			p.NodeProcessID: []*model.Partition{&model.Partition{
				PartitionIndexes: make([]int, 2),
				IsCopy:           false,
			}},
		}
	} else {
		if p.Partitions[item.TenantID][p.NodeProcessID] == nil {
			p.Partitions[item.TenantID][p.NodeProcessID] = []*model.Partition{&model.Partition{
				PartitionIndexes: make([]int, 2),
				IsCopy:           false,
			},
			}
		}
	}
	return p
}

func (p *Processor) handlePartitionIndex(item model.Item, isCopy bool) *Processor {
	lp := p.Partitions[item.TenantID][p.NodeProcessID]

	if len(lp[len(lp)-1].PartitionIndexes) == 0 {
		lp[len(lp)-1].PartitionIndexes = []int{
			p.lastIndexId, p.lastIndexId,
		}
		lp[len(lp)-1].IsCopy = isCopy
	} else {
		lp[len(lp)-1].PartitionIndexes[1] = p.lastIndexId
		lp[len(lp)-1].IsCopy = isCopy
	}
	return p
}

//Move moves between nodes
func (p *Processor) Move() *Processor {
	p.NodeIndex = (p.NodeIndex + 1) % p.NodeCount
	p.NodeProcessID = p.Nodes[p.NodeIndex].ProcessID
	return p
}

//NodeAddress returns current node's address
func (p *Processor) NodeAddress() string {
	return fmt.Sprintf("http://127.0.0.1:%v/items", p.Nodes[p.NodeIndex].Port)
}

//Insert ...
func (p *Processor) Insert(sendToServer sendToServerFn, item model.Item, w http.ResponseWriter, r *http.Request) *Processor {
	//Send Insert request first node
	lastIndexId, lastTenantId := sendToServer(item, w, r)

	p.lastIndexId = lastIndexId
	p.lastTenantId = strconv.Itoa(lastTenantId)

	p.Move().
		createPartition(item).
		handlePartitionIndex(item, false)

	lastIndexId, lastTenantId = sendToServer(item, w, r)
	p.lastIndexId = lastIndexId
	p.lastTenantId = strconv.Itoa(lastTenantId)

	p.Move().
		createPartition(item).
		handlePartitionIndex(item, true)

	fmt.Println(p.Partitions)
	return p
}

func (p *Processor) Merge() *Processor {
	return p
}
