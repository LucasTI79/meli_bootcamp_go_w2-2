package handlers

import (
	"errors"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	buyerService buyer.Service
}

func NewBuyerHandler(buyerService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		buyerService,
	}
}

// Get is the handlers to search for a buyer and return their details.
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

// GetAll is the handlers to search for all buyers.
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

		if buyers, err := handler.buyerService.GetAll(c); err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			web.Success(c, http.StatusOK, buyers)
			return
		}
	}
}

// Create is the handlers to create a buyer.
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

		if createdBuyer, err := handler.buyerService.Create(c, createBuyerRequest); err != nil {
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

// Update is the handlers to update a buyer details.
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
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if updatedBuyer, err := handler.buyerService.Update(c, id, updateBuyerRequest); err != nil {
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

// Delete is the handlers to delete a buyer.
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

		if err := handler.buyerService.Delete(c, id); err != nil {
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

func getIdFromUri(c *gin.Context) (id int, err error) {

	value, _ := c.Params.Get("id")
	id, err = strconv.Atoi(value)

	if err != nil {
		err = errors.New(fmt.Sprintf("Invalid id on request: %s", c.Request.RequestURI))
		return
	}

	return

}
