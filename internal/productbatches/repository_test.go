package productbatches_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/stretchr/testify/assert"
)

// var (
// expected = domain.ProductBatches{
// 	ID:                 1,
// 	BatchNumber:        123,
// 	CurrentQuantity:    10,
// 	CurrentTemperature: 25.5,
// 	DueDate:            "2023-07-10",
// 	InitialQuantity:    100,
// 	ManufacturingDate:  "2023-07-01",
// 	ManufacturingHour:  8,
// 	MinimumTemperature: 20.0,
// 	ProductID:          456,
// 	SectionID:          789,
// }
// 	payloadPB = domain.ProductBatches{
// 		BatchNumber:        123,
// 		CurrentQuantity:    10,
// 		CurrentTemperature: 25.5,
// 		DueDate:            "2023-07-10",
// 		InitialQuantity:    100,
// 		ManufacturingDate:  "2023-07-01",
// 		ManufacturingHour:  8,
// 		MinimumTemperature: 20.0,
// 		ProductID:          456,
// 		SectionID:          789,
// 	}
// )

func InitDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, context.Context) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.TODO()
	return db, mock, ctx
}
func TestRepositoryExistsProductBatch(t *testing.T) {
	t.Run("exist_productBatch", func(t *testing.T) {
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
		db, mock, ctx := InitDB(t)
		r := productbatches.NewRepository(db)
		rows := sqlmock.NewRows([]string{"batch_number"}).
			AddRow(expectedPB.BatchNumber)
		mock.ExpectQuery("SELECT batch_number FROM product_batches WHERE batch_number=?;").WithArgs(expectedPB.BatchNumber).WillReturnRows(rows)
		existsExistsProductBatch := r.ExistsProductBatch(ctx, expectedPB.BatchNumber)
		assert.True(t, existsExistsProductBatch)
	})
}
func TestGet(t *testing.T) {
	db, mock, ctx := InitDB(t)
	t.Run("Get", func(t *testing.T) {
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
		rows := sqlmock.NewRows([]string{"id"}).
			AddRow(expectedPB.ID)
		mock.ExpectQuery("SELECT * FROM product_batches WHERE id=?;").WithArgs(expectedPB.ID).WillReturnRows(rows)
		id := 1
		productBatchActual, err := r.Get(ctx, id)
		assert.Equal(t, productBatchActual, expectedPB)
		assert.Equal(t, err, nil)
	})
}

// func (r *repository) Get(ctx context.Context, id int) (domain.ProductBatches, error) {
// 	query := "SELECT * FROM product_batches WHERE id=?;"
// 	row := r.db.QueryRow(query, id)
// 	pb := domain.ProductBatches{}
// 	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
// 	if err != nil {
// 		return domain.ProductBatches{}, err
// 	}

// 	return pb, nil
// }
