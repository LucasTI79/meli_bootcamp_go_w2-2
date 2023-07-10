package carriers

import (
	"context"
	"database/sql"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Carrier, error)
	Exists(ctx context.Context, cid string) bool
	Save(ctx context.Context, w domain.Carrier) (int, error)
	GetLocalityById(ctx context.Context, localityId int) (domain.Locality, error)
	GetCountCarriersByLocalityId(ctx context.Context, localityId int) (int, error)
	GetCountAndDataByLocality(ctx context.Context) ([]dtos.DataLocalityAndCarrier, error)
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
	query := "SELECT id, cid, company_name, address, telephone, locality_id FROM carriers"
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
	query := "SELECT cid FROM carriers WHERE cid=?"
	row := r.db.QueryRow(query, cid)
	err := row.Scan(&cid)
	return err == nil
}

func (r *repository) Save(ctx context.Context, c domain.Carrier) (int, error) {
	query := "INSERT INTO carriers(cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"
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

func (r *repository) GetLocalityById(ctx context.Context, localityId int) (domain.Locality, error) {
	query := "SELECT localities.id, localities.province_name, localities.locality_name FROM localities WHERE id = ?"
	row := r.db.QueryRow(query, localityId)
	l := domain.Locality{}
	err := row.Scan(&l.ID, &l.ProvinceName, &l.LocalityName)
	if err != nil {
		return domain.Locality{}, err
	}
	return l, nil
}

func (r *repository) GetCountCarriersByLocalityId(ctx context.Context, localityId int) (int, error) {
	query := "SELECT COUNT(id) FROM carriers WHERE locality_id = ?"
	rows, err := r.db.Query(query, localityId)
	if err != nil {
		return 0, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) GetCountAndDataByLocality(ctx context.Context) ([]dtos.DataLocalityAndCarrier, error) {
	query := "SELECT l.id, l.locality_name, (SELECT count(id) FROM carriers c where c.locality_id = l.id) AS count_carrier FROM localities l LIMIT 10"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var data []dtos.DataLocalityAndCarrier
	for rows.Next() {
		d := dtos.DataLocalityAndCarrier{}
		_ = rows.Scan(&d.Id, &d.LocalityName, &d.CountCarrier)
		data = append(data, d)
	}
	return data, nil
}
