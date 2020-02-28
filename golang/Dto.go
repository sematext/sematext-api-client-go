package api

// TODO JIRA dto on update app doesn't have a model in documentation

import (
	"fmt"
)

// Dto TODO Doc Comment
type Dto struct {
	Description        string `json:"description,omitempty"`
	IgnorePercentage   int    `json:"ignorePercentage,omitempty"`
	MaxEvents          int    `json:"maxEvents,omitempty"`
	MaxLimitMB         int    `json:"maxLimitMB,omitempty"`
	Name               string `json:"name,omitempty"`
	Sampling           bool   `json:"sampling,omitempty"`
	SamplingPercentage int    `json:"samplingPercentage,omitempty"`
	Staggering         bool   `json:"staggering,omitempty"`
	Status             string `json:"status,omitempty"`
}

// Update TODO Doc comment
func (dto *Dto) Update(id string, client *Client, packet Dto) (*App, error) {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	genericAPIResponse, err := client.PutJSON(path, packet)
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
func (dto *Dto) IsValid() bool {

	return false // TODO implement field check
}
