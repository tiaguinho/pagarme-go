package pagarme

import (
	"fmt"
	"encoding/json"
)

//Recipient struct
type Recipient struct {
	ApiKey						string `json:"api_key"`
	TransferInterval	string `json:"transfer_interval"`
	TransferDay				string `json:"transfer_day,omitempty"`
	TransferEnabled		bool	 `json:"trnasfer_enabled"`
	BankAccountId			string `json:"bank_account_id"`
}

//NewRecipient returns a new Recipient struct with default values
func NewRecipient(apiKey string) (r Recipient) {
	r.ApiKey = apiKey
	r.TransferEnabled = true

	return
}

//Create create new recipient
func (r *Recipient) Create() (Response, error) {
	body, err := json.Marshal(r)
	if err != nil {
		return resp, err
	}

	return Do("POST", "recipients", body)
}
