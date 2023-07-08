package carriers

import (
	"context"
	"errors"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

var (
	ErrNotFound            = errors.New("carriers not found")
	ErrConflict            = errors.New("a carriers with this cid already exists")
	ErrUnprocessableEntity = errors.New("all fields are required")
)

type Service interface {
	Create(c *context.Context, dto dtos.CarrierRequestDTO) (*domain.Carrier, error)
	GetAll(c *context.Context) (*[]domain.Carrier, error)
	GetLocalityById(c *context.Context, localityId int) (*domain.Locality, error)
	GetCountCarriersByLocalityId(c *context.Context, localityId int) (int, error)
	GetCountAndDataByLocality(c *context.Context) (*[]dtos.DataLocalityAndCarrier, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Create(c *context.Context, dto dtos.CarrierRequestDTO) (*domain.Carrier, error) {

	exists := s.repository.Exists(*c, dto.CID)
	if exists {
		return nil, ErrConflict
	}

	var formatter domain.Carrier = domain.Carrier{
		ID:          0,
		CID:         dto.CID,
		CompanyName: dto.CompanyName,
		Address:     dto.Address,
		Telephone:   dto.Telephone,
		LocalityId:  dto.LocalityId,
	}

	id, err := s.repository.Save(*c, formatter)

	if err != nil {
		return nil, err
	}

	formatter.ID = id

	return &formatter, nil
}

func (s *service) GetAll(c *context.Context) (*[]domain.Carrier, error) {
	carriers, err := s.repository.GetAll(*c)

	if err != nil {
		return nil, err
	}

	return &carriers, nil
}

func (s *service) GetLocalityById(c *context.Context, localityId int) (*domain.Locality, error) {
	result, err := s.repository.GetLocalityById(*c, localityId)
	if err != nil {
		return nil, ErrNotFound
	}

	return &result, nil
}

func (s *service) GetCountCarriersByLocalityId(c *context.Context, localityId int) (int, error) {
	count, err := s.repository.GetCountCarriersByLocalityId(*c, localityId)
	if err != nil {
		return 0, ErrNotFound
	}

	return count, nil
}

func (s *service) GetCountAndDataByLocality(c *context.Context) (*[]dtos.DataLocalityAndCarrier, error) {
	dataAndCount, err := s.repository.GetCountAndDataByLocality(*c)

	if err != nil {
		return nil, err
	}
	return &dataAndCount, nil
}
