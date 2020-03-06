package api

import "fmt"

// UpdateAppInfo TODO Doc Comment
type UpdateAppInfo struct {
	Description        string `json:"description,omitempty"`
	IgnorePercentage   int    `json:"ignorePercentage,omitempty"`
	MaxEvents          int    `json:"maxEvents,omitempty"`
	MaxLimitMB         int    `json:"maxLimitMB,omitempty"`
	Name               string `json:"name,omitempty"`
	Sampling           bool   `json:"sampling,omitempty"`
	SamplingPercentage int    `json:"samplingPercentage,omitempty"`
	Staggering         bool   `json:"staggering,omitempty"`
	Status             string `json:"status,omitempty"` // TODO example: ACTIVE  Enum: [ ACTIVE, DISABLED ]
}

// Persist TODO Doc comment
func (updateAppInfo *UpdateAppInfo) Persist(id string, client *Client) (*App, error) {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	genericAPIResponse, err := client.PutJSON(path, updateAppInfo)
	if err != nil {
		return nil, err
	}
	var app *App
	app, err = genericAPIResponse.ExtractApp(id)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// IsValid TODO Doc comment
func (updateAppInfo *UpdateAppInfo) IsValid() (valid bool, err error) {

	// TODO - test minimum viable

	return true, nil
}
