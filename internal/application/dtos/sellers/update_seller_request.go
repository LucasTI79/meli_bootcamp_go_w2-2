package dtos

type UpdateSellerRequestDTO struct {
	CID         *int    `json:"cid"`
	CompanyName *string `json:"company_name"`
	Address     *string `json:"address"`
	Telephone   *string `json:"telephone"`
	LocalityID  *string `json:"locality_id"`
}
