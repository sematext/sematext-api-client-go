package api

// DataSeriesRequest TODO Doc Comment
type DataSeriesRequest struct {
	DefaultInterval int64                    `json:"defaultInterval"`
	End             string                   `json:"end"` // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
	Filters         map[int]DataSeriesFilter `json:"filters"`
	Granularity     string                   `json:"granularity"` // TODO - Data points interval granularity between two data points.Default value is “AUTO” - calculated based on selected time span. Not required while getting filters. Enum:[ AUTO, ONE_MINUTE, FIVE_MINUTES, HOUR, DAY, WEEK, MONTH ]
	Interval        string                   `json:"interval"`
	Metric          string                   `json:"metric"`
	Start           string                   `json:"start"` // TODO Start time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
}
