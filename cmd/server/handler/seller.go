package handler

import (
	"errors"
	"net/http"
	"strconv"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
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
//	@Success		200	{object}	web.response
//	@Router			/api/v1/sellers [get]
func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get sellers: %s", err.Error())
			return
		}

		if len(*sellers) == 0 {
			web.Error(c, http.StatusNoContent, "There are no sellers stored: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, *sellers)
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
//	@Param			Seller	body		dtos.CreateSellerRequestDTO	true	"Seller to Create"
//	@Success		201		{object}	web.response
//	@Router			/api/v1/sellers [post]
func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		createSellerRequestDTO := new(dtos.CreateSellerRequestDTO)
		if err := c.Bind(&createSellerRequestDTO); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "Error to read request: %s", err.Error())
			return
		}

		sellerDomain := &domain.Seller{
			CID:         createSellerRequestDTO.CID,
			CompanyName: createSellerRequestDTO.CompanyName,
			Address:     createSellerRequestDTO.Address,
			Telephone:   createSellerRequestDTO.Telephone,
		}

		if sellerDomain.CID <= 0.0 {
			web.Error(c, http.StatusBadRequest, "Field CID is required: %s", "")
			return
		}

		if sellerDomain.CompanyName == "" {
			web.Error(c, http.StatusBadRequest, "Field CompanyName is required: %s", "")
			return
		}

		if sellerDomain.Address == "" {
			web.Error(c, http.StatusBadRequest, "Field Address is required: %s", "")
			return
		}

		if sellerDomain.Telephone == "" {
			web.Error(c, http.StatusBadRequest, "Field Telephone is required: %s", "")
			return
		}

		sellerDomain, err := s.sellerService.Save(c, *sellerDomain)
		if err != nil {
			switch err {
			case seller.ErrConflict:
				web.Error(c, http.StatusConflict, "Seller with CID already exists: %s", err.Error())
				return
			default:
				web.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, *sellerDomain)
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
//	@Success		200	{object}	web.response
//	@Router			/api/v1/sellers/{id} [get]
func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid Seller ID: %s", err.Error())
			return
		}

		sellerResult, err := s.sellerService.Get(c, int(id))

		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "Seller not found: %s", err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, "Error to process the request, try again: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, *sellerResult)
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
//	@Success		200		{object}	web.response
//	@Router			/api/v1/sellers/{id} [patch]
func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		updateSellerRequestDTO := new(dtos.UpdateSellerRequestDTO)

		if err := c.Bind(&updateSellerRequestDTO); err != nil {
			web.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		sellerUpdated, err := s.sellerService.Update(c, int(id), updateSellerRequestDTO)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to update: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, sellerUpdated)
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
//	@Success		204	{object}	web.response
//	@Router			/api/v1/sellers/{id} [delete]
func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		err = s.sellerService.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}
		web.Success(c, http.StatusNoContent, nil)
	}
}
