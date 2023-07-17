package section_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/stretchr/testify/assert"
)
var(

	expectedSection = domain.Section{
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

	t.Run("GET Not ID", func(t *testing.T) {
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
	func TestRepositoryGetAll(t *testing.T) {
		type fields struct {
			db *sql.DB
		}
	
		db, mock, _ := sqlmock.New()
		ctx := context.TODO()
	
		t.Run("GET ALL - OK", func(t *testing.T) {
			expectedSections := &[]domain.Section{
				{
					ID:                 1,
					SectionNumber:      10,
					CurrentTemperature: 10,
					MinimumTemperature: 10,
					CurrentCapacity:    10,
					MinimumCapacity:    10,
					MaximumCapacity:    10,
					WarehouseID:        10,
					ProductTypeID:      10,
				},
				{
					ID:                 2,
					SectionNumber:      20,
					CurrentTemperature: 20,
					MinimumTemperature: 20,
					CurrentCapacity:    20,
					MinimumCapacity:    20,
					MaximumCapacity:    20,
					WarehouseID:        20,
					ProductTypeID:      20,
				},
			}
	
			r := section.NewRepository(fields{db}.db)
			rows := sqlmock.NewRows([]string{"id", "section_number", "current_temperature", "minimum_temperature", "current_capacity", "minimum_capacity", "maximum_capacity",
			"warehouse_id", "product_type_id"})
	
			for _, expectedSection := range *expectedSections {
				rows.AddRow(expectedSection.ID, expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID)
			}
			query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id=?"
			mock.ExpectQuery(query).WillReturnRows(rows)
			sectionReceived, err := r.GetAll(ctx)
	
			assert.Equal(t, *expectedSections, sectionReceived)
			assert.Nil(t, err)
		})
	
		t.Run("GET ALL Error", func(t *testing.T) {
	
			r := section.NewRepository(fields{db}.db)
			
			query := "SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id FROM sections WHERE id=?"
			mock.ExpectQuery(query).
				WithArgs().
				WillReturnError(sql.ErrNoRows)
	
			sectionReceived, err := r.GetAll(ctx)
	
			assert.Equal(t, []domain.Section(nil), sectionReceived)
			assert.NotNil(t, err)
		})
	}

func TestRepositoryExists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("EXISTS - True", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"section_number"}).
			AddRow(expectedSection.SectionNumber)
		query := "SELECT section_number FROM sections WHERE section_number=?"
		mock.ExpectQuery(query).
			WithArgs(expectedSection.SectionNumber).
			WillReturnRows(rows)

		sectionExists := r.Exists(ctx, expectedSection.SectionNumber)

		assert.True(t, sectionExists)
	})
	t.Run("EXISTS - False", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		rows := sqlmock.NewRows([]string{"section_number"}).
			AddRow(expectedSection.SectionNumber)
		query := "SELECT section_number FROM sections WHERE section_number=?"
		mock.ExpectQuery(query).
			WithArgs(expectedSection.SectionNumber).
			WillReturnRows(rows)

		sectionExists := r.Exists(ctx, 100)

		assert.False(t, sectionExists)
	})
}
func TestRepositoryDelete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()
	query := "DELETE FROM sections WHERE id=?"
	t.Run("DELETE - OK", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		rowsAffected := int64(1)
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.ID).
			WillReturnResult(sqlmock.NewResult(1, rowsAffected))
		err := r.Delete(ctx, expectedSection.ID)
		assert.Nil(t, err)
	})

	t.Run("DELETE - Error - Exec", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.ID).
			WillReturnError(sql.ErrNoRows)

		err := r.Delete(ctx, expectedSection.ID)
		assert.NotNil(t, err)
	})

	t.Run("DELETE - Error - RowlsAffected0", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Delete(ctx, expectedSection.ID)
		assert.NotNil(t, err)
	})

	t.Run("DELETE - Error - Prepare", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		mock.ExpectPrepare(regexp.QuoteMeta(query)).WillReturnError(sql.ErrConnDone)
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).
			WithArgs(expectedSection.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

		err := r.Delete(ctx, expectedSection.ID)
		assert.NotNil(t, err)
	})

	t.Run("DELETE - Error - Not Found", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := r.Delete(ctx, expectedSection.ID)

		assert.NotNil(t, err)
	})
}
func TestRepositorySave(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	t.Run("SAVE - OK", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		id, err := r.Save(ctx, expectedSection)
		assert.Equal(t, expectedSection.ID, id)
		assert.Nil(t, err)
	})

	t.Run("SAVE - Error - Exec", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).
			WillReturnError(sql.ErrNoRows)

		_, err := r.Save(ctx, expectedSection)

		assert.NotNil(t, err)
	})

	t.Run("SAVE - Error - RowlsAffected0", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedSection)

		assert.NotNil(t, err)
	})
	t.Run("SAVE - Error - Prepare", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		_, err := r.Save(ctx, expectedSection)

		assert.NotNil(t, err)
	})
}
func TestRepositoryUpdate(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()
	t.Run("UPDATE - OK", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID, expectedSection.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		err := r.Update(ctx, expectedSection)
		assert.Nil(t, err)
	})

	t.Run("UPDATE - Error - Exec", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)

		query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID, expectedSection.ID).
			WillReturnError(sql.ErrNoRows)
		err := r.Update(ctx, expectedSection)
		assert.NotNil(t, err)
	})
	t.Run("UPDATE - Error - RowlsAffected0", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID, expectedSection.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedSection)
		assert.NotNil(t, err)
	})
	t.Run("UPDATE - Error - Prepare", func(t *testing.T) {
		r := section.NewRepository(fields{db}.db)
		query := "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?"
		mock.ExpectPrepare(regexp.QuoteMeta(query))
		mock.ExpectExec(regexp.QuoteMeta(query)).
			WithArgs(expectedSection.SectionNumber, expectedSection.CurrentTemperature, expectedSection.MinimumTemperature, expectedSection.CurrentCapacity, expectedSection.MinimumCapacity, expectedSection.MaximumCapacity, expectedSection.WarehouseID, expectedSection.ProductTypeID, expectedSection.ID).
			WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))
		err := r.Update(ctx, expectedSection)

		assert.NotNil(t, err)
	})
}