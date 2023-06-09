package buyer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
)

type Service interface {
	Get(id int) (*domain.Buyer, error)
	GetAll() (*[]domain.Buyer, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (service *service) Get(id int) (*domain.Buyer, error) {
	// TODO: Remove this
	ctx := context.TODO()

	// Buscar o buyer pelo ID e trata o erro retornado pelo repository
	buyer, err := service.repository.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &buyer, nil
}

func (service *service) GetAll() (*[]domain.Buyer, error) {
	// TODO: Remove this
	ctx := context.TODO()

	buyers := make([]domain.Buyer, 0)

	buyers, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &buyers, nil

}
