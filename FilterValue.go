package api

// FilterValue TODO Doc Comment
type FilterValue struct {
	AggType    string   `json:"aggType"`
	FilterName string   `json:"filterName"`
	Key        string   `json:"key"`
	Label      string   `json:"label"`
	Name       string   `json:"name"`
	Values     []string `json:"values"`
}
