package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)
func TestGet(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		sectionExpected := &domain.Section{
			ID:                 1,
			SectionNumber:      65473,
			CurrentTemperature: 15,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    12,
			MaximumCapacity:    20,
			WarehouseID:        1234,
			ProductTypeID:      874893,
		}
		//Configurar o mock do service
		sectionServiceMock := new(mocks.SectionServiceMock)
		sectionServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(sectionExpected, nil)
		handler := handler.NewSection(sectionServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sections/:id", handler.Get())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(response, request)

		//Parsear response
		body, _ := ioutil.ReadAll(response.Body)

		var responseSectionDTO struct {
			Data domain.Section `json:"data"`
		}
		json.Unmarshal(body, &responseSectionDTO)

		//sectionActual := responseSectionDTO.Data

		//validar resultado
		// assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *sectionExpected, responseSectionDTO)

	})
}

func getTestSections() []domain.Section{
	return []domain.Section{
		{
			ID:                 1,
			SectionNumber:      65473,
			CurrentTemperature: 15,
			MinimumTemperature: 5,
			CurrentCapacity:    10,
			MinimumCapacity:    12,
			MaximumCapacity:    20,
			WarehouseID:        1234,
			ProductTypeID:      874893,
		},
		{
			ID:                 2,
			SectionNumber:      4653,
			CurrentTemperature: 20,
			MinimumTemperature: 50,
			CurrentCapacity:    1000,
			MinimumCapacity:    120,
			MaximumCapacity:    200,
			WarehouseID:        9878,
			ProductTypeID:      87489223,
		},
	}
}