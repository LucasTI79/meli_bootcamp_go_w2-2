package handlers

import (
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/web_utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SectionHandler struct {
	sectionService services.SectionService
}

func NewSection(s services.SectionService) *SectionHandler {
	return &SectionHandler{
		sectionService: s,
	}
}

// Method GetAll
// ListSections godoc
//
//	@Summary		List sections
//	@Tags			Sections
//	@Description	getAll sections
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/sections [get]
func (handler *SectionHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		sections, err := handler.sectionService.GetAll(&ctx)
		if err != nil {
			web_utils.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(*sections) == 0 {
			web_utils.Error(c, http.StatusNoContent, "There are no sellers stored")
			return
		}
		web_utils.Success(c, http.StatusOK, sections)
	}
}

// Method Get
// GetSections godoc
//
//	@Summary		Get SectionHandler
//	@Tags			Sections
//	@Description	Get the details of a SectionHandler
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of SectionHandler to be searched"
//	@Success		200	{object}	web_utils.response
//	@Router			/api/v1/sections/{id} [get]
func (handler *SectionHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		sectionResponse, err := handler.sectionService.Get(&ctx, int(id))
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting section %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusOK, sectionResponse)
	}
}

// Method Create
// CreateSections godoc
//
//	@Summary		Create SectionHandler
//	@Tags			Sections
//	@Description	Create section
//	@Accept			json
//	@Produce		json
//	@Param			SectionHandler	body		dtos.CreateSectionRequestDTO	true	"SectionHandler to Create"
//	@Success		201		{object}	web_utils.response
//	@Router			/api/v1/sections [post]
func (handler *SectionHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.CreateSectionRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if req.SectionNumber == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field SectionNumber is required.")
			return
		}

		if req.CurrentTemperature == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field CurrentTemperature is required.")
			return
		}

		if req.MinimumTemperature == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field MinimumTemperature is required.")
			return
		}

		if req.CurrentCapacity == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field CurrentCapacity is required.")
			return
		}

		if req.MaximumCapacity == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field MaximumCapacity is required.")
			return
		}

		if req.MinimumCapacity == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field MinimumCapacity is required.")
			return
		}

		if req.WarehouseID == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field WarehouseID is required.")
			return
		}

		if req.ProductTypeID == 0 {
			web_utils.Error(c, http.StatusUnprocessableEntity, "The field ProductTypeID is required.")
			return
		}

		ctx := c.Request.Context()
		sectionResponse, err := handler.sectionService.Save(&ctx, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			switch err {
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusCreated, sectionResponse)
	}
}

// Method Update
// UpdateSections godoc
//
//	@Summary		Update SectionHandler
//	@Tags			Sections
//	@Description	Update the details of a SectionHandler
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string			true	"ID of SectionHandler to be updated"
//	@Param			Sections	body		dtos.UpdateSectionRequestDTO	true	"Updated SectionHandler details"
//	@Success		200			{object}	web_utils.response
//	@Router			/api/v1/sections/{id} [patch]
func (handler *SectionHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		var req dtos.UpdateSectionRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web_utils.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		sectionResponse, err := handler.sectionService.Update(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID, id)
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			case services.ErrConflict:
				web_utils.Error(c, http.StatusConflict, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error updating section %s", err.Error()))
			}
			return
		}
		web_utils.Success(c, http.StatusOK, sectionResponse)
	}
}

// Method Delete
// DeleteSections godoc
//
//	@Summary		Delete SectionHandler
//	@Tags			Sections
//	@Description	Delete SectionHandler
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a SectionHandler to be excluded"
//	@Success		204	{object}	web_utils.response
//	@Router			/api/v1/sections/{id} [delete]
func (handler *SectionHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web_utils.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = handler.sectionService.Delete(&ctx, int(id))
		if err != nil {
			switch err {
			case services.ErrNotFound:
				web_utils.Error(c, http.StatusNotFound, err.Error())
			default:
				web_utils.Error(c, http.StatusInternalServerError, fmt.Sprintf("error deleting section %s", err.Error()))
			}
			return
		}

		web_utils.Success(c, http.StatusNoContent, nil)

	}

}
