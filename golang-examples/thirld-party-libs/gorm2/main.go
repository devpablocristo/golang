package main

import (
	"gormtest/infrastructure"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	infrastructure.LoadEnv()

	PORT := os.Getenv("SERVER_PORT")

	router := gin.Default()

	router.GET("/", helloWorld)

	router.Run(":" + PORT)
}

func helloWorld(c *gin.Context) {
	infrastructure.LoadEnv()
	infrastructure.NewDatabase()
	//c.String(http.StatusOK, "¡Hello World!")

	c.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
}

// package main

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type Event struct {
// 	ID        int       `json:"id"`
// 	StartTime time.Time `json:"start_time"`
// 	EndTime   time.Time `json:"end_time"`
// }

// func main() {
// 	//dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
// 	dsn := loadEnv(".")
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// migrar el esquema
// 	err = db.AutoMigrate(&Event{})
// 	if err != nil {
// 		panic("failed to migrate schema")
// 	}

// 	// crear un nuevo evento
// 	newEvent := &Event{
// 		StartTime: time.Now(),
// 		EndTime:   time.Now().Add(2 * time.Hour),
// 	}
// 	db.Create(newEvent)

// 	// obtener todos los eventos
// 	var events []Event
// 	db.Find(&events)
// 	fmt.Println(events)
// }