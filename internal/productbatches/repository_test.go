package productbatches_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/stretchr/testify/assert"
)

var (
	expectedQuery = `SELECT count\(pb\.id\) as ` + "`products_count`" + `, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id WHERE pb.section_id = \? GROUP BY pb.section_id`
	query         = `SELECT count\(pb\.id\) as ` + "`products_count`" + `, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id GROUP BY pb.section_id`
)

func TestRepositoryExistsProductBatch(t *testing.T) {

	expectedPB := domain.ProductBatches{
		ID:                 1,
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("exist_productBatch", func(t *testing.T) {

		r := productbatches.NewRepository(db)

		rows := sqlmock.NewRows([]string{"batch_number"}).
			AddRow(expectedPB.BatchNumber)

		mock.ExpectQuery(productbatches.ExistProductBatch).
			WithArgs(expectedPB.BatchNumber).
			WillReturnRows(rows)

		existsExistsProductBatch := r.ExistsProductBatch(ctx, expectedPB.BatchNumber)

		assert.True(t, existsExistsProductBatch)
	})
}

func TestRepositoryGet(t *testing.T) {

	expectedPB := domain.ProductBatches{
		ID:                 1,
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	ctx := context.TODO()

	t.Run("Get", func(t *testing.T) {

		r := productbatches.NewRepository(db)

		rows := sqlmock.NewRows([]string{"id", "batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}).
			AddRow(expectedPB.ID, expectedPB.BatchNumber, expectedPB.CurrentQuantity, expectedPB.CurrentTemperature, expectedPB.DueDate, expectedPB.InitialQuantity, expectedPB.ManufacturingDate, expectedPB.ManufacturingHour, expectedPB.MinimumTemperature, expectedPB.ProductID, expectedPB.SectionID)

		mock.ExpectQuery(productbatches.Get).WithArgs(expectedPB.ID).WillReturnRows(rows)

		productBatchActual, err := r.Get(ctx, expectedPB.ID)

		assert.Equal(t, expectedPB, productBatchActual)
		assert.Equal(t, err, nil)
	})
}
func TestRepositorySave(t *testing.T) {
	// type fields struct {
	// 	db *sql.DB
	// }
	db, mock, _ := sqlmock.New()
	ctx := context.TODO()
	t.Run("SAVE - OK", func(t *testing.T) {
		expectedPB := domain.ProductBatches{
			ID:                 1,
			BatchNumber:        123,
			CurrentQuantity:    10,
			CurrentTemperature: 25.5,
			DueDate:            "2023-07-10",
			InitialQuantity:    100,
			ManufacturingDate:  "2023-07-01",
			ManufacturingHour:  8,
			MinimumTemperature: 20.0,
			ProductID:          456,
			SectionID:          789,
		}

		r := productbatches.NewRepository(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedPB.BatchNumber, expectedPB.CurrentQuantity, expectedPB.CurrentTemperature, expectedPB.DueDate, expectedPB.InitialQuantity, expectedPB.ManufacturingDate, expectedPB.ManufacturingHour, expectedPB.MinimumTemperature, expectedPB.ProductID, expectedPB.SectionID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedPB)
		assert.Equal(t, expectedPB.ID, id)
		assert.Nil(t, err)
	})

	t.Run("SAVE - Error - Exec", func(t *testing.T) {
		expectedPB := domain.ProductBatches{
			ID:                 1,
			BatchNumber:        123,
			CurrentQuantity:    10,
			CurrentTemperature: 25.5,
			DueDate:            "2023-07-10",
			InitialQuantity:    100,
			ManufacturingDate:  "2023-07-01",
			ManufacturingHour:  8,
			MinimumTemperature: 20.0,
			ProductID:          456,
			SectionID:          789,
		}

		r := productbatches.NewRepository(db)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedPB.BatchNumber, expectedPB.CurrentQuantity, expectedPB.CurrentTemperature, expectedPB.DueDate, expectedPB.InitialQuantity, expectedPB.ManufacturingDate, expectedPB.ManufacturingHour, expectedPB.MinimumTemperature, expectedPB.ProductID, expectedPB.SectionID).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedPB)

		assert.NotNil(t, err)
	})

	t.Run("SAVE - Error - RowlsAffected0", func(t *testing.T) {
		expectedPB := domain.ProductBatches{
			ID:                 1,
			BatchNumber:        123,
			CurrentQuantity:    10,
			CurrentTemperature: 25.5,
			DueDate:            "2023-07-10",
			InitialQuantity:    100,
			ManufacturingDate:  "2023-07-01",
			ManufacturingHour:  8,
			MinimumTemperature: 20.0,
			ProductID:          456,
			SectionID:          789,
		}
		r := productbatches.NewRepository(db)
		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedPB.BatchNumber, expectedPB.CurrentQuantity, expectedPB.CurrentTemperature, expectedPB.DueDate, expectedPB.InitialQuantity, expectedPB.ManufacturingDate, expectedPB.ManufacturingHour, expectedPB.MinimumTemperature, expectedPB.ProductID, expectedPB.SectionID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedPB)
		assert.NotNil(t, err)
	})
	t.Run("SAVE - Error - Prepare", func(t *testing.T) {
		expectedPB := domain.ProductBatches{
			ID:                 1,
			BatchNumber:        123,
			CurrentQuantity:    10,
			CurrentTemperature: 25.5,
			DueDate:            "2023-07-10",
			InitialQuantity:    100,
			ManufacturingDate:  "2023-07-01",
			ManufacturingHour:  8,
			MinimumTemperature: 20.0,
			ProductID:          456,
			SectionID:          789,
		}
		r := productbatches.NewRepository(db)
		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)")).
			WithArgs(expectedPB.BatchNumber, expectedPB.CurrentQuantity, expectedPB.CurrentTemperature, expectedPB.DueDate, expectedPB.InitialQuantity, expectedPB.ManufacturingDate, expectedPB.ManufacturingHour, expectedPB.MinimumTemperature, expectedPB.ProductID, expectedPB.SectionID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedPB)
		assert.NotNil(t, err)
	})
}
func TestRepositorySectionProductsReports(t *testing.T) {
	db, mock, _ := sqlmock.New()
	t.Run("GET ALL - Section Products Reports", func(t *testing.T) {
		expectedReportProducts := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
			{
				SectionID:     2,
				SectionNumber: "20",
				ProductsCount: 200,
			},
		}
		r := productbatches.NewRepository(db)

		rows := sqlmock.NewRows([]string{"products_count", "section_id", "section_number"})

		for _, reportProduct := range expectedReportProducts {
			rows.AddRow(reportProduct.ProductsCount, reportProduct.SectionID, reportProduct.SectionNumber)
		}
		mock.ExpectQuery(query).WillReturnRows(rows)

		actualReportProducts, err := r.SectionProductsReports()

		assert.Equal(t, expectedReportProducts, actualReportProducts)
		assert.Nil(t, err)
	})

	t.Run("GET ALL - Error", func(t *testing.T) {
		r := productbatches.NewRepository(db)
		mock.ExpectQuery(query).
			WithArgs().
			WillReturnError(sql.ErrNoRows)

		actualReportProducts, err := r.SectionProductsReports()

		assert.Equal(t, []domain.ProductBySection(nil), actualReportProducts)
		assert.NotNil(t, err)
	})
}
func TestRepositorySectionProductsReportsBySection(t *testing.T) {
	db, mock, _ := sqlmock.New()
	t.Run("SectionID - TRUE", func(t *testing.T) {
		expectedReportProductsBySection := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
		}

		r := productbatches.NewRepository(db)
		rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"})
		for _, reportProduct := range expectedReportProductsBySection {
			rows.AddRow(reportProduct.SectionID, reportProduct.SectionNumber, reportProduct.ProductsCount)
		}
		mock.ExpectQuery(expectedQuery).WithArgs(expectedReportProductsBySection[0].SectionID).WillReturnRows(rows)

		actualReportProductsBySection, error := r.SectionProductsReportsBySection(1)

		assert.Equal(t, actualReportProductsBySection, expectedReportProductsBySection)
		assert.Nil(t, error)
	})
	t.Run("SectionID - FALSE", func(t *testing.T) {
		r := productbatches.NewRepository(db)
		rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"})
		mock.ExpectQuery(expectedQuery).WillReturnRows(rows)

		actualReportProductsBySection, error := r.SectionProductsReportsBySection(3)

		assert.Empty(t, actualReportProductsBySection, []domain.ProductBySection{})
		assert.NotNil(t, error)
	})
}
