package routes

import (
	"github.com/gin-gonic/gin"
	apps "dbo-test/controller"
	// "dbo-test/middleware"
)

func UsersRoutes(g *gin.RouterGroup) {

	g.GET("/", apps.GetUsers)

	g.POST("/create", apps.CreateUsers)

	g.PUT("/update/:id", apps.UpdateUsers)

	g.DELETE("/delete/:id", apps.DeleteUsers)

	g.GET("/:id", apps.FindUsers)

}
