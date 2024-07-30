package login

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	var err error

	conex := "pablo:rocky@tcp(127.0.0.1:3306)/pruebas?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", conex)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob("pla/*")

	/*
		sudo lsof -n -i :8080
		kill -9 <PID>
		sudo killall -9 main3
	*/
	router.Run(":8080")
}

/*
Get
http://localhost:8080/formulario-Employ
*/
func CrearEmployHTML(c *gin.Context) {
	//var resultado gin.H
	//var err error

	c.HTML(http.StatusOK, "formulario-Employ.tmpl", nil)
}

func CrearEmployJSON(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Employ

	if c.ShouldBindJSON(&e) != nil {
		resultado = gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, resultado)
		return
	}

	/*
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
	*/

	e.Crear()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	//c.JSON(http.StatusOK, "json")
	c.JSON(http.StatusOK, resultado)
}

/**
 *
 * @api {method} /path title
 * @apiName apiName
 * @apiGroup group
 * @apiVersion  major.minor.patch
 *
 *
 * @apiParam  {String} paramName description
 *
 * @apiSuccess (200) {type} name description
 *
 * @apiParamExample  {type} Request-Example:
 * {
 *     property : value
 * }
 *
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *     property : value
 * }
 *
 *
 */
func CrearEmployXML(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Employ

	if c.ShouldBindXML(&e) != nil {
		resultado = gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, resultado)
		return
	}

	/*
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
	*/

	e.Crear()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	//c.JSON(http.StatusOK, "json")
	c.JSON(http.StatusOK, resultado)
}

// Post
func CrearEmploy(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Employ

	if c.ShouldBind(&e) != nil {
		resultado = gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, resultado)
		return
	}

	e.Crear()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)
}

/*
Get
http://localhost:8080/obterner-Employ/1
*/
func ObtenerEmploy(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Employ
	var idParam string

	idParam = c.Param("id")

	e.IDEmploy, err = strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		resultado = gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, resultado)
		return
	}

	e.Leer()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)

	/*
		c.HTML(http.StatusOK, "mostrar-datos-Employ.tmpl", gin.H{
		"nombres": e.Leer(e.IDEmploy).Person.NombresPerson})
	*/
}
