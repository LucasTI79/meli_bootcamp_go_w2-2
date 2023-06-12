package dtos

type UpdateBuyerRequestDTO struct {
	ID int `uri:"id" binding:"required"`
	// TODO: Validate if has at least one field to update
	CardNumberID *string `json:"card_number_id" binding:"-"`
	FirstName    *string `json:"first_name" binding:"-"`
	LastName     *string `json:"last_name" binding:"-"`
}
