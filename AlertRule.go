package api

// AlertRule TODO Doc Comment
type AlertRule struct {
	AccountEmail                          string                        `json:"accountEmail"` // TODO readOnly: true
	AllowedAppTypes                       []int64                       `json:"allowedAppTypes"`
	AnalyzingTime                         string                        `json:"analyzingTime"`
	AppDisplayState                       string                        `json:"appDisplayState"` // TODO readOnly: true
	AppID                                 int64                         `json:"appId"`
	AppName                               string                        `json:"appName"`  // TODO readOnly: true
	AppState                              string                        `json:"appState"` // TODO readOnly: true
	AppType                               string                        `json:"appType"`  // TODO readOnly: true
	BackToNormalNeeded                    bool                          `json:"backToNormalNeeded"`
	ChartKey                              string                        `json:"chartKey"`
	Color                                 string                        `json:"color"`
	CreatorEmail                          string                        `json:"creatorEmail"`   // TODO readOnly: true
	DefaultAggType                        string                        `json:"defaultAggType"` // TODO readOnly: true
	Description                           string                        `json:"description"`
	DisallowedAppTypes                    []int64                       `json:"disallowedAppTypes"`
	Enabled                               bool                          `json:"enabled"`
	EstimateOperation                     string                        `json:"estimateOperation"` // TODO Enum: [] LESS, MORE, EQUAL, UN_EQUAL, LESS_OR_EQUAL, MORE_OR_EQUAL ]
	EstimateValue                         float64                       `json:"estimateValue"`
	FilterValues                          string                        `json:"filterValues"`
	FilterValuesObj                       []FilterValue                 `json:"filterValuesObj"`
	IgnoreRegularEventsEnabled            bool                          `json:"ignoreRegularEventsEnabled"`
	Integrations                          string                        `json:"integrations"`         // TODO readOnly: true
	LastDataReceivedDate                  int64                         `json:"lastDataReceivedDate"` // TODO readOnly: true
	LastSent                              int64                         `json:"lastSent"`             // TODO readOnly: true
	LastTriggered                         int64                         `json:"lastTriggered"`        // TODO readOnly: true
	Metadata                              map[string]string             `json:"metadata"`             // TODO mixed freeform in API?
	MetricKey                             string                        `json:"metricKey"`            // TODO readOnly: true
	MetricLabel                           string                        `json:"accometricLabeluntEmail"`
	MinDelayBetweenNotificationsInMinutes string                        `json:"minDelayBetweenNotificationsInMinutes"`
	Name                                  string                        `json:"name"`
	NotificationEmails                    []string                      `json:"notificationEmails"`
	NotificationIntegrations              []NotificationIntegration     `json:"notificationIntegrations"`
	NotificationsEnabled                  bool                          `json:"notificationsEnabled"`
	Query                                 string                        `json:"query"`
	ReportName                            string                        `json:"reportName"`
	RuleKey                               int64                         `json:"ruleKey"` // TODO readOnly: true
	RuleType                              string                        `json:"ruleType"`
	Runbook                               string                        `json:"runbook"`
	SavedQueryID                          int64                         `json:"savedQueryId"` // TODO readOnly: true
	Schedule                              []AlertRuleScheduleWeekdayDto `json:"schedule"`
	SematextService                       string                        `json:"sematextService"` // TODO readOnly: true
	SendToEmail                           string                        `json:"sendToEmail"`
	Timezone                              string                        `json:"timezone"`
	UseOnlyAlertRuleIntegrations          bool                          `json:"useOnlyAlertRuleIntegrations"`
	UserPermissions                       UserPermissions               `json:"userPermissions"`
	ValueColumnName                       string                        `json:"valueColumnName"` // TODO readOnly: true
	ValueName                             string                        `json:"valueName"`       // TODO readOnly: true
}
