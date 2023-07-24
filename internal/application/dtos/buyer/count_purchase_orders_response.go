package dtos

type GetNumberOfPurchaseOrdersByBuyerResponseDTO struct {
	BuyerID int `json:"buyer_id"`
	//LocalityName string `json:"locality_name"`
	PurchaseOrdersCount int `json:"purchase_orders_count"`
}
