package api

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
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
	ID                    int                  `json:"id"`
	Integration           ServiceIntegration   `json:"integration"` // TODO - Check integration for inclusion into resource
	LastDataReceivedDate  int64                `json:"lastDataReceivedDate"`
	LastDataSavedDate     int64                `json:"lastDataSavedDate"`
	LoggedInUserAppRole   string               `json:"loggedInUserAppRole"`
	MonthlyInvoiceAccount bool                 `json:"monthlyInvoiceAccount"`
	Name                  string               `json:"name"`
	OwnerEmail            string               `json:"ownerEmail"` // TODO - Check owner email present for inclusion into resource
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

// Load TODO Doc comment
func (app *App) Load(id int, client *Client) (*App, error) {

	fmt.Println("---------------------------------------")
	fmt.Println("App.Load Called")
	fmt.Println("---------------------------------------")

	var genericAPIResponse *GenericAPIResponse
	var err error

	path := fmt.Sprintf("/users-web/api/v3/apps/%d", id)
	fmt.Println("---------------------------------------")
	fmt.Println("path")
	fmt.Println(path)
	fmt.Println("---------------------------------------")

	if genericAPIResponse, err = client.GetJSON(path, nil); err != nil {
		return nil, err
	}

	fmt.Println("---------------------------------------")
	fmt.Println("genericAPIResponse")
	spew.Dump(genericAPIResponse)
	fmt.Println("---------------------------------------")

	if app, err = genericAPIResponse.ExtractApp(); err != nil {
		return nil, err
	}

	fmt.Println("---------------------------------------")
	fmt.Println("app")
	spew.Dump(app)
	fmt.Println("---------------------------------------")

	return app, nil

}

// Exists TODO Doc comment
func (app *App) Exists(id int, client *Client) (bool, error) {

	var err error
	if app, err = app.Load(id, client); err != nil {
		return false, err // TODO Return false on error?
	}
	exists := app != nil && !app.IsArchived()

	return exists, nil

}

// Persist TODO Doc comment
func (app *App) Persist(id int, client *Client, updateAppInfo UpdateAppInfo) (*App, error) {

	var genericAPIResponse *GenericAPIResponse
	var err error

	path := fmt.Sprintf("/spm-reports/api/v3/apps/%d", id)
	if genericAPIResponse, err = client.PutJSON(path, updateAppInfo); err != nil {
		return nil, err
	}

	if app, err = genericAPIResponse.ExtractAppByID(id); err != nil {
		return nil, err
	}

	return app, nil
}

// IsArchived TODO Doc comment
func (app *App) IsArchived() bool {
	return app.Status == "ARCHIVED"
}
