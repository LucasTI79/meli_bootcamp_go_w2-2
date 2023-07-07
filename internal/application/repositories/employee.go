package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type EmployeeRepository interface {
	GetAll(ctx context.Context) ([]entities.Employee, error)
	Get(ctx context.Context, id int) (entities.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e entities.Employee) (int, error)
	Update(ctx context.Context, e entities.Employee) error
	Delete(ctx context.Context, id int) error
}
