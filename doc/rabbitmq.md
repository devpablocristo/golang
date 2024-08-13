# RabbitMQ

RabbitMQ es un **broker de mensajes** que utiliza el protocolo AMQP (Advanced Message Queuing Protocol) para facilitar la comunicación asíncrona entre aplicaciones o servicios mediante el envío y recepción de mensajes. Este documento explica cómo funciona RabbitMQ, detalla los roles de productores y consumidores, compara RabbitMQ con REST y aborda la implementación de un CRUD utilizando RabbitMQ.

## Conceptos Básicos de RabbitMQ

### Broker de Mensajes
- **Intermediario:** RabbitMQ actúa como un intermediario entre aplicaciones, permitiendo la comunicación asíncrona entre productores (aplicaciones que envían mensajes) y consumidores (aplicaciones que los reciben).
- **Comunicación Asincrónica:** RabbitMQ permite que los productores y consumidores operen de manera independiente, sin necesidad de estar activos simultáneamente.

### Colas
- **Almacenamiento de Mensajes:** Los mensajes se almacenan en **colas** dentro de RabbitMQ hasta que un consumidor los procesa.
- **Persistencia:** Las colas pueden ser persistentes, sobreviviendo a reinicios del broker, o transitorias, desapareciendo si el broker se reinicia.

### Exchanges (Intercambio)
- **Distribución de Mensajes:** Los mensajes son enviados a un **exchange**, que decide a qué cola(s) deben ser enviados según las reglas de enrutamiento (bindings).
- **Tipos de Exchange:** Los tipos de exchange (direct, topic, fanout, headers) determinan cómo se enrutan los mensajes a las colas.

### Bindings
- **Conexión entre Exchange y Colas:** Los bindings definen las reglas para que los mensajes se enruten desde el exchange hasta la cola correspondiente.

## Productor (Producer)

Un productor es cualquier aplicación o servicio que envía mensajes a RabbitMQ. Los productores envían mensajes a un exchange sin necesidad de conocer la estructura de la cola o los consumidores.

### Rol del Productor
- **Publicación de Mensajes:** Publica mensajes en un exchange específico.
- **Desacoplamiento Completo:** No espera respuestas de los consumidores, lo que permite un desacoplamiento entre la producción y el consumo de mensajes.

### Funcionamiento de RabbitMQ como Productor

El proceso general para un productor en RabbitMQ incluye:

#### Paso 1: Conexión al Servidor RabbitMQ
El productor establece una conexión con el servidor RabbitMQ. Esta conexión es esencial para que el productor pueda interactuar con el servidor RabbitMQ.

```go
conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
if err != nil {
    log.Fatalf("Failed to connect to RabbitMQ: %v", err)
}
defer conn.Close()
```

#### Paso 2: Crear un Canal de Comunicación
Una vez establecida la conexión, el productor abre un canal. Un canal en RabbitMQ es un medio virtual a través del cual se envían y reciben mensajes. Es importante porque permite múltiples operaciones dentro de una sola conexión.

```go
ch, err := conn.Channel()
if err != nil {
    log.Fatalf("Failed to open a channel: %v", err)
}
defer ch.Close()
```

#### Paso 3: Declarar una Cola
Antes de enviar mensajes, el productor debe asegurarse de que la cola a la que quiere enviar los mensajes exista. Si no existe, se crea. Esto garantiza que el mensaje tenga un destino donde ser almacenado hasta que un consumidor lo procese.

```go
q, err := ch.QueueDeclare(
    "hello",  // Nombre de la cola
    false,    // durable
    false,    // delete when unused
    false,    // exclusive
    false,    // no-wait
    nil,      // arguments
)
if err != nil {
    log.Fatalf("Failed to declare a queue: %v", err)
}
```

#### Paso 4: Publicar un Mensaje en la Cola
El productor envía el mensaje a la cola especificada a través del exchange. Durante esta etapa, el productor no necesita preocuparse por la estructura de la cola ni por los consumidores.

```go
body := "Hello World!"
err = ch.Publish(
    "",        // exchange
    q.Name,    // routing key (nombre de la cola)
    false,     // mandatory
    false,     // immediate
    amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(body),
    })
if err != nil {
    log.Fatalf("Failed to publish a message: %v", err)
}
```

#### Paso 5: Confirmación de Envío (Opcional)
El productor puede esperar una confirmación de que el mensaje fue recibido correctamente por RabbitMQ. Esta etapa es opcional, pero puede ser útil en escenarios donde se necesita garantizar que el mensaje se ha entregado correctamente.

```go
err = ch.Confirm(false)
if err != nil {
    log.Fatalf("Failed to put channel into confirm mode: %v", err)
}

confirms := ch.NotifyPublish(make(chan amqp.Confirmation, 1))

select {
case confirm := <-confirms:
    if confirm.Ack {
        log.Println("Message sent successfully")
    } else {
        log.Println("Failed to deliver message")
    }
case <-time.After(5 * time.Second):
    log.Println("No confirmation received, message delivery uncertain")
}
```

#### Paso 6: Cerrar Canal y Conexión
Finalmente, el productor cierra el canal y la conexión para liberar recursos y asegurar que no haya fugas de conexión.

```go
defer ch.Close()
defer conn.Close()
```

## Consumidor (Consumer)

Un consumidor es cualquier aplicación o servicio que recibe y procesa mensajes de RabbitMQ. Los consumidores se suscriben a una cola específica y procesan los mensajes a medida que llegan.

### Rol del Consumidor
- **Consumo de Mensajes:** Se conecta a una cola y consume mensajes.
- **Reconocimiento de Mensajes:** Puede reconocer (ack) los mensajes, indicando que se procesaron correctamente.

### Funcionamiento de RabbitMQ como Consumidor

El proceso general para un consumidor en RabbitMQ incluye:

#### Paso 1: Conexión al Servidor RabbitMQ
Al igual que el productor, el consumidor establece una conexión con el servidor RabbitMQ para poder interactuar con él.

```go
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
if err != nil {
    log.Fatalf("Failed to connect to RabbitMQ: %v", err)
}
defer conn.Close()
```

#### Paso 2: Crear un Canal de Comunicación
El consumidor abre un canal en el que recibirá los mensajes desde RabbitMQ. Este canal es esencial para la comunicación continua entre el consumidor y el servidor.

```go
ch, err := conn.Channel()
if err != nil {
    log.Fatalf("Failed to open a channel: %v", err)
}
defer ch.Close()
```

#### Paso 3: Declarar una Cola
El consumidor se asegura de que la cola desde la que va a consumir mensajes exista. Si la cola no existe, se crea. Esto es importante porque el consumidor necesita un lugar desde donde leer los mensajes.

```go
q, err := ch.QueueDeclare(
    "hello", // nombre de la cola
    false,   // durable
    false,   // auto delete
    false,   // exclusive
    false,   // no-wait
    nil,     // argumentos
)
if err != nil {
    log.Fatalf("Failed to declare a queue: %v", err)
}
```

#### Paso 4: Consumir Mensajes
El consumidor se suscribe a la cola y empieza a consumir mensajes. Cada mensaje que llega a la cola se procesa de acuerdo con la lógica del consumidor.

```go
msgs, err := ch.Consume(
    q.Name, // nombre de la cola
    "",     // consumer tag
    true,   // auto-acknowledge
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // argumentos adicionales
)
if err != nil {
    log.Fatalf("Failed to register a consumer: %v", err)
}

for msg := range msgs {
    log.Printf("Received a message: %s", msg.Body)
    // Procesar el mensaje aquí
}
```

#### Paso 5: Reconocer Mensajes (Opcional)
El consumidor puede reconocer los mensajes (ack), lo que indica a RabbitMQ que el mensaje ha sido procesado correctamente y puede ser eliminado de la cola.

#### Paso 6: Cerrar Canal y Conexión
Después de procesar los mensajes, el consumidor cierra el canal y la conexión.

```go
defer ch.Close()
defer conn.Close()
```

## Diferencias con REST

RabbitMQ y REST son dos paradigmas de comunicación diferentes:

1. **Sincronía vs. Asincronía:**
   - **REST:** Protocolo síncrono basado en solicitudes HTTP, donde el cliente envía una solicitud y espera una respuesta.
   - **RabbitMQ:** Sistema asincrónico, donde los productores envían mensajes sin esperar una respuesta, y los consumidores procesan los mensajes de manera independiente.

2. **Desacoplamiento:**
   - **REST:** Los clientes dependen de la disponibilidad del servidor para recibir una respuesta.
   - **RabbitMQ:** Los productores y consumidores están completamente desacoplados. Los consumidores pueden estar desconectados o inactivos cuando se envía un mensaje y simplemente procesan los mensajes cuando están listos.

3. **Persistencia y Tolerancia

 a Fallos:**
   - **REST:** La comunicación es transitoria. Si una solicitud falla, generalmente se pierde a menos que se implemente una lógica de reintentos.
   - **RabbitMQ:** Los mensajes pueden ser persistentes y almacenarse hasta que se procesen, lo que proporciona una mayor tolerancia a fallos.

4. **Modelo de Interacción:**
   - **REST:** Utiliza un modelo de solicitud-respuesta, adecuado para operaciones directas e inmediatas.
   - **RabbitMQ:** Utiliza un modelo de mensajes basados en colas, adecuado para flujos de trabajo distribuidos, procesamiento de tareas en segundo plano y situaciones donde la latencia no es crítica.

## Implementación de CRUD con RabbitMQ

Para implementar un CRUD (Create, Read, Update, Delete) completo utilizando RabbitMQ, es necesario estructurar los mensajes para que indiquen claramente qué operación debe realizarse. Esto se puede lograr de diferentes maneras, como utilizando un campo en el mensaje que especifique la acción (`create`, `read`, `update`, `delete`) o utilizando diferentes exchanges o colas para cada operación.

## Formas Comunes de Usar RabbitMQ

RabbitMQ se puede usar de diferentes maneras, dependiendo de los requisitos del sistema:

1. **Mensajes de Comando:** Contienen un comando o acción que debe ejecutarse, como `create_user`.
2. **Mensajes de Evento:** Representan eventos que han ocurrido en el sistema, como `user_created`.
3. **Colas Dedicadas por Acción:** Cada acción tiene su propia cola, lo que simplifica el procesamiento en el lado del consumidor.
4. **Routing Keys y Exchanges:** Utilizan routing keys para enrutar mensajes a colas específicas, útil para lógica de enrutamiento más compleja.

## Patrones y Roles en RabbitMQ

1. **Productor (Producer):** Envía mensajes a un exchange en RabbitMQ, sin preocuparse por los consumidores.
2. **Consumidor (Consumer):** Recibe y procesa mensajes de una cola específica.
3. **Exchange Configurator:** Define cómo se enrutan los mensajes a las colas.
4. **Colas de Letra Muerta (DLQ):** Capturan mensajes que no pueden ser procesados correctamente.
5. **Patrones RPC (Remote Procedure Call):** Permiten a un servicio solicitar la ejecución de un procedimiento en otro servicio y recibir una respuesta.
6. **Retransmisión de Eventos (Event Relay):** Toma eventos de un sistema y los envía a otros sistemas o servicios.

## Consideraciones Finales

Al elegir el enfoque adecuado para tu aplicación, considera los siguientes aspectos:

1. **Requisitos de Negocio:** ¿Necesitas un sistema desacoplado y reactivo o uno con lógica de negocio clara y definida?
2. **Escalabilidad:** ¿Esperas que el sistema crezca significativamente? ¿Cómo manejarás el aumento en la carga de trabajo?
3. **Complejidad de Enrutamiento:** ¿Qué tan compleja es la lógica de enrutamiento necesaria para tu aplicación?
4. **Mantenibilidad:** ¿Qué enfoque es más fácil de mantener y expandir a medida que cambian los requisitos?

RabbitMQ facilita la construcción de sistemas distribuidos robustos y flexibles, adaptándose a diferentes casos de uso. La elección de la estrategia correcta dependerá de los objetivos específicos de tu proyecto y de las características de los sistemas que quieras integrar.