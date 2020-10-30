/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

// AppToken TODO Godoc commenting
type AppToken struct {
	ID        int    `json:"id,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt int64  `json:"createdAt,omitempty"`
	Readable  bool   `json:"readable,omitempty"`
	Writable  bool   `json:"writeable,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
}
