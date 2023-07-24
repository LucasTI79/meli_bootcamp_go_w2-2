package inbound_order_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	employee_mocks "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee/mocks"
	inbound_order "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_orders/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		expectedInboundOrders := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedInboundOrders, nil)
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedInboundOrders, *inboundOrdersReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, sql.ErrNoRows)
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, inboundOrdersReceived)
		assert.Equal(t, inbound_order.ErrNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedInboundOrders := &[]domain.InboundOrders{
			{
				ID:             1,
				OrderDate:      "teste",
				OrderNumber:    "teste",
				EmployeeID:     "teste",
				ProductBatchID: "teste",
				WarehouseID:    "teste",
			},
			{
				ID:             2,
				OrderDate:      "teste",
				OrderNumber:    "teste",
				EmployeeID:     "teste",
				ProductBatchID: "teste",
				WarehouseID:    "teste",
			},
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("GetAll", ctx).Return(*expectedInboundOrders, nil)
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedInboundOrders, *inboundOrdersReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("GetAll", ctx).Return([]domain.InboundOrders{}, errors.New("error"))
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersReceived, err := service.GetAll(&ctx)

		assert.Nil(t, inboundOrdersReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		inboundOrdersDeleted := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*inboundOrdersDeleted, nil)
		inboundOrdersRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, nil, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, nil)
		inboundOrdersRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(inbound_order.ErrNotFound)
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, inbound_order.ErrNotFound, err)
	})
	t.Run("delete_not_found", func(t *testing.T) {

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, inbound_order.ErrNotFound)
		//inboundOrdersRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(employee.ErrNotFound)
		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		err := service.Delete(&ctx, 1)

		assert.Equal(t, inbound_order.ErrNotFound, err)
	})

}

func TestCreate(t *testing.T) {
	// t.Run("create_conflict", func(t *testing.T) {
	// 	inboundOrdersCreated := &domain.InboundOrders{
	// 		ID:             1,
	// 		OrderDate:      "teste",
	// 		OrderNumber:    "teste",
	// 		EmployeeID:     "teste",
	// 		ProductBatchID: "teste",
	// 		WarehouseID:    "teste",
	// 	}

	// 	ctx := context.TODO()

	// 	inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
	// 	employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

	// 	inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

	// 	service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

	// 	inboundOrdersSaved, err := service.Save(&ctx, *inboundOrdersCreated)

	// 	assert.Equal(t, inbound_order.ErrConflict, err)
	// 	assert.Nil(t, inboundOrdersSaved)
	// })

	t.Run("create_error", func(t *testing.T) {
		inboundOrdersCreated := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		// inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		inboundOrdersRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.InboundOrders")).Return(0, errors.New("error"))

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersSaved, err := service.Save(&ctx, *inboundOrdersCreated)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, inboundOrdersSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedInboundOrdersCreated := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		employeeRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		// inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		inboundOrdersRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.InboundOrders")).Return(1, nil)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		inboundOrdersSaved, err := service.Save(&ctx, *expectedInboundOrdersCreated)

		assert.Equal(t, inboundOrdersSaved, expectedInboundOrdersCreated)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		originalInboundOrders := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		expectedInboundOrders := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "updated",
			OrderNumber:    "updated",
			EmployeeID:     "updated",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalInboundOrders, nil)
		inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		inboundOrdersRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.InboundOrders")).Return(nil)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		updateInbound_order, err := service.Update(&ctx, 1, &requestUpdateInboundOrders)

		assert.Equal(t, *updateInbound_order, *expectedInboundOrders)
		assert.Nil(t, err)
	})

	t.Run("update_non_existing", func(t *testing.T) {

		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, inbound_order.ErrNotFound)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)

		_, err := service.Update(&ctx, 1, &requestUpdateInboundOrders)

		assert.Equal(t, inbound_order.ErrNotFound, err)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {

		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := &domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, inbound_order.ErrNotFound)
		inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		inboundOrdersRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.InboundOrders")).Return(assert.AnError)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)
		UpdateInboundOrders, err := service.Update(&ctx, 1, requestUpdateInboundOrders)

		assert.Nil(t, UpdateInboundOrders)
		assert.Error(t, err)
	})

	t.Run("update_get_conflit_error", func(t *testing.T) {
		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := &domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, employee.ErrNotFound)
		inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)
		UpdateInboundOrders, err := service.Update(&ctx, 1, requestUpdateInboundOrders)

		assert.Nil(t, UpdateInboundOrders)
		assert.Error(t, err)
	})

	// t.Run("update_conflit_error_card_number_id", func(t *testing.T) {
	// 	orderDate := "updated"
	// 	orderNumber := "updated"
	// 	employeeID := "updated"

	// 	requestUpdateInboundOrders := &domain.RequestUpdateInboundOrders{
	// 		OrderDate:   &orderDate,
	// 		OrderNumber: &orderNumber,
	// 		EmployeeID:  &employeeID,
	// 	}

	// 	ctx := context.TODO()

	// 	inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
	// employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

	// 	inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, nil)
	// 	inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

	// 	service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)
	// 	UpdateInboundOrders, err := service.Update(&ctx, 1, requestUpdateInboundOrders)

	// 	assert.Nil(t, UpdateInboundOrders)
	// 	assert.Error(t, err)
	// })

	t.Run("update_repository_error", func(t *testing.T) {
		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := &domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		ctx := context.TODO()

		inboundOrdersRepositoryMock := mocks.NewInboundOrdersRepositoryMock()
		employeeRepositoryMock := new(employee_mocks.EmployeeRepositoryMock)

		inboundOrdersRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.InboundOrders{}, nil)
		inboundOrdersRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		inboundOrdersRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.InboundOrders")).Return(assert.AnError)

		service := inbound_order.NewService(inboundOrdersRepositoryMock, employeeRepositoryMock)
		UpdateInboundOrders, err := service.Update(&ctx, 1, requestUpdateInboundOrders)

		assert.Nil(t, UpdateInboundOrders)
		assert.Error(t, err)
	})
}
