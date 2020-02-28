package api

// UpdateAppInfo TODO Doc Comment
type UpdateAppInfo struct {
	Description        string `json:"description"`
	IgnorePercentage   int32  `json:"ignorePercentage"`
	MaxEvents          int64  `json:"maxEvents"`
	MaxLimitMB         int64  `json:"maxLimitMB"`
	Name               string `json:"name"`
	Sampling           bool   `json:"sampling"`
	SamplingPercentage int32  `json:"samplingPercentage"`
	Staggering         bool   `json:"staggering"`
	Status             string `json:"status"` // TODO example: ACTIVE  Enum: [ ACTIVE, DISABLED ]
}
