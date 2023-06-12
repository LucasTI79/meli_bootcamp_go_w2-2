package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
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
	return func(c *gin.Context) {
		employee, err := e.service.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get employee: %s", err.Error())
			return
		}

		if len(*employee) == 0 {
			web.Error(c, http.StatusNotFound, "There are no employee stored: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, *employee)
	}
}

func (e *Employee) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		createEmployee := domain.RequestCreateEmployee{}
		if err := c.Bind(&createEmployee); err != nil {
			web.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		employee := &domain.Employee{
			CardNumberID: createEmployee.CardNumberID,
			FirstName:    createEmployee.FirstName,
			LastName:     createEmployee.LastName,
			WarehouseID:  createEmployee.WarehouseID,
		}

		employee, err := e.service.Save(c, *employee)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
			return
		}

		web.Success(c, http.StatusCreated, *employee)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
