package handlers

import (
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/web_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WarehouseHandler struct {
	warehouseService services.WarehouseService
}

func NewWarehouse(s services.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		warehouseService: s,
	}
}

// GetOneWarehouse godoc
//
//	@Summary		Get warehouses
//	@Tags			Warehouses
//	@Description	get one warehouses by id
//	@Produce		json
//	@Param			id	path		int	true	"WarehouseHandler ID"
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/warehouses/{id} [get]
func (handler *WarehouseHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		if e != nil {
			web_utils.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		result, err := handler.warehouseService.GetOne(&ctx, warehouseId)

		if err != nil {
			web_utils.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web_utils.Success(c, http.StatusOK, result)
	}
}

// GetAllWarehouses godoc
//
//	@Summary		List warehouses
//	@Tags			Warehouses
//	@Description	get warehouses
//	@Produce		json
//	@Success		200	{object}	[]web_utils.response
//	@Router			/api/v1/warehouses [get]
func (handler *WarehouseHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		warehouses, err := handler.warehouseService.GetAll(&ctx)
		if err != nil {
			web_utils.Error(c, http.StatusInternalServerError, "Failed to get warehouses: %s", err.Error())
			return
		}

		if len(*warehouses) == 0 {
			web_utils.Error(c, http.StatusNoContent, "There are no warehouses stored")
			return
		}

		web_utils.Success(c, http.StatusOK, *warehouses)
	}
}

// createWarehouses godoc
//
//	@Summary		Create warehouses
//	@Tags			Warehouses
//	@Description	Create warehouses
//	@Accept			json
//	@Produce		json
//	@Param			WarehouseHandler	body		dtos.WarehouseRequestDTO	true	"warehouses to create"
//	@Success		200			{object}	web_utils.response
//	@Router			/api/v1/warehouses [post]
func (handler *WarehouseHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.WarehouseRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}
		ctx := c.Request.Context()
		if err := WarehouseFullRequestValidator(c, req); err != nil {
			web_utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		result, e := handler.warehouseService.Create(&ctx, req)
		if e != nil {
			switch e {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, e.Error())
				return
			default:
				web_utils.Error(c, http.StatusInternalServerError, e.Error())
				return
			}
		}
		web_utils.Success(c, http.StatusCreated, result)
	}
}

// @Summary		Update warehouses
// @Tags			Warehouses
// @Description	update warehouses
// @Accept			json
// @Produce		json
// @Param			id			path		int							true	"WarehouseHandler ID"
// @Param			WarehouseHandler	body		dtos.WarehouseRequestDTO	true	"WarehouseHandler to update"
// @Success		200			{object}	web_utils.response
// @Router			/api/v1/warehouses/{id} [patch]
func (handler *WarehouseHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		var req dtos.WarehouseRequestDTO

		if e != nil {
			web_utils.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}

		result, err := handler.warehouseService.Update(&ctx, warehouseId, req)

		if err != nil {
			web_utils.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web_utils.Success(c, http.StatusOK, result)
	}
}

// @Summary		Delete warehouses
// @Tags			Warehouses
// @Description	delete warehouses by id
// @Param			id	path	int	true	"WarehouseHandler ID"
// @Success		204
// @Router			/api/v1/warehouses/{id} [delete]
func (handler *WarehouseHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		if e != nil {
			web_utils.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		err := handler.warehouseService.Delete(&ctx, warehouseId)

		if err != nil {
			web_utils.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web_utils.Success(c, http.StatusNoContent, nil)
	}
}

func WarehouseFullRequestValidator(c *gin.Context, req dtos.WarehouseRequestDTO) error {
	if req.Address == "" {
		return errors.New("field address is required")
	}
	if req.Telephone == "" {
		return errors.New("field telephone is required")
	}
	if req.WarehouseCode == "" {
		return errors.New("field warehouse_code is required")
	}
	if req.MinimumCapacity == 0 {
		return errors.New("field minimum_capacity is required")
	}
	if req.MinimumTemperature == 0 {
		return errors.New("field minimum_temperature is required")
	}

	return nil
}
