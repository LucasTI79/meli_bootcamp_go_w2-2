package inbound_order_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	inbound_order "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_orders"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	// t.Run("get_ok", func(t *testing.T) {
	// 	expectedInboundOrders := domain.InboundOrders{
	// 		ID:             1,
	// 		OrderDate:      "teste",
	// 		OrderNumber:    "teste",
	// 		EmployeeID:     "teste",
	// 		ProductBatchID: "teste",
	// 		WarehouseID:    "teste",
	// 	}

	// 	r := inbound_order.NewRepository(fields{db}.db)

	// 	rows := sqlmock.NewRows([]string{"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id"}).
	// 		AddRow(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
	// 			expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID)

	// 	// id, order_date, order_number, employee_id, product_batch_id, warehouse_id
	// 	mock.ExpectQuery("SELECT * FROM inbound_orders WHERE id=?").
	// 		WithArgs(expectedInboundOrders.ID).
	// 		WillReturnRows(rows)

	// 	inboundOrdersReceived, err := r.Get(ctx, 1)

	// 	assert.Equal(t, expectedInboundOrders, inboundOrdersReceived)
	// 	assert.Nil(t, err)
	// })

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id"}).
			AddRow(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID)

		mock.ExpectQuery("SELECT * FROM inbound_orders WHERE id=?").
			WithArgs(expectedInboundOrders.ID).
			WillReturnRows(rows)

		productRecordReceived, err := r.Get(ctx, 99)

		assert.Equal(t, domain.InboundOrders{}, productRecordReceived)
		assert.NotNil(t, err)
	})
}

func TestRepositoryGetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	// t.Run("get_all_ok", func(t *testing.T) {

	// 	// expectedProductsRecords := []domain.ProductRecord{
	// 	// 	{
	// 	// 		ID:             1,
	// 	// 		LastUpdateDate: "Test",
	// 	// 		PurchasePrice:  1.1,
	// 	// 		SalePrice:      1.1,
	// 	// 		ProductId:      1,
	// 	// 	},
	// 	// 	{
	// 	// 		LastUpdateDate: "Test",
	// 	// 		PurchasePrice:  1.1,
	// 	// 		SalePrice:      1.1,
	// 	// 		ProductId:      1,
	// 	// 	},
	// 	// }

	// 	expectedInboundOrdersList := []domain.InboundOrders{
	// 		{
	// 			ID:             1,
	// 			OrderDate:      "teste",
	// 			OrderNumber:    "teste",
	// 			EmployeeID:     "teste",
	// 			ProductBatchID: "teste",
	// 			WarehouseID:    "teste",
	// 		},
	// 		{
	// 			ID:             2,
	// 			OrderDate:      "teste",
	// 			OrderNumber:    "teste",
	// 			EmployeeID:     "teste",
	// 			ProductBatchID: "teste",
	// 			WarehouseID:    "teste",
	// 		},
	// 	}

	// 	r := inbound_order.NewRepository(fields{db}.db)

	// 	rows := sqlmock.NewRows([]string{"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id"})

	// 	for _, expectedInboundOrders := range expectedInboundOrdersList {
	// 		rows.AddRow(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
	// 			expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID)
	// 	}
	// 	mock.ExpectQuery("SELECT * FROM inbound_orders").
	// 		WillReturnRows(rows)

	// 	inboundOrdersReceived, err := r.GetAll(ctx)

	// 	assert.Equal(t, expectedInboundOrdersList, inboundOrdersReceived)
	// 	assert.Nil(t, err)
	// })

	t.Run("get_all_error", func(t *testing.T) {

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT * FROM inbound_orders").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		inboundOrdersReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.InboundOrders(nil), inboundOrdersReceived)
		assert.NotNil(t, err)
	})
}

func TestRepositoryExists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	// t.Run("exists_true", func(t *testing.T) {

	// 	expectedInboundOrders := domain.InboundOrders{
	// 		ID:             1,
	// 		OrderDate:      "teste",
	// 		OrderNumber:    "teste",
	// 		EmployeeID:     "teste",
	// 		ProductBatchID: "teste",
	// 		WarehouseID:    "teste",
	// 	}

	// 	r := inbound_order.NewRepository(fields{db}.db)

	// 	rows := sqlmock.NewRows([]string{"id"}).
	// 		AddRow(expectedInboundOrders.ID)

	// 	mock.ExpectQuery("SELECT id FROM inbound_orders WHERE id=?").
	// 		WithArgs(expectedInboundOrders.ID).
	// 		WillReturnRows(rows)

	// 	strID := fmt.Sprint(expectedInboundOrders.ID)

	// 	inboundOrdersExists := r.Exists(ctx, strID)

	// 	assert.True(t, inboundOrdersExists)
	// })

	t.Run("exists_false", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(expectedInboundOrders.ID)

		mock.ExpectQuery("SELECT id FROM inbound_orders WHERE id=?").
			WithArgs(expectedInboundOrders.ID).
			WillReturnRows(rows)

		productExists := r.Exists(ctx, "2")

		assert.False(t, productExists)
	})

}

func TestRepositoryDelete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("delete_ok", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).
			WithArgs(expectedInboundOrders.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))

		err := r.Delete(ctx, expectedInboundOrders.ID)

		assert.Nil(t, err)
	})

	t.Run("delete_error_exec", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).
			WithArgs(expectedInboundOrders.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedInboundOrders.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_rowlsAffected0", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).
			WithArgs(expectedInboundOrders.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedInboundOrders.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_prepare", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).
			WithArgs(expectedInboundOrders.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedInboundOrders.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_not_found", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM inbound_orders WHERE id=?")).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := r.Delete(ctx, expectedInboundOrders.ID)

		assert.NotNil(t, err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	// t.Run("update_ok", func(t *testing.T) {

	// 	expectedInboundOrders := domain.InboundOrders{
	// 		ID:             1,
	// 		OrderDate:      "teste",
	// 		OrderNumber:    "teste",
	// 		EmployeeID:     "teste",
	// 		ProductBatchID: "teste",
	// 		WarehouseID:    "teste",
	// 	}

	// 	r := inbound_order.NewRepository(fields{db}.db)

	// 	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?"))
	// 	mock.ExpectExec(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?")).
	// 		WithArgs(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
	// 			expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
	// 		WillReturnResult(sqlmock.NewResult(1, 1))

	// 	err := r.Update(ctx, expectedInboundOrders)

	// 	assert.Nil(t, err)
	// })

	t.Run("update_error_exec", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})

	t.Run("update_error_rowlsAffected0", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})

	t.Run("update_error_prepare", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("UPDATE inbound_orders SET order_date=?, order_number=?, employee_id=?, product_batch_id=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedInboundOrders.ID, expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})
}

func TestRepositorySave(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("save_ok", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedInboundOrders)
		assert.Equal(t, expectedInboundOrders.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {

		expectedInboundOrders := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		r := inbound_order.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO inbound_orders(order_date, order_number, employee_id, product_batch_id, warehouse_id) VALUES (?,?,?,?,?)")).
			WithArgs(expectedInboundOrders.OrderDate, expectedInboundOrders.OrderNumber,
				expectedInboundOrders.EmployeeID, expectedInboundOrders.ProductBatchID, expectedInboundOrders.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedInboundOrders)

		assert.NotNil(t, err)
	})
}
