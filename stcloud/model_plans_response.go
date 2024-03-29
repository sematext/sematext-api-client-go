/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stcloud

type PlansResponse struct {
	Data *PlansResponseEntry `json:"data,omitempty"`
	Errors []ModelError `json:"errors,omitempty"`
	Message string `json:"message,omitempty"`
	Success bool `json:"success,omitempty"`
}
