package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type LocalityRepository interface {
	GetAll(ctx context.Context) ([]entities.Locality, error)
	Get(ctx context.Context, id int) (entities.Locality, error)
	Exists(ctx context.Context, id int) bool
	Save(ctx context.Context, locality entities.Locality) (int, error)
	Update(ctx context.Context, locality entities.Locality) error
	Delete(ctx context.Context, id int) error
	CountSellers(ctx context.Context, id int) (int, error)
}
