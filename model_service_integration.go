/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

// ServiceIntegration TODO Godoc Comment
type ServiceIntegration struct {
	AppTypeID           int64  `json:"appTypeId,omitempty"`
	AppTypeName         string `json:"appTypeName,omitempty"`
	DisplayName         string `json:"displayName,omitempty"`
	Enabled             bool   `json:"enabled,omitempty"`
	ExternalProductID   int64  `json:"externalProductId,omitempty"`
	ExternalProductName string `json:"externalProductName,omitempty"`
	ID                  int64  `json:"id,omitempty"`
	IntegrationType     string `json:"integrationType,omitempty"`
	Ordinal             int32  `json:"ordinal,omitempty"`
	ParentIntegrationID int64  `json:"parentIntegrationId,omitempty"`
	SematextService     string `json:"sematextService,omitempty"`
	Visible             bool   `json:"visible,omitempty"`
}
