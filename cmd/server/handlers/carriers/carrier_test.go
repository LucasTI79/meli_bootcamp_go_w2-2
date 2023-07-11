package carriers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	carrier_handler "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/carriers"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/carriers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/carriers/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		carriersFounds := &[]domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6700,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(carriersFounds, nil)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/carriers", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carriers", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Carrier `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseCarriers := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *carriersFounds, responseCarriers)
	})

	t.Run("empty_database", func(t *testing.T) {
		carriersFounds := &[]domain.Carrier{}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(carriersFounds, nil)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/carriers", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carriers", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		carriersFounds := &[]domain.Carrier{}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(carriersFounds, assert.AnError)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/carriers", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carriers", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})
}

func TestCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{
			ID:          1,
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}
		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("dtos.CarrierRequestDTO")).Return(expectedCarrier, nil)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		bodyReturn, _ := ioutil.ReadAll(res.Body)
		var responseDTO struct {
			Data *domain.Carrier `json:"data"`
		}
		json.Unmarshal(bodyReturn, &responseDTO)
		actualCarrier := responseDTO.Data

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedCarrier, *actualCarrier)
	})
	t.Run("create_fail", func(t *testing.T) {
		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("Create", mock.AnythingOfType("*context.Context")).Return(createCarrierRequestDTO, carriers.ErrUnprocessableEntity)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
	t.Run("create_conflit", func(t *testing.T) {
		expectedCarrier := &domain.Carrier{}
		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("dtos.CarrierRequestDTO")).Return(expectedCarrier, carriers.ErrConflict)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusConflict, res.Code)
	})
	t.Run("create_fail_cid_nil", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_company_name_nil", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:        "CID#1",
			Address:    "corrientes 800",
			Telephone:  "4567-4567",
			LocalityId: 6700,
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_address_nil", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Telephone:   "4567-4567",
			LocalityId:  6700,
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_telephone_nil", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			LocalityId:  6700,
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_locality_id_nil", func(t *testing.T) {

		createCarrierRequestDTO := dtos.CarrierRequestDTO{
			CID:         "CID#1",
			CompanyName: "some name",
			Address:     "corrientes 800",
			Telephone:   "4567-4567",
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/carriers", handler.Create())

		requestBody, _ := json.Marshal(createCarrierRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/carriers", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestGetReportCarriersByLocalities(t *testing.T) {
	t.Run("get_invalid_id", func(t *testing.T) {
		carrierServiceMock := new(mocks.CarrierServiceMock)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/localities/reportCarries?id=:id", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries?id=xyz", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("get_locality_not_found", func(t *testing.T) {

		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetLocalityById", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Locality{}, assert.AnError)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v/localities/reportCarries?id=:id", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries?id=1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("none_carrier_exist", func(t *testing.T) {

		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetLocalityById", mock.AnythingOfType("*context.Context")).Return(domain.Locality{}, nil)
		carrierServiceMock.On("GetCountCarriersByLocalityId", mock.AnythingOfType("*context.Context")).Return(0, assert.AnError)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/localities/reportCarries?id=:id", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries?id=1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
	t.Run("internal_server_error", func(t *testing.T) {
		localityExpected := &domain.Locality{
			ID:           1,
			ProvinceName: "Teste",
			LocalityName: "Teste",
		}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetLocalityById", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(localityExpected, nil)
		carrierServiceMock.On("GetCountCarriersByLocalityId", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil, errors.New("error"))
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/localities/reportCarries?id=:id", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries?id=1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
	t.Run("get_by_locality_id_ok", func(t *testing.T) {
		localityExpected := &domain.Locality{
			ID:           1,
			ProvinceName: "Teste",
			LocalityName: "Teste",
		}

		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetLocalityById", mock.AnythingOfType("*context.Context")).Return(localityExpected, nil)
		carrierServiceMock.On("GetCountCarriersByLocalityId", mock.AnythingOfType("*context.Context")).Return(23, nil)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/localities/reportCarries?id=:id", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries?id=1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data dtos.DataLocalityAndCarrier `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseFound := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, responseFound, responseFound)

	})
	t.Run("get_all_carriers_to_count", func(t *testing.T) {
		responsesFounds := &[]dtos.DataLocalityAndCarrier{
			{
				Id:           1,
				LocalityName: "Teste",
				CountCarrier: 23,
			}, {
				Id:           2,
				LocalityName: "Teste2",
				CountCarrier: 6,
			},
		}
		carrierServiceMock := new(mocks.CarrierServiceMock)
		carrierServiceMock.On("GetCountAndDataByLocality", mock.AnythingOfType("*context.Context")).Return(responsesFounds, nil)
		handler := carrier_handler.NewCarrier(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/localities/reportCarries", handler.GetReportCarriersByLocalities())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/localities/reportCarries", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *[]dtos.DataLocalityAndCarrier `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responses := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *responsesFounds, *responses)

	})
}
