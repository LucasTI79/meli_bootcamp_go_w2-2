package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Section queries
const (
	GetAllSections               = "SELECT sections.id, sections.section_number, sections.current_temperature, sections.minimum_temperature, sections.current_capacity, sections.minimum_capacity, sections.maximum_capacity, sections.warehouse_id, sections.product_type_id FROM sections"
	GetSectionByID               = "SELECT sections.id, sections.section_number, sections.current_temperature, sections.minimum_temperature, sections.current_capacity, sections.minimum_capacity, sections.maximum_capacity, sections.warehouse_id, sections.product_type_id FROM sections WHERE id=?"
	ExistsSectionBySectionNumber = "SELECT section_number FROM sections WHERE section_number=?"
	SaveSection                  = "INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, id_product_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	UpdateSection                = "UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, id_product_type=? WHERE id=?"
	DeleteSectionByID            = "DELETE FROM sections WHERE id=?"
)

type sectionRepository struct {
	db *sql.DB
}

func NewSectionRepository(db *sql.DB) repositories.SectionRepository {
	return &sectionRepository{
		db: db,
	}
}

func (r *sectionRepository) GetAll(ctx context.Context) ([]entities.Section, error) {
	rows, err := r.db.Query(GetAllSections)
	if err != nil {
		return nil, err
	}

	var sections []entities.Section

	for rows.Next() {
		s := entities.Section{}
		_ = rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
		sections = append(sections, s)
	}

	return sections, nil
}

func (r *sectionRepository) Get(ctx context.Context, id int) (entities.Section, error) {
	row := r.db.QueryRow(GetSectionByID, id)
	s := entities.Section{}
	err := row.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return entities.Section{}, err
	}

	return s, nil
}

func (r *sectionRepository) Exists(ctx context.Context, sectionNumber int) bool {
	row := r.db.QueryRow(ExistsSectionBySectionNumber, sectionNumber)
	err := row.Scan(&sectionNumber)
	return err == nil
}

func (r *sectionRepository) Save(ctx context.Context, s entities.Section) (int, error) {
	stmt, err := r.db.Prepare(SaveSection)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *sectionRepository) Update(ctx context.Context, s entities.Section) error {
	stmt, err := r.db.Prepare(UpdateSection)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.MaximumCapacity, &s.WarehouseID, &s.ProductTypeID, &s.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *sectionRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteSectionByID)
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
