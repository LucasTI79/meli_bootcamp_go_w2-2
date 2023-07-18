package dtos

type DataLocalityAndCarrier struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	CountCarrier int    `json:"count_carrier"`
}
