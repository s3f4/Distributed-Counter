package processor

type Processor struct {
	ServerIndex int
	ServerCount int
}

func (p *Processor) Move() *Processor {
	p.ServerIndex = (p.ServerIndex + 1) % p.ServerCount
	return p
}

func (p *Processor) InsertRoute() *Processor {
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
