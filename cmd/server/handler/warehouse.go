package handler

import (
	"errors"

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

func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {


	}
}

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

func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
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