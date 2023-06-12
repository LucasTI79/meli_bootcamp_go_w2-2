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
	Get(ctx context.Context, id int) (*domain.Buyer, error)
	GetAll(ctx context.Context) (*[]domain.Buyer, error)
	Create(ctx context.Context, buyer *domain.Buyer) (*domain.Buyer, error)
	Update(ctx context.Context, updateBuyerRequest *domain.UpdateBuyerRequestDTO) (*domain.Buyer, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (service *service) Get(ctx context.Context, id int) (*domain.Buyer, error) {
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

func (service *service) GetAll(ctx context.Context) (*[]domain.Buyer, error) {
	buyers := make([]domain.Buyer, 0)

	buyers, err := service.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &buyers, nil

}

func (service *service) Create(ctx context.Context, buyer *domain.Buyer) (*domain.Buyer, error) {
	id, err := service.repository.Save(ctx, *buyer)
	if err != nil {
		return nil, err
	}

	buyer.ID = id

	return buyer, nil
}

func (service *service) Update(ctx context.Context, updateBuyerRequest *domain.UpdateBuyerRequestDTO) (*domain.Buyer, error) {
	// Busca o buyer pelo ID
	buyer, err := service.Get(ctx, updateBuyerRequest.ID)
	if err != nil {
		return nil, err
	}

	// Sobrescreve os dados do buyer, se houver alteração no request
	if updateBuyerRequest.CardNumberID != nil {
		buyer.CardNumberID = *updateBuyerRequest.CardNumberID
	}

	if updateBuyerRequest.FirstName != nil {
		buyer.FirstName = *updateBuyerRequest.FirstName
	}

	if updateBuyerRequest.LastName != nil {
		buyer.LastName = *updateBuyerRequest.LastName
	}

	// Atualiza o buyer no banco de dados
	if err := service.repository.Update(ctx, *buyer); err != nil {
		return nil, err
	}

	return buyer, nil

}

func (service *service) Delete(ctx context.Context, id int) error {
	// Busca o buyer pelo ID
	if _, err := service.Get(ctx, id); err != nil {
		return err
	}

	return service.repository.Delete(ctx, id)
}
