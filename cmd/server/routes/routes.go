package routes

import (
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler/sections"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/buyers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/products"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/sellers"
	warehouse2 "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/warehouses"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
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
	repo := seller.NewRepository(r.db)
	service := seller.NewService(repo)
	handler := sellers.NewSeller(service)
	r.rg.POST("/sellers", handler.Create())
	r.rg.GET("/sellers", handler.GetAll())
	r.rg.GET("/sellers/:id", handler.Get())
	r.rg.PATCH("/sellers/:id", handler.Update())
	r.rg.DELETE("/sellers/:id", handler.Delete())
}

func (r *router) buildProductRoutes() {
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := products.NewProduct(service)
	r.rg.POST("/products", handler.Create())
	r.rg.GET("/products", handler.GetAll())
	r.rg.GET("/products/:id", handler.Get())
	r.rg.DELETE("/products/:id", handler.Delete())
	r.rg.PATCH("/products/:id", handler.Update())
}

func (r *router) buildSectionRoutes() {
	repo := section.NewRepository(r.db)
	service := section.NewService(repo)
	handler := sections.NewSection(service)
	r.rg.POST("/sections", handler.Create())
	r.rg.GET("/sections", handler.GetAll())
	r.rg.GET("/sections/:id", handler.Get())
	r.rg.DELETE("/sections/:id", handler.Delete())
	r.rg.PATCH("/sections/:id", handler.Update())
}

func (r *router) buildWarehouseRoutes() {
	repository := warehouse.NewRepository(r.db)
	service := warehouse.NewService(repository)
	handler := warehouse2.NewWarehouse(service)
	r.rg.POST("/warehouses", handler.Create())
	r.rg.GET("/warehouses", handler.GetAll())
	r.rg.GET("/warehouses/:id", handler.Get())
	r.rg.PATCH("/warehouses/:id", handler.Update())
	r.rg.DELETE("/warehouses/:id", handler.Delete())
}

func (r *router) buildEmployeeRoutes() {
	repo := employee.NewRepository(r.db)
	service := employee.NewService(repo)
	handler := handlers.NewEmployee(service)

	r.rg.POST("/employees", handler.Save())
	r.rg.GET("/employees", handler.GetAll())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.PATCH("/employees/:id", handler.Update())
	r.rg.DELETE("/employees/:id", handler.Delete())

}

func (r *router) buildBuyerRoutes() {
	buyerRepository := buyer.NewRepository(r.db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := buyers.NewBuyerHandler(buyerService)

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
