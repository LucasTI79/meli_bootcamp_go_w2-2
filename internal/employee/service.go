package employee

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
	ErrConflict = errors.New("409 Conflict: Employee with CardNumberID already exists")
)

type Service interface {
	Get(ctx *context.Context, id int) (*domain.Employee, error)
	GetAll(ctx *context.Context) (*[]domain.Employee, error)
	Save(ctx *context.Context, employee domain.Employee) (*domain.Employee, error)
	Update(ctx *context.Context, id int, reqUpdateEmployee *domain.RequestUpdateEmployee) (*domain.Employee, error)
	Delete(ctx *context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx *context.Context, id int) (*domain.Employee, error) {
	employees, err := s.repository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &employees, nil
}

func (s *service) GetAll(ctx *context.Context) (*[]domain.Employee, error) {
	employees := []domain.Employee{}

	employees, err := s.repository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}
	return &employees, nil
}

func (s *service) Save(ctx *context.Context, employee domain.Employee) (*domain.Employee, error) {

	existingEmployee := s.repository.Exists(*ctx, employee.CardNumberID)

	if existingEmployee {
		return nil, ErrConflict
	}

	id, err := s.repository.Save(*ctx, employee)
	if err != nil {
		return nil, err
	}

	employee.ID = id

	return &employee, nil
}

func (s *service) Update(ctx *context.Context, id int, reqUpdateEmployee *domain.RequestUpdateEmployee) (*domain.Employee, error) {
	existingEmployee, err := s.repository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	if reqUpdateEmployee.CardNumberID != nil {
		existingEmployeeSearch := s.repository.Exists(*ctx, *reqUpdateEmployee.CardNumberID)
		if existingEmployeeSearch && *reqUpdateEmployee.CardNumberID != existingEmployee.CardNumberID {
			return nil, ErrConflict
		}
		existingEmployee.CardNumberID = *reqUpdateEmployee.CardNumberID
	}

	if reqUpdateEmployee.FirstName != nil {
		existingEmployee.FirstName = *reqUpdateEmployee.FirstName
	}
	if reqUpdateEmployee.LastName != nil {
		existingEmployee.LastName = *reqUpdateEmployee.LastName
	}
	if reqUpdateEmployee.WarehouseID != nil {
		existingEmployee.WarehouseID = *reqUpdateEmployee.WarehouseID
	}

	err1 := s.repository.Update(*ctx, existingEmployee)
	if err1 != nil {
		return nil, err1
	}

	return &existingEmployee, nil
}

func (s *service) Delete(ctx *context.Context, id int) error {
	_, err := s.repository.Get(*ctx, id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return err
		}
	}

	err1 := s.repository.Delete(*ctx, id)
	if err1 != nil {
		return err
	}

	return nil
}
