package section_test

import (
	"context"
	"database/sql"
	"errors"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sections"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// var (
// 	expectedSection = domain.Section{
// 		ID:                 1,
// 		SectionNumber:      10,
// 		CurrentTemperature: 10,
// 		MinimumTemperature: 10,
// 		CurrentCapacity:    10,
// 		MinimumCapacity:    10,
// 		MaximumCapacity:    10,
// 		WarehouseID:        10,
// 		ProductTypeID:      10,
// 	}
// 	requestSection = dtos.CreateSectionRequestDTO{
// 		SectionNumber:      10,
// 		CurrentTemperature: 10,
// 		MinimumTemperature: 10,
// 		CurrentCapacity:    10,
// 		MinimumCapacity:    10,
// 		MaximumCapacity:    10,
// 		WarehouseID:        10,
// 		ProductTypeID:      10,
// 	}
// )

func TestGet(t *testing.T) {
	t.Run("GET - get_find_by_id_existent", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
		receivedSection, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedSection, *receivedSection)
		assert.Equal(t, nil, err)
	})

	t.Run("GET - get_find_by_id_non_existent", func(t *testing.T) {
		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Section{}, sql.ErrNoRows)
		sectionReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, sectionReceived)
		assert.Equal(t, section.ErrNotFound, err)
	})

	t.Run("GET - get_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Section{}, errors.New("error"))
		sectionReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, sectionReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("GET_ALL - getAll_find_all", func(t *testing.T) {
		expectedsections := &[]domain.Section{
			{
				ID:                 1,
				SectionNumber:      10,
				CurrentTemperature: 10,
				MinimumTemperature: 10,
				CurrentCapacity:    10,
				MinimumCapacity:    10,
				MaximumCapacity:    10,
				WarehouseID:        10,
				ProductTypeID:      10,
			},
			{
				ID:                 2,
				SectionNumber:      20,
				CurrentTemperature: 20,
				MinimumTemperature: 20,
				CurrentCapacity:    20,
				MinimumCapacity:    20,
				MaximumCapacity:    20,
				WarehouseID:        20,
				ProductTypeID:      20,
			},
		}

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("GetAll", ctx).Return(*expectedsections, nil)

		sectionsReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedsections, *sectionsReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("GET_ALL - getAll_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("GetAll", ctx).Return([]domain.Section{}, errors.New("error"))

		sectionsReceived, err := service.GetAll(&ctx)

		assert.Nil(t, sectionsReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("DELETE - delete_ok", func(t *testing.T) {

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()

		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("DELETE - delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()

		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(sql.ErrNoRows)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, section.ErrNotFound, err)
	})

	t.Run("DELETE - delete_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()

		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(errors.New("error"))

		err := service.Delete(&ctx, 1)

		assert.Equal(t, errors.New("error"), err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("CREATE - create_conflict", func(t *testing.T) {

		requestSection := dtos.CreateSectionRequestDTO{
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		sectionSaved, err := service.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)

		assert.Equal(t, section.ErrConflict, err)
		assert.Nil(t, sectionSaved)

	})

	t.Run("CREATE - create_error", func(t *testing.T) {

		requestSection := dtos.CreateSectionRequestDTO{
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		ctx := context.TODO()

		sectionRepositoryMock := new(mocks.SectionRepositoryMock)
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Section")).Return(0, errors.New("error"))

		service := section.NewService(sectionRepositoryMock)
		sectionSaved, err := service.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, sectionSaved)

	})

	t.Run("CREATE - create_error_get_section", func(t *testing.T) {

		requestSection := dtos.CreateSectionRequestDTO{
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Section")).Return(1, nil)
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Section{}, errors.New("error"))

		sectionSaved, err := service.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, sectionSaved)

	})

	t.Run("CREATE - create_ok", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		requestSection := dtos.CreateSectionRequestDTO{
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Section")).Return(1, nil)
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)

		sectionSaved, err := service.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)

		assert.Equal(t, sectionSaved, expectedSection)
		assert.Nil(t, err)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("UPDATE - update_existent", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		requestSection := dtos.UpdateSectionRequestDTO{
			SectionNumber:      &expectedSection.SectionNumber,
			CurrentTemperature: &expectedSection.CurrentTemperature,
			MinimumTemperature: &expectedSection.MinimumCapacity,
			CurrentCapacity:    &expectedSection.CurrentCapacity,
			MinimumCapacity:    &expectedSection.MinimumCapacity,
			MaximumCapacity:    &expectedSection.MinimumCapacity,
			WarehouseID:        &expectedSection.WarehouseID,
			ProductTypeID:      &expectedSection.ProductTypeID,
		}
		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Section")).Return(nil)

		sectionUpdate, err := service.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)

		assert.Equal(t, sectionUpdate, expectedSection)
		assert.Nil(t, err)
	})

	t.Run("UPDATE - update_non_existent", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		requestSection := dtos.UpdateSectionRequestDTO{
			SectionNumber:      &expectedSection.SectionNumber,
			CurrentTemperature: &expectedSection.CurrentTemperature,
			MinimumTemperature: &expectedSection.MinimumCapacity,
			CurrentCapacity:    &expectedSection.CurrentCapacity,
			MinimumCapacity:    &expectedSection.MinimumCapacity,
			MaximumCapacity:    &expectedSection.MinimumCapacity,
			WarehouseID:        &expectedSection.WarehouseID,
			ProductTypeID:      &expectedSection.ProductTypeID,
		}

		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Section")).Return(sql.ErrNoRows)

		sectionUpdate, err := service.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)

		assert.Equal(t, section.ErrNotFound, err)
		assert.Nil(t, sectionUpdate)
	})

	t.Run("UPDATE - update_unexpected_error", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		requestSection := dtos.UpdateSectionRequestDTO{
			SectionNumber:      &expectedSection.SectionNumber,
			CurrentTemperature: &expectedSection.CurrentTemperature,
			MinimumTemperature: &expectedSection.MinimumCapacity,
			CurrentCapacity:    &expectedSection.CurrentCapacity,
			MinimumCapacity:    &expectedSection.MinimumCapacity,
			MaximumCapacity:    &expectedSection.MinimumCapacity,
			WarehouseID:        &expectedSection.WarehouseID,
			ProductTypeID:      &expectedSection.ProductTypeID,
		}

		ctx := context.TODO()

		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Section")).Return(errors.New("error"))

		sectionUpdate, err := service.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, sectionUpdate)
	})

	t.Run("UPDATE - update_get_error", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		requestSection := dtos.UpdateSectionRequestDTO{
			SectionNumber:      &expectedSection.SectionNumber,
			CurrentTemperature: &expectedSection.CurrentTemperature,
			MinimumTemperature: &expectedSection.MinimumCapacity,
			CurrentCapacity:    &expectedSection.CurrentCapacity,
			MinimumCapacity:    &expectedSection.MinimumCapacity,
			MaximumCapacity:    &expectedSection.MinimumCapacity,
			WarehouseID:        &expectedSection.WarehouseID,
			ProductTypeID:      &expectedSection.ProductTypeID,
		}

		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Section{}, errors.New("error"))

		sectionUpdate, err := service.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, sectionUpdate)
	})

	t.Run("UPDDATE - conflit", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		sectionNumber := 10
		requestSection := dtos.UpdateSectionRequestDTO{
			SectionNumber:      &sectionNumber,
			CurrentTemperature: &expectedSection.CurrentTemperature,
			MinimumTemperature: &expectedSection.MinimumCapacity,
			CurrentCapacity:    &expectedSection.CurrentCapacity,
			MinimumCapacity:    &expectedSection.MinimumCapacity,
			MaximumCapacity:    &expectedSection.MinimumCapacity,
			WarehouseID:        &expectedSection.WarehouseID,
			ProductTypeID:      &expectedSection.ProductTypeID,
		}

		ctx := context.TODO()
		sectionRepositoryMock, service := InitMock()
		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Section{}, nil)
		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		sectionUpdate, err := service.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)

		assert.Equal(t, section.ErrConflict, err)
		assert.Nil(t, sectionUpdate)
	})
}

func InitMock() (*mocks.SectionRepositoryMock, section.Service) {
	sectionRepositoryMock := new(mocks.SectionRepositoryMock)
	service := section.NewService(sectionRepositoryMock)
	return sectionRepositoryMock, service
}
