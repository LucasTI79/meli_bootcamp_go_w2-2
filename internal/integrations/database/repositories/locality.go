package repositories

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type localityRepository struct {
	db *sql.DB
}

func NewLocalityRepository(db *sql.DB) repositories.LocalityRepository {
	return &localityRepository{
		db: db,
	}
}

func (r *localityRepository) GetAll(ctx context.Context) ([]entities.Locality, error) {
	query := "SELECT * FROM localities"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var localities []entities.Locality

	for rows.Next() {
		locality := entities.Locality{}
		_ = rows.Scan(&locality.ID, &locality.CountryName, &locality.ProvinceName, &locality.LocalityName)
		localities = append(localities, locality)
	}

	return localities, nil
}

func (r *localityRepository) Get(ctx context.Context, id string) (entities.Locality, error) {
	query := "SELECT * FROM localities WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	locality := entities.Locality{}
	err := row.Scan(&locality.ID, &locality.CountryName, &locality.ProvinceName, &locality.LocalityName)
	if err != nil {
		return entities.Locality{}, err
	}

	return locality, nil
}

func (r *localityRepository) Exists(ctx context.Context, id string) bool {
	query := "SELECT id FROM localities WHERE id=?;"
	row := r.db.QueryRow(query, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *localityRepository) Save(ctx context.Context, locality entities.Locality) (int, error) {
	query := "INSERT INTO localities(country_name, province_name, locality_name) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&locality.CountryName, &locality.ProvinceName, &locality.LocalityName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *localityRepository) Update(ctx context.Context, locality entities.Locality) error {
	query := "UPDATE localities SET country_name=?, province_name=?, locality_name=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&locality.CountryName, &locality.ProvinceName, &locality.LocalityName, &locality.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *localityRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM localities WHERE id = ?"
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
		return services.ErrNotFound
	}

	return nil
}

func (r *localityRepository) GetNumberOfSellers(ctx context.Context, id string) (int, error) {
	query := "SELECT COUNT(*) FROM localities WHERE id = ?;"

	count := 0
	err := r.db.QueryRow(query, id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
