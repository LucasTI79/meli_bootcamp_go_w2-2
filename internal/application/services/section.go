package services

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type SectionService interface {
	Save(ctx *context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
		warehouseID, productTypeID int) (*entities.Section, error)
	GetAll(ctx *context.Context) (*[]entities.Section, error)
	Get(ctx *context.Context, id int) (*entities.Section, error)
	Delete(ctx *context.Context, id int) error
	Update(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
		warehouseID, productTypeID *int, id int) (*entities.Section, error)
}

type sectionService struct {
	sectionRepository repositories.SectionRepository
}

func NewSectionService(r repositories.SectionRepository) SectionService {
	return &sectionService{
		sectionRepository: r,
	}
}

func (s *sectionService) Save(ctx *context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseID, productTypeID int) (*entities.Section, error) {
	existingSection := s.sectionRepository.Exists(*ctx, sectionNumber)

	if existingSection {
		return nil, ErrConflict
	}

	newSection := entities.Section{
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minimumTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minimumCapacity,
		MaximumCapacity:    maximumCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	sectionId, err := s.sectionRepository.Save(*ctx, newSection)
	if err != nil {
		return nil, err

	}

	savedSection, err := s.sectionRepository.Get(*ctx, sectionId)
	if err != nil {
		return nil, err
	}

	return &savedSection, nil
}

func (s *sectionService) GetAll(ctx *context.Context) (*[]entities.Section, error) {
	sections := make([]entities.Section, 0)
	sections, err := s.sectionRepository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}

	return &sections, nil
}

func (s *sectionService) Get(ctx *context.Context, id int) (*entities.Section, error) {
	section, err := s.sectionRepository.Get(*ctx, id)
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

func (s *sectionService) Delete(ctx *context.Context, id int) error {
	err := s.sectionRepository.Delete(*ctx, id)
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

func (s *sectionService) Update(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
	warehouseID, productTypeID *int, id int) (*entities.Section, error) {
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
