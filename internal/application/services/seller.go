package services

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type SellerService interface {
	GetAll(ctx *context.Context) (*[]entities.Seller, error)
	Get(ctx *context.Context, id int) (*entities.Seller, error)
	Save(ctx *context.Context, seller entities.Seller) (*entities.Seller, error)
	Update(ctx *context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*entities.Seller, error)
	Delete(ctx *context.Context, id int) error
}

type sellerService struct {
	sellerRepository repositories.SellerRepository
}

func NewSellerService(r repositories.SellerRepository) SellerService {
	return &sellerService{
		sellerRepository: r,
	}
}

func (s *sellerService) GetAll(ctx *context.Context) (*[]entities.Seller, error) {
	sellers := make([]entities.Seller, 0)

	sellers, err := s.sellerRepository.GetAll(*ctx)

	if err != nil {
		return nil, err
	}
	return &sellers, nil
}

func (s *sellerService) Get(ctx *context.Context, id int) (*entities.Seller, error) {
	seller, err := s.sellerRepository.Get(*ctx, id)
	if err != nil {
		return nil, err
	}

	return seller, nil
}

func (s *sellerService) Save(ctx *context.Context, seller entities.Seller) (*entities.Seller, error) {
	existingSeller := s.sellerRepository.Exists(*ctx, seller.CID)
	if existingSeller {
		return &entities.Seller{}, ErrConflict
	}

	id, err := s.sellerRepository.Save(*ctx, seller)
	if err != nil {
		return &entities.Seller{}, err
	}

	seller.ID = id

	return &seller, nil
}

func (s *sellerService) Update(ctx *context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*entities.Seller, error) {
	existingSeller, err := s.sellerRepository.Get(*ctx, id)
	if err != nil {
		return nil, err
	}

	if updateSellerRequest.CID != nil {
		existingSellerSearch := s.sellerRepository.Exists(*ctx, *updateSellerRequest.CID)
		if existingSellerSearch {
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

	if updateSellerRequest.LocalityID != nil {
		existingSeller.LocalityID = *updateSellerRequest.LocalityID
	}

	err1 := s.sellerRepository.Update(*ctx, *existingSeller)
	if err1 != nil {
		return nil, err1
	}

	return existingSeller, nil
}

func (s *sellerService) Delete(ctx *context.Context, id int) error {
	_, err := s.sellerRepository.Get(*ctx, id)
	if err != nil {
		return err
	}

	err = s.sellerRepository.Delete(*ctx, id)
	if err != nil {
		return err
	}

	return nil
}
