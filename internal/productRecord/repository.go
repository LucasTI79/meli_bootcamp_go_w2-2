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
	Exists(ctx context.Context, productId int) bool
	Save(ctx context.Context, p domain.ProductRecord) (int, error)
	Update(ctx context.Context, p domain.ProductRecord) error
	Delete(ctx context.Context, id int) error
	NumberRecords(ctx context.Context, id int) (int, error)
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
	query := "SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var product_records []domain.ProductRecord

	for rows.Next() {
		p := domain.ProductRecord{}
		_ = rows.Scan(&p.ID, &p.LastUpdateDate, &p.PurchasePrice, &p.SalePrice, &p.ProductId)
		product_records = append(product_records, p)
	}

	return product_records, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.ProductRecord, error) {
	query := "SELECT id, last_update_date, purchase_price, sale_price, product_id FROM product_records WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.ProductRecord{}
	err := row.Scan(&p.ID, &p.LastUpdateDate, &p.PurchasePrice, &p.SalePrice, &p.ProductId)
	if err != nil {
		return domain.ProductRecord{}, err
	}

	return p, nil
}

func (r *repository) Save(ctx context.Context, p domain.ProductRecord) (int, error) {
	query := "INSERT INTO product_records(last_update_date, purchase_price, sale_price, product_id) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.LastUpdateDate, p.PurchasePrice, p.SalePrice, p.ProductId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	query := "SELECT id FROM products_records WHERE id=?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) Update(ctx context.Context, p domain.ProductRecord) error {
	query := "UPDATE product_records SET last_update_date=?, purchase_price=?, sale_price=?, product_id=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.LastUpdateDate, p.PurchasePrice, p.SalePrice, p.ProductId, p.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM product_records WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) NumberRecords(ctx context.Context, product_id int) (int, error) {
	count := 0
	row := r.db.QueryRow("SELECT COUNT(*) from product_records where product_id =?", product_id)
	err := row.Scan(&count)

	return count, err
}
