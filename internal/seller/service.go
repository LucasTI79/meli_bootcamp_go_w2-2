package seller

import (
	"context"
	"errors"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrConflict = errors.New("409 Conflict: Seller with CID already exists")
)

type Service interface {
	GetAll(ctx context.Context) (*[]domain.Seller, error)
	Get(ctx context.Context, id int) (*domain.Seller, error)
	Save(ctx context.Context, seller domain.Seller) (*domain.Seller, error)
	Update(ctx context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*domain.Seller, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	sellerRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		sellerRepository: r,
	}
}

func (s *service) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)

	sellers, err := s.sellerRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &sellers, nil
}

func (s *service) Get(ctx context.Context, id int) (*domain.Seller, error) {
	seller, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *service) Save(ctx context.Context, seller domain.Seller) (*domain.Seller, error) {
	existingSeller := s.sellerRepository.Exists(ctx, seller.CID)
	if existingSeller {
		return nil, ErrConflict
	}

	id, err := s.sellerRepository.Save(ctx, seller)
	if err != nil {
		return nil, err
	}

	seller.ID = id

	return &seller, nil
}

func (s *service) Update(ctx context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*domain.Seller, error) {
	existingSeller, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if updateSellerRequest.CID != nil {
		existingSellerSearch := s.sellerRepository.Exists(ctx, *updateSellerRequest.CID)
		if existingSellerSearch && *updateSellerRequest.CID != existingSeller.CID {
			return nil, ErrConflict
		}
		existingSeller.CID = *updateSellerRequest.CID
	}
	if updateSellerRequest.CompanyName != nil {
		existingSeller.CompanyName = *updateSellerRequest.CompanyName
	}
	if updateSellerRequest.Address != nil {
		existingSeller.Address = *updateSellerRequest.Address
	}
	if updateSellerRequest.Telephone != nil {
		existingSeller.Telephone = *updateSellerRequest.Telephone
	}

	err1 := s.sellerRepository.Update(ctx, *existingSeller)
	if err1 != nil {
		return nil, err1
	}

	return existingSeller, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.sellerRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	err1 := s.sellerRepository.Delete(ctx, id)
	if err1 != nil {
		return err
	}

	return nil
}
