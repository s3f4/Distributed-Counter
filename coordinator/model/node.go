package model

//Node keeps nodes ports and process IDs
type Node struct {
	Port         int
	ProcessID    int
	LastIndexID  int
	LastTenantID string
}
