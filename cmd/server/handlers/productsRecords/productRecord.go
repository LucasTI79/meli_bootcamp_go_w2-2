package productsRecords

import (
	"net/http"

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
