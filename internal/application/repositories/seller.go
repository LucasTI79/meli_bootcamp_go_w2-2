package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// SellersRepository encapsulates the storage of a Seller.
type SellerRepository interface {
	GetAll(ctx context.Context) ([]entities.Seller, error)
	Get(ctx context.Context, id int) (*entities.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s entities.Seller) (int, error)
	Update(ctx context.Context, s entities.Seller) error
	Delete(ctx context.Context, id int) error
}
