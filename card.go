package pagarme

//Card struct
type Card struct {
	Number					string `json:"card_number,omitempty"`
	HolderName			string `json:"card_holder_name,omitempty"`
	ExpirationDate	string `json:"card_expiration_date,omitempty"`
	Cvv							string `json:"card_cvvi,omitempty"`
	Hash						string `json:"card_hash"`
}

//GenererateHash generate hash for the card
func (c *Card) GenerateHash() {
	
}
