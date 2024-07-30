
/*
	Get
	http://localhost:8080/obterner-Employ/1
*/
func ObtenerTodosLosEmploys(c *gin.Context) {
	var resultado gin.H
	//var err error
	var e Employ
	var todosE []Employ

	todosE = e.LeerTodos()

	resultado = gin.H{
		"resultado": todosE,
		"cantidad":  len(todosE),
	}

	c.JSON(http.StatusOK, resultado)
}

/*
	Get
	http://localhost:8080/eliminar-Employ/1
*/
func ActualizarEmploy(c *gin.Context) {
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

	e.Actualizar()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)
}

/*
	Get
	http://localhost:8080/eliminar-Employ/1
	Softdelete
*/
func EliminarEmploy(c *gin.Context) {
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

	e.Eliminar()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)

	/*c.HTML(http.StatusOK, "mostrar-datos-Employ.tmpl", gin.H{
	"nombres": e.Leer(e.IDEmploy).Person.NombresPerson})*/
}