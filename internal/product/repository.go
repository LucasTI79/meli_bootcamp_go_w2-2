package product

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
	ExistsByID(ctx context.Context, id int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	query := "SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	query := "SELECT id, description,expiration_rate,freezing_rate,height,length,net_weight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Exists(ctx context.Context, productCode string) bool {
	query := "SELECT product_code FROM products WHERE product_code=?;"
	row := r.db.QueryRow(query, productCode)
	err := row.Scan(&productCode)
	return err == nil
}

func (r *repository) ExistsByID(ctx context.Context, id int) bool {
	query := "SELECT id FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	query := "INSERT INTO products(description, expiration_rate, freezing_rate, height, length, net_weight, product_code, recommended_freezing_temperature,width, product_type_id, seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	query := "UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, net_weight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID, p.ID)
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
	query := "DELETE FROM products WHERE id=?"
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
