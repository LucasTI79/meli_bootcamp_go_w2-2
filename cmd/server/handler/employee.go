package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	service employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		service: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e.service.GetAll(ctx)

		// p, err := linha de cima
	}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		createEmployee := new(domain.RequestCreateEmployee)

		employee := &domain.Employee{
			CardNumberID: createEmployee.CardNumberID,
			FirstName:    createEmployee.FirstName,
			LastName:     createEmployee.LastName,
			WarehouseID:  createEmployee.WarehouseID,
		}

		e.service.Save(c, *employee)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
