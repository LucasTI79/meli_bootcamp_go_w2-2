package handler

import (
	"net/http"
	"strconv"

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
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid Employee ID: %s", err.Error())
			return
		}

		employee, err := e.service.Get(c, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get employee: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, *employee)
	}
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
		// createEmployee := domain.RequestCreateEmployee{}
		createEmployee := new(domain.RequestCreateEmployee)
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

		if employee.CardNumberID <= "" {
			web.Error(c, http.StatusBadRequest, "Field Card Number ID is required: %s", "")
			return
		}

		if employee.FirstName == "" {
			web.Error(c, http.StatusBadRequest, "Field First Name is required: %s", "")
		}

		if employee.LastName == "" {
			web.Error(c, http.StatusBadRequest, "Field Last Name is required: %s", "")
		}

		if employee.WarehouseID == 0 {
			web.Error(c, http.StatusBadRequest, "Field Ware House ID is required: %s", "")
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
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		ReqUpdateEmployee := new(domain.RequestUpdateEmployee)

		if err := c.Bind(&ReqUpdateEmployee); err != nil {
			web.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		employeeUpdate, err := e.service.Update(c, id, ReqUpdateEmployee)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to update: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, employeeUpdate)
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
