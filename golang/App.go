package api

import (
	"fmt"
)

// App TODO Doc Comment
type App struct {
	AjaxThreshold         int64                `json:"ajaxThreshold,omitempty"`
	AppType               string               `json:"appType"`
	AppTypeID             int64                `json:"appTypeId"`
	CreatorEmail          string               `json:"creatorEmail"`
	CreditCardExpiry      string               `json:"creditCardExpiry"`
	CreditCardNumber      string               `json:"creditCardNumber"`
	Description           string               `json:"description"`
	DisplayStatus         string               `json:"displayStatus"`
	FirstDataSavedDate    int64                `json:"firstDataSavedDate"`
	ID                    string               `json:"id"`
	Integration           ServiceIntegration   `json:"integration"` // TODO - Check integration for inclusion into resource
	LastDataReceivedDate  int64                `json:"lastDataReceivedDate"`
	LastDataSavedDate     int64                `json:"lastDataSavedDate"`
	LoggedInUserAppRole   string               `json:"loggedInUserAppRole"`
	MonthlyInvoiceAccount bool                 `json:"monthlyInvoiceAccount"`
	Name                  string               `json:"name"`
	OwnerEmail            string               `json:"ownerEmail"` // TODO - Check owner email for inclusion into resource
	OwningOrganization    BasicOrganizationDto `json:"owningOrganization"`
	PageLoadThreshold     int64                `json:"pageLoadThreshold"`
	PaymentMethodID       int64                `json:"paymentMethodId"`
	Plan                  Plan                 `json:"plan"`
	PrepaidAccount        bool                 `json:"prepaidAccount"`
	Status                string               `json:"status"` // TODO - Check for undocumented enum
	Token                 string               `json:"token"`
	TrialEndDate          int64                `json:"trialEndDate"`
	URLGroupLimit         int32                `json:"urlGroupLimit"`
	UserRoles             []UserRole           `json:"userRoles"` // TODO - Check userroles for inclusion into resource
}

// Retrieve TODO Doc comment
func (app *App) Retrieve(id string, client *Client) (*App, error) { // TODO Rework once API has a route for GET by id

	path := "/spm-reports/api/v3/apps"
	genericAPIResponse, err := client.GetJSON(path, nil)
	if err != nil {
		return nil, err
	}

	app, err = genericAPIResponse.ExtractAppByID(id)
	if err != nil {
		return nil, err
	}

	return app, nil

}

// Exists TODO Doc comment
func (app *App) Exists(id string, client *Client) (bool, error) { // TODO Rework once API has a route for GET by id

	path := "/spm-reports/api/v3/apps"
	genericAPIResponse, err := client.GetJSON(path, nil)
	if err != nil {
		return false, err
	}

	app, err = genericAPIResponse.ExtractAppByID(id)
	result := app != nil
	return result, nil

}

// Update TODO Doc comment
func (app *App) Update(id string, client *Client, dto Dto) (*App, error) {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	genericAPIResponse, err := client.PutJSON(path, dto)
	if err != nil {
		return nil, err
	}

	app, err = genericAPIResponse.ExtractAppByID(id)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Delete TODO Doc comment
func (app *App) Delete(id string, client *Client) error {

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%s", id)
	err := client.Delete(path)
	if err != nil {
		return err
	}
	return nil
}

// IsArchived TODO Doc comment
func (app *App) IsArchived() bool {
	return app.Status == "ARCHIVED"
}
