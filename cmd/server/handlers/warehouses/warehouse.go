package warehouses

import (
	"errors"
	"net/http"
	"strconv"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(s warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: s,
	}
}

// GetOneWarehouse godoc
//
//	@Summary		Get warehouses
//	@Tags			Warehouses
//	@Description	get one warehouses by id
//	@Produce		json
//	@Param			id	path		int	true	"Warehouse ID"
//	@Success		200	{object}	domain.Warehouse
//	@Router			/api/v1/warehouses/{id} [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		if e != nil {
			web.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		result, err := w.warehouseService.GetOne(&ctx, warehouseId)

		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, result)
	}
}

// GetAllWarehouses godoc
//
//	@Summary		List warehouses
//	@Tags			Warehouses
//	@Description	get warehouses
//	@Produce		json
//	@Success		200	{object}	[]domain.Warehouse
//	@Router			/api/v1/warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		warehouses, err := w.warehouseService.GetAll(&ctx)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get warehouses: %s", err.Error())
			return
		}

		if len(*warehouses) == 0 {
			web.Error(c, http.StatusNoContent, "There are no warehouses stored")
			return
		}

		web.Success(c, http.StatusOK, *warehouses)
	}
}

// createWarehouses godoc
//
//	@Summary		Create warehouses
//	@Tags			Warehouses
//	@Description	Create warehouses
//	@Accept			json
//	@Produce		json
//	@Param			Warehouse	body		dtos.WarehouseRequestDTO	true	"warehouses to create"
//	@Success		200			{object}	domain.Warehouse
//	@Router			/api/v1/warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.WarehouseRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}
		ctx := c.Request.Context()
		if err := WarehouseFullRequestValidator(c, req); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		result, e := w.warehouseService.Create(&ctx, req)

		if e != nil {
			web.Error(c, http.StatusConflict, e.Error())
			return
		}
		web.Success(c, http.StatusCreated, result)
	}
}

// @Summary		Update warehouses
// @Tags			Warehouses
// @Description	update warehouses
// @Accept			json
// @Produce		json
// @Param			id			path		int							true	"Warehouse ID"
// @Param			Warehouse	body		dtos.WarehouseRequestDTO	true	"Warehouse to update"
// @Success		200			{object}	domain.Warehouse
// @Router			/api/v1/warehouses/{id} [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		var req dtos.WarehouseRequestDTO

		if e != nil {
			web.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}

		result, err := w.warehouseService.Update(&ctx, warehouseId, req)

		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, result)
	}
}

// @Summary		Delete warehouses
// @Tags			Warehouses
// @Description	delete warehouses by id
// @Param			id	path	int	true	"Warehouse ID"
// @Success		204
// @Router			/api/v1/warehouses/{id} [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		ctx := c.Request.Context()
		if e != nil {
			web.Error(c, http.StatusBadRequest, "parameter id must be a integer")
			return
		}

		err := w.warehouseService.Delete(&ctx, warehouseId)

		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)
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
