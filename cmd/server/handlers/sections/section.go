package sections

import (
	"fmt"
	"net/http"
	"strconv"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sections"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

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
		ctx := c.Request.Context()
		sections, err := s.sectionService.GetAll(&ctx)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(*sections) == 0 {
			web.Error(c, http.StatusNoContent, "There are no sellers stored")
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
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := c.Request.Context()
		sectionResponse, err := s.sectionService.Get(&ctx, int(id))
		if err != nil {
			switch err {
			case section.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting section %s", err.Error()))
			}
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
//	@Param			Section	body		sections.CreateSectionRequestDTO	true	"Section to Create"
//	@Success		201		{object}	web.response
//	@Router			/api/v1/sections [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.CreateSectionRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if req.SectionNumber == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field SectionNumber is required.")
			return
		}

		if req.CurrentTemperature == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field CurrentTemperature is required.")
			return
		}

		if req.MinimumTemperature == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field MinimumTemperature is required.")
			return
		}

		if req.CurrentCapacity == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field CurrentCapacity is required.")
			return
		}

		if req.MaximumCapacity == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field MaximumCapacity is required.")
			return
		}

		if req.MinimumCapacity == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field MinimumCapacity is required.")
			return
		}

		if req.WarehouseID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field WarehouseID is required.")
			return
		}

		if req.ProductTypeID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "The field ProductTypeID is required.")
			return
		}

		ctx := c.Request.Context()
		sectionResponse, err := s.sectionService.Save(&ctx, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			switch err {
			case section.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
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
//	@Param			Sections	body		sections.UpdateSectionRequestDTO	true	"Updated Section details"
//	@Success		200			{object}	web.response
//	@Router			/api/v1/sections/{id} [patch]
func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		var req dtos.UpdateSectionRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		sectionResponse, err := s.sectionService.Update(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID, id)
		if err != nil {
			switch err {
			case section.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			case section.ErrConflict:
				web.Error(c, http.StatusConflict, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error updating section %s", err.Error()))
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
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}
		ctx := c.Request.Context()
		err = s.sectionService.Delete(&ctx, int(id))
		if err != nil {
			switch err {
			case section.ErrNotFound:
				web.Error(c, http.StatusNotFound, err.Error())
			default:
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error deleting section %s", err.Error()))
			}
			return
		}

		web.Success(c, http.StatusNoContent, nil)

	}

}
