
	router.GET("/formulario-nuevo-Employ", CrearEmployHTML)
	router.POST("/nuevo-Employ-JSON", CrearEmployJSON)
	router.POST("/nuevo-Employ-XML", CrearEmployXML)
	router.POST("/mostrar-Employ", CrearEmploy)
	router.GET("/obterner-Employ/:id", ObtenerEmploy)
	router.GET("/obterner-todos-Employs/", ObtenerTodosLosEmploys)
	router.PUT("/actualizar-Employ/:id", ActualizarEmploy)
	router.DELETE("/eliminar-Employ/:id", EliminarEmploy)