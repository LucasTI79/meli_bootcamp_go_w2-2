package dtos

type GetNumberOfRecordsResponseDTO struct {
	ProductID    int    `json:"product_id"`
	Description  string `json:"description"`
	RecordsCount int    `json:"records_count"`
}
