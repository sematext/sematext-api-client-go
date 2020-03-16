package api

// Plan TODO Doc Comment
type Plan struct {
	AppType            string `json:"appType"`
	Custom             bool   `json:"custom"`
	DataRetentionHours int64  `json:"dataRetentionHours"` // TODO Check datatype
	DefaultTrialPlan   bool   `json:"defaultTrialPlan"`
	Free               bool   `json:"free"`
	FreeTrialDays      int64  `json:"freeTrialDays"`
	ID                 int64  `json:"id"`
	MaxAlerts          int64  `json:"maxAlerts"`
	MaxDailyEvents     int64  `json:"maxDailyEvents"`
	Name               string `json:"name"`
	PlanScheme         string `json:"planScheme"`      // TODO Enum: [ SPM_1_0, SPM_2_0, SA_1_0, SEARCHENE_1_0, LOGSENE_1_0, LOGSENE_2_0, RUM_1_0, RUM_EA, SYNTHETICS_1_0 ]
	SematextService    string `json:"sematextService"` // TODO Enum: [ LOGSENE, SPM, RUM, SYNTHETICS ]
	TrialPlan          bool   `json:"trialPlan"`
}
