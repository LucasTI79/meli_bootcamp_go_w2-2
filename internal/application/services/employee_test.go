package services_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_employeeService_Get(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		expectedEmployee := &entities.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedEmployee, nil)
		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedEmployee, *employeeReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, sql.ErrNoRows)
		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, employeeReceived)
		assert.Equal(t, services.ErrNotFound, err)
	})
}

func Test_employeeService_GetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedEmployee := &[]entities.Employee{
			{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "Maria",
				LastName:     "Silva",
				WarehouseID:  1,
			},
			{
				ID:           2,
				CardNumberID: "234",
				FirstName:    "Joao",
				LastName:     "Silva",
				WarehouseID:  2,
			},
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("GetAll", ctx).Return(*expectedEmployee, nil)
		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedEmployee, *employeeReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("GetAll", ctx).Return([]entities.Employee{}, errors.New("error"))
		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeReceived, err := service.GetAll(&ctx)

		assert.Nil(t, employeeReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func Test_employeeService_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		employeeToDelete := &entities.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*employeeToDelete, nil)
		employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := services.NewEmployeeService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, nil, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, nil)
		employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(services.ErrNotFound)
		service := services.NewEmployeeService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, services.ErrNotFound, err)
	})
	t.Run("delete_not_found", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, services.ErrNotFound)
		//employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(employee.ErrEmployeeNotFound)
		service := services.NewEmployeeService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, services.ErrNotFound, err)
	})

}

func Test_employeeService_Create(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {
		employeeCreated := &entities.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *employeeCreated)

		assert.Equal(t, services.ErrConflict, err)
		assert.Nil(t, employeeSaved)
	})

	t.Run("create_error", func(t *testing.T) {
		employeeCreated := &entities.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Employee")).Return(0, errors.New("error"))

		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *employeeCreated)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, employeeSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedEmployeeCreate := &entities.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Employee")).Return(1, nil)

		service := services.NewEmployeeService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *expectedEmployeeCreate)

		assert.Equal(t, employeeSaved, expectedEmployeeCreate)
		assert.Nil(t, err)
	})
}

func Test_employeeService_Update(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		originalEmployee := &entities.Employee{
			ID:           1,
			CardNumberID: "1",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		expectedEmployee := &entities.Employee{
			ID:           1,
			CardNumberID: "2",
			FirstName:    "Test2",
			LastName:     "Test2",
			WarehouseID:  2,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalEmployee, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Employee")).Return(nil)

		service := services.NewEmployeeService(employeeRepositoryMock)

		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Equal(t, *UpdateEmployees, *expectedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existing", func(t *testing.T) {

		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, services.ErrNotFound)

		service := services.NewEmployeeService(employeeRepositoryMock)

		_, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Equal(t, services.ErrNotFound, err)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {

		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, services.ErrNotFound)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Employee")).Return(assert.AnError)

		service := services.NewEmployeeService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_get_conflit_error", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, services.ErrNotFound)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewEmployeeService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_conflit_error_card_number_id", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewEmployeeService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_repository_error", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &entities.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock(t)
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Employee{}, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Employee")).Return(assert.AnError)

		service := services.NewEmployeeService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})
}
