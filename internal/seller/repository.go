package seller

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Get(ctx context.Context, id int) (*domain.Seller, error)
	Exists(ctx context.Context, cid int) bool
	Save(ctx context.Context, s domain.Seller) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
}

const (
	GetAllSellers     = "SELECT sellers.id, sellers.cid, sellers.company_name, sellers.address, sellers.telephone, sellers.locality_id FROM sellers"
	GetSellerByID     = "SELECT sellers.id, sellers.cid, sellers.company_name, sellers.address, sellers.telephone, sellers.locality_id FROM sellers WHERE id=?"
	ExistsSellerByCID = "SELECT cid FROM sellers WHERE cid=?"
	SaveSeller        = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	UpdateSeller      = "UPDATE sellers SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"
	DeleteSellerByID  = "DELETE FROM sellers WHERE id=?"
)

type repository struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	sellers := make([]domain.Seller, 0)

	rows, err := r.db.Query(GetAllSellers)
	if err != nil {
		return sellers, err
	}

	for rows.Next() {
		s := domain.Seller{}
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

func (r *repository) Get(ctx context.Context, id int) (*domain.Seller, error) {
	row := r.db.QueryRow(GetSellerByID, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return &domain.Seller{}, err
		}
	}
	return &s, nil
}

func (r *repository) Exists(ctx context.Context, cid int) bool {
	row := r.db.QueryRow(ExistsSellerByCID, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, s domain.Seller) (int, error) {
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

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
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

func (r *repository) Delete(ctx context.Context, id int) error {
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
		return ErrNotFound
	}

	return nil
}
