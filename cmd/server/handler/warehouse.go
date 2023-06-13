package handler

import (
	"errors"
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
// @Summary Get warehouse
// @Tags Warehouses
// @Description get one warehouse by id
// @Produce  json
// @Param        id   path      int  true  "Warehouse ID"
// @Success 200 {object} domain.Warehouse
// @Router /warehouses/:id [get]
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))

		if e != nil {
			web.Error(c, 400, "parameter id must be a integer")
			return
		}

		result, err := w.warehouseService.GetOne(c, warehouseId)

		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}

		web.Success(c, 200, result)
	}
}

// GetAllWarehouses godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description get warehouses
// @Produce  json
// @Success 200 {object} []domain.Warehouse
// @Router /warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := w.warehouseService.GetAll(c)

		if err != nil {
			web.Error(c, 404, err.Error())
			return
		}

		web.Success(c, 200, result)
		return
	}
}

// createWarehouses godoc
// @Summary Create warehouse
// @Tags Warehouses
// @Description Create warehouse
// @Accept  json
// @Produce  json
// @Param dtos.WarehouseRequestDTO body request true "warehouse to create"
// @Success 200 {object} domain.Warehouse
// @Router /warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.WarehouseRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 400, "JSON format may be wrong")
			return
		}

		if err := WarehouseFullRequestValidator(c, req); err != nil {
			web.Error(c, 422, err.Error())
			return
		}

		result, e := w.warehouseService.Create(c, req)
		if e != nil {
			web.Error(c, 400, e.Error())
			return
		}

		web.Success(c, 201, result)
		return
	}
}

// @Summary Update warehouses
// @Tags Warehouses
// @Description update warehouses
// @Accept json
// @Produce json
// @Param        id   path      int  true  "Warehouse ID"
// @Param dtos.WarehouseRequestDTO body request true "Warehouse to update"
// @Success 200 {object} web.Response
// @Router /warehouses/:id [put]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouseId, e := strconv.Atoi(c.Param("id"))
		var req dtos.WarehouseRequestDTO

		if e != nil {
			web.Error(c, 400, "parameter id must be a integer")
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 400, "JSON format may be wrong")
			return
		}

		result, err := w.warehouseService.Update(c, warehouseId, req)

		if err != nil {
			web.Error(c, 400, e.Error())
			return
		}

		web.Success(c, 200, result)
	}
}

// @Summary Delete warehouse
// @Tags Warehouses
// @Description delete warehouse by id
// @Param        id   path      int  true  "Warehouse ID"
// @Success 204
// @Router /warehouses/:id [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
			warehouseId, e := strconv.Atoi(c.Param("id"))	
			if e != nil {
				web.Error(c, 400, "parameter id must be a integer")
				return
			}

			err := w.warehouseService.Delete(c, warehouseId)

			if err != nil {
				web.Error(c, 404, err.Error())
				return
			}

			web.Success(c, 204, nil)
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