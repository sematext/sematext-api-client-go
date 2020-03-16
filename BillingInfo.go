package api

import (
	"fmt"
)

// BillingInfo TODO Doc Comment
type BillingInfo struct {
	CreditCardID  int64  `json:"creditCardId"`
	PaymentMethod string `json:"paymentMethod"`
	PlanID        int64  `json:"planId"`
}

// Persist TODO Doc comment
func (billingInfo *BillingInfo) Persist(id string, client *Client) error {

	path := fmt.Sprintf("/users-web/api/v3/billing/info/%s", id)

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
