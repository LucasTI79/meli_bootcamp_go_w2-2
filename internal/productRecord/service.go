package productRecord

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors.
var (
	ErrNotFound = errors.New("productRecord not found")
	ErrConflict = errors.New("productRecord with ProductRecord Number already exists")
)

type Service interface {
	Save(ctx *context.Context, lastUpdateRate string, purchasePrice, salePrice float32, productId int) (*domain.ProductRecord, error)
	GetAll(ctx *context.Context) (*[]domain.ProductRecord, error)
	Get(ctx *context.Context, id int) (*domain.ProductRecord, error)
	Delete(ctx *context.Context, id int) error
	Update(ctx *context.Context, lastUpdateRate *string, purchasePrice, salePrice *float32, productId *int, id int) (*domain.ProductRecord, error)
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

func (s *service) Get(ctx *context.Context, id int) (*domain.ProductRecord, error) {
	productRecord, err := s.productRecordsRepository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &productRecord, nil
}

func (s *service) Save(ctx *context.Context, lastUpdateRate string, purchasePrice, salePrice float32, productId int) (*domain.ProductRecord, error) {
	existingProductRecord := s.productRecordsRepository.Exists(*ctx, productId)

	if existingProductRecord {
		return nil, ErrConflict
	}

	newProductRecord := domain.ProductRecord{
		LastUpdateRate: lastUpdateRate,
		PurchasePrice:  purchasePrice,
		SalePrice:      salePrice,
		ProductId:      productId,
	}

	productRecordId, err := s.productRecordsRepository.Save(*ctx, newProductRecord)
	if err != nil {
		return nil, err

	}

	savedProductRecord, err := s.productRecordsRepository.Get(*ctx, productRecordId)
	if err != nil {
		return nil, err
	}

	return &savedProductRecord, nil
}

func (s *service) Update(ctx *context.Context, lastUpdateRate *string, purchasePrice, salePrice *float32, productId *int, id int) (*domain.ProductRecord, error) {
	existingProductRecord, err := s.productRecordsRepository.Get(*ctx, id)
	if err != nil {
		return nil, err
	}

	if lastUpdateRate != nil {
		existingProductRecord.LastUpdateRate = *lastUpdateRate
	}
	if purchasePrice != nil {
		existingProductRecord.PurchasePrice = *purchasePrice
	}
	if salePrice != nil {
		existingProductRecord.SalePrice = *salePrice
	}

	if productId != nil {
		existingProductRecordSearch := s.productRecordsRepository.Exists(*ctx, *productId)
		if existingProductRecordSearch && *productId != existingProductRecord.ProductId {
			return nil, ErrConflict
		}
		existingProductRecord.ProductId = *productId
	}

	err1 := s.productRecordsRepository.Update(*ctx, existingProductRecord)
	if err1 != nil {
		switch err1 {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err1
		}
	}

	return &existingProductRecord, nil
}

func (s *service) Delete(ctx *context.Context, id int) error {
	err := s.productRecordsRepository.Delete(*ctx, id)
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
