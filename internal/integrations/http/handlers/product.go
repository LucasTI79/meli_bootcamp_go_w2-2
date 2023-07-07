package handlers

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/web_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestCreateProduct struct {
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

type RequestUpdateProduct struct {
	Description    *string  `json:"description"`
	ExpirationRate *int     `json:"expiration_rate"`
	FreezingRate   *int     `json:"freezing_rate"`
	Height         *float32 `json:"height"`
	Length         *float32 `json:"length"`
	Netweight      *float32 `json:"netweight"`
	ProductCode    *string  `json:"product_code"`
	RecomFreezTemp *float32 `json:"recommended_freezing_temperature"`
	Width          *float32 `json:"width"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id"`
}

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(p services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: p,
	}
}

// Method GetAll
// ListProducts godoc
//
//	@Summary		List products
//	@Tags			Products
//	@Description	getAll products
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/products [get]
func (handler *ProductHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		products, err := handler.productService.GetAll(&ctx)
		if err != nil {
			web_utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(*products) == 0 {
			web_utils.Success(c, http.StatusNoContent, nil)
		}
		web_utils.Success(c, http.StatusOK, products)
	}
}

// Method Get
// GetProducts godoc
//
//	@Summary		Get ProductHandler
//	@Tags			Products
//	@Description	Get the details of a Products
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of ProductHandler to be searched"
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/products/{id} [get]
func (handler *ProductHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		ctx := c.Request.Context()
		productResponse, err := handler.productService.Get(&ctx, int(id))
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting product %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusOK, productResponse)
	}
}

// Method Create
// CreateProducts godoc
//
//	@Summary		Create ProductHandler
//	@Tags			Products
//	@Description	Create ProductHandler
//	@Accept			json
//	@Produce		json
//	@Param			ProductHandler	body		RequestCreateProduct	true	"ProductHandler to Create"
//	@Success		201		{object}	web_utils.response
//	@Router			/api/v1/products [post]
func (handler *ProductHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestCreateProduct
		if err := c.ShouldBindJSON(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if req.Description == "" {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field Description is required.")
			return
		}

		if req.ExpirationRate == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field ExpirationRate is required.")
			return
		}

		if req.FreezingRate == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field FreezingRate is required.")
			return
		}

		if req.Height == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field Height is required.")
			return
		}

		if req.Length == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field Length is required.")
			return
		}

		if req.Netweight == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field Netweight is required.")
			return
		}

		if req.ProductCode == "" {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field ProductCode is required.")
			return
		}

		if req.RecomFreezTemp == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field RecomFreezTemp is required.")
			return
		}

		if req.Width == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field Width is required.")
			return
		}

		if req.ProductTypeID == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field ProductTypeID is required.")
			return
		}

		if req.SellerID == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field SellerID is required.")
			return
		}
		ctx := c.Request.Context()
		productResponse, err := handler.productService.Save(&ctx, req.Description, req.ExpirationRate, req.FreezingRate, req.Height,
			req.Length, req.Netweight, req.ProductCode, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID)
		if err != nil {
			switch err {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusCreated, productResponse)
	}
}

// Method Update
// UpdateProducts godoc
//
//	@Summary		Update ProductHandler
//	@Tags			Products
//	@Description	Update the details of a ProductHandler
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string			true	"ID of Products to be updated"
//	@Param			Products	body		RequestUpdateProduct	true	"Updated ProductHandler details"
//	@Success		200			{object}	web_utils.response
//	@Router			/api/v1/products/{id} [patch]
func (handler *ProductHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		var req RequestUpdateProduct
		if err := c.Bind(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		ctx := c.Request.Context()
		productResponse, err := handler.productService.Update(&ctx, req.Description, req.ExpirationRate, req.FreezingRate, req.Height,
			req.Length, req.Netweight, req.ProductCode, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID, id)
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error updating product %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusOK, productResponse)
	}
}

// Method Delete
// DeleteSections godoc
//
//	@Summary		Delete ProductHandler
//	@Tags			Products
//	@Description	Delete ProductHandler
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a ProductHandler to be excluded"
//	@Success		204	{object}	web_utils.response
//	@Router			/api/v1/products/{id} [delete]
func (handler *ProductHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = handler.productService.Delete(&ctx, int(id))
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error deleting product %s", err.Error()))
			}
			return
		}

		web_utils.Success(c, http.StatusNoContent, nil)
	}
}
