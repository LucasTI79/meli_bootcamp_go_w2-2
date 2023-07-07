package dtos

type DataLocalityAndCarrier struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name"`
	CountCarrier string `json:"count_carrier"`
}
