package api

// Invitation TODO Doc Comment
type Invitation struct {
	App           App    `json:"app"`
	Apps          []App  `json:"apps"`
	ID            int64  `json:"id"`            // TODO readOnly: true
	InviteDate    string `json:"inviteDate"`    // TODO readOnly: true, ISO 8063 or long as ($date-time)?
	InviteStatus  string `json:"inviteStatus"`  // TODO readOnly: true Enum:[ PENDING, ACCEPTED, CANCELLED, DECLINED ]
	InviteeEmail  string `json:"inviteeEmail"`  // example: guest@sematext.com
	InviteeRole   string `json:"inviteeRole"`   // TODO example: DEMO Enum:[ SUPER_USER, OWNER, ADMIN, USER, DEMO, ANONYMOUS ]
	InviteeStatus string `json:"inviteeStatus"` // TODO example: ACTIVE readOnly: true Enum:[ INACTIVE, ACTIVE ]
	InviterEmail  string `json:"inviterEmail"`  // TODO readOnly: true
	UUID          string `json:"uuid"`          // TODO readOnly: true
}
