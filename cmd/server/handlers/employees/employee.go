package employees

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound = errors.New("employee not found")
	ErrConflict = errors.New("409 Conflict: Employee with CardNumberID already exists")
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

		ctx := c.Request.Context()
		employee, err := e.service.Get(&ctx, id)

		if err != nil {
			if errors.Is(err, ErrNotFound) {
				web.Error(c, http.StatusNotFound, "Employee not found: %s", err.Error())
				return
			}
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
		ctx := c.Request.Context()
		employee, err := e.service.GetAll(&ctx)

		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get employee: %s", err.Error())
			return
		}

		if len(*employee) == 0 {
			web.Error(c, http.StatusNoContent, "There are no employee stored: ")
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
		if err := c.ShouldBindJSON(&createEmployee); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}

		employeeDomain := &domain.Employee{
			CardNumberID: createEmployee.CardNumberID,
			FirstName:    createEmployee.FirstName,
			LastName:     createEmployee.LastName,
			WarehouseID:  createEmployee.WarehouseID,
		}

		if employeeDomain.CardNumberID == "" {
			web.Error(c, http.StatusBadRequest, "Field Card Number ID is required: %s", "")
			return
		}

		if employeeDomain.FirstName == "" {
			web.Error(c, http.StatusBadRequest, "Field First Name is required: %s", "")
		}

		if employeeDomain.LastName == "" {
			web.Error(c, http.StatusBadRequest, "Field Last Name is required: %s", "")
		}

		if employeeDomain.WarehouseID == 0 {
			web.Error(c, http.StatusBadRequest, "Field Ware House ID is required: %s", "")
		}

		ctx := c.Request.Context()
		employeeDomain, err := e.service.Save(&ctx, *employeeDomain)
		if err != nil {
			switch err {
			case employee.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
				return
			}

		}

		web.Success(c, http.StatusCreated, *employeeDomain)
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

		if err := c.ShouldBindJSON(&ReqUpdateEmployee); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Error to read request: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		employeeUpdate, err := e.service.Update(&ctx, id, ReqUpdateEmployee)
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
		ctx := c.Request.Context()
		err = e.service.Delete(&ctx, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
