package domain

type Locality struct {
	ID           int    `json:"id" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
}
