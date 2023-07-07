package carriers

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Carrier, error)
	Exists(ctx context.Context, cid string) bool
	Save(ctx context.Context, w domain.Carrier) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Carrier, error) {
	query := "SELECT * FROM carriers"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var carriers []domain.Carrier

	for rows.Next() {
		c := domain.Carrier{}
		_ = rows.Scan(&c.ID, &c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)
		carriers = append(carriers, c)
	}

	return carriers, nil
}

func (r *repository) Exists(ctx context.Context, cid string) bool {
	query := "SELECT CID FROM carriers WHERE CID=?;"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, c domain.Carrier) (int, error) {
	query := "INSERT INTO carriers (CID, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&c.CID, &c.CompanyName, &c.Address, &c.Telephone, &c.LocalityId)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
