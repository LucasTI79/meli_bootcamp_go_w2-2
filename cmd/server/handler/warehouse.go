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
	return func(c *gin.Context) {}
}

func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.WarehouseRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Response(c, 400, nil)
			return
		}

		result, e := w.warehouseService.Create(req)
		if e != nil {
			web.Error(c, 400, e.Error())
		}

		web.Success(c, 201, result)
	}
}

func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
