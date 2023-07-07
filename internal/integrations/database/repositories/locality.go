package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Locality queries
const (
	GetAllLocalities         = "SELECT localities.id, localities.country_name, localities.province_name, localities.locality_name FROM localities"
	GetLocalityByID          = "SELECT localities.id, localities.country_name, localities.province_name, localities.locality_name FROM localities WHERE id = ?"
	ExistsLocalityByID       = "SELECT id FROM localities WHERE id=?"
	SaveLocality             = "INSERT INTO localities(country_name, province_name, locality_name) VALUES (?,?,?)"
	UpdateLocality           = "UPDATE localities SET country_name=?, province_name=?, locality_name=?  WHERE id=?"
	DeleteLocalityByID       = "DELETE FROM localities WHERE id = ?"
	CountLocalitySellersByID = "SELECT COUNT(*) from localities where id = ?"
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
	localities := make([]entities.Locality, 0)

	rows, err := r.db.Query(GetAllLocalities)
	if err != nil {
		return localities, err
	}

	for rows.Next() {
		locality := entities.Locality{}
		err := rows.Scan(&locality.ID, &locality.CountryName, &locality.ProvinceName, &locality.LocalityName)
		if err != nil {
			return localities, err
		}

		localities = append(localities, locality)
	}

	return localities, rows.Err()
}

func (r *localityRepository) Get(ctx context.Context, id int) (entities.Locality, error) {
	row := r.db.QueryRow(GetLocalityByID, id)
	locality := entities.Locality{}
	err := row.Scan(&locality.ID, &locality.CountryName, &locality.ProvinceName, &locality.LocalityName)
	if err != nil {
		return entities.Locality{}, err
	}

	return locality, nil
}

func (r *localityRepository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(ExistsLocalityByID, id)
	var foundId int
	err := row.Scan(&foundId)
	return err == nil
}

func (r *localityRepository) Save(ctx context.Context, locality entities.Locality) (int, error) {
	stmt, err := r.db.Prepare(SaveLocality)
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
	stmt, err := r.db.Prepare(UpdateLocality)
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

func (r *localityRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteLocalityByID)
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

func (r *localityRepository) CountSellers(ctx context.Context, id int) (int, error) {
	count := 0
	row := r.db.QueryRow(CountLocalitySellersByID, id)
	err := row.Scan(&count)

	return count, err
}
