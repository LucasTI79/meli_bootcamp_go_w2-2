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
	Create(buyer *domain.Buyer) (*domain.Buyer, error)
	Update(updateBuyerRequest *domain.UpdateBuyerRequestDTO) (*domain.Buyer, error)
	Delete(id int) error
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

func (service *service) Create(buyer *domain.Buyer) (*domain.Buyer, error) {
	// TODO: Remove this
	ctx := context.TODO()

	id, err := service.repository.Save(ctx, *buyer)
	if err != nil {
		return nil, err
	}

	buyer.ID = id

	return buyer, nil
}

func (service *service) Update(updateBuyerRequest *domain.UpdateBuyerRequestDTO) (*domain.Buyer, error) {
	// TODO: Remove this
	ctx := context.TODO()

	// Busca o buyer pelo ID
	buyer, err := service.Get(updateBuyerRequest.ID)
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

func (service *service) Delete(id int) error {
	// TODO: Remove this
	ctx := context.TODO()

	// Busca o buyer pelo ID
	if _, err := service.Get(id); err != nil {
		return err
	}

	return service.repository.Delete(ctx, id)
}
