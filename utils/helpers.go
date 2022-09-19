package utils

import(
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"dbo-test/models"
	"dbo-test/config"
	"dbo-test/database"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
	"errors"
	"strconv"
	"strings"
	"math"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

var jwtKey = []byte("dbo-test")

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PaginateLinks struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Meta struct {
	CurrentPage int         `json:"current_page"`
	From        int64       `json:"from"`
	LastPage    int         `json:"last_page"`
	Links       []MetaLinks `json:"links"`
	Path        string      `json:"path"`
	PerPage     int         `json:"per_page"`
	To          int64       `json:"to"`
	Total       int64       `json:"total"`
}

type MetaLinks struct {
	Url    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
}

type Paginate struct {
	Data  interface{}   `json:"data"`
	Links PaginateLinks `json:"links"`
	Meta  Meta          `json:"meta"`
}

type LoginData map[string]interface{}

type SortList []Sort

type SelectedList []string

type SearchList map[string]string

type Sort struct {
	Id   string `json:"id"`
	Desc bool   `json:"desc"`
}

type DefaultGetParam struct {
	Page     int          `json:"page"`
	Load     int          `json:"load"`
	Keyword  string       `json:"keyword"`
	Sorted   SortList     `json:"sorted"`
	Search   SearchList   `json:"search"`
	Selected SelectedList `json:"selected"`
}

func SendResponse(
	data interface{},
	message string,
	success bool,
	statuscode int,
	c *gin.Context,
)  {

  fmt.Println("data", statuscode)

	// UpdateAccess(c)

	//defer database.DBConn.Close()

  c.JSON(statuscode, gin.H{
		"data": data,
		"message": message,
		"status": statuscode,
	})

	return

}

func SendResponseResource(page int, load int, total int64, data interface{}, message string, success bool, statuscode int, c *gin.Context) {

	page, load, _ = CalculateOffset(page, load)

	lp := int(math.Ceil(float64(total) / float64(load)))

	f := ((page - 1) * load) + 1

	t := total

	if int64(load*page) <= total {
		t = int64(load * page)
	}

	fl := c.Request.Host + c.FullPath() + "?page=1&load=" + strconv.Itoa(load)
	ll := c.Request.Host + c.FullPath() + "?page=" + strconv.Itoa(lp) + "&load=" + strconv.Itoa(load)

	pl := c.Request.Host + c.FullPath() + "?page=" + strconv.Itoa(page-1) + "&load=" + strconv.Itoa(load)
	nl := c.Request.Host + c.FullPath() + "?page=" + strconv.Itoa(page+1) + "&load=" + strconv.Itoa(load)

	if lp == page {
		nl = ""
	}

	if page == 1 {
		pl = ""
	}

	ml := []MetaLinks{
		// {Url: "a", Label: "x", Active: false},
		// {Url: "a", Label: "x", Active: false},
	}

	m := Meta{CurrentPage: page, From: int64(f), LastPage: lp, Links: ml, Path: c.Request.Host + c.FullPath(), PerPage: load, To: t, Total: total}
	p := Paginate{Data: data, Links: PaginateLinks{First: fl, Last: ll, Prev: pl, Next: nl}, Meta: m}

	// UpdateAccess(c)

	//defer database.DBConn.Close()

	// return c.Status(statuscode).JSON(Response{Success: success, Message: message, Data: p})

	c.JSON(statuscode, gin.H{
		"data": p,
		"message": message,
		"status": success,
	})

	return

}

func CheckPasswordHash(password string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err == nil

}

func CheckPassword(providedPassword string) error {
	var user models.User
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims:= &JWTClaim{
		Email: email,
		// Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func GenerateLoginToken(context *gin.Context) {

	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := database.DBConn.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := CheckPassword(user.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err:= GenerateJWT(user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})

}

func GetLoginData(tokenString string) (LoginData, error) {

	var isi LoginData = make(LoginData)

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env("APP_KEY", "dbo-test")), nil
	})

	for key, val := range claims {
		isi[key] = val
	}

	if err != nil {
		return isi, err
	}

	return isi, nil

}

func GetIPAddress(c *gin.Context) string {

	return c.ClientIP()

}

func GetLoginToken(c *gin.Context) string {
	isi := c.Request.Header.Get("Authorization")
	if isi != "" && len(isi) >= 7 {
		return isi[7:]
	} else {
		return ""
	}
}

func LoggedUser(c *gin.Context) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		token := GetLoginToken(c)

		loginData, _ := GetLoginData(token)

		return db.Model(&models.User{}).Where("id = ?", loginData["id"])

	}

}

func ParseDefaultGetParam(c *gin.Context) DefaultGetParam {

	var getpage int

	var getload int

	var getkeyword string

	var getsorted SortList

	var getsearch map[string]string = make(map[string]string)

	var getselected SelectedList

	var mapsorted map[string]string = make(map[string]string)

	form, err := c.MultipartForm()

	var lastIndexSort int

	if err == nil {

		for key, val := range form.Value {

			if key == "page" {

				getpage, _ = strconv.Atoi(val[0])

			}

			if key == "load" {

				getload, _ = strconv.Atoi(val[0])

			}

			if key == "keyword" {

				getkeyword = val[0]

			}

			if len(key) >= 6 && key[0:6] == "search" {

				newkey := strings.ReplaceAll(key, "search", "")

				getsearch[newkey[1:len(newkey)-1]] = val[0]

			}

			if len(key) >= 6 && key[0:6] == "sorted" {

				sp := strings.Split(key, "[")

				var newsp []string

				for _, v := range sp {

					if v[len(v)-1:len(v)] == "]" {

						newsp = append(newsp, v[:len(v)-1])

					}

				}

				lastIndex, _ := strconv.Atoi(newsp[0])

				if lastIndex > lastIndexSort {

					lastIndexSort = lastIndex

				}

				mapsorted[newsp[0]+"-"+newsp[1]] = val[0]

			}

			if len(key) >= 8 && key[0:8] == "selected" {

				getselected = append(getselected, val[0])

			}

		}

	} else {

		// fmt.Println(err)

	}

	for index := 0; index <= lastIndexSort; index++ {

		desc := false

		if mapsorted[strconv.Itoa(index)+"-desc"] == "true" {

			desc = true

		}

		id := mapsorted[strconv.Itoa(index)+"-id"]

		if id != "" {

			getsorted = append(getsorted, Sort{
				Id:   id,
				Desc: desc,
			})

		}

	}

	isi := DefaultGetParam{
		Page:     getpage,
		Load:     getload,
		Keyword:  getkeyword,
		Sorted:   getsorted,
		Search:   getsearch,
		Selected: getselected,
	}

	return isi

}

func SortTable(sort SortList) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		qs := ""

		for _, data := range sort {
			if data.Id != "" {
				if data.Desc {
					qs += data.Id + " desc,"
				} else {
					qs += data.Id + ","
				}
			}
		}

		//fmt.Println("qs", qs)

		if qs == "" {
			return db
		}

		return db.Order(qs[:len(qs)-1])

	}

}

func SelectedTable(pk string, selected SelectedList) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		if len(selected) == 0 {
			return db
		}

		return db.Where(pk+" IN ?", selected)

	}

}

func KeywordTable(keyword string, field ...string) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		qs := ""

		for _, data := range field {
			if data != "" {
				qs += data + " ILIKE @keyword OR "
			}
		}

		if qs == "" || keyword == "" {
			return db
		}

		return db.Where("("+qs[:len(qs)-4]+")", map[string]interface{}{"keyword": "%" + keyword + "%"})

	}

}

func SearchTable(searchlist SearchList) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		qs := ""

		isi := make(map[string]interface{})

		for key, val := range searchlist {

			if key != "" && val != "" {

				qs += key + " ILIKE @" + val + " AND "

				isi[val] = "%" + val + "%"

			}

		}

		if qs == "" {

			return db

		}

		return db.Where(qs[:len(qs)-5], isi)

	}

}

func CalculateOffset(page int, pageSize int) (int, int, int) {

	if page == 0 {

		page = 1

	}

	switch {

	case pageSize > 100:

		pageSize = 100

	case pageSize <= 0:

		pageSize = 10

	}

	offset := (page - 1) * pageSize

	return page, pageSize, offset

}

func PaginateTable(page int, pageSize int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		_, _, offset := CalculateOffset(page, pageSize)

		return db.Offset(offset).Limit(pageSize)

	}

}
