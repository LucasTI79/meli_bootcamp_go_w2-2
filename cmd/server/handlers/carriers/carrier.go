package carriers

import (
	"errors"
	"net/http"

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
