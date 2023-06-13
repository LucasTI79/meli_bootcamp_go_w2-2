package domain

type Buyer struct {
	ID           int    `json:"id" extensions:"x-order=0"`
	CardNumberID string `json:"card_number_id" extensions:"x-order=1"`
	FirstName    string `json:"first_name" extensions:"x-order=2"`
	LastName     string `json:"last_name" extensions:"x-order=3"`
}
