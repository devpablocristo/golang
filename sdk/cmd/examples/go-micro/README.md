go-micro con: gin, logrus 

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