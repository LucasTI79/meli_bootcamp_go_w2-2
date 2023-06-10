package handler

import (
	"fmt"
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sections, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, sections)
	}
}

func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
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

		section, err := s.sectionService.Save(c, req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity,
			req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			if err.Error() == "section with Section Number already exists" {
				web.Error(c, http.StatusConflict, err.Error())
				return
			} else {
				web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving request %s", err.Error()))
				return
			}

		}
		web.Success(c, http.StatusCreated, section)
	}
}

func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
