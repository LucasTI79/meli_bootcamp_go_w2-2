package main

import (
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/routes"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/docs"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			MELI Bootcamp API - Sprint 1 - Grupo 2
//	@version		1.0
//	@description	This API Handle MELI - Sprint 1 - Grupo 2
//	@termsOfService	https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//	@contact.name	API Support
//	@contact.url	https://developers.mercadolibre.com.ar/support

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	localhost:8080
func main() {
	// NO MODIFICAR
	db, err := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123@/melisprint")
	if err != nil {
		panic(err)
	}

	eng := gin.Default()

	docs.SwaggerInfo.Host = "localhost:8080"
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := routes.NewRouter(eng, db)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}
}
