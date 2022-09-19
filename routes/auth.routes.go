package routes

import (
	"github.com/gin-gonic/gin"
	apps "dbo-test/controller"
	// "golang-test/middleware"
)

func AuthRoutes(g *gin.RouterGroup) {

	g.POST("/login", apps.Login)

	// auth.Post("/logout", middleware.Protected(), apps.Logout)
  //
	// auth.Post("/forgotPassword", apps.ForgotPassword)
  //
	// auth.Post("/resetPassword", apps.ResetPassword)
  //
	// auth.Post("/resetPasswordByPass", apps.ResetPasswordByPass)
  //
	// auth.Post("/updateProfile", middleware.Protected(), apps.UpdateProfile)
  //
	// auth.Post("/register", apps.CreateUsers)
  //
	// auth.Get("/confirmationEmail/:id", apps.ConfirmationEmail)
  //
	// auth.Post("/profile", middleware.Protected(), apps.Profile)
  //
	// auth.Post("/changeRole", middleware.Protected(), apps.ChangeRole)

}
