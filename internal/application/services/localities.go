package services

import (
	"context"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Errors
var (
	ErrNotFound = errors.New("locality not found")
	ErrConflict = errors.New("ID already exists")
)

type LocalityService interface {
	Get(ctx *context.Context, id int) (entities.Locality, error)
	GetAll(ctx *context.Context) ([]entities.Locality, error)
	Create(ctx *context.Context, locality entities.Locality) (entities.Locality, error)
	Update(ctx *context.Context, id int, updateLocalityRequest dtos.UpdateLocalityRequestDTO) (entities.Locality, error)
	Delete(ctx *context.Context, id int) error
	CountSellers(ctx *context.Context, id int) (int, error)
}

type localityService struct {
	localityRepository repositories.LocalityRepository
}

func NewLocalityService(r repositories.LocalityRepository) LocalityService {
	return &localityService{
		localityRepository: r,
	}
}

func (service *localityService) Get(ctx *context.Context, id int) (entities.Locality, error) {
	locality, err := service.localityRepository.Get(*ctx, id)
	if err != nil {
		return entities.Locality{}, err
	}

	return locality, nil
}

func (service *localityService) GetAll(ctx *context.Context) ([]entities.Locality, error) {
	localities := make([]entities.Locality, 0)

	localities, err := service.localityRepository.GetAll(*ctx)
	if err != nil {
		return localities, err
	}

	return localities, nil
}

func (service *localityService) Create(ctx *context.Context, locality entities.Locality) (entities.Locality, error) {
	existingLocality := service.localityRepository.Exists(*ctx, locality.ID)
	if existingLocality {
		return entities.Locality{}, ErrConflict
	}

	id, err := service.localityRepository.Save(*ctx, locality)
	if err != nil {
		return entities.Locality{}, err
	}

	locality.ID = id

	return locality, nil
}

func (service *localityService) Update(ctx *context.Context, id int, updateLocalityRequest dtos.UpdateLocalityRequestDTO) (entities.Locality, error) {
	existingLocality, err := service.localityRepository.Get(*ctx, id)
	if err != nil {
		return entities.Locality{}, err
	}

	existingLocalitySearch := service.localityRepository.Exists(*ctx, id)
	if existingLocalitySearch {
		return entities.Locality{}, ErrConflict
	}

	if updateLocalityRequest.CountryName != nil {
		existingLocality.CountryName = *updateLocalityRequest.CountryName
	}
	if updateLocalityRequest.ProvinceName != nil {
		existingLocality.ProvinceName = *updateLocalityRequest.ProvinceName
	}
	if updateLocalityRequest.LocalityName != nil {
		existingLocality.LocalityName = *updateLocalityRequest.LocalityName
	}

	if err = service.localityRepository.Update(*ctx, existingLocality); err != nil {
		return entities.Locality{}, err
	}

	return existingLocality, nil
}

func (service *localityService) Delete(ctx *context.Context, id int) error {
	_, err := service.localityRepository.Get(*ctx, id)
	if err != nil {
		return err
	}

	err = service.localityRepository.Delete(*ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *localityService) CountSellers(ctx *context.Context, id int) (int, error) {
	count, err := service.localityRepository.CountSellers(*ctx, id)
	if err != nil {
		return 0, err
	}

	return count, nil

}
