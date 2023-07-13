package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductBatchServiceMock struct {
	mock.Mock
}

func (p *ProductBatchServiceMock) Save(ctx *context.Context, product domain.ProductBatches) (*domain.ProductBatches, error) {
	args := p.Called(ctx, product)
	r0, r1 := args.Get(0).(*domain.ProductBatches), args.Error(1)
	return r0, r1
}

func (p *ProductBatchServiceMock) SectionProductsReports() ([]domain.ProductBySection, error) {
	args := p.Called()
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}
func (p *ProductBatchServiceMock) SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error) {
	args := p.Called(sectionID)
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}
