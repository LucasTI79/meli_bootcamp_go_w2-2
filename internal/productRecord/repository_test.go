package productRecord_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_ok", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).
			AddRow(expectedProductRecord.ID, expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId)

		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id=?").
			WithArgs(expectedProductRecord.ID).
			WillReturnRows(rows)

		productRecordReceived, err := r.Get(ctx, 1)

		assert.Equal(t, expectedProductRecord, productRecordReceived)
		assert.Nil(t, err)
	})

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"}).
			AddRow(expectedProductRecord.ID, expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId)

		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id=?").
			WithArgs(expectedProductRecord.ID).
			WillReturnRows(rows)

		productRecordReceived, err := r.Get(ctx, 99)

		assert.Equal(t, domain.ProductRecord{}, productRecordReceived)
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

		expectedProductsRecords := []domain.ProductRecord{
			{
				ID:             1,
				LastUpdateDate: "Test",
				PurchasePrice:  1.1,
				SalePrice:      1.1,
				ProductId:      1,
			},
			{
				LastUpdateDate: "Test",
				PurchasePrice:  1.1,
				SalePrice:      1.1,
				ProductId:      1,
			},
		}

		r := productRecord.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "last_update_date", "purchase_price", "sale_price", "product_id"})

		for _, expectedProductRecord := range expectedProductsRecords {
			rows.AddRow(expectedProductRecord.ID, expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId)
		}
		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records").
			WillReturnRows(rows)

		productsRecordsReceived, err := r.GetAll(ctx)

		assert.Equal(t, expectedProductsRecords, productsRecordsReceived)
		assert.Nil(t, err)
	})

	t.Run("get_all_error", func(t *testing.T) {

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		productsRecordsReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.ProductRecord(nil), productsRecordsReceived)
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

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"product_id"}).
			AddRow(expectedProductRecord.ProductId)

		mock.ExpectQuery("SELECT product_id FROM product_records WHERE product_id=?").
			WithArgs(expectedProductRecord.ProductId).
			WillReturnRows(rows)

		productRecordExists := r.Exists(ctx, expectedProductRecord.ProductId)

		assert.True(t, productRecordExists)
	})

	t.Run("exists_false", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"product_id"}).
			AddRow(expectedProductRecord.ProductId)

		mock.ExpectQuery("SELECT product_id FROM product_records WHERE product_id=?").
			WithArgs(expectedProductRecord.ProductId).
			WillReturnRows(rows)

		productExists := r.Exists(ctx, 2)

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

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).
			WithArgs(expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))

		err := r.Delete(ctx, expectedProductRecord.ID)

		assert.Nil(t, err)
	})

	t.Run("delete_error_exec", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).
			WithArgs(expectedProductRecord.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedProductRecord.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_rowlsAffected0", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).
			WithArgs(expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedProductRecord.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_prepare", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).
			WithArgs(expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedProductRecord.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_not_found", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM product_records WHERE id=?")).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := r.Delete(ctx, expectedProductRecord.ID)

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

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId, expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := r.Update(ctx, expectedProductRecord)

		assert.Nil(t, err)
	})

	t.Run("update_error_exec", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId, expectedProductRecord.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, expectedProductRecord)

		assert.NotNil(t, err)
	})

	t.Run("update_error_rowlsAffected0", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId, expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedProductRecord)

		assert.NotNil(t, err)
	})

	t.Run("update_error_prepare", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId, expectedProductRecord.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedProductRecord)

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

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedProductRecord)
		assert.Equal(t, expectedProductRecord.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedProductRecord)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedProductRecord)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {

		expectedProductRecord := domain.ProductRecord{
			ID:             1,
			LastUpdateDate: "Test2",
			PurchasePrice:  1.2,
			SalePrice:      1.2,
			ProductId:      2,
		}

		r := productRecord.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)")).
			WithArgs(expectedProductRecord.LastUpdateDate, expectedProductRecord.PurchasePrice, expectedProductRecord.SalePrice,
				expectedProductRecord.ProductId).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedProductRecord)

		assert.NotNil(t, err)
	})
}
