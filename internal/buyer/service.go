package buyer

import (
	"context"
	"database/sql"
	"errors"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound             = errors.New("buyer not found")
	ErrCardNumberDuplicated = errors.New("Credit card already exists for another user")
)

type Service interface {
	Get(ctx *context.Context, id int) (*domain.Buyer, error)
	GetAll(ctx *context.Context) (*[]domain.Buyer, error)
	Create(ctx *context.Context, createBuyerRequest *dtos.CreateBuyerRequestDTO) (*domain.Buyer, error)
	Update(ctx *context.Context, id int, updateBuyerRequest *dtos.UpdateBuyerRequestDTO) (*domain.Buyer, error)
	Delete(ctx *context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (service *service) Get(ctx *context.Context, id int) (*domain.Buyer, error) {
	buyer, err := service.repository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return &domain.Buyer{}, ErrNotFound
		default:
			return &domain.Buyer{}, err
		}
	}

	return &buyer, nil
}

func (service *service) GetAll(ctx *context.Context) (*[]domain.Buyer, error) {
	buyers := make([]domain.Buyer, 0)

	buyers, err := service.repository.GetAll(*ctx)
	if err != nil {
		return &buyers, err
	}

	return &buyers, nil

}

func (service *service) Create(ctx *context.Context, createBuyerRequest *dtos.CreateBuyerRequestDTO) (*domain.Buyer, error) {
	buyer := createBuyerRequest.ToDomain()

	cardNumberAlreadyExists := service.repository.Exists(*ctx, createBuyerRequest.CardNumberID)
	if cardNumberAlreadyExists {
		return &domain.Buyer{}, ErrCardNumberDuplicated
	}

	id, err := service.repository.Save(*ctx, *buyer)
	if err != nil {
		return &domain.Buyer{}, err
	}

	buyer.ID = id

	return buyer, nil
}

func (service *service) Update(ctx *context.Context, id int, updateBuyerRequest *dtos.UpdateBuyerRequestDTO) (*domain.Buyer, error) {
	// Busca o buyer pelo ID
	buyer, err := service.Get(ctx, id)
	if err != nil {
		return &domain.Buyer{}, err
	}

	// Sobrescreve os dados do buyer, se houver alteração no request
	if updateBuyerRequest.CardNumberID != nil {
		cardNumberAlreadyExists := service.repository.Exists(*ctx, *updateBuyerRequest.CardNumberID)
		if cardNumberAlreadyExists {
			return &domain.Buyer{}, ErrCardNumberDuplicated
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
		return &domain.Buyer{}, err
	}

	return buyer, nil

}

func (service *service) Delete(ctx *context.Context, id int) error {
	// Busca o buyer pelo ID
	if _, err := service.Get(ctx, id); err != nil {
		return err
	}

	return service.repository.Delete(*ctx, id)
}
