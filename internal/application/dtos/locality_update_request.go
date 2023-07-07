package dtos

type UpdateLocalityRequestDTO struct {
	CountryName  *string `json:"country_name"`
	ProvinceName *string `json:"province_name"`
	LocalityName *string `json:"locality_name"`
}
