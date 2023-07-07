package buyers

import (
	"errors"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	buyerService         buyer.Service
	purchaseOrderService services.PurchaseOrderService
}

func NewBuyerHandler(buyerService buyer.Service, purchaseOrderService services.PurchaseOrderService) *BuyerHandler {
	return &BuyerHandler{
		buyerService,
		purchaseOrderService,
	}
}

// Get is the handler to search for a buyer and return their details.
//
//	@Summary		Get Buyer
//	@Tags			Buyers
//	@Description	Get the details of a Buyer
//	@Produce		json
//	@Param			id	path		string	true	"ID of Buyer to be searched"
//	@Success		200	{object}	domain.Buyer
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/buyers/{id} [get]
func (handler *BuyerHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()

		if buyerResponse, err := handler.buyerService.Get(&ctx, id); err != nil {
			switch err {
			case buyer.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, buyerResponse)
			return
		}

	}
}

// GetAll is the handler to search for all buyers.
//
//	@Summary		List Buyers
//	@Tags			Buyers
//	@Description	Get the details of all buyers on the database.
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response{data=[]domain.Buyer}
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/buyers [get]
func (handler *BuyerHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		if buyers, err := handler.buyerService.GetAll(&ctx); err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			if len(*buyers) == 0 {
				web.Success(c, http.StatusNoContent, buyers)
				return
			}
			web.Success(c, http.StatusOK, buyers)
			return
		}
	}
}

// Create is the handler to create a buyer.
//
//	@Summary		Create Buyer
//	@Tags			Buyers
//	@Description	Save a buyer on the database.
//	@Accept			json
//	@Produce		json
//	@Param			Seller	body		dtos.CreateBuyerRequestDTO	true	"Buyer to Create"
//	@Success		201		{object}	domain.Buyer
//	@Failure		422		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/buyers [post]
func (handler *BuyerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		createBuyerRequest := new(dtos.CreateBuyerRequestDTO)
		if err := c.ShouldBindJSON(createBuyerRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := c.Request.Context()
		if createdBuyer, err := handler.buyerService.Create(&ctx, createBuyerRequest); err != nil {
			switch err {
			case buyer.ErrCardNumberDuplicated:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			web.Response(c, http.StatusCreated, createdBuyer)
			return
		}
	}
}

// Update is the handler to update a buyer details.
//
//	@Summary		Update Buyer
//	@Tags			Buyers
//	@Description	Update the details of a Buyer
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of Buyer to be updated"
//	@Param			Buyer	body		dtos.UpdateBuyerRequestDTO	true	"Updated Buyer details"
//	@Success		200		{object}	domain.Buyer
//	@Failure		400		{object}	web.errorResponse
//	@Failure		404		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/buyers/{id} [patch]
func (handler *BuyerHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		updateBuyerRequest := new(dtos.UpdateBuyerRequestDTO)
		if err := c.ShouldBind(updateBuyerRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := c.Request.Context()
		if updatedBuyer, err := handler.buyerService.Update(&ctx, id, updateBuyerRequest); err != nil {
			switch err {
			case buyer.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			case buyer.ErrCardNumberDuplicated:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, updatedBuyer)
			return
		}
	}
}

// Delete is the handler to delete a buyer.
//
//	@Summary		Delete Buyer
//	@Tags			Buyers
//	@Description	Delete Buyers
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"ID of a Buyer to be excluded"
//	@Success		204
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/buyers/{id} [delete]
func (handler *BuyerHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		if err := handler.buyerService.Delete(&ctx, id); err != nil {
			switch err {
			case buyer.ErrNotFound:
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

// CountPurchaseOrders is the handler search for a purchaseOrder and return the number of sellers
//
//	@Summary		CountPurchaseOrders
//	@Tags			PurchaseOrders
//	@Description	search for a purchaseOrder and return the number of sellers.
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of PurchaseOrder to be searched"
//	@Success		200		{object}	web.response{data=dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO}
//	@Failure		400		{object}	web.errorResponse
//	@Failure		404		{object}	web.errorResponse
//	@Failure		422		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/buyers/{id}/report-purchase-orders [get]
func (handler *BuyerHandler) CountPurchaseOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()

		count, err := handler.purchaseOrderService.CountByBuyerID(&ctx, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		response := dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO{
			BuyerID:             id,
			PurchaseOrdersCount: count,
		}

		web.Success(c, http.StatusOK, response)
		return

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
