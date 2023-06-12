package handler

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
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
			web.Error(c, 404, "no warehouses were found")
			return
		}

		web.Success(c, 200, result)
		return
	}
}

func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.WarehouseRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Response(c, 400, nil)
			return
		}

		if req.Address == "" {
			web.Error(c, 422, `field "address" is missing` )
			return 
		}
		if req.Telephone == "" {
			web.Error(c, 422, `field "telephone" is missing`)
			return
		}
		if req.WarehouseCode == "" {
			web.Error(c, 422, `field "warehousecode" is missing`)
			return
		}
		if req.MinimumCapacity == 0 {
			web.Error(c, 422, `field "minimumcapacity" is missing`)
			return
		}
		if req.MinimumTemperature == 0 {
			web.Error(c, 422, `field "minimumtemperature" is missing`)
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
