package pagarme

import (
	"encoding/json"
)

//Response struct
type Response struct {
	Object				string `json:"object"`
	ID						string `json:"id"`
	Event					string `json:"event"`
	CurrentStatus string `json:"current_status"`
	OldStatus			string `json:"old_status"`
	DesiredStatus string `json:"desired_status"`
}

//Cancel execute new request to cancel the transaction
func (r *Response) Cancel() error {
	
}
