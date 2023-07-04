package domain

type ProductRecord struct {
	ID             int     `json:"id"`
	LastUpdateRate string  `json:"lastUpdateRate"`
	PurchasePrice  float32 `json:"purchasePrice"`
	SalePrice      float32 `json:"salePrice"`
	ProductId      int     `json:"productId"`
}
