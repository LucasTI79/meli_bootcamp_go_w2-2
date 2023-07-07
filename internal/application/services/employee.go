package services

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type EmployeeService interface {
	Get(ctx *context.Context, id int) (*entities.Employee, error)
	GetAll(ctx *context.Context) (*[]entities.Employee, error)
	Save(ctx *context.Context, employee entities.Employee) (*entities.Employee, error)
	Update(ctx *context.Context, id int, reqUpdateEmployee *entities.RequestUpdateEmployee) (*entities.Employee, error)
	Delete(ctx *context.Context, id int) error
}

type employeeService struct {
	repository repositories.EmployeeRepository
}

func NewEmployeeService(r repositories.EmployeeRepository) EmployeeService {
	return &employeeService{
		repository: r,
	}
}

func (s *employeeService) Get(ctx *context.Context, id int) (*entities.Employee, error) {
	employees, err := s.repository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &employees, nil
}

func (s *employeeService) GetAll(ctx *context.Context) (*[]entities.Employee, error) {
	employees := []entities.Employee{}

	employees, err := s.repository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}
	return &employees, nil
}

func (s *employeeService) Save(ctx *context.Context, employee entities.Employee) (*entities.Employee, error) {

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

func (s *employeeService) Update(ctx *context.Context, id int, reqUpdateEmployee *entities.RequestUpdateEmployee) (*entities.Employee, error) {
	existingEmployee, err := s.repository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
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

	err = s.repository.Update(*ctx, existingEmployee)
	if err != nil {
		return nil, err
	}

	return &existingEmployee, nil
}

func (s *employeeService) Delete(ctx *context.Context, id int) error {
	_, err := s.repository.Get(*ctx, id)
	if err != nil {
		return ErrNotFound
	}

	err = s.repository.Delete(*ctx, id)
	if err != nil {
		return err
	}
	return nil
}
