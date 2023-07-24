package employee_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedEmployee, nil)
		service := employee.NewService(employeeRepositoryMock)

		employeeReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedEmployee, *employeeReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, sql.ErrNoRows)
		service := employee.NewService(employeeRepositoryMock)

		employeeReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, employeeReceived)
		assert.Equal(t, employee.ErrNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedEmployee := &[]domain.Employee{
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

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("GetAll", ctx).Return(*expectedEmployee, nil)
		service := employee.NewService(employeeRepositoryMock)

		employeeReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedEmployee, *employeeReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("GetAll", ctx).Return([]domain.Employee{}, errors.New("error"))
		service := employee.NewService(employeeRepositoryMock)

		employeeReceived, err := service.GetAll(&ctx)

		assert.Nil(t, employeeReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		employeeToDelete := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*employeeToDelete, nil)
		employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := employee.NewService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, nil, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(employee.ErrNotFound)
		service := employee.NewService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, employee.ErrNotFound, err)
	})
	t.Run("delete_not_found", func(t *testing.T) {

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, employee.ErrNotFound)
		//employeeRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(employee.ErrNotFound)
		service := employee.NewService(employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, employee.ErrNotFound, err)
	})

}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {
		employeeCreated := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := employee.NewService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *employeeCreated)

		assert.Equal(t, employee.ErrConflict, err)
		assert.Nil(t, employeeSaved)
	})

	t.Run("create_error", func(t *testing.T) {
		employeeCreated := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Employee")).Return(0, errors.New("error"))

		service := employee.NewService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *employeeCreated)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, employeeSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedEmployeeCreate := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Employee")).Return(1, nil)

		service := employee.NewService(employeeRepositoryMock)

		employeeSaved, err := service.Save(&ctx, *expectedEmployeeCreate)

		assert.Equal(t, employeeSaved, expectedEmployeeCreate)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		originalEmployee := &domain.Employee{
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

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "2",
			FirstName:    "Test2",
			LastName:     "Test2",
			WarehouseID:  2,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalEmployee, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Employee")).Return(nil)

		service := employee.NewService(employeeRepositoryMock)

		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Equal(t, *UpdateEmployees, *expectedEmployee)
		assert.Nil(t, err)
	})

	t.Run("update_non_existing", func(t *testing.T) {

		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, employee.ErrNotFound)

		service := employee.NewService(employeeRepositoryMock)

		_, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Equal(t, employee.ErrNotFound, err)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {

		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, employee.ErrNotFound)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Employee")).Return(assert.AnError)

		service := employee.NewService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_get_conflict_error", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, employee.ErrNotFound)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := employee.NewService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_conflict_error_card_number_id", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := employee.NewService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})

	t.Run("update_repository_error", func(t *testing.T) {
		newCardNumberID := "2"
		newFirstName := "Test2"
		newLastName := "Test2"
		newWarehouseID := 2

		updateEmployeeRequest := &domain.RequestUpdateEmployee{
			CardNumberID: &newCardNumberID,
			FirstName:    &newFirstName,
			LastName:     &newLastName,
			WarehouseID:  &newWarehouseID,
		}

		ctx := context.TODO()

		employeeRepositoryMock := mocks.NewEmployeeRepositoryMock()
		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		employeeRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		employeeRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Employee")).Return(assert.AnError)

		service := employee.NewService(employeeRepositoryMock)
		UpdateEmployees, err := service.Update(&ctx, 1, updateEmployeeRequest)

		assert.Nil(t, UpdateEmployees)
		assert.Error(t, err)
	})
}
