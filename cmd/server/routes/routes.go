package routes

import (
	"database/sql"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(eng *gin.Engine, db *sql.DB) Router {
	return &router{eng: eng, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildSellerRoutes()
	r.buildProductRoutes()
	r.buildSectionRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	// Example
	// repo := seller.NewRepository(r.db)
	// service := seller.NewService(repo)
	// handler := handler.NewSeller(service)
	// r.r.GET("/seller", handler.GetAll)
}

func (r *router) buildProductRoutes() {}

func (r *router) buildSectionRoutes() {}

func (r *router) buildWarehouseRoutes() {}

func (r *router) buildEmployeeRoutes() {}

func (r *router) buildBuyerRoutes() {
	buyerRepository := buyer.NewRepository(r.db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := handler.NewBuyer(buyerService)

	// Create custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(dtos.UpdateBuyerRequestValidation, dtos.UpdateBuyerRequestDTO{})
	}

	buyerRoutes := r.rg.Group("/buyers/")
	buyerRoutes.GET(":id", buyerHandler.Get())
	buyerRoutes.GET("", buyerHandler.GetAll())
	buyerRoutes.POST("", buyerHandler.Create())
	buyerRoutes.PATCH(":id", buyerHandler.Update())
	buyerRoutes.DELETE(":id", buyerHandler.Delete())
}
