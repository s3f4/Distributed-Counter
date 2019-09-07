package processor

import (
	"coordinator/model"
)

type Processor struct {
	NodeIndex int
	NodeCount int
	Nodes     []*model.Node
	Lookup    map[string]map[int][]int // for example TenantMap[tenantID][firstItemID,lastItemID,2nd partition firstItemId,lastItemId ....]
}

func (p *Processor) Move() *Processor {
	p.NodeIndex = (p.NodeIndex + 1) % p.NodeCount
	return p
}

func (p *Processor) InsertRoute(item model.Item) *Processor {
	p.Move()
	return p
}

func (p *Processor) ReplicateRoute(item interface{}) *Processor {
	p.Move().Move()
	return p
}
func (p *Processor) ReReplicate() *Processor {
	return p
}
func (p *Processor) Merge() *Processor {
	return p
}
