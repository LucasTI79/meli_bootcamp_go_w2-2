package handlers

import (
	"errors"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/web_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	sellerService services.SellerService
}

func NewSeller(s services.SellerService) *SellerHandler {
	return &SellerHandler{
		sellerService: s,
	}
}

// Method GetAll
// ListSellers godoc
//
//	@Summary		List sellers
//	@Tags			Sellers
//	@Description	getAll sellers
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/sellers [get]
func (handler *SellerHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		sellers, err := handler.sellerService.GetAll(&ctx)
		if err != nil {
			web_utils.Error(c, http.StatusInternalServerError, "Failed to get sellers: %s", err.Error())
			return
		}

		if len(*sellers) == 0 {
			web_utils.Error(c, http.StatusNoContent, "There are no sellers stored")
			return
		}

		web_utils.Success(c, http.StatusOK, *sellers)
	}
}

// Method Create
// CreateSellers godoc
//
//	@Summary		Create Sellers
//	@Tags			Sellers
//	@Description	Create sellers
//	@Accept			json
//	@Produce		json
//	@Param			SellerHandler	body		dtos.CreateSellerRequestDTO	true	"SellerHandler to Create"
//	@Success		201		{object}	web_utils.response
//	@Router			/api/v1/sellers [post]
func (handler *SellerHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		createSellerRequestDTO := new(dtos.CreateSellerRequestDTO)
		if err := c.ShouldBindJSON(createSellerRequestDTO); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, "Error to read request: %s", err.Error())
			return
		}

		sellerDomain := &entities.Seller{
			CID:         createSellerRequestDTO.CID,
			CompanyName: createSellerRequestDTO.CompanyName,
			Address:     createSellerRequestDTO.Address,
			Telephone:   createSellerRequestDTO.Telephone,
			LocalityID:  createSellerRequestDTO.LocalityID,
		}

		ctx := c.Request.Context()
		sellerDomain, err := handler.sellerService.Save(&ctx, *sellerDomain)
		if err != nil {
			switch err {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web_utils.Error(c, http.StatusInternalServerError, "Error to save request: %s", err.Error())
				return
			}
		}

		web_utils.Success(c, http.StatusCreated, *sellerDomain)
	}
}

// Method Get
// GetSellers godoc
//
//	@Summary		Get Sellers
//	@Tags			Sellers
//	@Description	Get the details of a Sellers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of Sellers to be searched"
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/sellers/{id} [get]
func (handler *SellerHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid SellerHandler ID: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		sellerResult, err := handler.sellerService.Get(&ctx, int(id))

		if err != nil {
			if errors.Is(err, services.ErrNotFound) {
				web_utils.Error(c, http.StatusNotFound, "SellerHandler not found: %s", err.Error())
				return
			}
			web_utils.Error(c, http.StatusInternalServerError, "Error to process the request, try again: %s", err.Error())
			return
		}

		web_utils.Success(c, http.StatusOK, *sellerResult)
	}
}

// Method Update
// UpdateSellers godoc
//
//	@Summary		Update Sellers
//	@Tags			Sellers
//	@Description	Update the details of a Sellers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of Sellers to be updated"
//	@Param			Sellers	body		dtos.UpdateSellerRequestDTO	true	"Updated Sellers details"
//	@Success		200		{object}	web_utils.response
//	@Router			/api/v1/sellers/{id} [patch]
func (handler *SellerHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		updateSellerRequestDTO := new(dtos.UpdateSellerRequestDTO)

		if err := c.Bind(&updateSellerRequestDTO); err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		ctx := c.Request.Context()
		sellerUpdated, err := handler.sellerService.Update(&ctx, int(id), updateSellerRequestDTO)
		if err != nil {
			switch err {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
				return
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
				return
			}

		}

		web_utils.Success(c, http.StatusOK, sellerUpdated)
	}
}

// Method Delete
// DeleteSellers godoc
//
//	@Summary		Delete Sellers
//	@Tags			Sellers
//	@Description	Delete Sellers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a Sellers to be excluded"
//	@Success		204	{object}	web_utils.response
//	@Router			/api/v1/sellers/{id} [delete]
func (handler *SellerHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = handler.sellerService.Delete(&ctx, int(id))
		if err != nil {
			web_utils.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}
		web_utils.Success(c, http.StatusNoContent, nil)
	}
}
