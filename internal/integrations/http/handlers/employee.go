package handlers

import (
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/web_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	employeeService services.EmployeeService
}

func NewEmployeeHandler(e services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: e,
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
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/employees/{id} [get]
func (handler *EmployeeHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid EmployeeHandler ID: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		employee, err := handler.employeeService.Get(&ctx, id)

		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				web_utils.Error(c, http.StatusNotFound, "EmployeeHandler not found: %s", err.Error())
				return
			}
			web_utils.Error(c, http.StatusInternalServerError, "Failed to get employee: %s", err.Error())
			return
		}

		web_utils.Success(c, http.StatusOK, *employee)
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
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/employees [get]
func (handler *EmployeeHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		employee, err := handler.employeeService.GetAll(&ctx)

		if err != nil {
			web_utils.Error(c, http.StatusInternalServerError, "Failed to get employee: %s", err.Error())
			return
		}

		if len(*employee) == 0 {
			web_utils.Error(c, http.StatusNoContent, "There are no employee stored: ")
			return
		}

		web_utils.Success(c, http.StatusOK, *employee)
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
//	@Param			Employees	body		entities.RequestCreateEmployee	true	"EmployeeHandler to Create"
//	@Success		200			{object}	web_utils.response
//	@Router			/api/v1/employees [post]
func (handler *EmployeeHandler) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		// createEmployee := entities.RequestCreateEmployee{}
		createEmployee := new(entities.RequestCreateEmployee)
		if err := c.ShouldBindJSON(&createEmployee); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}

		employeeDomain := &entities.Employee{
			CardNumberID: createEmployee.CardNumberID,
			FirstName:    createEmployee.FirstName,
			LastName:     createEmployee.LastName,
			WarehouseID:  createEmployee.WarehouseID,
		}

		if employeeDomain.CardNumberID == "" {
			web_utils.Error(c, http.StatusBadRequest, "Field Card Number ID is required: %s", "")
			return
		}

		if employeeDomain.FirstName == "" {
			web_utils.Error(c, http.StatusBadRequest, "Field First Name is required: %s", "")
		}

		if employeeDomain.LastName == "" {
			web_utils.Error(c, http.StatusBadRequest, "Field Last Name is required: %s", "")
		}

		if employeeDomain.WarehouseID == 0 {
			web_utils.Error(c, http.StatusBadRequest, "Field Ware House ID is required: %s", "")
		}

		ctx := c.Request.Context()
		employeeDomain, err := handler.employeeService.Save(&ctx, *employeeDomain)
		if err != nil {
			switch err {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web_utils.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
				return
			}

		}

		web_utils.Success(c, http.StatusCreated, *employeeDomain)
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
//	@Param			Employees	body		entities.RequestUpdateEmployee	true	"Updated Employeesers details"
//	@Success		200			{object}	web_utils.response
//	@Router			/api/v1/employees/{id} [patch]
func (handler *EmployeeHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		ReqUpdateEmployee := new(entities.RequestUpdateEmployee)

		if err := c.ShouldBindJSON(&ReqUpdateEmployee); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, "Error to read request: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		employeeUpdate, err := handler.employeeService.Update(&ctx, id, ReqUpdateEmployee)
		if err != nil {
			web_utils.Error(c, http.StatusNotFound, "Error to update: %s", err.Error())
			return
		}

		web_utils.Success(c, http.StatusOK, employeeUpdate)
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
//	@Success		204	{object}	web_utils.response
//	@Router			/api/v1/employees/{id} [delete]
func (handler *EmployeeHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = handler.employeeService.Delete(&ctx, int(id))
		if err != nil {
			web_utils.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}

		web_utils.Success(c, http.StatusNoContent, nil)
	}
}
