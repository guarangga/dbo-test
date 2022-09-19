package apps

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"dbo-test/database"
	"dbo-test/utils"
	"dbo-test/models"
	// "github.com/dgrijalva/jwt-go"
	// "golang-test/rules"
	// "fmt"
	"time"
	"net/http"
)

func Login(c *gin.Context) {

	var user models.User

	var email, password string

	form, e := c.MultipartForm()

	if e == nil {

		for key, val := range form.Value {

			if key == "email" {

				email = val[0]

			}

			if key == "password" {

				password = val[0]

			}

		}

	}

	tx := database.DBConn.Model(&models.User{}).Begin()

	if err := tx.Where(&models.User{Email: email}).First(&user).Error; err != nil {

		tx.Rollback()

		utils.SendResponse(nil, "Invalid email or password", false, http.StatusUnauthorized, c)

		// return

	}

	if !utils.CheckPasswordHash(user.Password, password) {

		tx.Rollback()

		utils.SendResponse(nil, "Login email or password is invalid", false, http.StatusUnauthorized, c)

		// return

	}

	token, error := utils.GenerateJWT(user.Email)

	if error != nil {

		tx.Rollback()

		utils.SendResponse(nil, "Failed to generate token", false, http.StatusUnauthorized, c)

		// return

	}

	// loginData, error := utils.GetLoginData(token)
	//
	// if error != nil {
	//
	// 	tx.Rollback()
	//
	// 	utils.SendResponse(c.Error(error), "Failed to login", false, http.StatusUnauthorized, c)
	//
	// 	return
	//
	// }
	//
	// loginData["token_type"] = "Bearer"
	// loginData["access_token"] = token

	user.Token = token
	user.LoginAt = time.Now()
	user.LogoutAt = time.Time{}
	user.Ip = utils.GetIPAddress(c)
	// user.MacAddress = utils.GetMacAdress()
	// user.UserAgent = utils.GetUserAgent(c)

	if err := tx.Where(&models.User{Email: email}).Save(&user).Error; err != nil {

		tx.Rollback()

		utils.SendResponse(nil, "Failed to login", false, http.StatusUnauthorized, c)

		// return

	}

	tx.Commit()

	// return utils.SendResponse(loginData, "Success to login", true, fiber.StatusOK, c)

	// result = gin.H{
	// 	"data": loginData,
	// 	"message": "Success to login",
	// 	"status": http.StatusOK,
	// }
	//
	// c.JSON(http.StatusOK, result)

	if err := tx.Where(&models.User{}).First(&user).Error; err != nil {

		tx.Rollback()

		utils.SendResponse(user, "Success to login", true, http.StatusOK, c)

		// return

	}

}

func Logout(c *gin.Context) {

	var user models.User

	tx := database.DBConn.Model(&models.User{}).Begin()

	res := tx.Scopes(utils.LoggedUser(c)).Where("token != ?", "").First(&user).RowsAffected

	if res == 0 {

		tx.Rollback()

		utils.SendResponse(nil, "User not found", false, http.StatusInternalServerError, c)

		// return

	}

	user.Token = ""

	user.LogoutAt = time.Now()

	if err := tx.Scopes(utils.LoggedUser(c)).Save(&user).Error; err != nil {

		tx.Rollback()

		utils.SendResponse(err, "Failed to logout", false, http.StatusInternalServerError, c)

		// return

	}

	tx.Commit()

	utils.SendResponse(nil, "Success to logout", false, http.StatusInternalServerError, c)

	// return

}
