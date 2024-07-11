package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"
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

	// Simular solicitudes
	for i := 0; i < 20; i++ {
		go makeRequest(i)
	}

	// Esperar a que todas las solicitudes se completen
	time.Sleep(10 * time.Second)
}

func makeRequest(id int) {
	// Simular una función que puede fallar
	err := hystrix.Do("my_command", func() error {
		fmt.Printf("Solicitud %d: Intentando realizar la solicitud...\n", id)
		// Simular fallo en las primeras 10 solicitudes
		if id < 10 {
			return errors.New("servicio no disponible")
		}
		// Simular éxito en las siguientes solicitudes
		fmt.Printf("Solicitud %d: Solicitud exitosa!\n", id)
		return nil
	}, func(err error) error {
		// Función de fallback
		fmt.Printf("Solicitud %d: Fallback ejecutado. Error: %v\n", id, err)
		return nil
	})

	if err != nil {
		fmt.Printf("Solicitud %d: La solicitud falló después de varios intentos: %v\n", id, err)
	}
}
