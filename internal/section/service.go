package section

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
	ErrConflict = errors.New("section with Section Number already exists")
)

type Service interface {
	Save(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
		warehouseID, productTypeID int) (*domain.Section, error)
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (*domain.Section, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
		warehouseID, productTypeID *int, id int) (*domain.Section, error)
}

type service struct {
	sectionRepository Repository
}

func NewService(r Repository) Service {
	return &service{
		sectionRepository: r,
	}
}

func (s *service) Save(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseID, productTypeID int) (*domain.Section, error) {
	existingSection := s.sectionRepository.Exists(ctx, sectionNumber)

	if existingSection {
		return nil, ErrConflict
	}

	newSection := domain.Section{
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	sectionId, err := s.sectionRepository.Save(ctx, newSection)
	if err != nil {
		return nil, err

	}

	savedSection, err := s.sectionRepository.Get(ctx, sectionId)
	if err != nil {
		return nil, err
	}

	return &savedSection, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	sections, err := s.sectionRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (s *service) Get(ctx context.Context, id int) (*domain.Section, error) {
	section, err := s.sectionRepository.Get(ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &section, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.sectionRepository.Delete(ctx, id)
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

func (s *service) Update(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
	warehouseID, productTypeID *int, id int) (*domain.Section, error) {
	existingSection, err := s.sectionRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if sectionNumber != nil {
		existingSectionSearch := s.sectionRepository.Exists(ctx, *sectionNumber)
		if existingSectionSearch && *sectionNumber != existingSection.SectionNumber {
			return nil, ErrConflict
		}
		existingSection.SectionNumber = *sectionNumber
	}
	if currentTemperature != nil {
		existingSection.CurrentTemperature = *currentTemperature
	}
	if minimumTemperature != nil {
		existingSection.MinimumTemperature = *minimumTemperature
	}
	if currentCapacity != nil {
		existingSection.CurrentCapacity = *currentCapacity
	}
	if minimumCapacity != nil {
		existingSection.MinimumCapacity = *minimumCapacity
	}
	if maximumCapacity != nil {
		existingSection.MaximumCapacity = *maximumCapacity
	}
	if warehouseID != nil {
		existingSection.WarehouseID = *warehouseID
	}
	if productTypeID != nil {
		existingSection.ProductTypeID = *productTypeID
	}

	err1 := s.sectionRepository.Update(ctx, existingSection)
	if err1 != nil {
		switch err1 {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err1
		}
	}

	return &existingSection, nil
}
