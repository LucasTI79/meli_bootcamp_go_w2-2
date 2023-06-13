package handler

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

func NewBuyer(buyerService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		buyerService,
	}
}

func (handler *BuyerHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if buyerResponse, err := handler.buyerService.Get(c, id); err != nil {
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

func (handler *BuyerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		createBuyerRequest := new(dtos.CreateBuyerRequestDTO)
		if err := c.ShouldBindJSON(createBuyerRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if createdBuyer, err := handler.buyerService.Create(c, createBuyerRequest); err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			web.Response(c, http.StatusCreated, createdBuyer)
			return
		}
	}
}

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
