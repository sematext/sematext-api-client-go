package api

// AlertNotificationRequest TODO Doc Comment
type AlertNotificationRequest struct {
	DefaultInterval int64  `json:"defaultInterval"`
	End             string `json:"end"` // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
	Interval        string `json:"interval"`
	Start           string `json:"start"` // End time of interval. Can be expressed as timestamp in milliseconds or UTC date in yyyy-MM-dd HH:mm:ss format
}
