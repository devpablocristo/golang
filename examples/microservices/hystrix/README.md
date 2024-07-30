### Código Completo

```go
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
```

### Explicación Detallada

#### 1. Configuración del Circuit Breaker

```go
hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
	Timeout:                1000,
	MaxConcurrentRequests:  100,
	ErrorPercentThreshold:  50,
	RequestVolumeThreshold: 10,
	SleepWindow:            5000,
})
```

- **Timeout**: El tiempo máximo permitido para la ejecución de la función antes de considerarla un fallo (1000 ms).
- **MaxConcurrentRequests**: El número máximo de solicitudes concurrentes permitidas (100).
- **ErrorPercentThreshold**: El porcentaje de errores permitidos antes de que el circuito se abra (50%).
- **RequestVolumeThreshold**: El número mínimo de solicitudes necesarias antes de que Hystrix considere abrir el circuito (10).
- **SleepWindow**: El tiempo que el circuito permanecerá abierto antes de intentar cerrarse y permitir nuevamente las solicitudes (5000 ms).

#### 2. Simulación de Solicitudes

```go
for i := 0; i < 20; i++ {
	go makeRequest(i)
}

// Esperar a que todas las solicitudes se completen
time.Sleep(10 * time.Second)
```

- Se simulan 20 solicitudes concurrentes utilizando un bucle `for`.
- Cada solicitud se ejecuta en una gorutina separada llamando a la función `makeRequest`.
- `time.Sleep(10 * time.Second)` asegura que el programa espere lo suficiente para que todas las solicitudes se completen.

#### 3. Función Principal (Protected Function)

```go
func makeRequest(id int) {
	err := hystrix.Do("my_command", func() error {
		fmt.Printf("Solicitud %d: Intentando realizar la solicitud...\n", id)
		if id < 10 {
			return errors.New("servicio no disponible")
		}
		fmt.Printf("Solicitud %d: Solicitud exitosa!\n", id)
		return nil
	}, func(err error) error {
		fmt.Printf("Solicitud %d: Fallback ejecutado. Error: %v\n", id, err)
		return nil
	})

	if err != nil {
		fmt.Printf("Solicitud %d: La solicitud falló después de varios intentos: %v\n", id, err)
	}
}
```

- **Función Principal**: La primera función pasada a `hystrix.Do` simula una operación que puede fallar. Las primeras 10 solicitudes fallan, mientras que las siguientes 10 tienen éxito.
- **Función de Fallback**: La segunda función pasada a `hystrix.Do` es una función de fallback que se ejecuta si la función principal falla o si el circuito está abierto. Imprime un mensaje indicando que se ha ejecutado el fallback y muestra el error recibido.
- **Manejo de Errores**: Si `hystrix.Do` devuelve un error, se imprime un mensaje indicando que la solicitud falló después de varios intentos.

### Estados del Circuit Breaker

1. **Estado Cerrado (Closed)**:
   - El circuito permite que todas las solicitudes pasen a la función protegida.
   - Si las solicitudes fallan, se registran los fallos.
   - En el código, cuando se hacen las solicitudes iniciales (id < 10), los fallos se registran.

2. **Estado Abierto (Open)**:
   - Si el número de fallos supera el umbral de error (`ErrorPercentThreshold`), el circuito se abre.
   - En este estado, el circuito bloquea todas las solicitudes subsiguientes durante el tiempo especificado en `SleepWindow`.
   - Las solicitudes que se realicen en este estado se manejan inmediatamente mediante la función de fallback.
   - En el código, si las primeras 10 solicitudes fallan y cumplen el umbral de errores, el circuito se abrirá y las solicitudes posteriores se manejarán mediante la función de fallback.

3. **Estado Semi-Abierto (Half-Open)**:
   - Después de que expira el `SleepWindow`, el circuito permite una cantidad limitada de solicitudes de prueba.
   - Si estas solicitudes tienen éxito, el circuito se cierra nuevamente.
   - Si fallan, el circuito vuelve a abrirse.
   - En el código, después de que expire el `SleepWindow`, el circuito permitirá algunas solicitudes de prueba. Si estas solicitudes tienen éxito (id >= 10), el circuito se cerrará nuevamente.

Este ejemplo muestra cómo `hystrix-go` gestiona automáticamente los estados del circuit breaker y cómo se manejan las solicitudes y los fallos en cada estado.