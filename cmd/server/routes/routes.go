package routes

import (
	"database/sql"

	handlers "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/localities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/purchaseOrder"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/buyers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/carriers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/inbound_orders"
	productbatcheshandler "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/product_batches_handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/products"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/productsRecords"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/sections"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/sellers"
	warehouse2 "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/warehouses"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	carrier "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/carriers"
	inbound_order "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord"
	prodBatches "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/employees"
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
	r.buildProductBatchesRoutes()
	r.buildWarehouseRoutes()
	r.buildEmployeeRoutes()
	r.buildBuyerRoutes()
	r.buildInboundOrdersRoutes()
	r.buildCarriersRoutes()
	r.buildProductRecordsRoutes()
	r.buildPurchaseOrderRoutes()
	r.buildLocalityRoutes()

}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildSellerRoutes() {
	repo := seller.NewSellerRepository(r.db)
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
func (r *router) buildProductBatchesRoutes() {
	repo := prodBatches.NewRepository(r.db)
	productRepo := product.NewRepository(r.db)
	sectionRepo := section.NewRepository(r.db)
	service := prodBatches.NewService(repo, productRepo, sectionRepo)
	handler := productbatcheshandler.NewProductBatches(service)
	r.rg.POST("/productBatches", handler.Create())
	r.rg.GET("sections/reportProducts/:id", handler.Get())
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
	handler := employees.NewEmployee(service)

	r.rg.POST("/employees", handler.Save())
	r.rg.GET("/employees", handler.GetAll())
	r.rg.GET("/employees/:id", handler.Get())
	r.rg.PATCH("/employees/:id", handler.Update())
	r.rg.DELETE("/employees/:id", handler.Delete())

}

func (r *router) buildBuyerRoutes() {
	buyerRepository := buyer.NewBuyerRepository(r.db)
	buyerService := buyer.NewService(buyerRepository)

	purchaseOrdersRepository := purchaseOrder.NewPurchaseOrderRepository(r.db)
	purchaseOrderService := purchaseOrder.NewPurchaseOrderService(purchaseOrdersRepository, buyerRepository)

	buyerHandler := buyers.NewBuyerHandler(buyerService, purchaseOrderService)

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
	buyerRoutes.GET(":id/report-purchase-orders", buyerHandler.CountPurchaseOrders())
}

func (r *router) buildLocalityRoutes() {
	localityRepository := locality.NewLocalityRepository(r.db)
	localityService := locality.NewLocalityService(localityRepository)
	localityHandler := handlers.NewLocalityHandler(localityService)

	localityRoutes := r.rg.Group("/localities/")
	localityRoutes.GET(":id", localityHandler.Get())
	localityRoutes.GET("", localityHandler.GetAll())
	localityRoutes.POST("", localityHandler.Create())
	localityRoutes.PATCH(":id", localityHandler.Update())
	localityRoutes.DELETE(":id", localityHandler.Delete())
	localityRoutes.GET(":id/reportSellers", localityHandler.CountSellers())
}

func (r *router) buildPurchaseOrderRoutes() {
	purchaseOrderRepository := purchaseOrder.NewPurchaseOrderRepository(r.db)
	buyerRepository := buyer.NewBuyerRepository(r.db)
	purchaseOrderService := purchaseOrder.NewPurchaseOrderService(purchaseOrderRepository, buyerRepository)
	purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderService)

	purchaseOrderRoutes := r.rg.Group("/purchase-orders/")
	purchaseOrderRoutes.GET(":id", purchaseOrderHandler.Get())
	purchaseOrderRoutes.GET("", purchaseOrderHandler.GetAll())
	purchaseOrderRoutes.POST("", purchaseOrderHandler.Create())
	purchaseOrderRoutes.PATCH(":id", purchaseOrderHandler.Update())
	purchaseOrderRoutes.DELETE(":id", purchaseOrderHandler.Delete())
}

func (r *router) buildProductRecordsRoutes() {
	productRecordRepository := productRecord.NewRepository(r.db)
	productRecordService := productRecord.NewService(productRecordRepository)
	productRepository := product.NewRepository(r.db)
	productService := product.NewService(productRepository)
	handler := productsRecords.NewProductRecord(productRecordService, productService)

	r.rg.POST("/productRecords", handler.Create())
	r.rg.GET("/productRecords", handler.GetAll())
	r.rg.GET("/productRecords/:id", handler.Get())
	r.rg.DELETE("/productRecords/:id", handler.Delete())
	r.rg.PATCH("/productRecords/:id", handler.Update())
	r.rg.GET("/products/reportRecords", handler.NumberRecords())

}

func (r *router) buildCarriersRoutes() {
	repo := carrier.NewRepository(r.db)
	service := carrier.NewService(repo)
	handler := carriers.NewCarrier(service)
	r.rg.POST("/carriers", handler.Create())
	r.rg.GET("/carriers", handler.GetAll())
	r.rg.GET("/localities/reportCarries", handler.GetReportCarriersByLocalities())
}

func (r *router) buildInboundOrdersRoutes() {
	repo := inbound_order.NewRepository(r.db)
	repoEmployee := employee.NewRepository(r.db)
	service := inbound_order.NewService(repo, repoEmployee)
	handler := inbound_orders.NewInboundOrders(service)

	r.rg.POST("/inbound-orders", handler.Save())
	r.rg.GET("/inbound-orders", handler.GetAll())
	r.rg.GET("/inbound-orders/:id", handler.Get())
	r.rg.PATCH("/inbound-orders/:id", handler.Update())
	r.rg.DELETE("/inbound-orders/:id", handler.Delete())
	r.rg.GET("/reportInboundOrders", handler.CountInboundOrders())
	r.rg.GET("/reportInboundOrders/:id", handler.CountInboundOrdersByID())

}
