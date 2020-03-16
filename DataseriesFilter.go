package api

// DataSeriesFilter TODO Doc Comment
type DataSeriesFilter struct {
	Aggregation string   `json:"aggregation"` // TODO Enum: [ NONE, SUM, AVG, MIN, MAX ]
	MultiValue  bool     `json:"multiValue"`
	Values      []string `json:"values"`
}
