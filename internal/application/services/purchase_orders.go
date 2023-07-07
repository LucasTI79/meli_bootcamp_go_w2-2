package services

import (
	"context"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

//// Errors
//var (
//	ErrNotFound = errors.New("purchaseOrder not found")
//	ErrConflict = errors.New("ID already exists")
//)

type PurchaseOrderService interface {
	Get(ctx *context.Context, id int) (entities.PurchaseOrder, error)
	GetAll(ctx *context.Context) ([]entities.PurchaseOrder, error)
	Create(ctx *context.Context, purchaseOrder entities.PurchaseOrder) (entities.PurchaseOrder, error)
	Update(ctx *context.Context, id int, updatePurchaseOrderRequest dtos.UpdatePurchaseOrderRequestDTO) (entities.PurchaseOrder, error)
	Delete(ctx *context.Context, id int) error
	CountByBuyerID(ctx *context.Context, buyerID int) (int, error)
}

type purchaseOrderService struct {
	purchaseOrderRepository repositories.PurchaseOrderRepository
}

func NewPurchaseOrderService(r repositories.PurchaseOrderRepository) PurchaseOrderService {
	return &purchaseOrderService{
		purchaseOrderRepository: r,
	}
}

func (service *purchaseOrderService) Get(ctx *context.Context, id int) (entities.PurchaseOrder, error) {
	purchaseOrder, err := service.purchaseOrderRepository.Get(*ctx, id)
	if err != nil {
		return entities.PurchaseOrder{}, err
	}

	return purchaseOrder, nil
}

func (service *purchaseOrderService) GetAll(ctx *context.Context) ([]entities.PurchaseOrder, error) {
	localities := make([]entities.PurchaseOrder, 0)

	localities, err := service.purchaseOrderRepository.GetAll(*ctx)
	if err != nil {
		return localities, err
	}

	return localities, nil
}

func (service *purchaseOrderService) Create(ctx *context.Context, purchaseOrder entities.PurchaseOrder) (entities.PurchaseOrder, error) {
	existingPurchaseOrder := service.purchaseOrderRepository.Exists(*ctx, purchaseOrder.ID)
	if existingPurchaseOrder {
		return entities.PurchaseOrder{}, ErrConflict
	}

	id, err := service.purchaseOrderRepository.Save(*ctx, purchaseOrder)
	if err != nil {
		return entities.PurchaseOrder{}, err
	}

	purchaseOrder.ID = id

	return purchaseOrder, nil
}

func (service *purchaseOrderService) Update(ctx *context.Context, id int, updatePurchaseOrderRequest dtos.UpdatePurchaseOrderRequestDTO) (entities.PurchaseOrder, error) {
	existingPurchaseOrder, err := service.purchaseOrderRepository.Get(*ctx, id)
	if err != nil {
		return entities.PurchaseOrder{}, err
	}

	existingPurchaseOrderSearch := service.purchaseOrderRepository.Exists(*ctx, id)
	if existingPurchaseOrderSearch {
		return entities.PurchaseOrder{}, ErrConflict
	}

	if updatePurchaseOrderRequest.OrderNumber != nil {
		existingPurchaseOrder.OrderNumber = *updatePurchaseOrderRequest.OrderNumber
	}

	if updatePurchaseOrderRequest.OrderDate != nil {
		existingPurchaseOrder.OrderDate = *updatePurchaseOrderRequest.OrderDate
	}

	if updatePurchaseOrderRequest.TrackingCode != nil {
		existingPurchaseOrder.TrackingCode = *updatePurchaseOrderRequest.TrackingCode
	}

	if updatePurchaseOrderRequest.BuyerID != nil {
		existingPurchaseOrder.BuyerID = *updatePurchaseOrderRequest.BuyerID
	}

	if updatePurchaseOrderRequest.CarrierID != nil {
		existingPurchaseOrder.CarrierID = *updatePurchaseOrderRequest.CarrierID
	}

	if updatePurchaseOrderRequest.OrderStatusID != nil {
		existingPurchaseOrder.OrderStatusID = *updatePurchaseOrderRequest.OrderStatusID
	}

	if updatePurchaseOrderRequest.WarehouseID != nil {
		existingPurchaseOrder.WarehouseID = *updatePurchaseOrderRequest.WarehouseID
	}

	if updatePurchaseOrderRequest.ProductRecordID != nil {
		existingPurchaseOrder.ProductRecordID = *updatePurchaseOrderRequest.ProductRecordID
	}

	if err = service.purchaseOrderRepository.Update(*ctx, existingPurchaseOrder); err != nil {
		return entities.PurchaseOrder{}, err
	}

	return existingPurchaseOrder, nil
}

func (service *purchaseOrderService) Delete(ctx *context.Context, id int) error {
	_, err := service.purchaseOrderRepository.Get(*ctx, id)
	if err != nil {
		return err
	}

	err = service.purchaseOrderRepository.Delete(*ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *purchaseOrderService) CountByBuyerID(ctx *context.Context, buyerID int) (int, error) {
	count, err := service.purchaseOrderRepository.CountByBuyerID(*ctx, buyerID)
	if err != nil {
		return 0, err
	}

	return count, nil

}
