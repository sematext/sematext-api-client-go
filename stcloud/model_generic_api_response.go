/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stcloud

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// GenericAPIResponse is a generic wrapper class for all API responses
type GenericAPIResponse struct {
	// Contains actual data when response is successful. Key and Value is specific to each endpoint
	Data    *interface{} `json:"data,omitempty"`
	Errors  []ModelError `json:"errors,omitempty"`
	Message string       `json:"message,omitempty"`
	Success bool         `json:"success,omitempty"`
}

// ExtractApp - TODO Doc Comment
func (genericAPIResponse *GenericAPIResponse) ExtractApp() (*App, error) {

	// TODO - Shift this to TFP?

	var dataField map[string]interface{}
	var appsField []interface{}
	var appField map[string]interface{}
	var exists bool
	var app App

	dataField = (*genericAPIResponse.Data).(map[string]interface{})

	if appsField, exists = dataField["apps"].([]interface{}); exists {

		mapstructure.Decode(appsField[0], &app)
		return &app, nil

	} else if appField, exists = dataField["app"].(map[string]interface{}); exists {

		mapstructure.Decode(appField, &app)
		return &app, nil

	} else {

		return nil, fmt.Errorf("Unexpected missing apps or app field in API response")

	}

}

// ExtractAppTokens - TODO Doc Comment
func (genericAPIResponse *GenericAPIResponse) ExtractAppTokens() (*[]AppToken, error) {

	// TODO - Shift this to TFP?

	var dataField map[string]interface{}
	var appTokenField []interface{}
	var appTokenList []AppToken
	var exists bool

	dataField = (*genericAPIResponse.Data).(map[string]interface{})
	if appTokenField, exists = dataField["tokens"].([]interface{}); exists {
		mapstructure.Decode(appTokenField, appTokenList)
		return &appTokenList, nil
	}
	return nil, fmt.Errorf("Unexpected missing tokens field in API response")

}
