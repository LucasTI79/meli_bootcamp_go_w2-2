package localities

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"

	"github.com/gin-gonic/gin"
)

type LocalityHandler struct {
	localityService services.LocalityService
}

func NewLocalityHandler(localityService services.LocalityService) *LocalityHandler {
	return &LocalityHandler{
		localityService,
	}
}

// Get is the handler to search for a locality and return their details.
//
//	@Summary		Get Locality
//	@Tags			Localities
//	@Description	Get the details of a Locality
//	@Produce		json
//	@Param			id	path		string	true	"ID of Locality to be searched"
//	@Success		200	{object}	entities.Locality
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/localities/{id} [get]
func (handler *LocalityHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		if locality, err := handler.localityService.Get(&ctx, id); err != nil {
			switch err {
			case services.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		} else {
			web.Response(c, http.StatusOK, locality)
			return
		}

	}
}

// GetAll is the handler to search for all localities.
//
//	@Summary		List Localities
//	@Tags			Localities
//	@Description	Get the details of all localities on the database.
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web.response{data=[]entities.Locality}
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/localities [get]
func (handler *LocalityHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()
		if localities, err := handler.localityService.GetAll(&ctx); err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		} else {
			if len(localities) == 0 {
				web.Success(c, http.StatusNoContent, localities)
				return
			}
			web.Success(c, http.StatusOK, localities)
			return
		}
	}
}

// Create is the handler to create a locality.
//
//	@Summary		Create Locality
//	@Tags			Localities
//	@Description	Save a locality on the database.
//	@Accept			json
//	@Produce		json
//	@Param			Seller	body		entities.Locality	true	"Locality to Create"
//	@Success		201		{object}	entities.Locality
//	@Failure		422		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/localities [post]
func (handler *LocalityHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var createLocalityRequest entities.Locality
		if err := c.ShouldBindJSON(&createLocalityRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := c.Request.Context()
		if createdLocality, err := handler.localityService.Create(&ctx, createLocalityRequest); err != nil {
			switch err {
			case services.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
				return
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			web.Response(c, http.StatusCreated, createdLocality)
			return
		}
	}
}

// Update is the handler to update a locality details.
//
//	@Summary		Update Locality
//	@Tags			Localities
//	@Description	Update the details of a Locality
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of Locality to be updated"
//	@Param			Locality	body		dtos.UpdateLocalityRequestDTO	true	"Updated Locality details"
//	@Success		200		{object}	entities.Locality
//	@Failure		400		{object}	web.errorResponse
//	@Failure		404		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/localities/{id} [patch]
func (handler *LocalityHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		var updateLocalityRequest dtos.UpdateLocalityRequestDTO
		if err := c.ShouldBindJSON(&updateLocalityRequest); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		ctx := c.Request.Context()
		if updatedLocality, err := handler.localityService.Update(&ctx, id, updateLocalityRequest); err != nil {
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
			web.Response(c, http.StatusOK, updatedLocality)
			return
		}
	}
}

// Delete is the handler to delete a locality.
//
//	@Summary		Delete Locality
//	@Tags			Localities
//	@Description	Delete Localities
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"ID of a Locality to be excluded"
//	@Success		204
//	@Failure		400	{object}	web.errorResponse
//	@Failure		404	{object}	web.errorResponse
//	@Failure		500	{object}	web.errorResponse
//	@Router			/api/v1/localities/{id} [delete]
func (handler *LocalityHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		if err := handler.localityService.Delete(&ctx, id); err != nil {
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

// CountSellers is the handler search for a locality and return the number of sellers
//
//	@Summary		CountSellers
//	@Tags			Localities
//	@Description	search for a locality and return the number of sellers.
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"ID of Locality to be searched"
//	@Success		200		{object}	web.response{data=dtos.GetNumberOfSellersResponseDTO}
//	@Failure		400		{object}	web.errorResponse
//	@Failure		404		{object}	web.errorResponse
//	@Failure		422		{object}	web.errorResponse
//	@Failure		500		{object}	web.errorResponse
//	@Router			/api/v1/localities/count [get]
func (handler *LocalityHandler) CountSellers() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := getIdFromUri(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()

		locality, err := handler.localityService.Get(&ctx, id)
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, err.Error())
			}
			return
		}

		count, err := handler.localityService.CountSellers(&ctx, id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		response := dtos.GetNumberOfSellersResponseDTO{
			LocalityID:   locality.ID,
			LocalityName: locality.LocalityName,
			SellersCount: count,
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
