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

// Invitation TODO Godoc comment
type Invitation struct {

	// For invite request, only app.ID needs to be set.
	App *App `json:"app,omitempty"`

	// For invite request, only apps.ID needs to be set.
	Apps          []App     `json:"apps,omitempty"`
	ID            int64     `json:"id,omitempty"`
	InviteDate    time.Time `json:"inviteDate,omitempty"`
	InviteStatus  string    `json:"inviteStatus,omitempty"`
	InviteeEmail  string    `json:"inviteeEmail,omitempty"`
	InviteeRole   string    `json:"inviteeRole,omitempty"`
	InviteeStatus string    `json:"inviteeStatus,omitempty"`
	InviterEmail  string    `json:"inviterEmail,omitempty"`
	UUID          string    `json:"uuid,omitempty"`
}