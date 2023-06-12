package dtos

import "github.com/go-playground/validator/v10"

type UpdateBuyerRequestDTO struct {
	ID           int
	CardNumberID *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
}

// Check if it has at least one field to update
func UpdateBuyerRequestValidation(sl validator.StructLevel) {
	dto := sl.Current().Interface().(UpdateBuyerRequestDTO)
	if dto.CardNumberID == nil && dto.FirstName == nil && dto.LastName == nil {
		sl.ReportError(dto.CardNumberID, "CardNumberID", "card_number_id", "atleastone", "")
		sl.ReportError(dto.FirstName, "FistName", "first_name", "atleastone", "")
		sl.ReportError(dto.LastName, "LastName", "last_name", "atleastone", "")
	}
}
