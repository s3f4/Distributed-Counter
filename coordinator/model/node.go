package model

//Node keeps nodes ports and process IDs
type Node struct {
	Port      int
	ProcessID int
	DataCount int // this property is keeping node's total data count
}
