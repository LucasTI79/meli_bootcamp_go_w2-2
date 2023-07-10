package carriers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/carriers"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carrier struct {
	carrierService carriers.Service
}

func NewCarrier(s carriers.Service) *Carrier {
	return &Carrier{
		carrierService: s,
	}
}

// GetAllCarriers godoc
//
//	@Summary		List carriers
//	@Tags			Carriers
//	@Description	get carriers
//	@Produce		json
//	@Success		200	{object}	[]domain.Carrier
//	@Router			/api/v1/carriers [get]
func (carrier *Carrier) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		carriers, err := carrier.carrierService.GetAll(&ctx)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get carriers: %s", err.Error())
			return
		}

		if len(*carriers) == 0 {
			web.Error(c, http.StatusNoContent, "There are no carriers stored")
			return
		}

		web.Success(c, http.StatusOK, *carriers)
	}
}

// createCarriers godoc
//
//	@Summary		Create carriers
//	@Tags			Carriers
//	@Description	Create carriers
//	@Accept			json
//	@Produce		json
//	@Param			Carrier	body		dtos.CarrierRequestDTO	true	"carrier to create"
//	@Success		200			{object}	domain.Carrier
//	@Router			/api/v1/carriers [post]
func (carrier *Carrier) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dtos.CarrierRequestDTO
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, "JSON format may be wrong")
			return
		}
		ctx := c.Request.Context()
		if err := CarrierFullRequestValidator(c, req); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		result, e := carrier.carrierService.Create(&ctx, req)

		if e != nil {
			web.Error(c, http.StatusConflict, e.Error())
			return
		}
		web.Success(c, http.StatusCreated, result)
	}
}

func CarrierFullRequestValidator(c *gin.Context, req dtos.CarrierRequestDTO) error {
	if req.CID == "" {
		return errors.New("field cid is required")
	}
	if req.CompanyName == "" {
		return errors.New("field company_name is required")
	}
	if req.Address == "" {
		return errors.New("field address is required")
	}
	if req.Telephone == "" {
		return errors.New("field telephone is required")
	}
	if req.LocalityId == 0 {
		return errors.New("field locality_id is required")
	}

	return nil
}

func (carrier *Carrier) GetReportCarriersByLocalities() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		type Response struct {
			locality_id   int
			locality_name string
			carries_count int
		}

		if c.Query("id") == "" {
			data, err := carrier.carrierService.GetCountAndDataByLocality(&ctx)
			if err != nil {
				web.Error(c, http.StatusNoContent, "data not found")
				return
			}
			web.Success(c, http.StatusOK, data)

		} else {
			localityId, e := strconv.Atoi(c.Query("id"))
			if e != nil {
				web.Error(c, http.StatusBadRequest, "parameter id must be a integer")
				return
			}
			locality, err := carrier.carrierService.GetLocalityById(&ctx, localityId)
			if err != nil {
				web.Error(c, http.StatusNotFound, "locality not found")
				return
			}
			count, err := carrier.carrierService.GetCountCarriersByLocalityId(&ctx, localityId)
			if err != nil {
				web.Error(c, http.StatusNotFound, "none carrier exists with this location_id")
				return
			}
			response := &Response{
				locality_id:   localityId,
				locality_name: locality.LocalityName,
				carries_count: *count,
			}
			web.Success(c, http.StatusOK, *response)
		}
	}
}
