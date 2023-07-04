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
