package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestProduct struct {
	Description    string  `json:"description"`
	ExpirationRate int     `json:"expiration_rate"`
	FreezingRate   int     `json:"freezing_rate"`
	Height         float32 `json:"height"`
	Length         float32 `json:"length"`
	Netweight      float32 `json:"netweight"`
	ProductCode    string  `json:"product_code"`
	RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
	Width          float32 `json:"width"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       int     `json:"seller_id"`
}

type Product struct {
	productService product.Service
}

func NewProduct(p product.Service) *Product {
	return &Product{
		productService: p,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, products)
	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		productResponse, err := p.productService.Get(c, int(id))
		if err != nil {
			switch err {
			case product.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting product %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusOK, productResponse)
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestProduct
		if err := c.Bind(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if req.Description == "" {
			web.Error(c, http.StatusUnprocessableEntity, "The field Description is required.")
			return
		}

		if req.ExpirationRate == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field ExpirationRate is required.")
			return
		}

		if req.FreezingRate == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field FreezingRate is required.")
			return
		}

		if req.Height == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field Height is required.")
			return
		}

		if req.Length == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field Length is required.")
			return
		}

		if req.Netweight == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field Netweight is required.")
			return
		}

		if req.ProductCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, "The field ProductCode is required.")
			return
		}

		if req.RecomFreezTemp == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field RecomFreezTemp is required.")
			return
		}

		if req.Width == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field Width is required.")
			return
		}

		if req.ProductTypeID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field ProductTypeID is required.")
			return
		}

		if req.SellerID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field SellerID is required.")
			return
		}

		productResponse, err := p.productService.Save(c, req.Description, req.ExpirationRate, req.FreezingRate, req.Height,
			req.Length, req.Netweight, req.ProductCode, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID)
		if err != nil {
			switch err {
			case product.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusCreated, productResponse)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		var req requestProduct
		if err := c.Bind(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		sectionResponse, err := p.productService.Update(c, req.Description, req.ExpirationRate, req.FreezingRate, req.Height,
			req.Length, req.Netweight, req.ProductCode, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID, id)
		if err != nil {
			switch err {
			case product.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error updating product %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusOK, sectionResponse)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		err = p.productService.Delete(c, int(id))
		if err != nil {
			switch err {
			case product.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error deleting product %s", err.Error()))
			}
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
