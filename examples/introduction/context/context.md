# ¿Qué es context?

El concepto de `context` en Go es fundamentalmente una forma de manejar el estado y el control de ejecución en un programa, especialmente cuando se trata de operaciones concurrentes y en sistemas distribuidos.

## Definición y Propósito

- **Contexto como Interfaz**: En Go, `context` es una interfaz definida en el paquete estándar `context`. Esta interfaz proporciona métodos para controlar la cancelación de operaciones, establecer límites de tiempo para la ejecución de procesos (plazos y timeouts), y transportar datos a través de la ejecución de una aplicación.

- **Manejo de Cancelaciones**: Una de las principales razones para usar contextos es manejar la cancelación de operaciones de forma eficaz. Esto es especialmente útil en programas concurrentes, donde múltiples hilos de ejecución (goroutines) necesitan ser coordinados y posiblemente cancelados en respuesta a ciertos eventos, como la finalización de tareas dependientes, errores, o interrupciones de usuario.

- **Establecimiento de Plazos y Timeouts**: Los contextos permiten especificar plazos o timeouts después de los cuales una operación debería ser cancelada automáticamente. Esto ayuda a evitar que las operaciones se bloqueen indefinidamente y facilita la creación de programas más robustos y con mejor respuesta a fallos.

- **Propagación de Datos**: Además de controlar la cancelación y el tiempo de vida de las operaciones, los contextos pueden llevar datos a través de las fronteras de las llamadas de función. Esto es útil para pasar información relevante como identificadores de solicitud, detalles de autenticación, o preferencias de configuración a lo largo de una cadena de procesamiento, sin tener que modificar la firma de cada función involucrada.

## Cómo Funciona

- **Creación y Derivación**: Un contexto se crea con una función de base como `context.Background()` o `context.TODO()`. A partir de ahí, se derivan nuevos contextos utilizando funciones como `context.WithCancel`, `context.WithDeadline`, `context.WithTimeout`, y `context.WithValue`. Estos contextos derivados heredan el estado del contexto padre, pero pueden ser cancelados o tener datos adjuntos de forma independiente.

- **Cancelación Propagada**: Cuando un contexto se cancela, ya sea manualmente a través de una función de cancelación o automáticamente debido a un plazo o timeout, esta cancelación se propaga a todos los contextos derivados de él. Esto permite una coordinación efectiva de la cancelación a través de múltiples goroutines y operaciones.

- **Uso a través de Goroutines y API**: Los contextos se pasan explícitamente de una función a otra como el primer argumento (por convención). Esto asegura que la cancelación, los plazos y los datos específicos estén disponibles a lo largo de toda la operación, desde el inicio hasta las funciones más profundas en la cadena de llamadas.

## Context Base

La utilización más simple del paquete `context` en Go podría ser el uso de `context.Background()` para crear un contexto base que no hace nada especial: no se cancela, no tiene un deadline, y no lleva valores. Este contexto luego se pasa a una función que lo acepta como argumento pero no utiliza ninguna de las características especiales del contexto, cumpliendo con una interfaz que espera un contexto pero sin aplicar control de concurrencia o cancelación.

```go
package main

import (
    "context"
    "fmt"
)

// Una función simple que acepta un contexto y un mensaje, pero solo imprime el mensaje.
func printWithContext(ctx context.Context, message string) {
    fmt.Println(message)
}

func main() {
    // Crear un contexto básico
    ctx := context.Background()
    
    // Pasar el contexto y un mensaje a la función
    printWithContext(ctx, "Hello, world del context!")
}
```

Este ejemplo es extremadamente simplificado y en la práctica, el valor de `context` viene de su uso para cancelar operaciones, establecer deadlines, o pasar valores específicos a través de las fronteras de las llamadas de función en operaciones más complejas y en ejecuciones concurrentes.

## Propósito context.Background

`context.Background()` tiene un propósito muy específico y valioso en Go, aunque a primera vista parezca que "no se puede hacer nada con él". Es la raíz de cualquier árbol de contextos dentro de una aplicación y sirve como el punto de partida para crear contextos más específicos, o "contextos hijos", con funcionalidades adicionales como cancelación, deadlines, y almacenamiento de valores específicos. Aquí detallamos sus principales usos y por qué es importante:

### Punto de Partida Universal

- **Contexto Base:** `context.Background()` proporciona un contexto base cuando no hay otro contexto más específico disponible. Es especialmente útil al inicio de una aplicación, en la función `main`, o al comenzar una nueva goroutine para la cual no existe un contexto entrante.

### Creación de Contextos Específicos

- **Crear Contextos Hijos:** A partir de `context.Background()`, puedes derivar contextos con características adicionales. Por ejemplo, puedes usar `context.WithCancel`, `context.WithDeadline`, y `context.WithTimeout` para crear contextos que pueden ser cancelados o que expiran después de un cierto tiempo. Esto es esencial para manejar la cancelación de operaciones y timeouts en procesos concurrentes.

### Neutralidad

- **Contexto Neutral:** Cuando estás escribiendo una biblioteca o un paquete que será usado en diversos contextos, comenzar con `context.Background()` permite que el usuario de tu biblioteca decida cómo y cuándo introducir contextos más específicos, proporcionando flexibilidad.

### Casos de Uso

- **Operaciones de Larga Duración:** Para operaciones que son esenciales para la vida útil de toda la aplicación y que no se espera que sean canceladas (por ejemplo, procesos de inicialización o servicios que corren indefinidamente).

### Ejemplo Práctico

Supongamos que tienes una aplicación que inicia varios servidores internos (HTTP, gRPC, etc.) y servicios de fondo al arrancar:

```go
func main() {
    ctx := context.Background()

    // Iniciar un servidor HTTP
    go startHTTPServer(ctx)

    // Iniciar un servicio de fondo
    go startBackgroundService(ctx)

    // Esperar señales del sistema o condiciones de salida
    waitForShutdown(ctx)
}
```

En este escenario, `context.Background()` actúa como el contexto raíz del cual pueden derivarse otros contextos más específicos, según las necesidades de cada goroutine o servicio. Por ejemplo, podrías querer cancelar todas las operaciones iniciadas por `main` en caso de una señal de apagado; esto se facilita pasando un contexto cancelable derivado de `context.Background()` a tus funciones.

Aunque `context.Background()` por sí solo no proporciona funcionalidades de cancelación, deadlines, o almacenamiento de valores, su valor radica en ser el ancestro universal de todos los contextos en una aplicación Go. Facilita la estructura y el manejo de operaciones concurrentes de manera ordenada y predecible, permitiendo la derivación de contextos más específicos según se requiera.


## `context.Background()` y `context.TODO()`

Son dos funciones del paquete `context` en Go que sirven para crear contextos base. Estos contextos son puntos de partida para la derivación de nuevos contextos con más funcionalidades, como cancelación,

 deadlines, timeouts y valores adicionales.

### `context.Background()`

`context.Background()` devuelve un contexto vacío. Este contexto no se cancela, no tiene valores y no tiene deadline. Es el contexto más "puro" y sirve como el punto de partida para crear contextos más específicos en una aplicación. Es utilizado como el contexto raíz de una cadena de operaciones cuando no hay otro contexto más apropiado que se deba utilizar.

**Cuándo y por qué se usa `context.Background()`**:

- **Inicio de una aplicación/servicio**: Ideal para usar en el `main()` de tu programa o al iniciar goroutines a nivel de la aplicación, especialmente cuando no hay un contexto entrante de otra operación.
- **Pruebas**: A menudo se usa en tests para inicializar contextos necesarios para ejecutar operaciones que requieren un contexto.
- **Operaciones de larga duración**: Para operaciones que se espera que se ejecuten durante toda la vida útil de la aplicación y no necesitan ser canceladas.

### `context.TODO()`

`context.TODO()` también devuelve un contexto vacío, similar a `context.Background()`. La diferencia es semántica y sirve como una indicación de que el contexto debe ser definido más adelante. `context.TODO()` se utiliza en lugares del código donde aún no está claro qué contexto usar o si el código eventualmente necesitará un contexto más específico.

**Cuándo, cómo y por qué se usa `context.TODO()`**:

- **Código en Desarrollo**: Durante el desarrollo inicial, cuando aún estás decidiendo cómo manejar la cancelación o el paso de valores.
- **Refactorización**: Cuando estás refactorizando código existente para usar contextos pero aún no has decidido cómo se integrarán en toda la aplicación.
- **Marcador de posición**: Como una señal a ti mismo o a otros desarrolladores de que el contexto necesita una revisión y probablemente debería ser reemplazado por un contexto más específico en el futuro.

## Inmutabilidad

Los contextos en Go son inmutables, lo que significa que una vez que un contexto es creado, no puede ser modificado. Cada vez que necesitas agregar información o cambiar el comportamiento de un contexto (por ejemplo, añadiendo un timeout, un deadline, o valores específicos), lo que en realidad haces es crear un nuevo contexto basado en el anterior. Este nuevo contexto hereda las propiedades del contexto original, además de las modificaciones o adiciones que hayas hecho.

La inmutabilidad de los contextos tiene varias ventajas importantes:

### Simplificación del manejo de concurrencia

Dado que los contextos son inmutables, puedes pasarlos de manera segura entre goroutines sin preocuparte por problemas de concurrencia, como condiciones de carrera. No hay riesgo de que una goroutine modifique el contexto de manera que afecte a otras goroutines que podrían estar utilizándolo.

### Seguridad y previsibilidad

Al ser inmutables, el comportamiento de un contexto es predecible. Sabes que no cambiará una vez creado, lo que facilita el razonamiento sobre tu código y reduce la probabilidad de efectos secundarios inesperados.

### Encadenamiento y derivación

Cuando derivas un nuevo contexto de uno existente, estás creando una cadena de contextos. Esto es útil para cancelaciones y timeouts, ya que cancelar un contexto padre automáticamente cancela todos los contextos derivados de él, lo cual es una forma efectiva de propagar señales de cancelación a través de tu aplicación.

### Ejemplo práctico

Veamos cómo se aplica esto en código. Supongamos que tienes un servidor web que maneja solicitudes de usuarios. Para una solicitud particular, quieres establecer un timeout para asegurarte de que no se tarde demasiado en responder.

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Crear un contexto con un timeout
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // Asegurarse de cancelar el contexto para liberar recursos

	// Simular un trabajo que podría tardar más tiempo del permitido por el timeout
	work(ctx)

	fmt.Fprintf(w, "Respuesta enviada al cliente")
}

func work(ctx context.Context) {
	select {
	case <-time.After(200 * time.Millisecond): // Simular trabajo
		fmt.Println("Trabajo completado")
	case <-ctx.Done(): // El contexto fue cancelado o expiró el timeout
		fmt.Println("Trabajo cancelado:", ctx.Err())
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
```

En este ejemplo, aunque `context.WithTimeout` modifica el comportamiento del contexto original (`context.Background()`), lo hace creando un nuevo contexto que es una versión modificada del original. Si la operación `work` tarda demasiado, el contexto se cancela (debido al timeout), y la ejecución puede detenerse antes de que la operación se complete, demostrando cómo los contextos inmutables pueden usarse para controlar el flujo de operaciones en aplicaciones concurrentes.

### Ejemplos de Uso

#### `context.Background()` en una aplicación web

```go
func main() {
    ctx := context.Background()
    setupServer(ctx)
}

func setupServer(ctx context.Context) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Aquí podríamos derivar un nuevo contexto para cada solicitud
        doWork(ctx)
    })
    http.ListenAndServe(":8080", nil)
}
```

#### `context.TODO()` en un código en proceso de integrar contextos

```go
func fetchDataFromDB() {
    // Supongamos que este código necesita ser refactorizado para usar contextos
    ctx := context.TODO() // Marcador de que este contexto necesita ser reemplazado
    // Imagina una llamada a la base de datos aquí usando ctx
}
```

La elección entre `context.Background()` y `context.TODO()` se reduce a la intención y la claridad del código. Usa `context.Background()` como un contexto raíz o cuando estés seguro de que una operación no necesita ser cancelada. Utiliza `context.TODO()` como una señal de que el contexto está pendiente de revisión y podría necesitar ser especificado con más detalle en el futuro. Ambos contextos sirven como puntos de partida claros y limpios para la creación de contextos más complejos en aplicaciones Go.