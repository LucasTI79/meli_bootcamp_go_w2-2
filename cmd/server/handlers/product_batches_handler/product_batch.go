package productbatcheshandler

import (
	"errors"
	"net/http"
	"strconv"

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

		result, err := p.productBatchesService.SectionProductsReports()
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}

		idParam := c.Request.URL.Query().Get("id")
		if idParam == "" {
			web.Success(c, http.StatusOK, result)
			return
		}
		sectionID, err := strconv.Atoi(idParam)
		if err != nil{
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		result, err = p.productBatchesService.SectionProductsReportsBySection(sectionID)
		if err != nil{
			if errors.Is(err, productbatches.ErrNotFoundSection){
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
		}
		web.Success(c, http.StatusOK, result)
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
		productBatchReq := domain.ProductBatches{
			BatchNumber:        req.BatchNumber,
			CurrentQuantity:    req.CurrentQuantity,
			CurrentTemperature: req.CurrentTemperature,
			DueDate:            req.DueDate,
			InitialQuantity:    req.InitialQuantity,
			ManufacturingDate:  req.ManufacturingDate,
			ManufacturingHour:  req.ManufacturingHour,
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
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, productBatchRes)
	}
}
