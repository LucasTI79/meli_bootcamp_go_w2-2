package productsRecords

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type RequestCreateProductRecord struct {
	LastUpdateRate string  `json:"lastUpdateRate"`
	PurchasePrice  float32 `json:"purchasePrice"`
	SalePrice      float32 `json:"salePrice"`
	ProductId      int     `json:"productId"`
}

type RequestUpdateProductRecord struct {
	LastUpdateRate *string  `json:"lastUpdateRate"`
	PurchasePrice  *float32 `json:"purchasePrice"`
	SalePrice      *float32 `json:"salePrice"`
	ProductId      *int     `json:"productId"`
}

type ProductRecord struct {
	productRecordService productRecord.Service
}

func NewProductRecord(p productRecord.Service) *ProductRecord {
	return &ProductRecord{
		productRecordService: p,
	}
}

// Method GetAll
// ListProductsRecords godoc
//
//	@Summary		List productsRecords
//	@Tags			ProductsRecords
//	@LastUpdateRate	getAll productsRecords
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response
//	@Router			/api/v1/productsRecords [get]
func (p *ProductRecord) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		productsRecords, err := p.productRecordService.GetAll(&ctx)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(*productsRecords) == 0 {
			web.Success(c, http.StatusNoContent, nil)
		}
		web.Success(c, http.StatusOK, productsRecords)
	}
}

// Method Get
// GetProductsRecords godoc
//
//	@Summary		Get ProductRecord
//	@Tags			ProductsRecords
//	@LastUpdateRate	Get the details of a ProductsRecords
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of ProductRecord to be searched"
//	@Success		200	{object}	web.response
//	@Router			/api/v1/productsRecords/{id} [get]
func (p *ProductRecord) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		ctx := c.Request.Context()
		productRecordResponse, err := p.productRecordService.Get(&ctx, int(id))
		if err != nil {
			switch err {
			case productRecord.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting productRecord %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusOK, productRecordResponse)
	}
}

// Method Create
// CreateProductsRecords godoc
//
//	@Summary		Create ProductRecord
//	@Tags			ProductsRecords
//	@LastUpdateRate	Create ProductRecord
//	@Accept			json
//	@Produce		json
//	@Param			ProductRecord	body		requestCreateProductRecord	true	"ProductRecord to Create"
//	@Success		201		{object}	web.response
//	@Router			/api/v1/productsRecords [post]
func (p *ProductRecord) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestCreateProductRecord
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if req.LastUpdateRate == "" {
			web.Error(c, http.StatusUnprocessableEntity, "The field LastUpdateRate is required.")
			return
		}

		if req.PurchasePrice == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field PurchasePrice is required.")
			return
		}

		if req.SalePrice == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field SalePrice is required.")
			return
		}

		if req.ProductId == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field ProductId is required.")
			return
		}

		ctx := c.Request.Context()
		productRecordResponse, err := p.productRecordService.Save(&ctx, req.LastUpdateRate, req.PurchasePrice, req.SalePrice, req.ProductId)
		if err != nil {
			switch err {
			case productRecord.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusCreated, productRecordResponse)
	}
}

// Method Update
// UpdateProductsRecords godoc
//
//	@Summary		Update ProductRecord
//	@Tags			ProductsRecords
//	@LastUpdateRate	Update the details of a ProductRecord
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string			true	"ID of ProductsRecords to be updated"
//	@Param			ProductsRecords	body		requestUpdateProductRecord	true	"Updated ProductRecord details"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/productsRecords/{id} [patch]
func (p *ProductRecord) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		var req RequestUpdateProductRecord
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		ctx := c.Request.Context()
		productRecordResponse, err := p.productRecordService.Update(&ctx, req.LastUpdateRate, req.PurchasePrice, req.SalePrice, req.ProductId, id)
		if err != nil {
			switch err {
			case productRecord.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			case productRecord.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error updating productRecord %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusOK, productRecordResponse)
	}
}

// Method Delete
// DeleteProductsRecords godoc
//
//	@Summary		Delete ProductRecord
//	@Tags			ProductsRecords
//	@Description	Delete ProductRecord
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a ProductRecord to be excluded"
//	@Success		204	{object}	web.response
//	@Router			/api/v1/productsRecords/{id} [delete]
func (p *ProductRecord) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = p.productRecordService.Delete(&ctx, int(id))
		if err != nil {
			switch err {
			case productRecord.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error deleting productRecord %s", err.Error()))
			}
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
