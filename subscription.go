package pagarme

import (
	"fmt"
	"encoding/json"
)

//Subscription struct
type Subscription struct {
	ApiKey				string	 `json:"api_key"`
	PaymentMethod	string	 `json:"payment_method"`
	CardHash			string	 `json:"card_hash"`
	PostbackURL		string	 `json:"postback_url"`
	Customer			Customer `json:"customer"`
}

//NewSubscription return new Subscription struct
func NewSubscription(apiKey string) (s Subscription) {
	s.ApiKey = apiKey
	s.PaymentMethod = "credit_card"

	return
}

//Create create a new subscription
func (s *Subscription) Create() (Response, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return resp, err
	}

	return Do("POST", "subscriptions", body)
}

//Cancel cancel a subscription
func (s *Subscription) Cancel(id string) (Response, error) {
	body, err := json.Marshal(s)
	if err != nil {
		return resp, err
	}

	return Do("POST", fmt.Sprintf("subscriptions/%s/cancel", id), body)
}
