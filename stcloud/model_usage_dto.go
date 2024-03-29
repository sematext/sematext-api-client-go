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

type UsageDto struct {
	DailyUsage []DailyDto `json:"dailyUsage,omitempty"`
	DailyVolumeMB int64 `json:"dailyVolumeMB,omitempty"`
	End time.Time `json:"end,omitempty"`
	FailedCount int64 `json:"failedCount,omitempty"`
	IngestedCount int64 `json:"ingestedCount,omitempty"`
	IngestedVolume int64 `json:"ingestedVolume,omitempty"`
	LimitChangeEvents []LimitChangeEventDto `json:"limitChangeEvents,omitempty"`
	MaxAllowedMB int64 `json:"maxAllowedMB,omitempty"`
	MaxLimitMB int64 `json:"maxLimitMB,omitempty"`
	Start time.Time `json:"start,omitempty"`
	StoredCount int64 `json:"storedCount,omitempty"`
	StoredVolume int64 `json:"storedVolume,omitempty"`
	VolumeChangeEvents []LimitChangeEventDto `json:"volumeChangeEvents,omitempty"`
}
