package api

// BasicOrganizationDto TODO Doc Comment
type BasicOrganizationDto struct {
	AuthMethods []BasicAuthMethodDto `json:"authMethods"`
	Name        string               `json:"name"`
	Status      string               `json:"status"` // TODO Enum:[ ACTIVE, IN_REGISTRATION, DISABLED, EXPIRED, INVITED, DEMO ]
	UUID        string               `json:"uuid"`
}
