package model

type Item struct {
	ID        string
	DataIndex int // DataIndex is using in LookupTable
	TenantID  string
}
