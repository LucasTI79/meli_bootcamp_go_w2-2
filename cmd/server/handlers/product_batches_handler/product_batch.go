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

var (
	ErrInvalidID              = errors.New("Invalid ID").Error()
	ErrSectionProductsReports = errors.New("error returning product reports by Section").Error()
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
// GetSectionProductsReports godoc
//
//	@Summary		SectionProductsReports / SectionProductsReportsBySection
//	@Tags			ProductBatch
//	@Description	Get all products by section / Section Products Reports By Section
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	false "ID of a Products Reports to search"
//	@Success		200	{object}	[]domain.ProductBySection
//	@Router			/api/v1/sections/reportProducts/{id} [get]
func (p *ProductBatches) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		if idParam == "" {
			result, err := p.productBatchesService.SectionProductsReports()
			if err != nil {
				web.Error(c, http.StatusInternalServerError, ErrSectionProductsReports)
				return
			}
			web.Success(c, http.StatusOK, result)
			return
		}

		sectionID, err := strconv.Atoi(idParam)
		if err != nil {
			web.Error(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		sectionProductsReportsBySection, err := p.productBatchesService.SectionProductsReportsBySection(int(sectionID))
		if err != nil {
			if errors.Is(err, productbatches.ErrNotFoundSection) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, sectionProductsReportsBySection)
	}
}

// Method Create
// CreateProductBatch godoc

// @Summary		Create ProductBatch
// @Tags			ProductBatch
// @Description	Create ProductBatch
// @Accept			json
// @Produce		json
// @Param			ProductBatch	body		productbatchesdto.CreateProductBatchesDTO	true	"ProductBatch to Create"
// @Success		201		{object}	web.response
// @Router			/api/v1/productBatches [post]
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
