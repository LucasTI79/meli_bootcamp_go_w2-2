package productbatches

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

var (
	ErrNotFound = errors.New("product Batches not found")
	ErrConflict = errors.New("product batches with batch_number already exists")
)

type IService interface {
	Get(ctx *context.Context, id int) (*domain.ProductBatches, error)
	Save(ctx *context.Context, product domain.ProductBatches) (*domain.ProductBatches, error)
}
type Service struct {
	productRepository IRepository
}

func NewService(r IRepository) IService {
	return &Service{
		productRepository: r,
	}
}

func (s *Service) Get(ctx *context.Context, id int) (*domain.ProductBatches, error) {
	productBatches, err := s.productRepository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &productBatches, nil
}

func (s *Service) Save(ctx *context.Context, product domain.ProductBatches) (*domain.ProductBatches, error) {
	batchNumberExist := s.productRepository.Exists(*ctx, product.BatchNumber)
	if batchNumberExist {
		return nil, ErrConflict
	}

	productBatchesID, err := s.productRepository.Save(*ctx, product)
	if err != nil {
		return nil, err
	}
	savedProductBatches, err := s.productRepository.Get(*ctx, productBatchesID)
	if err != nil {
		return nil, err
	}
	return &savedProductBatches, nil
}
