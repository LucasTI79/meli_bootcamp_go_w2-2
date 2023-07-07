package dtos

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type CreateBuyerRequestDTO struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
}

func (dto *CreateBuyerRequestDTO) ToDomain() *entities.Buyer {
	return &entities.Buyer{
		CardNumberID: dto.CardNumberID,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
	}
}
