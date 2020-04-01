package api

// SavedQuery TODO Doc Comment
type SavedQuery struct {
	AlertRule         AlertRule       `json:"alertRule"`
	AllowModification bool            `json:"allowModification"` // TODO example: false readOnly: true
	ApplicationID     string          `json:"applicationId"`
	ApplicationName   string          `json:"applicationName"`  // TODO readOnly: true
	ApplicationToken  string          `json:"applicationToken"` // TODO readOnly: true
	CreatorEmail      string          `json:"creatorEmail"`     // TODO readOnly: true
	ID                string          `json:"id"`               // TODO readOnly: true
	LabelColor        string          `json:"labelColor"`
	LogseneAlertType  map[int]string  `json:"logseneAlertType"` // TODO Check datatype
	OwnerEmail        string          `json:"ownerEmail"`       // TODO readOnly: true
	QueryName         string          `json:"queryName"`
	QueryString       string          `json:"queryString"`
	UserPermissions   UserPermissions `json:"userPermissions"`
}
