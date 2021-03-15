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

type Subscription struct {
	Addresses       string           `json:"addresses,omitempty"`
	CreatedBy       string           `json:"createdBy,omitempty"`
	DashboardID     int64            `json:"dashboardId,omitempty"`
	DashboardName   string           `json:"dashboardName,omitempty"`
	Enabled         bool             `json:"enabled,omitempty"`
	Filters         string           `json:"filters,omitempty"`
	Frequency       string           `json:"frequency,omitempty"`
	ID              int64            `json:"id,omitempty"`
	NextSendDate    time.Time        `json:"nextSendDate,omitempty"`
	Subject         string           `json:"subject,omitempty"`
	SystemID        int64            `json:"systemId,omitempty"`
	SystemName      string           `json:"systemName,omitempty"`
	Text            string           `json:"text,omitempty"`
	TimeRange       string           `json:"timeRange,omitempty"`
	UserPermissions *UserPermissions `json:"userPermissions,omitempty"`
}
