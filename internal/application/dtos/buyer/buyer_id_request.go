package dtos

type BuyerIDRequestDTO struct {
	ID int `uri:"id" binding:"required"`
}
