package product_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("get_ok", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "net_weight",
			"product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"}).
			AddRow(expectedProduct.ID, expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID)

		mock.ExpectQuery("SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products WHERE id=?").
			WithArgs(expectedProduct.ID).
			WillReturnRows(rows)

		productReceived, err := r.Get(ctx, 1)

		assert.Equal(t, expectedProduct, productReceived)
		assert.Nil(t, err)
	})

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "net_weight",
			"product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"}).
			AddRow(expectedProduct.ID, expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID)

		mock.ExpectQuery("SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products WHERE id=?").
			WithArgs(expectedProduct.ID).
			WillReturnRows(rows)

		productReceived, err := r.Get(ctx, 99)

		assert.Equal(t, domain.Product{}, productReceived)
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

		expectedProducts := []domain.Product{
			{
				ID:             1,
				Description:    "Test",
				ExpirationRate: 1,
				FreezingRate:   1,
				Height:         1.1,
				Length:         1.1,
				Netweight:      1.1,
				ProductCode:    "Teste",
				RecomFreezTemp: 1.1,
				Width:          1.1,
				ProductTypeID:  1,
				SellerID:       1,
			},
			{
				Description:    "Teste",
				ExpirationRate: 1,
				FreezingRate:   1,
				Height:         1.1,
				Length:         1.1,
				Netweight:      1.1,
				ProductCode:    "Teste",
				RecomFreezTemp: 1.1,
				Width:          1.1,
				ProductTypeID:  1,
				SellerID:       1,
			},
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "net_weight",
			"product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"})

		for _, expectedProduct := range expectedProducts {
			rows.AddRow(expectedProduct.ID, expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID)
		}
		mock.ExpectQuery("SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products").
			WillReturnRows(rows)

		productsReceived, err := r.GetAll(ctx)

		assert.Equal(t, expectedProducts, productsReceived)
		assert.Nil(t, err)
	})

	t.Run("get_all_error", func(t *testing.T) {

		r := product.NewRepository(fields{db}.db)

		mock.ExpectQuery("SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products").
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		productsReceived, err := r.GetAll(ctx)

		assert.Equal(t, []domain.Product(nil), productsReceived)
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

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"product_code"}).
			AddRow(expectedProduct.ProductCode)

		mock.ExpectQuery("SELECT product_code FROM products WHERE product_code=?").
			WithArgs(expectedProduct.ProductCode).
			WillReturnRows(rows)

		productExists := r.Exists(ctx, expectedProduct.ProductCode)

		assert.True(t, productExists)
	})

	t.Run("exists_false", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"product_code"}).
			AddRow(expectedProduct.ProductCode)

		mock.ExpectQuery("SELECT product_code FROM products WHERE product_code=?").
			WithArgs(expectedProduct.ProductCode).
			WillReturnRows(rows)

		productExists := r.Exists(ctx, "Test2")

		assert.False(t, productExists)
	})
}

func TestRepositoryExistsByID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("exists_byID_true", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(expectedProduct.ID)

		mock.ExpectQuery("SELECT id FROM products WHERE id=?").
			WithArgs(expectedProduct.ID).
			WillReturnRows(rows)

		productExists := r.ExistsByID(ctx, expectedProduct.ID)

		assert.True(t, productExists)
	})

	t.Run("exists_byID_false", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(expectedProduct.ProductCode)

		mock.ExpectQuery("SELECT id FROM products WHERE id=?").
			WithArgs(expectedProduct.ProductCode).
			WillReturnRows(rows)

		productExists := r.ExistsByID(ctx, 99)

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

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedProduct.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))

		err := r.Delete(ctx, expectedProduct.ID)

		assert.Nil(t, err)
	})

	t.Run("delete_error_exec", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedProduct.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedProduct.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_rowlsAffected0", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedProduct.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedProduct.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_not_found", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedProduct.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := r.Delete(ctx, expectedProduct.ID)

		assert.NotNil(t, err)
	})

	t.Run("delete_error_prepare", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedProduct.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedProduct.ID)

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

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID, expectedProduct.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := r.Update(ctx, expectedProduct)

		assert.Nil(t, err)
	})

	t.Run("update_error_exec", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID, expectedProduct.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Update(ctx, expectedProduct)

		assert.NotNil(t, err)
	})

	t.Run("update_error_rowlsAffected0", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID, expectedProduct.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedProduct)

		assert.NotNil(t, err)
	})

	t.Run("update_error_prepare", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID, expectedProduct.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedProduct)

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

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedProduct)
		assert.Equal(t, expectedProduct.ID, id)
		assert.Nil(t, err)
	})

	t.Run("save_error_exec", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedProduct)

		assert.NotNil(t, err)
	})

	t.Run("save_error_rowlsAffected0", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedProduct)

		assert.NotNil(t, err)
	})

	t.Run("save_error_prepare", func(t *testing.T) {

		expectedProduct := domain.Product{
			ID:             1,
			Description:    "Test2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         1.2,
			Length:         1.2,
			Netweight:      1.2,
			ProductCode:    "Test2",
			RecomFreezTemp: 1.2,
			Width:          1.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		r := product.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code,recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedProduct.Description, expectedProduct.ExpirationRate, expectedProduct.FreezingRate,
				expectedProduct.Height, expectedProduct.Length, expectedProduct.Netweight, expectedProduct.ProductCode,
				expectedProduct.RecomFreezTemp, expectedProduct.Width, expectedProduct.ProductTypeID, expectedProduct.SellerID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedProduct)

		assert.NotNil(t, err)
	})
}
