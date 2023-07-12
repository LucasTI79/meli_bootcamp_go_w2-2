package productbatches

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

type IRepository interface {
	SectionProductsReports() ([]domain.ProductBySection, error)
	SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error)
	Save(ctx context.Context, product domain.ProductBatches) (int, error)
	Get(ctx context.Context, id int) (domain.ProductBatches, error)
	ExistsProductBatch(ctx context.Context, batchNumber int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) ExistsProductBatch(ctx context.Context, batchNumber int) bool {
	query := "SELECT batch_number FROM product_batches WHERE batch_number=?;"
	row := r.db.QueryRow(query, batchNumber)
	err := row.Scan(&batchNumber)
	return err == nil
}
func (r *repository) Get(ctx context.Context, id int) (domain.ProductBatches, error) {
	query := "SELECT * FROM product_batches WHERE id=?;"
	row := r.db.QueryRow(query, id)
	pb := domain.ProductBatches{}
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.ProductID, &pb.SectionID)
	if err != nil {
		return domain.ProductBatches{}, err
	}

	return pb, nil
}

func (r *repository) Save(ctx context.Context, product domain.ProductBatches) (int, error) {
	query := "INSERT INTO product_batches ( batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?,?,?,?,?,?,?,?,?,?)"
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

func (r *repository) SectionProductsReports() ([]domain.ProductBySection, error) {
	rows, err := r.db.Query("SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id GROUP BY pb.section_id")
	if err != nil {
		return nil, err
	}
	var productsBySection []domain.ProductBySection
	for rows.Next() {
		var product domain.ProductBySection
		err := rows.Scan(&product.ProductsCount, &product.SectionID, &product.SectionNumber)
		if err != nil {
			return nil, err
		}
		productsBySection = append(productsBySection, product)
	}
	return productsBySection, nil
}

func (r *repository) SectionProductsReportsBySection(sectionID int) ([]domain.ProductBySection, error) {
	rows, err := r.db.Query("SELECT count(pb.id) as `products_count`, pb.section_id, s.section_number FROM product_batches pb JOIN sections s ON pb.section_id = s.id WHERE pb.section_id = ? GROUP BY pb.section_id", sectionID)
	if err != nil {
		panic(err)
	}
	found := false
	var productBySection []domain.ProductBySection
	for rows.Next() {
		var pb domain.ProductBySection
		err := rows.Scan(&pb.SectionID, &pb.SectionNumber, &pb.ProductsCount)
		if err != nil {
			return productBySection, err
		}
		productBySection = append(productBySection, pb)
		found = true
	}
	if !found {
		return productBySection, sql.ErrNoRows
	}
	return productBySection, nil
}
