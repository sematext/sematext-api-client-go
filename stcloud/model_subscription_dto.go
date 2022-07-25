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
	"time"
)

type SubscriptionDto struct {
	AdditionalParams string           `json:"additionalParams,omitempty"`
	Addresses        string           `json:"addresses,omitempty"`
	Enable           bool             `json:"enable,omitempty"`
	Filters          string           `json:"filters,omitempty"`
	Frequency        string           `json:"frequency,omitempty"`
	Id               int64            `json:"id,omitempty"`
	ReportName       string           `json:"reportName,omitempty"`
	SendTime         time.Time        `json:"sendTime,omitempty"`
	Subject          string           `json:"subject,omitempty"`
	SystemId         int64            `json:"systemId,omitempty"`
	Text             string           `json:"text,omitempty"`
	TimeRange        string           `json:"timeRange,omitempty"`
	UserPermissions  *UserPermissions `json:"userPermissions,omitempty"`
}
