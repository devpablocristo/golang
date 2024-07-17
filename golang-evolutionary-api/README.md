# API Evolutiva

## Paso 1: API Simple con Servidor HTTP

**Temas cubiertos:**
1. **Paquete main**:
   - Estructura básica de un programa en Golang.
   - Función `main` como punto de entrada.

2. **Manejo de paquetes**:
   - Importación de paquetes (`fmt` y `net/http`).

3. **Definición de rutas y manejadores**:
   - Uso de `http.HandleFunc` para definir rutas.
   - Creación de manejadores de solicitudes HTTP.

4. **Servidor HTTP básico**:
   - Uso de `http.ListenAndServe` para iniciar un servidor HTTP en el puerto 8080.

5. **Escritura de respuestas HTTP**:
   - Uso de `http.ResponseWriter` para enviar respuestas al cliente.
   - Uso de `fmt.Fprintf` para escribir la respuesta.

## Paso 2: Múltiples Rutas en el Servidor HTTP

**Temas adicionales cubiertos:**
1. **Definición de múltiples rutas**:
   - Uso de `http.HandleFunc` para definir varias rutas (e.g., `/hello` y `/bye`).

2. **Creación de múltiples manejadores**:
   - Definición de varias funciones manejadoras (`hello` y `bye`).

3. **Respuesta con diferentes mensajes**:
   - Envío de diferentes respuestas según la ruta solicitada.

## Paso 3: Manejo de Métodos HTTP

**Temas adicionales cubiertos:**
1. **Manejo de múltiples métodos HTTP**:
   - Uso del campo `r.Method` del objeto `http.Request` para identificar el método HTTP de la solicitud.
   - Implementación de lógica específica para cada método HTTP (GET, POST, PUT, DELETE).

2. **Uso de switch para métodos HTTP**:
   - Estructuración del código con `switch` para manejar diferentes métodos HTTP de manera ordenada y clara.

3. **Manejo de errores HTTP**:
   - Uso de `http.Error` para enviar respuestas de error apropiadas cuando se utiliza un método HTTP no soportado.
   - Envío de códigos de estado HTTP (`http.StatusMethodNotAllowed`) para indicar que un método HTTP no es permitido.

## Paso 4: Uso del Framework Gin para Manejo de Rutas y Métodos HTTP

**Temas adicionales cubiertos:**
1. **Introducción al Framework Gin**:
   - Configuración básica del framework Gin.
   - Beneficios de usar un framework en lugar del paquete `net/http` estándar.

2. **Definición de rutas con Gin**:
   - Uso de `router.GET`, `router.POST`, `router.PUT`, y `router.DELETE` para definir rutas y métodos HTTP.
   - Comparación entre la definición de rutas en `net/http` y Gin.

3. **Creación de manejadores con Gin**:
   - Uso de `gin.Context` para manejar solicitudes y respuestas.
   - Métodos de `gin.Context` para enviar respuestas (`c.String`).

4. **Inicialización y ejecución del servidor con Gin**:
   - Uso de `router.Run` para iniciar el servidor en un puerto específico.

## Paso 5: Manejo de Rutas y Métodos HTTP con Parámetros

**Temas adicionales cubiertos:**
1. **Manejo de parámetros en la URL**:
   - Uso de `c.Param` para acceder a parámetros en la URL.
   - Definición de rutas con parámetros dinámicos.

2. **Validación de entradas**:
   - Validación de parámetros de entrada.
   - Manejo de errores relacionados con parámetros inválidos o ausentes.

## Paso 6: Manejo de Datos en el Cuerpo de la Solicitud (Request Body)

**Temas adicionales cubiertos:**
1. **Manejo de datos en el cuerpo de la solicitud**:
   - Acceso a los datos enviados en el cuerpo de la solicitud (Request Body).
   - Manejo de datos enviados en solicitudes POST.

2. **Deserialización de JSON**:
   - Cómo leer y deserializar datos JSON del cuerpo de la solicitud.
   - Validación de datos recibidos.

## Paso 7: Manejo de Datos en el Cuerpo de la Solicitud Usando `io.ReadAll`

**Temas adicionales cubiertos:**
1. **Lectura directa del cuerpo de la solicitud**:
   - Uso de `io.ReadAll` para leer datos directamente del cuerpo de la solicitud (Request Body).
   - Manejo de errores al leer el cuerpo de la solicitud.

2. **Conversión de datos a cadenas**:
   - Conversión de datos leídos del cuerpo de la solicitud a cadenas (`string(body)`).

## Paso 8: Logging y Validación Mejorada de JSON

**Temas adicionales cubiertos:**
1. **Logging**:
   - Uso de `fmt.Println` y `log.Println` para imprimir mensajes en la consola.
   - Uso de `log.Fatal` para manejar errores fatales al iniciar el servidor.

2. **Validación mejorada de JSON**:
   - Uso de `c.BindJSON` para deserializar JSON.
   - Validación de la presencia de campos específicos en el JSON.
   - Manejo de errores cuando los campos requeridos están ausentes.

## Paso 9: Refactorización con Métodos de una Estructura

**Temas adicionales cubiertos:**
1. **Refactorización con métodos de una estructura**:
   - Uso de una estructura (`handler`) para agrupar métodos manejadores.
   - Creación de métodos para la estructura para manejar solicitudes.

2. **Inicialización y uso de la estructura manejadora**:
   - Creación de una función `newHandler` para inicializar la estructura manejadora.
   - Asignación de métodos de la estructura como manejadores de rutas.

## Paso 10: Introducción a Use Cases y Entidades de Dominio

**Temas adicionales cubiertos:**
1. **Introducción a los casos de uso (Use Cases)**:
   - Creación de una capa de casos de uso para manejar la lógica de negocio.
   - Separación de la lógica de negocio de los manejadores HTTP.

2. **Manejo de entidades de dominio**:
   - Definición y uso de una entidad de dominio (`item`).
   - Creación y manipulación de instancias de entidades de dominio.

3. **Manejo avanzado de errores**:
   - Definición de errores globales (e.g., `errNotFound`).
   - Manejo y propagación de errores en la capa de casos de uso y los manejadores.

4. **Rutas y métodos adicionales**:
   - Implementación de rutas para operaciones de CRUD (Create, Read).
   - Manejo de solicitudes POST y GET para crear y listar ítems.

## Paso 11: Introducción a la Capa de Repositorio

**Temas adicionales cubiertos:**
1. **Introducción a la capa de repositorio**:
   - Creación de una capa de repositorio para manejar la persistencia de datos.
   - Separación de la lógica de acceso a datos de la lógica de negocio.

2. **Integración de la capa de repositorio con los casos de uso**:
   - Modificación de los casos de uso para interactuar con la capa de repositorio.
   - Propagación de errores desde la capa de repositorio a la capa de casos de uso.

3. **Validación de datos en la capa de repositorio**:
   - Validación de datos antes de guardar en la capa de repositorio.
   - Manejo de errores específicos en la capa de repositorio.

## Paso 12: Introducción a la Abstracción de Interfaces (Ports)

**Temas adicionales cubiertos:**
1. **Introducción a la abstracción de interfaces (Ports)**:
   - Creación de interfaces para abstraer la capa de casos de uso.
   - Implementación de las interfaces por parte de la capa de casos de uso.
   - Uso de interfaces para inyección de dependencias.

2. **Separación de la lógica de negocio y la infraestructura**:
   - Definición de interfaces para los casos de uso.
   - Implementación concreta de las interfaces por parte de los casos de uso.
   - Inyección de dependencias utilizando las interfaces.

## Paso 13: Organización del Código en Paquetes

**Temas adicionales cubiertos:**
1. **Organización del código en paquetes**:
   - Separación de la lógica en paquetes específicos (`repository`, `handler`, `usecase`, `domain`).
   - Creación de estructuras de directorios para organizar mejor el código.

2. **Uso de importaciones entre paquetes**:
   - Cómo importar y utilizar código de otros paquetes.
   - Definición de dependencias claras entre las capas de la aplicación.

3. **Inyección de dependencias a través de paquetes**:
   - Uso de funciones constructoras para inicializar estructuras con sus dependencias.
   - Inyección de dependencias utilizando interfaces y paquetes específicos.

## Paso 14: Refactorización y Modularización Avanzada

**Temas adicionales cubiertos:**
1. **Refactorización y modularización avanzada**:
   - Separación del código en módulos más pequeños y específicos.
   - Uso de un paquete `config` para definir errores globales.
   - Reorganización de los paquetes y módulos para mejorar la estructura del proyecto.

2. **Separación de las capas del proyecto**:
   - Creación de diferentes paquetes para entidades de

 dominio, casos de uso y repositorios.
   - Definición clara de interfaces y dependencias entre los módulos.

3. **Configuración y uso del paquete `config`**:
   - Definición de errores globales y configuración en un paquete separado.
   - Importación y uso de configuraciones globales desde el paquete `config`.

## Paso 15: Integración con MySQL y Refactorización Avanzada

**Nota: el map se pasa a llamar de `Repository` a `MapRepository` por diferenciarlo de `Mysql`.**

**Temas adicionales cubiertos:**
1. **Integración con MySQL**:
   - Configuración de un cliente MySQL.
   - Creación de un repositorio MySQL.
   - Uso de SQL para las operaciones CRUD.

2. **Refactorización y modularización avanzada**:
   - Separación de lógica de configuración de bases de datos.
   - Refactorización de casos de uso, repositorios y controladores (handlers).

## Paso 16: Implementación Completa de Operaciones CRUD

**Temas adicionales cubiertos:**
1. **Implementación completa de operaciones CRUD (Crear, Leer, Actualizar, Eliminar)**:
   - Adición de métodos para actualizar y eliminar elementos.
   - Integración de las nuevas operaciones en los casos de uso y los repositorios.

2. **Ampliación de la API para soportar operaciones CRUD**:
   - Adición de rutas para actualizar y eliminar elementos.
   - Manejo de solicitudes PUT y DELETE en los controladores.