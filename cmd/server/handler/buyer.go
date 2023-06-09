package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/gin-gonic/gin"
)

type BuyerHandler struct {
	buyerService buyer.Service
}

func NewBuyer(buyerService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		buyerService,
	}
}

func (handler *BuyerHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Parse id from url
		var buyerRequest domain.BuyerIDRequestDTO
		if err := c.ShouldBindUri(&buyerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if buyerResponse, err := handler.buyerService.Get(buyerRequest.ID); err != nil {
			switch err {
			case buyer.ErrNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"buyer": buyerResponse})
			return
		}

	}
}

func (handler *BuyerHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (handler *BuyerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (handler *BuyerHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (handler *BuyerHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
