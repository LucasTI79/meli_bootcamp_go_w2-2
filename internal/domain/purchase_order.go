package domain

type PurchaseOrder struct {
	ID              int    `json:"id" extensions:"x-order=0"`
	OrderNumber     string `json:"order_number" extensions:"x-order=1"`
	OrderDate       string `json:"order_date" extensions:"x-order=2" format:"2006-01-02"`
	TrackingCode    string `json:"tracking_code" extensions:"x-order=3"`
	BuyerID         int    `json:"buyer_id" extensions:"x-order=4"`
	CarrierID       int    `json:"carrier_id" extensions:"x-order=5"`
	OrderStatusID   int    `json:"order_status_id" extensions:"x-order=6"`
	WarehouseID     int    `json:"warehouse_id" extensions:"x-order=7"`
	ProductRecordID int    `json:"product_record_id" extensions:"x-order=8"`
}
