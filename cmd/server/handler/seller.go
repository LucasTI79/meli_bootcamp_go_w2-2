package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "Failed to get sellers: %s", err.Error())
			return
		}

		if len(sellers) == 0 {
			web.Error(c, http.StatusNotFound, "There are no sellers stored: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		if req.CID <= 0.0 {
			web.Error(c, http.StatusBadRequest, "Field CID is required: %s", "")
			return
		}

		if req.CompanyName == "" {
			web.Error(c, http.StatusBadRequest, "Field CompanyName is required: %s", "")
			return
		}

		if req.Address == "" {
			web.Error(c, http.StatusBadRequest, "Field Address is required: %s", "")
			return
		}

		if req.Telephone == "" {
			web.Error(c, http.StatusBadRequest, "Field Telephone is required: %s", "")
			return
		}

		seller, err := s.sellerService.Save(c, req.CID, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Error to save request: %s", err.Error())
			return
		}

		web.Success(c, http.StatusCreated, seller)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid Seller ID: %s", err.Error())
			return
		}

		seller, err := s.sellerService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Seller not found: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, seller)
	}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		var updatedSeller domain.Seller

		if err := c.Bind(&updatedSeller); err != nil {
			web.Error(c, http.StatusBadRequest, "Error to read request: %s", err.Error())
			return
		}

		sellerUpdated, err := s.sellerService.Update(c, int(id), updatedSeller)
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to update: %s", err.Error())
			return
		}

		web.Success(c, http.StatusOK, sellerUpdated)
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Invalid ID: %s", err.Error())
			return
		}

		err = s.sellerService.Delete(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "Error to delete: %s", err.Error())
			return
		}
	}
}
