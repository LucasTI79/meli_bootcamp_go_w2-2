package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Buyer queries
const (
	GetAllBuyers    = "SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name FROM buyers"
	GetBuyerByID    = "SELECT buyers.id, buyers.card_number_id, buyers.first_name, buyers.last_name FROM buyers WHERE id = ?"
	ExistsBuyerByID = "SELECT card_number_id FROM buyers WHERE card_number_id=?;"
	SaveBuyer       = "INSERT INTO buyers(card_number_id,first_name,last_name) VALUES (?,?,?)"
	UpdateBuyer     = "UPDATE buyers SET first_name=?, last_name=?  WHERE id=?"
	DeleteBuyerByID = "DELETE FROM buyers WHERE id = ?"
)

type buyerRepository struct {
	db *sql.DB
}

func NewBuyerRepository(db *sql.DB) repositories.BuyerRepository {
	return &buyerRepository{
		db: db,
	}
}

func (r *buyerRepository) GetAll(ctx context.Context) ([]entities.Buyer, error) {
	rows, err := r.db.Query(GetAllBuyers)
	if err != nil {
		return nil, err
	}

	var buyers []entities.Buyer

	for rows.Next() {
		b := entities.Buyer{}
		_ = rows.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
		buyers = append(buyers, b)
	}

	return buyers, nil
}

func (r *buyerRepository) Get(ctx context.Context, id int) (entities.Buyer, error) {
	row := r.db.QueryRow(GetBuyerByID, id)
	b := entities.Buyer{}
	err := row.Scan(&b.ID, &b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return entities.Buyer{}, err
	}

	return b, nil
}

func (r *buyerRepository) Exists(ctx context.Context, cardNumberID string) bool {
	row := r.db.QueryRow(ExistsBuyerByID, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *buyerRepository) Save(ctx context.Context, b entities.Buyer) (int, error) {
	stmt, err := r.db.Prepare(SaveBuyer)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b.CardNumberID, &b.FirstName, &b.LastName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *buyerRepository) Update(ctx context.Context, b entities.Buyer) error {
	stmt, err := r.db.Prepare(UpdateBuyer)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&b.FirstName, &b.LastName, &b.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *buyerRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteBuyerByID)
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
