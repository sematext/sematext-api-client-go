package api

import "fmt"

// UpdateAppInfo TODO Doc Comment
type UpdateAppInfo struct {
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"` // TODO example: ACTIVE  Enum: [ ACTIVE, DISABLED ]
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
