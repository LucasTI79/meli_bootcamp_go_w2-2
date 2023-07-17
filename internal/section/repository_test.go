package section_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGet(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("GET - OK", func(t *testing.T) {
		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}

		r := section.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity",
			"warehouse_id", "product_type_id"}).
			AddRow(expectedSection.ID, expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID)
		query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id=?"
		mock.ExpectQuery(query).WithArgs(expectedSection.ID).WillReturnRows(rows)

		sectionReceived, err := r.Get(ctx, 1)

		assert.Equal(t, *expectedSection, sectionReceived)
		assert.Nil(t, err)
	})

	t.Run("get_non_existent_by_id", func(t *testing.T) {

		expectedSection := &domain.Section{
			ID:                 1,
			SectionNumber:      10,
			CurrentTemperature: 10,
			MinimumTemperature: 10,
			CurrentCapacity:    10,
			MinimumCapacity:    10,
			MaximumCapacity:    10,
			WarehouseID:        10,
			ProductTypeID:      10,
		}
		r := section.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity",
			"warehouse_id", "product_type_id"}).
			AddRow(expectedSection.ID, expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID)
		query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id=?"
		mock.ExpectQuery(query).WithArgs(expectedSection.ID).WillReturnRows(rows)

		sectionReceived, err := r.Get(ctx, 66)

		assert.Equal(t, domain.Section{}, sectionReceived)
		assert.NotNil(t, err)
	})
}