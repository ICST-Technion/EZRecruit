package queries

// Sortable can be extended to query structures to support sorting logic.
type Sortable struct {
	SortLabels []string `form:"sortBy" json:"sortBy"`
}

// Filterable can be extended to query structures to support sorting logic.
type Filterable struct {
	FilterLabels []string `form:"filterBy" json:"filterBy"`
}
