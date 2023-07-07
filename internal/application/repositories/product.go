package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// ProductRepository encapsulates the storage of a Product.
type ProductRepository interface {
	GetAll(ctx context.Context) ([]entities.Product, error)
	Get(ctx context.Context, id int) (entities.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p entities.Product) (int, error)
	Update(ctx context.Context, p entities.Product) error
	Delete(ctx context.Context, id int) error
}
