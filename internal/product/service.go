package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
	ErrConflict = errors.New("section with Product Number already exists")
)

type Service interface {
	Save(ctx context.Context, description string, expiration_rate, freezing_rate int, height, length, netweight float32, product_code string,
		recommended_freezing_temperature, width float32, product_type_id, seller_id int) (*domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (*domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	productRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		productRepository: r,
	}
}

func (s *service) Save(ctx context.Context, description string, expiration_rate, freezing_rate int, height, length, netweight float32, product_code string,
	recommended_freezing_temperature, width float32, product_type_id, seller_id int) (*domain.Product, error) {
	existingProduct := s.productRepository.Exists(ctx, description)

	if existingProduct {
		return nil, ErrConflict
	}

	newProduct := domain.Product{
		Description:    description,
		ExpirationRate: expiration_rate,
		FreezingRate:   freezing_rate,
		Height:         height,
		Length:         length,
		Netweight:      netweight,
		ProductCode:    product_code,
		RecomFreezTemp: recommended_freezing_temperature,
		Width:          width,
		ProductTypeID:  product_type_id,
		SellerID:       seller_id,
	}

	productId, err := s.productRepository.Save(ctx, newProduct)
	if err != nil {
		return nil, err

	}

	savedProduct, err := s.productRepository.Get(ctx, productId)
	if err != nil {
		return nil, err
	}

	return &savedProduct, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Get(ctx context.Context, id int) (*domain.Product, error) {
	product, err := s.productRepository.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &product, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.productRepository.Delete(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}
	return nil
}
