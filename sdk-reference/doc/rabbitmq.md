rabbitmq
Producer envía un mensaje a RabbitMQ.
Exchange enruta el mensaje a la cola apropiada.
Consumer suscrito a la cola recibe el mensaje.
Consumer convierte el mensaje en una estructura y llama a los métodos de los Usecases.
Usecases procesan los datos, ejecutan la lógica de negocio, y opcionalmente envían una respuesta de vuelta.

***

Resumen del Flujo de Trabajo
- Cliente envía una solicitud de login al servicio auth utilizando gRPC.
- Servicio auth valida las credenciales básicas y, si necesita más información del usuario, envía una solicitud a través de RabbitMQ a la cola user_request_queue.
- Servicio user escucha en la cola user_request_queue, procesa la solicitud y envía una respuesta con el UUID del usuario a la cola user_response_queue.
- Servicio auth consume la respuesta desde la cola user_response_queue y envía una respuesta al cliente gRPC.

¿Cuándo Usar Este Enfoque?
- gRPC para Comunicación Síncrona: Se utiliza cuando necesitas una respuesta inmediata y fuerte tipado, como en la autenticación de usuarios.
- RabbitMQ para Comunicación Asíncrona: Se utiliza para manejar solicitudes de larga duración o cuando necesitas desacoplar el procesamiento de mensajes entre servicios, permitiendo que el servicio user maneje la lógica de usuario de manera independiente.


****

RabbitMQ se utiliza comúnmente en varias arquitecturas y patrones de diseño en sistemas distribuidos para manejar la mensajería asincrónica entre microservicios. Las formas más habituales de utilizar RabbitMQ incluyen:

### 1. **Cola de Trabajo (Work Queue) o Tareas en Segundo Plano (Background Tasks)**

**Uso Común**: Este es uno de los patrones más comunes. Las colas de trabajo permiten distribuir tareas pesadas o de larga duración a varios trabajadores. Por ejemplo, un microservicio de frontend puede enviar trabajos a una cola para que otros microservicios los procesen en segundo plano, como el procesamiento de imágenes, la generación de informes, o la indexación de datos.

**Cómo Funciona**:
- **Producer** envía mensajes a una cola en RabbitMQ.
- **Multiple Consumers** están suscritos a esta cola y cada uno toma un mensaje para procesarlo.
- RabbitMQ distribuye los mensajes a los consumidores de manera balanceada (Round-Robin).
- Esto permite que varias instancias del microservicio trabajen en paralelo, mejorando la eficiencia y la velocidad del procesamiento.

**Ejemplo**:
- **Microservicio de imágenes**: Un productor (como un servicio web) recibe una solicitud para procesar una imagen. Envía una tarea a la cola de RabbitMQ. Varios consumidores escuchan esta cola y realizan el procesamiento de la imagen.

### 2. **Publicación/Suscripción (Pub/Sub)**

**Uso Común**: El patrón Pub/Sub se utiliza cuando se necesita que un mensaje sea entregado a múltiples consumidores. Este patrón es útil para aplicaciones que requieren notificaciones en tiempo real o actualizaciones a múltiples servicios cuando ocurre un evento.

**Cómo Funciona**:
- **Producer** publica un mensaje a un **Exchange** de tipo `fanout` o `topic` en RabbitMQ.
- El **Exchange** enruta el mensaje a todas las colas suscritas sin considerar la clave de enrutamiento (en el caso de `fanout`) o utilizando un patrón coincidente (en el caso de `topic`).
- **Multiple Consumers** están suscritos a las colas que reciben los mensajes publicados.

**Ejemplo**:
- **Notificaciones de eventos**: Un servicio de eventos publica un mensaje cada vez que ocurre un evento (como una nueva orden o una actualización de inventario). Varios servicios consumidores, como facturación, inventario, y notificación de usuario, reciben este mensaje y actúan en consecuencia.

### 3. **Enrutamiento de Mensajes (Direct Exchange Routing)**

**Uso Común**: El patrón de enrutamiento directo se utiliza cuando los mensajes deben ser enviados a una cola específica basada en una clave de enrutamiento. Esto es útil para sistemas donde diferentes tipos de mensajes deben ser procesados por diferentes consumidores.

**Cómo Funciona**:
- **Producer** envía un mensaje a un **Exchange** de tipo `direct` con una clave de enrutamiento específica.
- El **Exchange** enruta el mensaje solo a las colas que están vinculadas con esa clave de enrutamiento.
- **Consumer** suscrito a la cola específica recibe y procesa el mensaje.

**Ejemplo**:
- **Procesamiento de registros**: Un servicio que genera registros de diferentes niveles (INFO, ERROR, DEBUG) envía estos registros a un `direct exchange` con claves de enrutamiento correspondientes. Diferentes servicios de monitoreo están suscritos a colas específicas para recibir solo los registros que les interesan.

### 4. **Temas de Enrutamiento (Topic Exchange Routing)**

**Uso Común**: Este patrón es una extensión del enrutamiento directo, pero permite más flexibilidad mediante el uso de comodines (`*` y `#`) en la clave de enrutamiento. Es útil cuando necesitas que un consumidor reciba un subconjunto de mensajes.

**Cómo Funciona**:
- **Producer** envía un mensaje a un **Exchange** de tipo `topic` con una clave de enrutamiento que puede contener patrones.
- El **Exchange** enruta el mensaje a todas las colas cuyas claves de vinculación coinciden con el patrón de la clave de enrutamiento del mensaje.
- **Consumers** suscritos a las colas específicas reciben los mensajes basados en patrones.

**Ejemplo**:
- **Sistema de notificación**: Una aplicación de mensajería podría usar claves de enrutamiento como `user.signup`, `user.login`, `order.created`, etc. Un servicio podría estar interesado en todas las acciones de `user.*` para realizar auditorías o métricas.

### 5. **Colas de Respuesta (Reply Queue) y Solicitud/Respuesta (Request/Reply)**

**Uso Común**: El patrón solicitud/respuesta se utiliza cuando se necesita comunicación bidireccional entre servicios. Un servicio envía una solicitud y espera una respuesta específica de otro servicio.

**Cómo Funciona**:
- **Producer** envía un mensaje a una cola de solicitud y espera una respuesta en una cola de respuesta.
- El mensaje incluye una propiedad `ReplyTo` que indica a qué cola debe enviarse la respuesta.
- **Consumer** recibe la solicitud, procesa la lógica de negocio, y envía una respuesta a la cola especificada en la propiedad `ReplyTo`.

**Ejemplo**:
- **Servicios de autenticación**: Un servicio de frontend envía una solicitud de autenticación a un backend que verifica las credenciales y responde con un token de autenticación a una cola de respuesta.

### 6. **Patrón de Colas Retrasadas (Dead-Letter Queues y Retries)**

**Uso Común**: Las colas de mensajes muertos (DLQs) y los patrones de reintento se utilizan para manejar errores o mensajes que no pueden ser procesados inmediatamente. Esto es útil en sistemas donde los mensajes deben ser procesados eventualmente, incluso si fallan inicialmente.

**Cómo Funciona**:
- **Producer** envía mensajes a una cola normal.
- Si un **Consumer** no puede procesar el mensaje, se redirige a una **Dead-Letter Queue (DLQ)** o a una cola de reintento.
- El mensaje se puede reintentar después de un retraso o ser procesado manualmente más tarde.

**Ejemplo**:
- **Procesamiento de pagos**: Un sistema de procesamiento de pagos intenta procesar una transacción. Si falla debido a un error temporal (por ejemplo, conexión de red), el mensaje se envía a una cola de reintento para un segundo intento.

### Conclusión

La forma más habitual de utilizar RabbitMQ depende de los requisitos específicos de tu aplicación:

- **Cola de trabajo (Work Queue)**: Para procesamiento distribuido de tareas.
- **Publicación/Suscripción (Pub/Sub)**: Para notificaciones en tiempo real y eventos.
- **Enrutamiento de mensajes (Direct o Topic Exchange)**: Para enrutar mensajes a consumidores específicos basados en tipos de mensajes o patrones.
- **Solicitud/Respuesta (Request/Reply)**: Para comunicación bidireccional síncrona-asíncrona.
- **Colas retrasadas y DLQs**: Para manejo de errores y reintentos.

La elección del patrón correcto depende de cómo deseas que los diferentes servicios se comuniquen y manejen los mensajes, la carga de trabajo, y la arquitectura de tu sistema distribuido.


La forma más habitual de utilizar RabbitMQ en arquitecturas de microservicios es la **Cola de Trabajo (Work Queue)**, también conocida como **Cola de Tareas en Segundo Plano**. Este patrón es el más común porque permite:

1. **Desacoplamiento de Servicios**: Los servicios pueden enviar tareas a la cola para que otros servicios las procesen sin necesidad de una comunicación directa. Esto facilita la escalabilidad y el mantenimiento del sistema.

2. **Procesamiento Asincrónico**: Las tareas pueden ser procesadas en segundo plano sin bloquear la ejecución del servicio que las envió. Esto es útil para operaciones que pueden llevar mucho tiempo, como procesamiento de imágenes, envío de correos electrónicos, generación de informes, etc.

3. **Distribución de Carga**: RabbitMQ distribuye automáticamente las tareas entre los consumidores disponibles, lo que permite balancear la carga de trabajo y mejorar la utilización de los recursos del sistema.

4. **Escalabilidad Horizontal**: Es fácil añadir más consumidores para procesar tareas cuando la carga aumenta, simplemente ejecutando más instancias del servicio consumidor.

### Ejemplo de Uso Típico de la Cola de Trabajo con RabbitMQ

Supongamos un sistema de procesamiento de imágenes donde un servicio web recibe las solicitudes para procesar imágenes y los trabajadores las procesan:

1. **Productor (Producer)**: Un servicio web que recibe una solicitud para procesar una imagen y envía un mensaje a la cola de RabbitMQ con los detalles de la imagen (por ejemplo, la ubicación del archivo).

2. **Cola de Trabajo (Work Queue)**: Una cola en RabbitMQ que almacena las tareas de procesamiento de imágenes.

3. **Consumidores (Consumers)**: Múltiples instancias de un servicio de procesamiento de imágenes que están suscritas a la cola y procesan las imágenes de manera concurrente.

#### Código Simplificado para una Cola de Trabajo con RabbitMQ

**Producer (Productor):**

```go
package main

import (
	"log"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"image_processing_queue", // Queue name
		true,   // Durable
		false,  // Delete when unused
		false,  // Exclusive
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	body := "Image details or path"
	err = ch.Publish(
		"",     // Exchange
		q.Name, // Routing key
		false,  // Mandatory
		false,  // Immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Printf(" [x] Sent %s", body)
}
```

**Consumer (Consumidor):**

```go
package main

import (
	"log"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"image_processing_queue", // Queue name
		true,   // Durable
		false,  // Delete when unused
		false,  // Exclusive
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer
		true,   // Auto-Ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			// Aquí se procesaría la imagen
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
```

### Conclusión

La **Cola de Trabajo (Work Queue)** es la forma más habitual de utilizar RabbitMQ debido a su capacidad para procesar tareas de forma asincrónica, distribuir la carga de trabajo entre múltiples consumidores y mejorar la escalabilidad del sistema. Este patrón es ideal para aplicaciones que requieren procesamiento en segundo plano, manejo de cargas de trabajo distribuidas, y sistemas desacoplados donde la fiabilidad y la persistencia de mensajes son importantes.