package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Seller queries
const (
	GetAllSellers     = "SELECT sellers.id, sellers.cid, sellers.company_name, sellers.address, sellers.telephone, sellers.locality_id FROM sellers"
	GetSellerByID     = "SELECT sellers.id, sellers.cid, sellers.company_name, sellers.address, sellers.telephone, sellers.locality_id FROM sellers WHERE id=?"
	ExistsSellerByCID = "SELECT cid FROM sellers WHERE cid=?"
	SaveSeller        = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	UpdateSeller      = "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	DeleteSellerByID  = "DELETE FROM sellers WHERE id=?"
)

type sellerRepository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) repositories.SellerRepository {
	return &sellerRepository{
		db: db,
	}
}

func (r *sellerRepository) GetAll(ctx context.Context) ([]entities.Seller, error) {
	sellers := make([]entities.Seller, 0)

	rows, err := r.db.Query(GetAllSellers)
	if err != nil {
		return sellers, err
	}

	for rows.Next() {
		s := entities.Seller{}
		err := rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		if err != nil {
			return sellers, err
		}
		sellers = append(sellers, s)
	}

	tre := rows.Err()
	if tre != nil {
		return sellers, tre
	}

	return sellers, rows.Err()
}

func (r *sellerRepository) Get(ctx context.Context, id int) (*entities.Seller, error) {
	row := r.db.QueryRow(GetSellerByID, id)
	s := entities.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, services.ErrNotFound
		default:
			return &entities.Seller{}, err
		}
	}
	return &s, nil
}

func (r *sellerRepository) Exists(ctx context.Context, cid int) bool {
	row := r.db.QueryRow(ExistsSellerByCID, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *sellerRepository) Save(ctx context.Context, s entities.Seller) (int, error) {
	stmt, err := r.db.Prepare(SaveSeller)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *sellerRepository) Update(ctx context.Context, s entities.Seller) error {
	stmt, err := r.db.Prepare(UpdateSeller)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *sellerRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteSellerByID)
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
