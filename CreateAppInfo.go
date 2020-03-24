package api

import "encoding/json"

// CreateAppInfo TODO Doc Comment
type CreateAppInfo struct {
	AppType       string      `json:"appType"`
	DiscountCode  string      `json:"discountCode,omitempty"`
	InitialPlanID int         `json:"initialPlanId"`
	MetaData      AppMetadata `json:"metaData,omitempty"`
	Name          string      `json:"name"`
}

// Persist TODO Doc comment
func (createAppInfo *CreateAppInfo) Persist(client *Client) (*App, error) {

	path := "/spm-reports/api/v3/apps"

	genericAPIResponse, err := client.PostJSON(path, createAppInfo)
	if err != nil {
		return nil, err
	}

	var extract interface{}
	extract, err = json.Marshal(genericAPIResponse.Data.Apps)

	var apps []App

	err = json.Unmarshal(extract.([]byte), &apps) // TODO check this conversion
	if err != nil {
		return nil, err
	}

	return &apps[0], nil
}

// IsValid TODO Doc comment
func (createAppInfo *CreateAppInfo) IsValid() bool {

	// TODO - test minimum viable

	return true
}
