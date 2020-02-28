package api

// AlertRuleScheduleWeekdayDto TODO Doc Comment
type AlertRuleScheduleWeekdayDto struct {
	Day       string                          `json:"day"`
	Index     int32                           `json:"index"`
	Intervals []AlertRuleScheduleTimeRangeDto `json:"intervals"`
	Label     string                          `json:"end"`
	Type      string                          `json:"type"`
}
