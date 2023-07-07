package services_test

//
//import (
//	"context"
//	"database/sql"
//	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
//	"testing"
//
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section/mocks"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//func Test_sectionService_Get(t *testing.T) {
//	t.Run("GET - get_find_by_id_existent", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
//		receivedSection, err := sellerService.Get(&ctx, 1)
//
//		assert.Equal(t, *expectedSection, *receivedSection)
//		assert.Equal(t, nil, err)
//	})
//
//	t.Run("GET - get_find_by_id_non_existent", func(t *testing.T) {
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Section{}, sql.ErrNoRows)
//		sectionReceived, err := sellerService.Get(&ctx, 1)
//
//		assert.Nil(t, sectionReceived)
//		assert.Equal(t, services.ErrNotFound, err)
//	})
//
//	t.Run("GET - get_unexpected_error", func(t *testing.T) {
//
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Section{}, assert.AnError)
//		sectionReceived, err := sellerService.Get(&ctx, 1)
//
//		assert.Nil(t, sectionReceived)
//		assert.Equal(t, assert.AnError, err)
//	})
//}
//
//func Test_sectionService_GetAll(t *testing.T) {
//	t.Run("GET_ALL - getAll_find_all", func(t *testing.T) {
//		expectedSections := &[]entities.Section{
//			{
//				ID:                 1,
//				SectionNumber:      10,
//				CurrentTemperature: 10,
//				MinimumTemperature: 10,
//				CurrentCapacity:    10,
//				MinimumCapacity:    10,
//				MaximumCapacity:    10,
//				WarehouseID:        10,
//				ProductTypeID:      10,
//			},
//			{
//				ID:                 2,
//				SectionNumber:      20,
//				CurrentTemperature: 20,
//				MinimumTemperature: 20,
//				CurrentCapacity:    20,
//				MinimumCapacity:    20,
//				MaximumCapacity:    20,
//				WarehouseID:        20,
//				ProductTypeID:      20,
//			},
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("GetAll", ctx).Return(*expectedSections, nil)
//
//		sectionsReceived, err := sellerService.GetAll(&ctx)
//
//		assert.Equal(t, *expectedSections, *sectionsReceived)
//		assert.Equal(t, nil, err)
//	})
//
//	t.Run("GET_ALL - getAll_unexpected_error", func(t *testing.T) {
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("GetAll", ctx).Return([]entities.Section{}, assert.AnError)
//
//		sectionsReceived, err := sellerService.GetAll(&ctx)
//
//		assert.Nil(t, sectionsReceived)
//		assert.Equal(t, assert.AnError, err)
//	})
//}
//
//func Test_sectionService_Delete(t *testing.T) {
//	t.Run("DELETE - delete_ok", func(t *testing.T) {
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//
//		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)
//
//		err := sellerService.Delete(&ctx, 1)
//
//		assert.Nil(t, err)
//	})
//
//	t.Run("DELETE - delete_non_existent", func(t *testing.T) {
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//
//		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(sql.ErrNoRows)
//
//		err := sellerService.Delete(&ctx, 1)
//
//		assert.Equal(t, services.ErrNotFound, err)
//	})
//
//	t.Run("DELETE - delete_unexpected_error", func(t *testing.T) {
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//
//		sectionRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(assert.AnError)
//
//		err := sellerService.Delete(&ctx, 1)
//
//		assert.Equal(t, assert.AnError, err)
//	})
//}
//
//func Test_sectionService_Create(t *testing.T) {
//	t.Run("CREATE - create_conflict", func(t *testing.T) {
//
//		requestSection := dtos.CreateSectionRequestDTO{
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)
//
//		sectionSaved, err := sellerService.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)
//
//		assert.Equal(t, services.ErrConflict, err)
//		assert.Nil(t, sectionSaved)
//
//	})
//
//	t.Run("CREATE - create_error", func(t *testing.T) {
//
//		requestSection := dtos.CreateSectionRequestDTO{
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock := new(mocks.SectionRepositoryMock)
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Section")).Return(0, assert.AnError)
//
//		sellerService := services.NewBuyerService(sectionRepositoryMock)
//		sectionSaved, err := sellerService.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)
//
//		assert.Equal(t, assert.AnError, err)
//		assert.Nil(t, sectionSaved)
//
//	})
//
//	t.Run("CREATE - create_error_get_section", func(t *testing.T) {
//
//		requestSection := dtos.CreateSectionRequestDTO{
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Section")).Return(1, nil)
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Section{}, assert.AnError)
//
//		sectionSaved, err := sellerService.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)
//
//		assert.Equal(t, assert.AnError, err)
//		assert.Nil(t, sectionSaved)
//
//	})
//
//	t.Run("CREATE - create_ok", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		requestSection := dtos.CreateSectionRequestDTO{
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Section")).Return(1, nil)
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
//
//		sectionSaved, err := sellerService.Save(&ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID)
//
//		assert.Equal(t, sectionSaved, expectedSection)
//		assert.Nil(t, err)
//
//	})
//}
//
//func Test_sectionService_Update(t *testing.T) {
//	t.Run("UPDATE - update_existent", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		requestSection := dtos.UpdateSectionRequestDTO{
//			SectionNumber:      &expectedSection.SectionNumber,
//			CurrentTemperature: &expectedSection.CurrentTemperature,
//			MinimumTemperature: &expectedSection.MinimumCapacity,
//			CurrentCapacity:    &expectedSection.CurrentCapacity,
//			MinimumCapacity:    &expectedSection.MinimumCapacity,
//			MaximumCapacity:    &expectedSection.MinimumCapacity,
//			WarehouseID:        &expectedSection.WarehouseID,
//			ProductTypeID:      &expectedSection.ProductTypeID,
//		}
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Section")).Return(nil)
//
//		sectionUpdate, err := sellerService.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)
//
//		assert.Equal(t, sectionUpdate, expectedSection)
//		assert.Nil(t, err)
//	})
//
//	t.Run("UPDATE - update_non_existent", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		requestSection := dtos.UpdateSectionRequestDTO{
//			SectionNumber:      &expectedSection.SectionNumber,
//			CurrentTemperature: &expectedSection.CurrentTemperature,
//			MinimumTemperature: &expectedSection.MinimumCapacity,
//			CurrentCapacity:    &expectedSection.CurrentCapacity,
//			MinimumCapacity:    &expectedSection.MinimumCapacity,
//			MaximumCapacity:    &expectedSection.MinimumCapacity,
//			WarehouseID:        &expectedSection.WarehouseID,
//			ProductTypeID:      &expectedSection.ProductTypeID,
//		}
//
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Section")).Return(sql.ErrNoRows)
//
//		sectionUpdate, err := sellerService.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)
//
//		assert.Equal(t, services.ErrNotFound, err)
//		assert.Nil(t, sectionUpdate)
//	})
//
//	t.Run("UPDATE - update_unexpected_error", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		requestSection := dtos.UpdateSectionRequestDTO{
//			SectionNumber:      &expectedSection.SectionNumber,
//			CurrentTemperature: &expectedSection.CurrentTemperature,
//			MinimumTemperature: &expectedSection.MinimumCapacity,
//			CurrentCapacity:    &expectedSection.CurrentCapacity,
//			MinimumCapacity:    &expectedSection.MinimumCapacity,
//			MaximumCapacity:    &expectedSection.MinimumCapacity,
//			WarehouseID:        &expectedSection.WarehouseID,
//			ProductTypeID:      &expectedSection.ProductTypeID,
//		}
//
//		ctx := context.TODO()
//
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedSection, nil)
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
//		sectionRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Section")).Return(assert.AnError)
//
//		sectionUpdate, err := sellerService.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)
//
//		assert.Equal(t, assert.AnError, err)
//		assert.Nil(t, sectionUpdate)
//	})
//
//	t.Run("UPDATE - update_get_error", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		requestSection := dtos.UpdateSectionRequestDTO{
//			SectionNumber:      &expectedSection.SectionNumber,
//			CurrentTemperature: &expectedSection.CurrentTemperature,
//			MinimumTemperature: &expectedSection.MinimumCapacity,
//			CurrentCapacity:    &expectedSection.CurrentCapacity,
//			MinimumCapacity:    &expectedSection.MinimumCapacity,
//			MaximumCapacity:    &expectedSection.MinimumCapacity,
//			WarehouseID:        &expectedSection.WarehouseID,
//			ProductTypeID:      &expectedSection.ProductTypeID,
//		}
//
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Section{}, assert.AnError)
//
//		sectionUpdate, err := sellerService.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)
//
//		assert.Equal(t, assert.AnError, err)
//		assert.Nil(t, sectionUpdate)
//	})
//
//	t.Run("UPDDATE - conflit", func(t *testing.T) {
//		expectedSection := &entities.Section{
//			ID:                 1,
//			SectionNumber:      10,
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//		sectionNumber := 10
//		requestSection := dtos.UpdateSectionRequestDTO{
//			SectionNumber:      &sectionNumber,
//			CurrentTemperature: &expectedSection.CurrentTemperature,
//			MinimumTemperature: &expectedSection.MinimumCapacity,
//			CurrentCapacity:    &expectedSection.CurrentCapacity,
//			MinimumCapacity:    &expectedSection.MinimumCapacity,
//			MaximumCapacity:    &expectedSection.MinimumCapacity,
//			WarehouseID:        &expectedSection.WarehouseID,
//			ProductTypeID:      &expectedSection.ProductTypeID,
//		}
//
//		ctx := context.TODO()
//		sectionRepositoryMock, sellerService := InitMock()
//		sectionRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Section{}, nil)
//		sectionRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)
//
//		sectionUpdate, err := sellerService.Update(ctx, requestSection.SectionNumber, requestSection.CurrentTemperature, requestSection.MinimumTemperature, requestSection.CurrentCapacity, requestSection.MinimumCapacity, requestSection.MaximumCapacity, requestSection.WarehouseID, requestSection.ProductTypeID, 1)
//
//		assert.Equal(t, services.ErrConflict, err)
//		assert.Nil(t, sectionUpdate)
//	})
//}
