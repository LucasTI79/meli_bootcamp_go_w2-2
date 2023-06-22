package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type requestCreateSection struct {
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type requestUpdateSection struct {
	SectionNumber      *int `json:"section_number"`
	CurrentTemperature *int `json:"current_temperature"`
	MinimumTemperature *int `json:"minimum_temperature"`
	CurrentCapacity    *int `json:"current_capacity"`
	MinimumCapacity    *int `json:"minimum_capacity"`
	MaximumCapacity    *int `json:"maximum_capacity"`
	WarehouseID        *int `json:"warehouse_id"`
	ProductTypeID      *int `json:"product_type_id"`
}

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
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
//	@Success		200	{object}	web.response
//	@Router			/api/v1/sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c)
		if err != nil {
			web.ErrorString(c, http.StatusInternalServerError, err.ErrorString())
			return
		}
		web.Success(c, http.StatusOK, sections)
	}
}

// Method Get
// GetSections godoc
//
//	@Summary		Get Section
//	@Tags			Sections
//	@Description	Get the details of a Section
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of Section to be searched"
//	@Success		200	{object}	web.response
//	@Router			/api/v1/sections/{id} [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, errors.ErrInvalidId)
			return
		}
		sectionResponse, err := s.sectionService.Get(c, id)
		if err != nil {
			if errors.Is(err, domain.ErrNoRows) {
				web.Error(c, http.StatusNotFound, errors.ErrNotFound)
				return
			}

			web.Error(c, http.StatusInternalServerError, errors.ErrGettingSectionById)
			return
		}
		web.Success(c, http.StatusOK, sectionResponse)
	}
}

// Method Create
// CreateSections godoc
//
//	@Summary		Create Section
//	@Tags			Sections
//	@Description	Create section
//	@Accept			json
//	@Produce		json
//	@Param			Section	body		requestCreateSection	true	"Section to Create"
//	@Success		201		{object}	web.response
//	@Router			/api/v1/sections [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestCreateSection
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, errors.ErrTryAgain)
			return
		}

		if req.SectionNumber == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field SectionNumber is required.")
			return
		}

		if req.CurrentTemperature == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field CurrentTemperature is required.")
			return
		}

		if req.MinimumTemperature == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field MinimumTemperature is required.")
			return
		}

		if req.CurrentCapacity == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field CurrentCapacity is required.")
			return
		}

		if req.MaximumCapacity == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field MaximumCapacity is required.")
			return
		}

		if req.MinimumCapacity == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field MinimumCapacity is required.")
			return
		}

		if req.WarehouseID == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field WarehouseID is required.")
			return
		}

		if req.ProductTypeID == 0 {
			web.ErrorString(c, http.StatusUnprocessableEntity, "The field ProductTypeID is required.")
			return
		}

		sectionResponse, err := s.sectionService.Save(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			switch err {
			case section.ErrConflict:
				web.ErrorString(c, http.StatusConflict, err)
			default:
				web.ErrorString(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.ErrorString()))
			}
			return
		}
		web.Success(c, http.StatusCreated, sectionResponse)
	}
}

// Method Update
// UpdateSections godoc
//
//	@Summary		Update Section
//	@Tags			Sections
//	@Description	Update the details of a Section
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string			true	"ID of Section to be updated"
//	@Param			Sections	body		requestUpdateSection	true	"Updated Section details"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, errors.ErrInvalidId)
			return
		}
		var req requestUpdateSection
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, errors.ErrTryAgain)
			return
		}
		sectionResponse, err := s.sectionService.Update(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID, id)
		if err != nil {
			switch err {
			case section.ErrNotFound:
				web.Error(c, http.StatusNotFound, err)
			case section.ErrConflict:
				web.Error(c, http.StatusConflict, err)
			default:
				web.ErrorString(c, http.StatusInternalServerError, fmt.Sprintf("error updating section %s", err.Error()))
			}
			return
		}
		web.Success(c, http.StatusOK, sectionResponse)
	}
}

// Method Delete
// DeleteSections godoc
//
//	@Summary		Delete Section
//	@Tags			Sections
//	@Description	Delete Section
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"ID of a Section to be excluded"
//	@Success		204	{object}	web.response
//	@Router			/api/v1/sections/{id} [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, errors.ErrInvalidId)
			return
		}
		err = s.sectionService.Delete(c, int(id))
		if err != nil {
			if errors.Is(err, section.ErrNotFound) {
				web.Error(c, http.StatusNotFound, errors.ErrNotFound)
				return
			}
			web.Error(c, http.StatusInternalServerError, errors.ErrDeletingSection)
			return
		}

		web.Success(c, http.StatusNoContent, nil)

	}

}
