package pagarme

import (
	"fmt"
	"encoding/json"
)

//Plan struct
type Plan struct {
	ApiKey					string	`json:"api_key"`
	Name						string	`json:"name,omitempty"`
	Amount					int32		`json:"amount,omitempty"`
	Days						int32		`json:"days,omitempty"`
	TrialDays				int32		`json:"trial_days,omitempty"`
	PaymentMethods	string	`json:"payment_methods"`
	Charges					int32		`json:"charges,omitempty"`
	Installments		int32		`json:"installments,omitempty"`
}

//NewPlan return new Plan struct with default values
func NewPlan(apiKey string) (p Plan) {
	p.ApiKey = apiKey
	p.PaymentMethods = "credit_card"

	return
}

//Create create a new plan
func (p *Plan) Create() (Response, error) {
	body, err := json.Marshal(p)
	if err != nil {
		return resp, err
	}

	return Do("POST", "plans", body)
}
