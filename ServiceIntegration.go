package api

// ServiceIntegration TODO Doc Comment
type ServiceIntegration struct {
	AppTypeID           int64  `json:"appTypeId"`
	AppTypeName         string `json:"appTypeName"`
	DisplayName         string `json:"displayName"`
	Enabled             bool   `json:"enabled"`
	ExternalProductID   int64  `json:"externalProductId"`
	ExternalProductName string `json:"externalProductName"`
	ID                  int64  `json:"id"`
	IntegrationType     string `json:"integrationType"`
	Ordinal             int32  `json:"ordinal"`
	ParentIntegrationID int64  `json:"parentIntegrationId"`
	SematextService     string `json:"sematextService"`
	Visible             bool   `json:"visible"`
}
