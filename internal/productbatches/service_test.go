package productbatches_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product/product_mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section/section_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expectedProductBatch = domain.ProductBatches{
		ID:                 1,
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
	payload = domain.ProductBatches{
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
)

func TestCreate(t *testing.T) {
	t.Run("CREATE - Status OK", func(t *testing.T) {
		ctx := context.TODO()

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)

		productBatchesRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductBatches")).Return(1, nil)
		productBatchesRepositoryMock.On("ExistsProductBatch", ctx, payload.BatchNumber).Return(false)
		productRepository.On("ExistsByID", ctx, payload.ProductID).Return(true)
		sectionRepository.On("ExistsByID", ctx, payload.SectionID).Return(true)
		productBatchesRepositoryMock.On("Get", ctx, 1).Return(expectedProductBatch, nil)

		productBatchSaved, err := service.Save(&ctx, payload)
		assert.Equal(t, *productBatchSaved, expectedProductBatch)
		assert.Nil(t, err)
	})
	t.Run("CREATE - Conflits - Batch Number Already exist", func(t *testing.T) {
		ctx := context.TODO()

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)

		productBatchesRepositoryMock.On("ExistsProductBatch", ctx, payload.BatchNumber).Return(true)
		productRepository.On("ExistsByID", ctx, payload.ProductID).Return(true)
		sectionRepository.On("ExistsByID", ctx, payload.SectionID).Return(true)
		productBatchesRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductBatches")).Return(domain.ProductBatches{}, productbatches.ErrConflict)

		_, err := service.Save(&ctx, payload)
		assert.Equal(t, productbatches.ErrConflict, err)
	})
	t.Run("CREATE - Err GET", func(t *testing.T) {
		ctx := context.TODO()

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)

		productBatchesRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductBatches")).Return(1, nil)
		productBatchesRepositoryMock.On("ExistsProductBatch", ctx, payload.BatchNumber).Return(false)
		productRepository.On("ExistsByID", ctx, payload.ProductID).Return(true)
		sectionRepository.On("ExistsByID", ctx, payload.SectionID).Return(true)
		productBatchesRepositoryMock.On("Get", ctx, 1).Return(domain.ProductBatches{}, assert.AnError)

		productBatchSaved, err := service.Save(&ctx, payload)

		assert.Equal(t, assert.AnError, err)
		assert.Nil(t, productBatchSaved)

	})
	t.Run("CREATE - Err Save", func(t *testing.T) {
		ctx := context.TODO()

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)

		productBatchesRepositoryMock.On("ExistsProductBatch", ctx, payload.BatchNumber).Return(false)
		productRepository.On("ExistsByID", ctx, payload.ProductID).Return(true)
		sectionRepository.On("ExistsByID", ctx, payload.SectionID).Return(true)
		productBatchesRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductBatches")).Return(0, assert.AnError)

		productBatchesRepositoryMock.On("Get", ctx, 1).Return(expectedProductBatch, nil)

		_, err := service.Save(&ctx, payload)

		assert.Equal(t, assert.AnError, err)
	})

}
func TestSectionProductsReports(t *testing.T) {
	t.Run("SectionProductsReports - Internal Server Error", func(t *testing.T) {

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)
		productBatchesRepositoryMock.On("SectionProductsReports").Return([]domain.ProductBySection{}, assert.AnError)

		_, err := service.SectionProductsReports()

		assert.Equal(t, assert.AnError, err)
	})
	t.Run("SectionProductsReports - OK", func(t *testing.T) {
		expectedReportProducts := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
			{
				SectionID:     2,
				SectionNumber: "20",
				ProductsCount: 200,
			},
		}
		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)
		productBatchesRepositoryMock.On("SectionProductsReports").Return(expectedReportProducts, nil)

		sectionProductsReportsActual, err := service.SectionProductsReports()

		assert.Equal(t, sectionProductsReportsActual, expectedReportProducts)
		assert.Equal(t, nil, err)
	})
}

func TestSectionProductsReportsBySection(t *testing.T) {
	t.Run("Not Found - SectionProductsReportsBySection", func(t *testing.T) {
		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)
		productBatchesRepositoryMock.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return([]domain.ProductBySection{}, sql.ErrNoRows)
		id := 1
		received, err := service.SectionProductsReportsBySection(id)

		assert.Nil(t, received)
		assert.Equal(t, productbatches.ErrNotFoundSection, err)
	})
	t.Run("Internal Server Error - SectionProductsReportsBySection", func(t *testing.T) {

		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)
		productBatchesRepositoryMock.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return([]domain.ProductBySection{}, assert.AnError)
		id := 1
		received, err := service.SectionProductsReportsBySection(id)

		assert.Nil(t, received)
		assert.Equal(t, assert.AnError, err)
	})
	t.Run("OK - SectionProductsReportsBySection", func(t *testing.T) {
		expectedReportProductsBySection := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
		}
		productBatchesRepositoryMock := mocks.NewProductBatchesRepositoryMock()
		productRepository := new(product_mocks.ProductRepositoryMock)
		sectionRepository := new(section_mocks.SectionRepositoryMock)

		service := productbatches.NewService(productBatchesRepositoryMock, productRepository, sectionRepository)
		productBatchesRepositoryMock.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return(expectedReportProductsBySection, nil)
		id := 1
		received, err := service.SectionProductsReportsBySection(id)

		assert.Equal(t, expectedReportProductsBySection, received)
		assert.Equal(t, nil, err)
	})

}
