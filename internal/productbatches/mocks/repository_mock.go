package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockIRepository struct {
	mock.Mock
}
func NewProductBatchesRepositoryMock() *MockIRepository {
	return &MockIRepository{}
}

func (m *MockIRepository) SectionProductsReports() ([]domain.ProductBySection, error) {
	args := m.Called()
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}

func (m *MockIRepository) SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error) {
	args := m.Called(sectionID)
	return args.Get(0).([]domain.ProductBySection), args.Error(1)
}

func (m *MockIRepository) Save(ctx context.Context, product domain.ProductBatches) (int, error) {
	args := m.Called(ctx, product)
	return args.Int(0), args.Error(1)
}

func (m *MockIRepository) Get(ctx context.Context, id int) (domain.ProductBatches, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.ProductBatches), args.Error(1)
}

func (m *MockIRepository) ExistsProductBatch(ctx context.Context, batchNumber int) bool {
	args := m.Called(ctx, batchNumber)
	return args.Bool(0)
}
func (m *MockIRepository) ExistsByID(ctx context.Context, id int) bool{
	args := m.Called(ctx, id)
	return args.Bool(0)
}

