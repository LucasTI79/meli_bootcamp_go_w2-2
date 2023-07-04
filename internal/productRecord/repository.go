package productRecord

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.ProductRecord, error)
	Get(ctx context.Context, id int) (domain.ProductRecord, error)
	//Exists(ctx context.Context, productCode string) bool
	//Save(ctx context.Context, p domain.Product) (int, error)
	//Update(ctx context.Context, p domain.Product) error
	//Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.ProductRecord, error) {
	query := "SELECT * FROM productsRecords;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var productsRecords []domain.ProductRecord

	for rows.Next() {
		p := domain.ProductRecord{}
		_ = rows.Scan(&p.ID, &p.LastUpdateRate, &p.PurchasePrice, &p.SalePrice, &p.ProductId)
		productsRecords = append(productsRecords, p)
	}

	return productsRecords, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.ProductRecord, error) {
	query := "SELECT * FROM productsRecords WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.ProductRecord{}
	err := row.Scan(&p.ID, &p.LastUpdateRate, &p.PurchasePrice, &p.SalePrice, &p.ProductId)
	if err != nil {
		return domain.ProductRecord{}, err
	}

	return p, nil
}
