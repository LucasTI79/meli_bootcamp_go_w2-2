package purchase_orders

import (
	"errors"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseOrderHandler struct {
	purchaseOrderService services.PurchaseOrderService
}

func NewPurchaseOrderHandler(purchaseOrderService services.PurchaseOrderService) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{
		purchaseOrderService,
	}
}

// Get is the handler to search for a purchaseOrder and return their details.
//
//	@Summary		Get PurchaseOrder
//	@Tags			PurchaseOrders
//	@Description	Get the details of a PurchaseOrder
//	@Produce		json
//	@Param			id	path		string	true	"ID of PurchaseOrder to be searched"
//	@Success		200	{object}	entities.PurchaseOrder
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/purchase-orders/{id} [get]
func (handler *PurchaseOrderHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		if purchaseOrder, err := handler.purchaseOrderService.Get(&ctx, id); err != nil {
			switch err {
			case services.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, purchaseOrder)
			return
		}

	}
}

// GetAll is the handler to search for all purchase-orders.
//
//	@Summary		List PurchaseOrders
//	@Tags			PurchaseOrders
//	@Description	Get the details of all purchase-orders on the database.
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response{data=[]entities.PurchaseOrder}
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/purchase-orders [get]
func (handler *PurchaseOrderHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		if purchaseOrders, err := handler.purchaseOrderService.GetAll(&ctx); err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			if len(purchaseOrders) == 0 {
				web.Success(c, http.StatusNoContent, purchaseOrders)
				return
			}
			web.Success(c, http.StatusOK, purchaseOrders)
			return
		}
	}
}

// Create is the handler to create a purchaseOrder.
//
//	@Summary		Create PurchaseOrder
//	@Tags			PurchaseOrders
//	@Description	Save a purchaseOrder on the database.
//	@Accept			json
//	@Produce		json
//	@Param			Seller	body		entities.PurchaseOrder	true	"PurchaseOrder to Create"
//	@Success		201		{object}	entities.PurchaseOrder
//	@Failure		422		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/purchase-orders [post]
func (handler *PurchaseOrderHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createPurchaseOrderRequestDTO dtos.CreatePurchaseOrderRequestDTO
		if err := c.ShouldBindJSON(&createPurchaseOrderRequestDTO); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		createPurchaseOrderRequest := createPurchaseOrderRequestDTO.ToDomain()

		ctx := c.Request.Context()
		if createdPurchaseOrder, err := handler.purchaseOrderService.Create(&ctx, createPurchaseOrderRequest); err != nil {
			switch err {
			case services.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			web.Response(c, http.StatusCreated, createdPurchaseOrder)
			return
		}
	}
}

// Update is the handler to update a purchaseOrder details.
//
//	@Summary		Update PurchaseOrder
//	@Tags			PurchaseOrders
//	@Description	Update the details of a PurchaseOrder
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of PurchaseOrder to be updated"
//	@Param			PurchaseOrder	body		dtos.UpdatePurchaseOrderRequestDTO	true	"Updated PurchaseOrder details"
//	@Success		200		{object}	entities.PurchaseOrder
//	@Failure		400		{object}	web.errorResponse
//	@Failure		404		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/purchase-orders/{id} [patch]
func (handler *PurchaseOrderHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		var updatePurchaseOrderRequest dtos.UpdatePurchaseOrderRequestDTO
		if err := c.ShouldBindJSON(&updatePurchaseOrderRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := c.Request.Context()
		if updatedPurchaseOrder, err := handler.purchaseOrderService.Update(&ctx, id, updatePurchaseOrderRequest); err != nil {
			switch err {
			case services.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			case services.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, updatedPurchaseOrder)
			return
		}
	}
}

// Delete is the handler to delete a purchaseOrder.
//
//	@Summary		Delete PurchaseOrder
//	@Tags			PurchaseOrders
//	@Description	Delete PurchaseOrders
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"ID of a PurchaseOrder to be excluded"
//	@Success		204
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/purchase-orders/{id} [delete]
func (handler *PurchaseOrderHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		if err := handler.purchaseOrderService.Delete(&ctx, id); err != nil {
			switch err {
			case services.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusNoContent, nil)
			return
		}
	}
}

func getIdFromUri(c *gin.Context) (id int, err error) {

	value, _ := c.Params.Get("id")
	id, err = strconv.Atoi(value)

	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid id on request: %s", c.Request.RequestURI))
		return
	}

	return

}
