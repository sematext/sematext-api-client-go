package api

import (
	"fmt"
)

// BillingInfo TODO Doc Comment
type BillingInfo struct {
	PlanID int `json:"planId"`
}

// Persist TODO Doc comment
func (billingInfo *BillingInfo) Persist(appID int, client *Client) error {

	path := fmt.Sprintf("/users-web/api/v3/billing/info/%d", appID)

	_, err := client.PutJSON(path, billingInfo)
	if err != nil {
		return err // TODO Extract information from genericAPIResponse
	}

	// TODO Check http status and other return errors.

	return nil
}

// IsValid TODO Doc comment
func (billingInfo *BillingInfo) IsValid() (valid bool, err error) {

	// TODO - test minimum viable

	return true, nil
}
