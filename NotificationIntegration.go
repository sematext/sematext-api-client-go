package api

// NotificationIntegration TODO Doc Comment
type NotificationIntegration struct {
	Applicability   string         `json:"applicability"` // TODO Enum: [ NONE, USE_ALWAYS ]
	CreateDate      string         `json:"createDate"`    // TODO ($date-time) ISO8063 or long?
	CreatedByOwner  bool           `json:"createdByOwner"`
	CreatorID       int64          `json:"creatorId"`
	ID              int64          `json:"id"`
	IntegrationType string         `json:"integrationType"` // TODO Enum:[ PAGER_DUTY, NAGIOS, WEB_HOOKS, WEB_HOOKS_TEMPLATE, HIP_CHAT, EMAIL_LIST, TEMPORARY_EMAIL_LIST ]
	Name            string         `json:"name"`
	Params          map[int]string `json:"params"` //TODO Check key appropriate type?
	State           string         `json:"state"`  // TODO Enum: [ ACTIVE, DISABLED, DELETED ]
	UserID          int64          `json:"userId"`
}
