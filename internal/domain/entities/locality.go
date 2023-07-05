package entities

type Locality struct {
	ID           int    `json:"id" binding:"required"`
	CountryName  string `json:"country_name" binding:"required"`
	ProvinceName string `json:"province_name" binding:"required"`
	LocalityName string `json:"locality_name" binding:"required"`
}
