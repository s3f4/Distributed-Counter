package model

//LookupTable ...
type LookupTable struct {
	ServerIndex int
	TenantMap   map[int][]int // for example TenantMap[tenantID][firstItemID,lastItemID,2nd partition firstItemId,lastItemId ....]
}
