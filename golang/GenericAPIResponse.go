package api

import "fmt"

// GenericAPIResponse TODO Doc Comment
type GenericAPIResponse struct {
	Errors  []Error `json:"errors"`
	Message string  `json:"message"`
	Success bool    `json:"success"`
	Data    struct {
		Apps []App `json:"apps,omitempty"`
		App  App   `json:"app,omitempty"`
	} `json:"data"`
}

// TODO JIRA TASK - GET App by ID
// TODO JIRA TASK - DELETE App by ID
// TODO JIRA TASK - RESET App STATUS to archived or similar.

// ExtractAppByID - TODO Doc Comment
func (genericAPIResponse *GenericAPIResponse) ExtractAppByID(id string) (*App, error) {

	if &genericAPIResponse.Data == nil {
		return nil, fmt.Errorf("Missing Data field")
	}

	if &genericAPIResponse.Data.Apps == nil {
		return nil, fmt.Errorf("Missing Apps field")
	}

	if len(genericAPIResponse.Data.Apps) == 0 {
		return nil, nil
	}

	for i := range genericAPIResponse.Data.Apps {
		if genericAPIResponse.Data.Apps[i].ID == id {
			app := genericAPIResponse.Data.Apps[i]
			return &app, nil
		}
	}

	return nil, fmt.Errorf("App %s not found", id)

}

// ExtractApp - TODO Doc Comment
func (genericAPIResponse *GenericAPIResponse) ExtractApp(id string) (*App, error) {

	if &genericAPIResponse.Data.App != nil {
		return nil, fmt.Errorf("Missing App field")
	}

	return &genericAPIResponse.Data.App, nil

}
