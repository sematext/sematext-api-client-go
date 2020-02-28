package api

// BillingInfo TODO Doc Comment
type BillingInfo struct {
	CreditCardID  int64  `json:"creditCardId"`
	PaymentMethod string `json:"paymentMethod"`
	PlanID        int64  `json:"planId"`
}
