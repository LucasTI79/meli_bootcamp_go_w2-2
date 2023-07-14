package productbatches

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
)

var (
	ErrNotFound        = errors.New("product Batches not found")
	ErrNotFoundSection = errors.New("informed section does not exist in the system")
	ErrConflict        = errors.New("product batches with batch_number already exists")
)

type IService interface {
	Save(ctx *context.Context, product domain.ProductBatches) (*domain.ProductBatches, error)
	SectionProductsReports() ([]domain.ProductBySection, error)
	SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error)
}
type Service struct {
	productBatchRepository IRepository
	productRepo            product.Repository
	sectionRepo            section.Repository
}

func NewService(r IRepository, productRepo product.Repository, sectionRepo section.Repository) IService {
	return &Service{
		productBatchRepository: r,
		productRepo:            productRepo,
		sectionRepo:            sectionRepo,
	}
}

func (s *Service) SectionProductsReports() ([]domain.ProductBySection, error) {
	sectionProductsReports, err := s.productBatchRepository.SectionProductsReports()
	if err != nil {
		return sectionProductsReports, err
	}
	return sectionProductsReports, nil
}
func (s *Service) SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error) {
	sectionProductsBySection, err := s.productBatchRepository.SectionProductsReportsBySection(sectionID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFoundSection
		default:
			return nil, err
		}
	}
	return sectionProductsBySection, nil
}

func (s *Service) Save(ctx *context.Context, product domain.ProductBatches) (*domain.ProductBatches, error) {
	batchNumberExist := s.productBatchRepository.ExistsProductBatch(*ctx, product.BatchNumber)
	productIdExist := s.productRepo.ExistsByID(*ctx, product.ProductID)
	sectionIdExist := s.sectionRepo.ExistsByID(*ctx, product.SectionID)

	if batchNumberExist || !productIdExist || !sectionIdExist {
		//return nil, fmt.Errorf("batch_number exists: %t, product_id exists: %t, section_id exists: %t", batchNumberExist, productIdExist, sectionIdExist)
		return &domain.ProductBatches{}, ErrConflict
	}

	productBatchesID, err := s.productBatchRepository.Save(*ctx, product)
	if err != nil {
		return nil, err
	}
	savedProductBatches, err := s.productBatchRepository.Get(*ctx, productBatchesID)
	if err != nil {
		return nil, err
	}
	return &savedProductBatches, nil
}
