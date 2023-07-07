package services

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type BuyerService interface {
	Get(ctx *context.Context, id int) (*entities.Buyer, error)
	GetAll(ctx *context.Context) (*[]entities.Buyer, error)
	Create(ctx *context.Context, createBuyerRequest *dtos.CreateBuyerRequestDTO) (*entities.Buyer, error)
	Update(ctx *context.Context, id int, updateBuyerRequest *dtos.UpdateBuyerRequestDTO) (*entities.Buyer, error)
	Delete(ctx *context.Context, id int) error
}

type buyerService struct {
	repository repositories.BuyerRepository
}

func NewBuyerService(repository repositories.BuyerRepository) BuyerService {
	return &buyerService{
		repository,
	}
}

func (service *buyerService) Get(ctx *context.Context, id int) (*entities.Buyer, error) {
	buyer, err := service.repository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return &entities.Buyer{}, ErrNotFound
		default:
			return &entities.Buyer{}, err
		}
	}

	return &buyer, nil
}

func (service *buyerService) GetAll(ctx *context.Context) (*[]entities.Buyer, error) {
	buyers := make([]entities.Buyer, 0)

	buyers, err := service.repository.GetAll(*ctx)
	if err != nil {
		return &buyers, err
	}

	return &buyers, nil

}

func (service *buyerService) Create(ctx *context.Context, createBuyerRequest *dtos.CreateBuyerRequestDTO) (*entities.Buyer, error) {
	buyer := createBuyerRequest.ToDomain()

	cardNumberAlreadyExists := service.repository.Exists(*ctx, createBuyerRequest.CardNumberID)
	if cardNumberAlreadyExists {
		return &entities.Buyer{}, ErrConflict
	}

	id, err := service.repository.Save(*ctx, *buyer)
	if err != nil {
		return &entities.Buyer{}, err
	}

	buyer.ID = id

	return buyer, nil
}

func (service *buyerService) Update(ctx *context.Context, id int, updateBuyerRequest *dtos.UpdateBuyerRequestDTO) (*entities.Buyer, error) {
	// Busca o buyer pelo ID
	buyer, err := service.Get(ctx, id)
	if err != nil {
		return &entities.Buyer{}, err
	}

	// Sobrescreve os dados do buyer, se houver alteração no request
	if updateBuyerRequest.CardNumberID != nil {
		cardNumberAlreadyExists := service.repository.Exists(*ctx, *updateBuyerRequest.CardNumberID)
		if cardNumberAlreadyExists {
			return &entities.Buyer{}, ErrConflict
		}

		buyer.CardNumberID = *updateBuyerRequest.CardNumberID
	}

	if updateBuyerRequest.FirstName != nil {
		buyer.FirstName = *updateBuyerRequest.FirstName
	}

	if updateBuyerRequest.LastName != nil {
		buyer.LastName = *updateBuyerRequest.LastName
	}

	if err := service.repository.Update(*ctx, *buyer); err != nil {
		return &entities.Buyer{}, err
	}

	return buyer, nil

}

func (service *buyerService) Delete(ctx *context.Context, id int) error {
	// Busca o buyer pelo ID
	if _, err := service.Get(ctx, id); err != nil {
		return err
	}

	return service.repository.Delete(*ctx, id)
}
