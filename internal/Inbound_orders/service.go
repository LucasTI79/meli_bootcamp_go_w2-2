package inbound_order

import (
	"context"
	"errors"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
)

// Errors
var (
	ErrNotFound            = errors.New("inbound orders not found")
	ErrConflict            = errors.New("409 Conflict: inbound orders already exists")
	ErrUnprocessableEntity = errors.New("all fields are required")
)

type Service interface {
	Get(ctx *context.Context, id int) (*domain.InboundOrders, error)
	GetAll(ctx *context.Context) (*[]domain.InboundOrders, error)
	Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error)
	Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error)
	Delete(ctx *context.Context, id int) error
	CountInboundOrders(ctx *context.Context) ([]domain.EmployeeInboundOrdersCount, error)
	CountInboundOrdersByID(ctx *context.Context, employeeID int) (domain.EmployeeInboundOrdersCount, error)
}

type service struct {
	inboundOrdersRepository Repository
	employeeRepository      employee.Repository
}

func NewService(r Repository, employeeRepository employee.Repository) Service {
	return &service{
		inboundOrdersRepository: r,
		employeeRepository:      employeeRepository,
	}
}

func (s *service) Get(ctx *context.Context, id int) (*domain.InboundOrders, error) {
	inboundOrders, err := s.inboundOrdersRepository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}
	return &inboundOrders, nil
}

func (s *service) GetAll(ctx *context.Context) (*[]domain.InboundOrders, error) {
	inboundOrders := []domain.InboundOrders{}

	inboundOrders, err := s.inboundOrdersRepository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}
	return &inboundOrders, nil
}

func (s *service) Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error) {

	employeeIDint, _ := strconv.Atoi(inboundOrders.EmployeeID)
	_, existEmployee := s.employeeRepository.Get(*ctx, employeeIDint)

	if existEmployee != nil {
		return nil, ErrConflict
	}

	id, err := s.inboundOrdersRepository.Save(*ctx, inboundOrders)
	if err != nil {
		return nil, err
	}

	inboundOrders.ID = id

	return &inboundOrders, nil
}

func (s *service) Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error) {
	existingInboundOrders, err := s.inboundOrdersRepository.Get(*ctx, id)
	if err != nil {
		return nil, ErrNotFound
	}

	if reqUpdateInboundOrders.OrderDate != nil {
		existingInboundOrders.OrderDate = *reqUpdateInboundOrders.OrderDate
	}
	if reqUpdateInboundOrders.OrderNumber != nil {
		existingInboundOrders.OrderNumber = *reqUpdateInboundOrders.OrderNumber
	}
	if reqUpdateInboundOrders.EmployeeID != nil {
		existingInboundOrders.EmployeeID = *reqUpdateInboundOrders.EmployeeID
	}

	if reqUpdateInboundOrders.ProductBatchID != nil {
		existingInboundOrders.ProductBatchID = *reqUpdateInboundOrders.ProductBatchID
	}
	if reqUpdateInboundOrders.WarehouseID != nil {
		existingInboundOrders.WarehouseID = *reqUpdateInboundOrders.WarehouseID
	}

	err = s.inboundOrdersRepository.Update(*ctx, existingInboundOrders)
	if err != nil {
		return nil, err
	}

	return &existingInboundOrders, nil
}

func (s *service) Delete(ctx *context.Context, id int) error {
	_, err := s.inboundOrdersRepository.Get(*ctx, id)
	if err != nil {
		return ErrNotFound
	}

	err = s.inboundOrdersRepository.Delete(*ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) CountInboundOrders(ctx *context.Context) ([]domain.EmployeeInboundOrdersCount, error) {
	allEmployees, err := s.employeeRepository.GetAll(*ctx)
	if err != nil {
		return nil, err
	}

	allInboundOrders, err := s.inboundOrdersRepository.GetAll(*ctx)

	if err != nil {
		return nil, err
	}

	inboundOrdersCount := []domain.EmployeeInboundOrdersCount{}

	count := 0
	for _, employee := range allEmployees {
		for _, inboundOrder := range allInboundOrders {
			inboundOrderEmployeeID, _ := strconv.Atoi(inboundOrder.EmployeeID)
			if inboundOrderEmployeeID == employee.ID {
				count++
			}
		}
		employeeCount := domain.EmployeeInboundOrdersCount{
			ID:                 employee.ID,
			CardNumberID:       employee.CardNumberID,
			FirstName:          employee.FirstName,
			LastName:           employee.LastName,
			WarehouseID:        employee.WarehouseID,
			InboundOrdersCount: count,
		}
		inboundOrdersCount = append(inboundOrdersCount, employeeCount)
	}

	return inboundOrdersCount, nil
}

func (s *service) CountInboundOrdersByID(ctx *context.Context, employeeID int) (domain.EmployeeInboundOrdersCount, error) {
	employee, err := s.employeeRepository.Get(*ctx, employeeID)
	if err != nil {
		return domain.EmployeeInboundOrdersCount{}, err
	}

	allInboundOrders, err := s.inboundOrdersRepository.GetAll(*ctx)

	if err != nil {
		return domain.EmployeeInboundOrdersCount{}, err
	}

	count := 0

	for _, inboundOrder := range allInboundOrders {
		inboundOrderEmployeeID, _ := strconv.Atoi(inboundOrder.EmployeeID)
		if inboundOrderEmployeeID == employee.ID {
			count++
		}
	}
	employeeCount := domain.EmployeeInboundOrdersCount{
		ID:                 employee.ID,
		CardNumberID:       employee.CardNumberID,
		FirstName:          employee.FirstName,
		LastName:           employee.LastName,
		WarehouseID:        employee.WarehouseID,
		InboundOrdersCount: count,
	}

	return employeeCount, nil
}
