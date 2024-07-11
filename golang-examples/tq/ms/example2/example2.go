package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configuración del Circuit Breaker
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:                1000, // Tiempo máximo de ejecución en milisegundos
		MaxConcurrentRequests:  100,  // Número máximo de solicitudes concurrentes permitidas
		ErrorPercentThreshold:  50,   // Umbral de porcentaje de errores para abrir el circuito
		RequestVolumeThreshold: 10,   // Volumen mínimo de solicitudes antes de evaluar si abrir el circuito
		SleepWindow:            5000, // Tiempo en milisegundos para esperar antes de intentar cerrar el circuito
	})

	// Configurar Gin
	router := gin.Default()

	// Ruta para manejar solicitudes
	router.GET("/request", func(c *gin.Context) {
		err := makeRequest()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Solicitud exitosa"})
		}
	})

	// Iniciar el servidor
	router.Run(":8080")
}

func makeRequest() error {
	// Simular una función que puede fallar
	return hystrix.Do("my_command", func() error {
		fmt.Println("Intentando realizar la solicitud...")
		// Simular fallo
		if time.Now().Unix()%2 == 0 {
			return errors.New("servicio no disponible")
		}
		// Simular éxito
		fmt.Println("Solicitud exitosa!")
		return nil
	}, func(err error) error {
		// Función de fallback
		fmt.Printf("Fallback ejecutado. Error: %v\n", err)
		return errors.New("servicio no disponible (fallback)")
	})
}
