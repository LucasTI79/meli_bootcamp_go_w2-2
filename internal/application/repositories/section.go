package repositories

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// SectionRepository encapsulates the storage of a section.
type SectionRepository interface {
	GetAll(ctx context.Context) ([]entities.Section, error)
	Get(ctx context.Context, id int) (entities.Section, error)
	Exists(ctx context.Context, sectionNumber int) bool
	Save(ctx context.Context, s entities.Section) (int, error)
	Update(ctx context.Context, s entities.Section) error
	Delete(ctx context.Context, id int) error
}
