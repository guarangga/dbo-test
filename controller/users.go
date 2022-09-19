package apps

import (
	"github.com/gin-gonic/gin"
	"dbo-test/database"
	"dbo-test/models"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
	"fmt"
	// "gorm.io/gorm"
)

var result gin.H

func GetUsers(c *gin.Context) {

	tx := database.DBConn.Model(&models.User{})

	var user []models.User

	err := tx.Find(&user).Error

	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "Empty data",
		}
	} else {
		result = gin.H{
			"status": 200,
			"data": user,
		}
	}

	c.JSON(http.StatusOK, result)

}

func FindUsers(c *gin.Context) {

	tx := database.DBConn.Model(&models.User{})

	var user []models.User

	id := c.Param("id")

	res := tx.Where("id = ?", id).First(&user).RowsAffected

	if res == 0 {

		tx.Rollback()

		result = gin.H{
			"status": 404,
			"result": "data not found",
		}

	}

	err := tx.Find(&user).Error

	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "Empty data",
		}
	} else {
		result = gin.H{
			"status": 200,
			"data": user,
		}
	}

	c.JSON(http.StatusOK, result)

}

func CreateUsers(c *gin.Context) {

	tx := database.DBConn.Model(&models.User{})

	var (
		result gin.H
		user models.User
	)

	var name, email, password, ip string

	form, e := c.MultipartForm()

	if e == nil {

		for key, val := range form.Value {

			if key == "name" {

				name = val[0]

			}

			if key == "email" {

				email = val[0]

			}

			if key == "password" {

				password = val[0]

			}

			if key == "ip" {

				ip = val[0]

			}

		}

	}

	user.Id = uuid.NewV4()
	user.Name = name
	user.Email = email
	user.Ip = ip
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	err := tx.Create(&user).Error

	if err != nil {
		result = gin.H{
			"message": err,
			"status": 400,
			"result": "create failed",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "successfully created data",
			"data": user,
		}
	}

	c.JSON(http.StatusOK, result)

}

func UpdateUsers(c *gin.Context) {

	var (
		user    models.User
		result  gin.H
	)

	id := c.Param("id")

	tx := database.DBConn.Model(&models.User{})

	res := tx.Where("id = ?", id).First(&user).RowsAffected

	if res == 0 {

		tx.Rollback()

		result = gin.H{
			"status": 404,
			"result": "data not found",
		}

	}

	var name, email string

	form, e := c.MultipartForm()

	if e == nil {

		for key, val := range form.Value {

			if key == "name" {

				name = val[0]

			}

			if key == "email" {

				email = val[0]

			}

		}

	}

	err := tx.First(&user).Error

	println(err)

	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}

	user.Name = name

	user.Email = email

	fmt.Println("ID : ", id)

	err = tx.Save(&user).Error

	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "update failed",
			"message": c.Error(err),
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "successfully updated data",
			"data": user,
		}
	}

	c.JSON(http.StatusOK, result)

}

func DeleteUsers(c *gin.Context) {

	var (
		user models.User
		result gin.H
	)

	id := c.Param("id")

	tx := database.DBConn.Model(&models.User{})

	res := tx.Where("id = ?", id).First(&user).RowsAffected

	if res == 0 {

		tx.Rollback()

		result = gin.H{
			"status": 404,
			"result": "data not found",
		}

	}

	err := tx.First(&user).Error

	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}

	err = tx.Delete(&user).Error

	fmt.Println(err)

	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "delete failed",
			"message": c.Error(err),
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "Data deleted successfully",
			"data": user,
		}
	}

	c.JSON(http.StatusOK, result)

}
// func GetUsers(c *gin.Context) error {
//
// 	db := c.Value("database").(*gorm.DB)
// 	pageStr := c.DefaultQuery("page", "1")
// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}
// 	var userCount int64
// 	if err := db.Table("users").Count(&userCount).Error; err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}
// 	const usersPerPage = 15
// 	pageCount := int(math.Ceil(float64(userCount) / float64(usersPerPage)))
// 	if pageCount == 0 {
// 		pageCount = 1
// 	}
// 	if page < 1 || page > pageCount {
// 		c.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}
// 	offset := (page - 1) * usersPerPage
// 	users := []User{}
// 	if err := db.Limit(usersPerPage).Offset(offset).Find(&users).Error; err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}
// 	var prevPage, nextPage string
// 	if page > 1 {
// 		prevPage = fmt.Sprintf("%d", page-1)
// 	}
// 	if page < pageCount {
// 		nextPage = fmt.Sprintf("%d", page+1)
// 	}
//
// 	pages := make([]int, pageCount)
// 	for i := 0; i < pageCount; i++ {
// 		pages[i] = i + 1
// 	}
//
// 	c.HTML(http.StatusOK, "users/index.html", gin.H{
// 		"users":     users,
// 		"pageCount": pageCount,
// 		"page":      page,
// 		"prevPage":  prevPage,
// 		"nextPage":  nextPage,
// 		"pages":     pages,
// 	})
//
// }
