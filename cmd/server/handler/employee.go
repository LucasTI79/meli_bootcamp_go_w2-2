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

// Method Get
// GetEmployees godoc
//
//	@Summary		Get Employees
//	@Tags			Employees
//	@Description	Get the details of a Employees
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of Employees to be searched"
//	@Success		200	{object}	web.response
//	@Router			/api/v1/employees/{id} [get]
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

// Method GetAll
// ListEmployees godoc
//
//	@Summary		List Employees
//	@Tags			Employees
//	@Description	getAll employees
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Router			/api/v1/employees [get]
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

// Method Save
// CreateEmployees godoc
//
//	@Summary		Create Employees
//	@Tags			Employees
//	@Description	Create employees
//	@Accept			json
//	@Produce		json
//	@Param			Employees	body		domain.RequestCreateEmployee	true	"Employee to Create"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/employees [post]
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

		if employee.CardNumberID == "" {
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

// Method Update
// UpdateEmployees godoc
//
//	@Summary		Update Employees
//	@Tags			Employees
//	@Description	Update the details of a Employees
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string							true	"ID of Employees to be updated"
//	@Param			Employees	body		domain.RequestUpdateEmployee	true	"Updated Employeesers details"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/employees/{id} [patch]
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

// Method Delete
// DeleteEmployees godoc
//
//	@Summary		Delete Employees
//	@Tags			Employees
//	@Description	Delete Employees
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a Employees to be excluded"
//	@Success		204	{object}	web.response
//	@Router			/api/v1/employees/{id} [delete]
func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		err = e.service.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
