package productbatches_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/stretchr/testify/assert"
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

func TestGet(t *testing.T) {

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
