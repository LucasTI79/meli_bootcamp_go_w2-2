package dtos

type GetNumberOfSellersResponseDTO struct {
	LocalityID   string `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}
