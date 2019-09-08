package model

//Partition ...
type Partition struct {
	PartitionIndexes []int // Partition holds first index and last index
	IsCopy           bool
}
