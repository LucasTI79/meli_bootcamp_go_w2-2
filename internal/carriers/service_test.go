package carriers_test

import (
	"context"
	"database/sql"
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

	t.Run("create_internal_server_error", func(t *testing.T) {

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
		carrieRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Carrier")).Return(0, errors.New("error connecting to server"))
		service := carriers.NewService(carrieRepositoryMock)

		carrieSaved, err := service.Create(&ctx, createCarrierRequestDTO)

		assert.Equal(t, carriers.ErrInternalServerError, err)
		assert.Nil(t, carrieSaved)

	})

	t.Run("create_error_cid_exists", func(t *testing.T) {

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

	t.Run("create_internal_server_error", func(t *testing.T) {
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
		carrieRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Carrier")).Return(0, errors.New("error connecting to server"))
		service := carriers.NewService(carrieRepositoryMock)

		carrieSaved, err := service.Create(&ctx, createCarrierRequestDTO)

		assert.Nil(t, carrieSaved)
		assert.Equal(t, errors.New("error connecting to server"), err)
	})
}

func TestGetLocalityById(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		localityExpected := &domain.Locality{
			ID:           1,
			ProvinceName: "Teste",
			LocalityName: "Teste",
		}
		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetLocalityById", ctx, mock.AnythingOfType("int")).Return(*localityExpected, nil)

		service := carriers.NewService(carrieRepositoryMock)
		locality, err := service.GetLocalityById(&ctx, 1)
		assert.Equal(t, *localityExpected, *locality)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetLocalityById", ctx, mock.AnythingOfType("int")).Return(domain.Locality{}, sql.ErrNoRows)

		service := carriers.NewService(carrieRepositoryMock)

		localityReceived, err := service.GetLocalityById(&ctx, 1)

		assert.Nil(t, localityReceived)
		assert.Equal(t, errors.New("carriers not found"), err)
	})
	t.Run("get_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetLocalityById", ctx, mock.AnythingOfType("int")).Return(domain.Locality{}, errors.New("carriers not found"))

		service := carriers.NewService(carrieRepositoryMock)

		localityReceived, err := service.GetLocalityById(&ctx, 1)

		assert.Nil(t, localityReceived)
		assert.Equal(t, errors.New("carriers not found"), err)
	})
}

func TestGetCountCarriersByLocalityId(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		intExpected := 1
		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetCountCarriersByLocalityId", ctx, mock.AnythingOfType("int")).Return(1, nil)

		service := carriers.NewService(carrieRepositoryMock)
		int, err := service.GetCountCarriersByLocalityId(&ctx, 1)
		assert.Equal(t, *int, intExpected)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_not_found", func(t *testing.T) {

		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetCountCarriersByLocalityId", ctx, mock.AnythingOfType("int")).Return(0, sql.ErrNoRows)

		service := carriers.NewService(carrieRepositoryMock)

		idReceived, err := service.GetCountCarriersByLocalityId(&ctx, 1)

		assert.Nil(t, idReceived)
		assert.Equal(t, carriers.ErrNotFound, err)

	})
}

func TestGetCountAndDataByLocality(t *testing.T) {

	t.Run("find_existent", func(t *testing.T) {
		responsesFounds := []dtos.DataLocalityAndCarrier{
			{
				Id:           1,
				LocalityName: "Teste",
				CountCarrier: 23,
			}, {
				Id:           2,
				LocalityName: "Teste2",
				CountCarrier: 6,
			},
		}
		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetCountAndDataByLocality", ctx).Return(responsesFounds, nil)

		service := carriers.NewService(carrieRepositoryMock)
		returnExpected, err := service.GetCountAndDataByLocality(&ctx)
		assert.Equal(t, *returnExpected, responsesFounds)
		assert.Equal(t, nil, err)
	})
	t.Run("find_non_existent", func(t *testing.T) {
		responsesFounds := []dtos.DataLocalityAndCarrier{}
		ctx := context.TODO()
		carrieRepositoryMock := new(mocks.CarrierRepositoryMock)
		carrieRepositoryMock.On("GetCountAndDataByLocality", ctx).Return(responsesFounds, errors.New("carriers not found"))

		service := carriers.NewService(carrieRepositoryMock)

		count, err := service.GetCountAndDataByLocality(&ctx)

		assert.Nil(t, count)
		assert.Equal(t, errors.New("carriers not found"), err)
	})
}
