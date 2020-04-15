/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

// AlertRuleScheduleWeekdayDto TODO godoc comment
type AlertRuleScheduleWeekdayDto struct {
	Day       string                          `json:"day,omitempty"`
	Index     int32                           `json:"index,omitempty"`
	Intervals []AlertRuleScheduleTimeRangeDto `json:"intervals,omitempty"`
	Label     string                          `json:"label,omitempty"`
	Type_     string                          `json:"type,omitempty"`
}
