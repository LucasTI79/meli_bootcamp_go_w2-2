package productbatches

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

type IRepository interface {
	Get(ctx context.Context, id int) (domain.ProductBatches, error)
	Save(ctx context.Context, product domain.ProductBatches) (int, error)
	Exists(ctx context.Context, batchNumber int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, batchNumber int) bool {
	query := "SELECT batch_number FROM product_batches WHERE batch_number=?;"
	row := r.db.QueryRow(query, batchNumber)
	err := row.Scan(&batchNumber)
	return err == nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.ProductBatches, error) {
	query := "SELECT * FROM product_batches WHERE id=?;"
	row := r.db.QueryRow(query, id)
	productBatches := domain.ProductBatches{}
	err := row.Scan(
		&productBatches.ID,
		&productBatches.BatchNumber,
		&productBatches.CurrentQuantity,
		&productBatches.CurrentTemperature,
		&productBatches.DueDate,
		&productBatches.InitialQuantity,
		&productBatches.ManufacturingDate,
		&productBatches.ManufacturingHour,
		&productBatches.MinimumTemperature,
		&productBatches.ProductID,
		&productBatches.SectionID,
	)
	if err != nil {
		return domain.ProductBatches{}, err
	}

	return productBatches, nil
}

func (r *repository) Save(ctx context.Context, product domain.ProductBatches) (int, error) {
	// Query SQL para inserir um ProductBatches na tabela
	query := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(
		&product.BatchNumber,
		&product.CurrentQuantity,
		&product.CurrentTemperature,
		&product.DueDate,
		&product.InitialQuantity,
		&product.ManufacturingDate,
		&product.ManufacturingHour,
		&product.MinimumTemperature,
		&product.ProductID,
		&product.SectionID,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
