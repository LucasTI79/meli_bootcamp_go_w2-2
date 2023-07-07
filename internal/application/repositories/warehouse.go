package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// WarehouseRepository encapsulates the storage of a warehouses.
type WarehouseRepository interface {
	GetAll(ctx context.Context) ([]entities.Warehouse, error)
	Get(ctx context.Context, id int) (entities.Warehouse, error)
	Exists(ctx context.Context, warehouseCode string) bool
	Save(ctx context.Context, w entities.Warehouse) (int, error)
	Update(ctx context.Context, w entities.Warehouse) error
	Delete(ctx context.Context, id int) error
}
