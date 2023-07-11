package warehouse_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_ok", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}).
			AddRow(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode,
				expectedWarehouse.MinimumCapacity, expectedWarehouse.MinimumTemperature)

		mock.ExpectQuery("SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouses WHERE id=?").
			WithArgs(expectedWarehouse.ID).
			WillReturnRows(rows)

		warehouseReceived, err := r.Get(ctx, 1)

		assert.Equal(t, expectedWarehouse, warehouseReceived)
		assert.Nil(t, err)
	})

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}).
			AddRow(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode,
				expectedWarehouse.MinimumCapacity, expectedWarehouse.MinimumTemperature)

		mock.ExpectQuery("SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouses WHERE id=?").
			WithArgs(expectedWarehouse.ID).
			WillReturnRows(rows)

		warehouseReceived, err := r.Get(ctx, 11)

		assert.Equal(t, domain.Warehouse{}, warehouseReceived)
		assert.NotNil(t, err)
	})
}

func TestRepositoryGetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_all_ok", func(t *testing.T) {

		expectedWarehouses := []domain.Warehouse{
			{
				ID:                 1,
				Address:            "Rua Teste",
				Telephone:          "11938473125",
				WarehouseCode:      "CX-2281-TCD",
				MinimumCapacity:    12,
				MinimumTemperature: 18,
			},
			{
				ID:                 1,
				Address:            "Rua Teste",
				Telephone:          "11938473125",
				WarehouseCode:      "CX-2281-TCD",
				MinimumCapacity:    12,
				MinimumTemperature: 18,
			},
		}

		r := warehouse.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"})

		for _, expectedWarehouse := range expectedWarehouses {
			rows.AddRow(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode,
				expectedWarehouse.MinimumCapacity, expectedWarehouse.MinimumTemperature)
		}
		mock.ExpectQuery("SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouses").
			WillReturnRows(rows)

		warehousesReceived, err := r.GetAll(ctx)

		assert.Equal(t, expectedWarehouses, warehousesReceived)
		assert.Nil(t, err)
	})

	t.Run("get_all_error", func(t *testing.T) {

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT id, address, telephone, warehouse_code, minimum_capacity, minimum_temperature FROM warehouses").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		warehousesReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.Warehouse(nil), warehousesReceived)
		assert.NotNil(t, err)
	})
}

func TestRepositoryExists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("exists_true", func(t *testing.T) {
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"warehouse_code"}).
			AddRow(expectedWarehouse.WarehouseCode)

		mock.ExpectQuery("SELECT warehouse_code FROM warehouses WHERE warehouse_code=?").
			WithArgs(expectedWarehouse.WarehouseCode).
			WillReturnRows(rows)

		expectedW := r.Exists(ctx, expectedWarehouse.WarehouseCode)

		assert.True(t, expectedW)
	})

	t.Run("exists_false", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"warehouse_code"}).
			AddRow(expectedWarehouse.WarehouseCode)

		mock.ExpectQuery("SELECT warehouse_code FROM warehouses WHERE warehouse_code=?").
			WithArgs(expectedWarehouse.WarehouseCode).
			WillReturnRows(rows)

		expectedW := r.Exists(ctx, "Teste")

		assert.False(t, expectedW)
	})
}

func TestRepositoryDelete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("delete_ok", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?")).
			WithArgs(expectedWarehouse.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))

		err := r.Delete(ctx, expectedWarehouse.ID)

		assert.Nil(t, err)
	})

	t.Run("delete_error_exec", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?")).
			WithArgs(expectedWarehouse.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedWarehouse.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_rowlsAffected0", func(t *testing.T) {
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?")).
			WithArgs(expectedWarehouse.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedWarehouse.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_prepare", func(t *testing.T) {
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?")).
			WithArgs(expectedWarehouse.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedWarehouse.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_not_found", func(t *testing.T) {
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses WHERE id=?")).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := r.Delete(ctx, expectedWarehouse.ID)

		assert.NotNil(t, err)
	})
}

func TestRepositoryUpdate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("update_ok", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := r.Update(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})

	t.Run("update_error_exec", func(t *testing.T) {
		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})

	t.Run("update_error_rowlsAffected0", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Update(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})

	t.Run("update_error_prepare", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Update(ctx, expectedWarehouse)

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

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)")).
			WithArgs(expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedWarehouse)
		assert.Equal(t, expectedWarehouse.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {

		expectedWarehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		r := warehouse.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)")).
			WithArgs(expectedWarehouse.ID, expectedWarehouse.Address, expectedWarehouse.Telephone, expectedWarehouse.WarehouseCode, expectedWarehouse.MinimumCapacity,
				expectedWarehouse.MinimumTemperature).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedWarehouse)

		assert.NotNil(t, err)
	})
}
