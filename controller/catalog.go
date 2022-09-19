package apps

import (
	"github.com/gin-gonic/gin"
	"dbo-test/database"
	"dbo-test/models"
	"dbo-test/utils"
	"net/http"
	"github.com/satori/go.uuid"
	"fmt"
	"strconv"
	// "gorm.io/gorm"
)

var resultCatalog gin.H

func GetCatalog(c *gin.Context) {

	tx := database.DBConn.Model(&models.Catalog{})

	var catalog []models.Catalog

	// err := tx.Find(&catalog).Error

	var count int64

  param := utils.ParseDefaultGetParam(c)

	tx.Scopes(utils.SelectedTable("id", param.Selected)).Scopes(utils.KeywordTable(param.Keyword, "name", "email")).Count(&count)

	tx.Scopes(utils.SelectedTable("id", param.Selected)).Scopes(utils.PaginateTable(param.Page, param.Load)).Scopes(utils.SortTable(param.Sorted)).Scopes(utils.KeywordTable(param.Keyword, "name", "email")).Scopes(utils.SearchTable(param.Search)).Find(&catalog)

	utils.SendResponseResource(param.Page, param.Load, count, catalog, "Success to get catalogs", true, http.StatusOK, c)

	// if err != nil {
	// 	utils.SendResponse(nil, "Empty data", false, http.StatusInternalServerError, c)
	// 	// return
	// } else {
	// 	utils.SendResponse(catalog, "Success get data", false, http.StatusOK, c)
	// 	// return
	// }

	// c.JSON(http.StatusOK, resultCatalog)

}

func FindCatalog(c *gin.Context) {

	tx := database.DBConn.Model(&models.Catalog{})

	var catalog []models.Catalog

	id := c.Param("id")

	res := tx.Where("id = ?", id).First(&catalog).RowsAffected

	if res == 0 {

		tx.Rollback()

		utils.SendResponse(nil, "data not found", false, http.StatusInternalServerError, c)
		// return

	}

	err := tx.Find(&catalog).Error

	if err != nil {
		utils.SendResponse(nil, "Empty data", false, http.StatusInternalServerError, c)
		// return
	} else {
		utils.SendResponse(catalog, "Success get data", false, http.StatusOK, c)
		// return
	}

	// c.JSON(http.StatusOK, resultCatalog)

}

func CreateCatalog(c *gin.Context) {

	tx := database.DBConn.Model(&models.Catalog{})

	var (
		catalog models.Catalog
	)

	var code, short_description, long_description, status, notes, uom string

	var price int64

	form, e := c.MultipartForm()

	if e == nil {

		for key, val := range form.Value {

			if key == "code" {

				code = val[0]

			}

			if key == "short_description" {

				short_description = val[0]

			}

			if key == "long_description" {

				long_description = val[0]

			}

			if key == "status" {

				status = val[0]

			}

			if key == "notes" {

				notes = val[0]

			}

			if key == "uom" {

				uom = val[0]

			}

			if key == "price" {

				x, _ := strconv.ParseInt(val[0], 10, 64)

				price = x

			}

		}

	}

	catalog.Id = uuid.NewV4()
	catalog.Code = code
	catalog.ShortDescription = short_description
	catalog.LongDescription = long_description
	catalog.Status = status
	catalog.Notes = notes
	catalog.Uom = uom
	catalog.Price = price

	err := tx.Create(&catalog).Error

	if err != nil {
		utils.SendResponse(nil, "create failed", false, http.StatusInternalServerError, c)
		// return
	} else {
		utils.SendResponse(catalog, "Success created data", false, http.StatusOK, c)
		// return
	}

	// c.JSON(http.StatusOK, resultCatalog)

}

func UpdateCatalog(c *gin.Context) {

	var (
		catalog    models.Catalog
	)

	id := c.Param("id")

	tx := database.DBConn.Model(&models.Catalog{})

	res := tx.Where("id = ?", id).First(&catalog).RowsAffected

	if res == 0 {

		tx.Rollback()

		utils.SendResponse(nil, "data not found", false, http.StatusInternalServerError, c)
		// return

	}

	var code, short_description, long_description, status, notes, uom string

	var price int64

	form, e := c.MultipartForm()

	if e == nil {

		for key, val := range form.Value {

			if key == "code" {

				code = val[0]

			}

			if key == "short_description" {

				short_description = val[0]

			}

			if key == "long_description" {

				long_description = val[0]

			}

			if key == "status" {

				status = val[0]

			}

			if key == "notes" {

				notes = val[0]

			}

			if key == "uom" {

				uom = val[0]

			}

			if key == "price" {

				x, _ := strconv.ParseInt(val[0], 10, 64)

				price = x

			}

		}

	}

	err := tx.First(&catalog).Error

	// println(err)

	if err != nil {
		utils.SendResponse(nil, "data not found", false, http.StatusInternalServerError, c)
		// return
	}

	catalog.Id = uuid.NewV4()
	catalog.Code = code
	catalog.ShortDescription = short_description
	catalog.LongDescription = long_description
	catalog.Status = status
	catalog.Notes = notes
	catalog.Uom = uom
	catalog.Price = price

	fmt.Println("ID : ", id)

	err = tx.Save(&catalog).Error

	if err != nil {
		utils.SendResponse(nil, "update failed", false, http.StatusInternalServerError, c)
		// return
	} else {
		utils.SendResponse(catalog, "Success updated data", false, http.StatusOK, c)
		// return
	}

	// c.JSON(http.StatusOK, resultCatalog)

}

func DeleteCatalog(c *gin.Context) {

	var (
		catalog models.Catalog
	)

	id := c.Param("id")

	tx := database.DBConn.Model(&models.Catalog{})

	res := tx.Where("id = ?", id).First(&catalog).RowsAffected

	if res == 0 {

		tx.Rollback()

		resultCatalog = gin.H{
			"status": 404,
			"resultCatalog": "data not found",
		}
		utils.SendResponse(nil, "data not found", false, http.StatusInternalServerError, c)
		// return

	}

	err := tx.First(&catalog).Error

	if err != nil {
		utils.SendResponse(nil, "data not found", false, http.StatusInternalServerError, c)
		// return
	}

	err = tx.Delete(&catalog).Error

	fmt.Println(err)

	if err != nil {
		utils.SendResponse(nil, "delete failed", false, http.StatusInternalServerError, c)
		// return
	} else {
		utils.SendResponse(nil, "Success delete data", false, http.StatusOK, c)
		// return
	}

	// c.JSON(http.StatusOK, resultCatalog)

}
