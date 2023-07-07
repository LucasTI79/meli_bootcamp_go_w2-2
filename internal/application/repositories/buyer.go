package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type BuyerRepository interface {
	GetAll(ctx context.Context) ([]entities.Buyer, error)
	Get(ctx context.Context, id int) (entities.Buyer, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b entities.Buyer) (int, error)
	Update(ctx context.Context, b entities.Buyer) error
	Delete(ctx context.Context, id int) error
}
