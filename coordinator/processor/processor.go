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
	Partitions    map[string]map[int]map[bool][]*model.Partition // tenantIds partitions
}

type sendToServerFn func(model.Item, http.ResponseWriter, *http.Request) (int, int, error)
type countGetFn func(int, int, int, http.ResponseWriter, *http.Request) (interface{}, error)

func (p *Processor) createPartition(item model.Item) *Processor {
	if p.Partitions[item.TenantID] == nil {
		p.Partitions[item.TenantID] = map[int]map[bool][]*model.Partition{
			p.NodeProcessID: map[bool][]*model.Partition{
				true:  []*model.Partition{},
				false: []*model.Partition{},
			},
		}
	} else {

		if p.Partitions[item.TenantID][p.NodeProcessID] == nil {
			p.Partitions[item.TenantID][p.NodeProcessID] = map[bool][]*model.Partition{
				true:  []*model.Partition{},
				false: []*model.Partition{},
			}
		}

		if p.Partitions[item.TenantID][p.NodeProcessID][true] == nil {
			p.Partitions[item.TenantID][p.NodeProcessID][true] = []*model.Partition{}
		}

		if p.Partitions[item.TenantID][p.NodeProcessID][false] == nil {
			p.Partitions[item.TenantID][p.NodeProcessID][false] = []*model.Partition{}
		}
	}
	return p
}

func (p *Processor) handlePartitionIndex(item model.Item, LastIndexID int, isCopy bool) *Processor {
	if p.GetCurrentNode().LastTenantID != item.TenantID {
		newPartition := &model.Partition{
			PartitionIndexes: []int{
				LastIndexID, LastIndexID,
			},
		}

		p.Partitions[item.TenantID][p.NodeProcessID][isCopy] = append(p.Partitions[item.TenantID][p.NodeProcessID][isCopy], newPartition)
	} else {
		p.Partitions[item.TenantID][p.NodeProcessID][isCopy][len(p.Partitions[item.TenantID][p.NodeProcessID][isCopy])-1].PartitionIndexes[1] = p.GetCurrentNode().LastIndexID + 1
	}

	return p
}

//GetPortByProcessID ...
func (p *Processor) GetPortByProcessID(ProcessID int) int {
	port := 0
	for i := range p.Nodes {
		if p.Nodes[i].ProcessID == ProcessID {
			port = p.Nodes[i].Port
		}
	}
	return port
}

//GetCurrentNode ..
func (p *Processor) GetCurrentNode() *model.Node {
	return p.Nodes[p.NodeIndex]
}

//SetProcessID ...
func (p *Processor) SetProcessID() *Processor {
	p.NodeProcessID = p.Nodes[p.NodeIndex].ProcessID
	return p
}

//Move moves between nodes
func (p *Processor) Move() *Processor {
	p.NodeIndex = (p.NodeIndex + 1) % p.NodeCount
	return p
}

//NodePostAddress returns current node's address
func (p *Processor) NodePostAddress() string {
	return fmt.Sprintf("http://127.0.0.1:%v/items", p.Nodes[p.NodeIndex].Port)
}

//NodeCountAddress returns current node's address
func (p *Processor) NodeCountAddress(port int, startIndex int, endIndex int) string {
	fmt.Println(fmt.Sprintf("http://127.0.0.1:%v/items/%v/%v", port, startIndex, endIndex))
	return fmt.Sprintf("http://127.0.0.1:%v/items/%v/%v", port, startIndex, endIndex)
}

//Insert ...
func (p *Processor) Insert(sendToServer sendToServerFn, item model.Item, w http.ResponseWriter, r *http.Request) error {
	//Send Insert request first node
	p.SetProcessID()
	lastIndexID, lastTenantID, err := sendToServer(item, w, r)
	if err != nil {
		return err
	}

	p.createPartition(item).
		handlePartitionIndex(item, lastIndexID, false)

	p.GetCurrentNode().LastIndexID = lastIndexID
	p.GetCurrentNode().LastTenantID = strconv.Itoa(lastTenantID)
	p.Move()

	p.SetProcessID()
	lastIndexID, lastTenantID, err = sendToServer(item, w, r)
	if err != nil {
		return err
	}

	p.createPartition(item).
		handlePartitionIndex(item, lastIndexID, true)

	p.GetCurrentNode().LastIndexID = lastIndexID
	p.GetCurrentNode().LastTenantID = strconv.Itoa(lastTenantID)

	p.Move()
	return nil
}

//GetResults ...
func (p *Processor) GetResults(partitionArray map[int][]*model.Partition, get countGetFn, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var results []interface{}
	for processID := range partitionArray {
		if len(partitionArray[processID]) > 0 {
			partitions := partitionArray[processID]
			port := p.GetPortByProcessID(processID)
			for i := range partitions {
				result, err := get(port,
					partitions[i].PartitionIndexes[0],
					partitions[i].PartitionIndexes[1],
					w, r)
				results = append(results, result)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return results, nil
}

//Count ...
func (p *Processor) Count(TenantID string, get countGetFn, w http.ResponseWriter, r *http.Request) (interface{}, error) {

	partitionArray := make(map[int][]*model.Partition, 0)
	copyArray := make(map[int][]*model.Partition, 0)

	for processID, nodes := range p.Partitions[TenantID] {
		for isCopy := range nodes {
			if !isCopy {
				partitionArray[processID] = nodes[isCopy]
			} else {
				copyArray[processID] = nodes[isCopy]
			}
		}
	}

	result, err := p.GetResults(partitionArray, get, w, r)
	if err != nil {
		result, err = p.GetResults(copyArray, get, w, r)
		if err != nil {
			return err, nil
		}
		return result, nil
	}
	return result, nil

	// lastIndexId, lastTenantId := countGetFn(TenantID, w, r)

	// p.lastIndexId = lastIndexId
	// p.lastTenantId = strconv.Itoa(lastTenantId)

	// p.Move().
	// 	createPartition(item).
	// 	handlePartitionIndex(item, false)

	// lastIndexId, lastTenantId = sendToServer(item, w, r)
	// p.lastIndexId = lastIndexId
	// p.lastTenantId = strconv.Itoa(lastTenantId)

	// p.Move().
	// 	createPartition(item).
	// 	handlePartitionIndex(item, true)

	// fmt.Println(p.Partitions)
}

func (p *Processor) Merge() *Processor {
	return p
}
