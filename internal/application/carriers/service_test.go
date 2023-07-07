package carriers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/carriers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/carriers/mocks"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedCarriers := &[]domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6700,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}
		ctx := context.TODO()

		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetAll", ctx).Return(*expectedCarriers, nil)

		service := carriers.NewService(carrieRepositoryMock)
		carriersReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedCarriers, *carriersReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetAll", ctx).Return([]domain.Carrier{}, errors.New("error"))

		service := carriers.NewService(carrieRepositoryMock)
		warehousesReceived, err := service.GetAll(&ctx)

		assert.Nil(t, warehousesReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		ctx := context.TODO()

		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := carriers.NewService(carrieRepositoryMock)
		carrieSaved, err := service.Create(&ctx, createCarrierRequestDTO)

		assert.Equal(t, carriers.ErrConflict, err)
		assert.Nil(t, carrieSaved)

	})

	t.Run("create_error", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		ctx := context.TODO()

		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)
		carrieRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Carrie")).Return(0, errors.New("a carriers with this cid already exists"))
		service := carriers.NewService(carrieRepositoryMock)

		carrieSaved, err := service.Create(&ctx, createCarrierRequestDTO)

		assert.Equal(t, errors.New("a carriers with this cid already exists"), err)
		assert.Nil(t, carrieSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}
		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		ctx := context.TODO()

		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		carrieRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Carrier")).Return(1, nil)
		service := carriers.NewService(carrieRepositoryMock)

		carrieSaved, err := service.Create(&ctx, createCarrierRequestDTO)

		assert.Equal(t, carrieSaved, expectedCarrier)
		assert.Nil(t, err)
	})
}
