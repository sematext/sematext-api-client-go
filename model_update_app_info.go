/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

import "fmt"

// UpdateAppInfo TODO Golang comment
type UpdateAppInfo struct {
	Description        string `json:"description,omitempty"`
	IgnorePercentage   int32  `json:"ignorePercentage,omitempty"`
	MaxEvents          int64  `json:"maxEvents,omitempty"`
	MaxLimitMB         int64  `json:"maxLimitMB,omitempty"`
	Name               string `json:"name,omitempty"`
	Sampling           bool   `json:"sampling,omitempty"`
	SamplingPercentage int32  `json:"samplingPercentage,omitempty"`
	Staggering         bool   `json:"staggering,omitempty"`
	Status             string `json:"status,omitempty"`
}

// Persist TODO Doc comment
func (updateAppInfo *UpdateAppInfo) Persist(id int, client *Client) (*App, error) {

	path := fmt.Sprintf("/users-web/api/v3/apps/%d", id)
	genericAPIResponse, err := client.PutJSON(path, updateAppInfo)
	if err != nil {
		return nil, err
	}
	var app *App
	if app, err = genericAPIResponse.ExtractApp(); err != nil {
		return nil, err
	}

	return app, nil
}

// IsValid TODO Doc comment
func (updateAppInfo *UpdateAppInfo) IsValid() (valid bool, err error) {

	// TODO - test minimum viable

	return true, nil
}
