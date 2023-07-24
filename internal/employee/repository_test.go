package employee_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_ok", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
			AddRow(expectedEmployee.ID, expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID)

		mock.ExpectQuery("SELECT *. FROM employees WHERE id=?").
			WithArgs(expectedEmployee.ID).
			WillReturnRows(rows)

		productReceived, err := r.Get(ctx, 1)

		assert.Equal(t, expectedEmployee, productReceived)
		assert.Nil(t, err)
	})

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
			AddRow(expectedEmployee.ID, expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID)

		mock.ExpectQuery("SELECT * FROM employees WHERE id=?").
			WithArgs(expectedEmployee.ID).
			WillReturnRows(rows)

		employeeReceived, err := r.Get(ctx, 99)

		assert.Equal(t, domain.Employee{}, employeeReceived)
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

		expectedEmployees := []domain.Employee{
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

		r := employee.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})

		for _, expectedEmployee := range expectedEmployees {
			rows.AddRow(expectedEmployee.ID, expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID)
		}
		mock.ExpectQuery("SELECT *. FROM employees").
			WillReturnRows(rows)

		employeeReceived, err := r.GetAll(ctx)

		assert.Equal(t, expectedEmployees, employeeReceived)
		assert.Nil(t, err)
	})

	t.Run("get_all_error", func(t *testing.T) {

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT *. FROM employees").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		employeeReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.Employee(nil), employeeReceived)
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

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"card_number_id"}).
			AddRow(expectedEmployee.CardNumberID)

		mock.ExpectQuery("SELECT card_number_id FROM employees WHERE card_number_id=?").
			WithArgs(expectedEmployee.CardNumberID).
			WillReturnRows(rows)

		employeeExists := r.Exists(ctx, expectedEmployee.CardNumberID)

		assert.True(t, employeeExists)
	})

	t.Run("exists_false", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"card_number_id"}).
			AddRow(expectedEmployee.CardNumberID)

		mock.ExpectQuery("SELECT card_number_id FROM employees WHERE card_number_id=?").
			WithArgs(expectedEmployee.CardNumberID).
			WillReturnRows(rows)

		employeeExists := r.Exists(ctx, "Test2")

		assert.False(t, employeeExists)
	})
}

func TestRepositoryDelete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("delete_ok", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
			WithArgs(expectedEmployee.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))

		err := r.Delete(ctx, expectedEmployee.ID)

		assert.Nil(t, err)
	})

	t.Run("delete_error_exec", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
			WithArgs(expectedEmployee.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedEmployee.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_rowlsAffected0", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
			WithArgs(expectedEmployee.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedEmployee.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_prepare", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
			WithArgs(expectedEmployee.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedEmployee.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_not_found", func(t *testing.T) {

		expectedEmployee := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := r.Delete(ctx, expectedEmployee.ID)

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

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID, expectedEmployee.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := r.Update(ctx, expectedEmployee)

		assert.Nil(t, err)
	})

	t.Run("update_error_exec", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID, expectedEmployee.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, expectedEmployee)

		assert.NotNil(t, err)
	})

	t.Run("update_error_rowlsAffected0", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID, expectedEmployee.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedEmployee)

		assert.NotNil(t, err)
	})

	t.Run("update_error_prepare", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET card_number_id=?, first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID, expectedEmployee.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedEmployee)

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

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedEmployee)
		assert.Equal(t, expectedEmployee.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedEmployee)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedEmployee)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {

		expectedEmployee := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		r := employee.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code,recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
			WithArgs(expectedEmployee.CardNumberID, expectedEmployee.FirstName, expectedEmployee.LastName,
				expectedEmployee.WarehouseID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedEmployee)

		assert.NotNil(t, err)
	})
}
