/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stcloud

type Plan struct {
	AppType string `json:"appType,omitempty"`
	Custom bool `json:"custom,omitempty"`
	DataRetentionHours float64 `json:"dataRetentionHours,omitempty"`
	DefaultTrialPlan bool `json:"defaultTrialPlan,omitempty"`
	Free bool `json:"free,omitempty"`
	FreeTrialDays int64 `json:"freeTrialDays,omitempty"`
	Id int64 `json:"id,omitempty"`
	MaxAlerts int64 `json:"maxAlerts,omitempty"`
	MaxDailyEvents int64 `json:"maxDailyEvents,omitempty"`
	Name string `json:"name,omitempty"`
	PlanScheme string `json:"planScheme,omitempty"`
	SematextService string `json:"sematextService,omitempty"`
	TrialPlan bool `json:"trialPlan,omitempty"`
}
