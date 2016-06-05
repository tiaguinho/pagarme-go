package pagarme

import (
	"fmt"
	"encoding/json"
)

//Transaction struct
type Transaction struct {
	ApiKey					string						`json:"api_key"`
	Amount					int32							`json:"amount"`
	CardHash				string						`json:"card_hash"`
	Installments		int32							`json:"installments,omitempty"`
	PaymentMethod		string						`json:"payment_method"`
	PostbackURL			string						`json:"postback_url,omitempty"`
	SoftDescriptor	string						`json:"soft_descriptor,omitempty"`
	Capture					bool							`json:"capture,omitempty"`
	Metadata				map[string]string `json:"metadata,omitempty"`
}

//NewTransaction returns a new Transaction struct with default values
func NewTransaction(apiKey string) (t Transaction) {
	t.Installments	= 1
	t.PaymentMethod = "credit_card"
	t.Capture				= true

	return
}

//Card execute a new request to transactions
//with card_hash string
func (t *Transaction) Card(amount int32, cardHash string) (resp Response, error) {
	t.Amount	 = amount
	t.CardHash = cardHash

	body, err := json.Marshal(t)
	if err != nil {
		return resp, err
	}

	return Do("POST", "transactions", body)
}

//Boleto execute a new request to transactions
//with payment_method = boleto
func (t *Transaction) Boleto(amount int32) (resp Response, error) {
	t.Amount = amount
	t.PaymentMethod = "boleto"

	body, err := json.Marshal(t)
	if err != nil {
		return resp, err
	}

	return Do("POST", "transactions", body)
}

//TODO split transaction
//TODO antifraude

//Capture execute new request to capture the transaction
func (t *Transaction) Capture(amount int32, token string) (resp Response, error) {
	r.Amount = amount

	body, err := json.Marshal(t)
	if err != nil {
		return resp, err
	}

	return Do("POST", fmt.Sprintf("transactions/%s/capture", token), body)
}

//Cancel execute new request to cancel one transaction
func (t *Transaction) Cancel(id string) (resp Response, error) {
	body, err := json.Marshal(t)
	if err != nil {
		return resp, err
	}

	return Do("POST", fmt.Sprintf("transactions/%s/refund", id), nil)
}
