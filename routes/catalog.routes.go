package routes

import (
	"github.com/gin-gonic/gin"
	apps "dbo-test/controller"
	// "dbo-test/middleware"
)

func CatalogRoutes(g *gin.RouterGroup) {

	g.GET("/", apps.GetCatalog)

	g.POST("/create", apps.CreateCatalog)

	g.PUT("/update/:id", apps.UpdateCatalog)

	g.DELETE("/delete/:id", apps.DeleteCatalog)

	g.GET("/:id", apps.FindCatalog)

}
