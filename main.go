package main

import (
  "fmt"
  // "net/http"
  "github.com/gin-gonic/gin"
  "dbo-test/database"
  "dbo-test/migrations"
  "dbo-test/routes"
  "dbo-test/middleware"
)

// var (
// 	DBConn   = database.DBConn
//   validate *validator.Validate
// )

func main() {

  r := gin.Default()

  fmt.Println("Starting...")

  database.InitDB()

  migrations.Migrate()

  setupRoutes()

  r.Run()

}

func setupRoutes() {

  router := gin.Default()

  v1 := router.Group("/v1")
  {
   v1.Use(middleware.Auth())
   routes.UsersRoutes(v1.Group("/user"))
   routes.CatalogRoutes(v1.Group("/catalog"))
  }

  v2 := router.Group("/v1")
  {
   routes.AuthRoutes(v2.Group("/auth"))
  }


  // routes.UsersRoutes()

  // routes.AuthRoutes()

  router.Run(":8080")

}
