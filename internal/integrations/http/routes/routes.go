package routes

import (
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/database/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/handlers"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

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
	repo := repositories.NewSellerRepository(r.db)
	service := services.NewSellerService(repo)
	handler := handlers.NewSeller(service)
	r.rg.POST("/sellers", handler.Create())
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.PATCH("/sellers/:id", handler.Update())
	r.rg.DELETE("/sellers/:id", handler.Delete())
}

func (r *router) buildProductRoutes() {
	repo := repositories.NewProductRepository(r.db)
	service := services.NewProductService(repo)
	handler := handlers.NewProductHandler(service)
	r.rg.POST("/products", handler.Create())
	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:id", handler.Get())
	r.rg.DELETE("/products/:id", handler.Delete())
	r.rg.PATCH("/products/:id", handler.Update())
}

func (r *router) buildSectionRoutes() {
	repo := repositories.NewSectionRepository(r.db)
	service := services.NewSectionService(repo)
	handler := handlers.NewSection(service)
	r.rg.POST("/sections", handler.Create())
	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.DELETE("/sections/:id", handler.Delete())
	r.rg.PATCH("/sections/:id", handler.Update())
}

func (r *router) buildWarehouseRoutes() {
	repository := repositories.NewWarehouseRepository(r.db)
	service := services.NewWarehouseService(repository)
	handler := handlers.NewWarehouse(service)
	r.rg.POST("/warehouses", handler.Create())
	r.rg.GET("/warehouses", handler.GetAll())
	r.rg.GET("/warehouses/:id", handler.Get())
	r.rg.PATCH("/warehouses/:id", handler.Update())
	r.rg.DELETE("/warehouses/:id", handler.Delete())
}

func (r *router) buildEmployeeRoutes() {
	repo := repositories.NewEmployeeRepository(r.db)
	service := services.NewEmployeeService(repo)
	handler := handlers.NewEmployeeHandler(service)

	r.rg.POST("/employees", handler.Save())
	r.rg.GET("/employees", handler.GetAll())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.PATCH("/employees/:id", handler.Update())
	r.rg.DELETE("/employees/:id", handler.Delete())

}

func (r *router) buildBuyerRoutes() {
	buyerRepository := repositories.NewBuyerRepository(r.db)
	buyerService := services.NewBuyerService(buyerRepository)
	buyerHandler := handlers.NewBuyerHandler(buyerService)

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

func (r *router) buildLocalityRoutes() {
	localityRepository := repositories.NewLocalityRepository(r.db)
	localityService := services.NewLocalityService(localityRepository)
	localityHandler := handlers.NewLocalityHandler(localityService)

	localityRoutes := r.rg.Group("/localities/")
	localityRoutes.GET(":id", localityHandler.Get())
	localityRoutes.GET("", localityHandler.GetAll())
	localityRoutes.POST("", localityHandler.Create())
	localityRoutes.PATCH(":id", localityHandler.Update())
	localityRoutes.DELETE(":id", localityHandler.Delete())
	localityRoutes.GET(":id/reportSellers", localityHandler.CountSellers())
}
