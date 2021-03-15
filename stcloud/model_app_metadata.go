/*
 * Sematext Cloud API
 *
 * API Explorer provides access and documentation for Sematext REST API. The REST API requires the API Key to be sent as part of `Authorization` header. E.g.: `Authorization : apiKey e5f18450-205a-48eb-8589-7d49edaea813`.
 *
 * API version: v3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stcloud

type AppMetadata struct {
	AwsCloudWatchAccessKey string `json:"awsCloudWatchAccessKey,omitempty"`
	AwsCloudWatchSecretKey string `json:"awsCloudWatchSecretKey,omitempty"`
	AwsFetchFrequency      string `json:"awsFetchFrequency,omitempty"`
	AwsRegion              string `json:"awsRegion,omitempty"`
	// Comma separated list of AWS types monitored by created app
	SubTypes []string `json:"subTypes,omitempty"`
}
