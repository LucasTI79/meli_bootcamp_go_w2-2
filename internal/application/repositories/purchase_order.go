package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type PurchaseOrderRepository interface {
	GetAll(ctx context.Context) ([]entities.PurchaseOrder, error)
	Get(ctx context.Context, id int) (entities.PurchaseOrder, error)
	Exists(ctx context.Context, id int) bool
	Save(ctx context.Context, purchaseOrder entities.PurchaseOrder) (int, error)
	Update(ctx context.Context, purchaseOrder entities.PurchaseOrder) error
	Delete(ctx context.Context, id int) error
	CountByBuyerID(ctx context.Context, buyerID int) (int, error)
}
