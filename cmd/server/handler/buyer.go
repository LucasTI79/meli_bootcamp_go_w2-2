package handler

import (
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	buyerService buyer.Service
}

// Errors
var (
	ErrMissingIdOnRequest = errors.New("id is required")
)

func NewBuyer(buyerService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		buyerService,
	}
}

func (handler *BuyerHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Parse id from url
		var buyerRequest dtos.BuyerIDRequestDTO
		if err := c.ShouldBindUri(&buyerRequest); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		if buyerResponse, err := handler.buyerService.Get(c, buyerRequest.ID); err != nil {
			switch err {
			case buyer.ErrNotFound:
				web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, buyerResponse)
			return
		}

	}
}

func (handler *BuyerHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		if buyers, err := handler.buyerService.GetAll(c); err != nil {
			web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			return
		} else {
			web.Success(c, http.StatusOK, buyers)
			return
		}
	}
}

func (handler *BuyerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Parse buyer from body
		createBuyerRequestDTO := new(dtos.CreateBuyerRequestDTO)
		if err := c.ShouldBindJSON(createBuyerRequestDTO); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "error: %s", err.Error())
			return
		}

		// Parse DTO to entity
		buyer := &domain.Buyer{
			CardNumberID: createBuyerRequestDTO.CardNumberID,
			FirstName:    createBuyerRequestDTO.FirstName,
			LastName:     createBuyerRequestDTO.LastName,
		}

		if buyer, err := handler.buyerService.Create(c, buyer); err != nil {
			web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			return
		} else {
			web.Response(c, http.StatusCreated, buyer)
			return
		}
	}
}

func (handler *BuyerHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		updateBuyerRequest := new(dtos.UpdateBuyerRequestDTO)

		// Parse id from url
		if err := c.ShouldBindUri(updateBuyerRequest); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", ErrMissingIdOnRequest.Error())
			return
		}

		// Parse buyer from body
		if err := c.ShouldBindJSON(updateBuyerRequest); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", err.Error())
			return
		}

		if updatedBuyer, err := handler.buyerService.Update(c, updateBuyerRequest); err != nil {
			switch err {
			case buyer.ErrNotFound:
				web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			}
			return
		} else {
			web.Success(c, http.StatusOK, updatedBuyer)
			return
		}
	}
}

func (handler *BuyerHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Parse id from url
		var buyerRequest dtos.BuyerIDRequestDTO
		if err := c.ShouldBindUri(&buyerRequest); err != nil {
			web.Error(c, http.StatusBadRequest, "error: %s", ErrMissingIdOnRequest.Error())
			return
		}

		if err := handler.buyerService.Delete(c, buyerRequest.ID); err != nil {
			switch err {
			case buyer.ErrNotFound:
				web.Error(c, http.StatusNotFound, "error: %s", err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, "error: %s", err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusNoContent, nil)
			return
		}
	}
}
