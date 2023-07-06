package productbatcheshandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	dto "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/productbatchesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	productbatches "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatches struct {
	productBatchesService productbatches.IService
}

func NewProductBatches(p productbatches.IService) *ProductBatches {
	return &ProductBatches{
		productBatchesService: p,
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
func (p *ProductBatches) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		ctx := c.Request.Context()
		productBatchesResponse, err := p.productBatchesService.Get(&ctx, id)
		if err != nil {
			if errors.Is(err, productbatches.ErrNotFound) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error getting product batches%s", err.Error()))
			return
		}
		web.Success(c, http.StatusOK, productBatchesResponse)
	}
}

// Method Create
// CreateSections godoc

// @Summary		Create Section
// @Tags			Sections
// @Description	Create section
// @Accept			json
// @Produce		json
// @Param			Section	body		requestCreateSection	true	"Section to Create"
// @Success		201		{object}	web.response
// @Router			/api/v1/sections [post]
func (p *ProductBatches) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateProductBatchesDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		dueDate, err := time.Parse(time.DateOnly, req.DueDate)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		manufacturingDate, err := time.Parse(time.DateOnly, req.ManufacturingDate)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		// manufacturingHour, err := time.Parse(time.TimeOnly, req.ManufacturingHour)
		// if err != nil{
		// 	web.Error(c, http.StatusBadRequest, err.Error())
		// 	return
		// }
		manufacturingHour := time.Now()
		productBatchReq := domain.ProductBatches{
			BatchNumber:        req.BatchNumber,
			CurrentQuantity:    req.CurrentQuantity,
			CurrentTemperature: req.CurrentTemperature,
			DueDate:            dueDate,
			InitialQuantity:    req.InitialQuantity,
			ManufacturingDate:  manufacturingDate,
			ManufacturingHour:  manufacturingHour,
			MinimumTemperature: req.MinimumTemperature,
			ProductID:          req.ProductID,
			SectionID:          req.SectionID,
		}

		ctx := c.Request.Context()
		productBatchRes, err := p.productBatchesService.Save(&ctx, productBatchReq)

		if err != nil {
			if errors.Is(err, productbatches.ErrConflict) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, fmt.Sprintf("error saving product batches %s", err.Error()))
			return
		}
		web.Success(c, http.StatusCreated, productBatchRes)
	}
}
