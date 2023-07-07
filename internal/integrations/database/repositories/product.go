package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Product queries
const (
	GetAllProducts             = "SELECT products.id, products.description, products.expiration_rate, products.freezing_rate, products.height, products.length, products.netweight, products.product_code, products.recommended_freezing_temperature, products.width, products.product_type_id, products.seller_id FROM products"
	GetProductByID             = "SELECT products.id, products.description, products.expiration_rate, products.freezing_rate, products.height, products.length, products.netweight, products.product_code, products.recommended_freezing_temperature, products.width, products.product_type_id, products.seller_id FROM products WHERE id=?"
	ExistsProductByProductCode = "SELECT product_code FROM products WHERE product_code=?"
	SaveProduct                = "INSERT INTO products(description,expiration_rate,freezing_rate,height,lenght,netweight,product_code,recommended_freezing_temperature,width,id_product_type,id_seller) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	UpdateProduct              = "UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, lenght=?, netweight=?, product_code=?, recommended_freezing_temperature=?, width=?, id_product_type=?, id_seller=?  WHERE id=?"
	DeleteProductByID          = "DELETE FROM products WHERE id = ?"
)

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) repositories.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetAll(ctx context.Context) ([]entities.Product, error) {
	rows, err := r.db.Query(GetAllProducts)
	if err != nil {
		return nil, err
	}

	var products []entities.Product

	for rows.Next() {
		p := entities.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	return products, nil
}

func (r *productRepository) Get(ctx context.Context, id int) (entities.Product, error) {
	row := r.db.QueryRow(GetProductByID, id)
	p := entities.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		return entities.Product{}, err
	}

	return p, nil
}

func (r *productRepository) Exists(ctx context.Context, productCode string) bool {
	row := r.db.QueryRow(ExistsProductByProductCode, productCode)
	err := row.Scan(&productCode)
	return err == nil
}

func (r *productRepository) Save(ctx context.Context, p entities.Product) (int, error) {
	stmt, err := r.db.Prepare(SaveProduct)
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

func (r *productRepository) Update(ctx context.Context, p entities.Product) error {
	stmt, err := r.db.Prepare(UpdateProduct)
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

func (r *productRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteProductByID)
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
		return services.ErrNotFound
	}

	return nil
}
