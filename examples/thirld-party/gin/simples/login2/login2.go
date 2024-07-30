package main

import "github.com/gin-gonic/gin"
import "github.com/gin-gonic/gin/binding"

//https://github.com/gin-gonic/gin/issues/170
//There can not be two handlers with the same path.

// Binding from JSON
type LoginJSON struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Binding from form values
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json LoginJSON

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.
		if json.User == "manu" && json.Password == "123" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	})

	// Example for binding a HTML form (user=manu&password=123)
	r.POST("/loginHTML", func(c *gin.Context) {
		var form LoginForm

		c.BindWith(&form, binding.Form) // You can also specify which binder to use. We support binding.Form, binding.JSON and binding.XML.
		if form.User == "manu" && form.Password == "123" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}
