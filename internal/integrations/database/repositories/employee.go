package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Employee queries
const (
	GetAllEmployees              = "SELECT employees.id, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id FROM employees"
	GetEmployeeByID              = "SELECT employees.id, employees.card_number_id, employees.first_name, employees.last_name, employees.warehouse_id FROM employees WHERE id=?"
	ExistsEmployeeByCardNumberID = "SELECT card_number_id FROM employees WHERE card_number_id=?"
	SaveEmployee                 = "INSERT INTO employees(card_number_id,first_name,last_name) VALUES (?,?,?)"
	UpdateEmployee               = "UPDATE employees SET first_name=?, last_name=?  WHERE id=?"
	DeleteEmployeeByID           = "DELETE FROM employees WHERE id = ?"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) repositories.EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

func (r *employeeRepository) GetAll(ctx context.Context) ([]entities.Employee, error) {
	rows, err := r.db.Query(GetAllEmployees)
	if err != nil {
		return nil, err
	}

	var employees []entities.Employee

	for rows.Next() {
		e := entities.Employee{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *employeeRepository) Get(ctx context.Context, id int) (entities.Employee, error) {
	row := r.db.QueryRow(GetEmployeeByID, id)
	e := entities.Employee{}
	err := row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return entities.Employee{}, err
	}

	return e, nil
}

func (r *employeeRepository) Exists(ctx context.Context, cardNumberID string) bool {
	row := r.db.QueryRow(ExistsEmployeeByCardNumberID, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *employeeRepository) Save(ctx context.Context, e entities.Employee) (int, error) {
	stmt, err := r.db.Prepare(SaveEmployee)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *employeeRepository) Update(ctx context.Context, e entities.Employee) error {
	stmt, err := r.db.Prepare(UpdateEmployee)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&e.FirstName, &e.LastName, &e.WarehouseID, &e.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *employeeRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteEmployeeByID)
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
