package productRecord

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("productRecord not found")
	ErrConflict = errors.New("productRecord with ProductRecord Number already exists")
)

type Service interface {
	//	Save(ctx *context.Context, description string, expiration_rate, freezing_rate int, height, length, netweight float32, product_code string,
	//		recommended_freezing_temperature, width float32, product_type_id, seller_id int) (*domain.Product, error)
	GetAll(ctx *context.Context) (*[]domain.ProductRecord, error)
	// Get(ctx *context.Context, id int) (*domain.Product, error)
	// Delete(ctx *context.Context, id int) error
	// Update(ctx *context.Context, description *string, expiration_rate, freezing_rate *int, height, length, netweight *float32, product_code *string,
	//
	//	recommended_freezing_temperature, width *float32, product_type_id, seller_id *int, id int) (*domain.Product, error)
}

type service struct {
	productRecordsRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		productRecordsRepository: r,
	}
}

func (s *service) GetAll(ctx *context.Context) (*[]domain.ProductRecord, error) {
	productsRecords, err := s.productRecordsRepository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}

	return &productsRecords, nil
}
