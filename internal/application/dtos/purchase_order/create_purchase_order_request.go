package dtos

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

type CreatePurchaseOrderRequestDTO struct {
	OrderNumber     string `json:"order_number"  binding:"required"`
	OrderDate       string `json:"order_date"  binding:"required"`
	TrackingCode    string `json:"tracking_code"  binding:"required"`
	BuyerID         int    `json:"buyer_id"  binding:"required"`
	CarrierID       int    `json:"carrier_id"  binding:"required"`
	OrderStatusID   int    `json:"order_status_id"  binding:"required"`
	WarehouseID     int    `json:"warehouse_id"  binding:"required"`
	ProductRecordID int    `json:"product_record_id"  binding:"required"`
}

func (dto *CreatePurchaseOrderRequestDTO) ToDomain() domain.PurchaseOrder {
	return domain.PurchaseOrder{
		OrderNumber:     dto.OrderNumber,
		OrderDate:       dto.OrderDate,
		TrackingCode:    dto.TrackingCode,
		BuyerID:         dto.BuyerID,
		CarrierID:       dto.CarrierID,
		OrderStatusID:   dto.OrderStatusID,
		WarehouseID:     dto.WarehouseID,
		ProductRecordID: dto.ProductRecordID,
	}
}
