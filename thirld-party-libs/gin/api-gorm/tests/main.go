package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

/*  */
type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

/*  */
type Empleado struct {
	ID           uint   `json:"id"`
	Puesto       string `json:"puesto"`
	Movil        string `json:"movil"`
	NumeroLegajo uint   `json:"numerolegajo"`
}

/*  */
type Usuario struct {
	ID                 uint   `json:"id"`
	NombreUsuario      string `json:"nombreusuario"`
	ContraseniaUsuario string `json:"contraseniausuario"`
}

func main() {

	// NOTE: See we're using = to assign the global var
	//         	instead of := which would assign it only in this function
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	db, _ = gorm.Open("mysql", "pablo:rocky@tcp(127.0.0.1:3306)/pruebas?charset=utf8&parseTime=True&loc=Local")

	//db, err := gorm.Open("mysql", "pablo:rocky@/pruebas?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//db.AutoMigrate(&Person{})

	//db.AutoMigrate(&Empleado{})

	db.AutoMigrate(&Usuario{})

	r := gin.Default()
	//r.GET("/people/", GetPeople)
	//r.GET("/people/:id", GetPerson)
	r.POST("/people", CreatePerson)
	r.POST("/crear", CrearUsuario)
	//r.PUT("/people/:id", UpdatePerson)
	//r.DELETE("/people/:id", DeletePerson)

	r.Run(":8080")
}

/* func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
} */

/* func UpdatePerson(c *gin.Context) {

	var person Person
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)

	db.Save(&person)
	c.JSON(200, person)

} */

func CreatePerson(c *gin.Context) {

	var person Person
	c.BindJSON(&person)

	db.Create(&person)
	c.JSON(200, person)
}

/*
curl -i -X  POST http://localhost:8080/crear -d '{ "NombreUsuario": "pablo", "ContraseñaUsuario":"lalala" }'
*/
func CrearUsuario(c *gin.Context) {
	var usuario Usuario

	/*
		Bind comprueba el Content-Type para seleccionar un motor de enlace (binding engine) automáticamente. Según el encabezado "Content-Type" se utilizan diferentes enlaces:

		"application/json" --> JSON binding
		"application/xml"  --> XML binding

		de lo contrario -> devuelve un error. Analiza el cuerpo de la solicitud como JSON si Content-Type == "application / json" usa JSON o XML como entrada JSON. Decodifica la carga json en la estructura especificada como puntero. Escribe un error 400 y establece el encabezado de tipo de contenido "texto / plano" en la respuesta si la entrada no es válida.
	*/
	// 	BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
	c.BindJSON(&usuario)

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.NewRecord(user) // => returns `true` as primary key is blank
	//db.Create(&user)
	//db.NewRecord(user) // => return `false` after `user` created

	db.Create(&usuario)
	c.JSON(200, usuario)
}

/* func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
} */

/* func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}

} */
