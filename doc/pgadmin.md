## Documentación pgAdmin

### Conexión con PostgreSQL

pgAdmin no se conecta automáticamente a tu servidor PostgreSQL. Necesitas configurarlo manualmente. Aquí te explico cómo hacerlo:

#### Paso 1: Acceder a pgAdmin

1. Abre pgAdmin en tu navegador en `http://localhost:8081` (o el puerto que hayas configurado).

#### Paso 2: Iniciar Sesión en pgAdmin

1. Usa las credenciales configuradas en el archivo `docker-compose.yml`:
   - **Correo electrónico**: admin@admin.com
   - **Contraseña**: admin

#### Paso 3: Configurar una Nueva Conexión al Servidor PostgreSQL

1. **Añadir un nuevo servidor**:
   - En el panel de la izquierda, haz clic derecho en "Servers" y selecciona "Create" -> "Server...".

2. **Configurar la conexión del servidor**:
   - En la pestaña "General":
     - **Name**: `Postgres Local`
   - En la pestaña "Connection":
     - **Host name/address**: `postgres` (nombre del servicio en `docker-compose.yml`).
     - **Port**: `5432`.
     - **Maintenance database**: `my_db`.
     - **Username**: `admin`.
     - **Password**: `admin`.

3. **Guardar la configuración**:
   - Haz clic en "Save".

#### Paso 4: Verificar la Conexión

1. Una vez guardada la configuración, deberías ver tu nuevo servidor en el panel de la izquierda.
2. Haz clic en él para expandir y ver las bases de datos y otros objetos dentro de PostgreSQL.

### Ejemplo Visual de Configuración

#### General

- **Name**: `Postgres Local`

#### Connection

- **Host name/address**: `postgres`
- **Port**: `5432`
- **Maintenance database**: `my_db`
- **Username**: `admin`
- **Password**: `admin`

### Instrucciones para Crear la Base de Datos y la Tabla en pgAdmin

#### Paso 1: Crear la Base de Datos

1. En pgAdmin, abre el Query Tool y ejecuta:
   ```sql
   CREATE DATABASE my_db;
   ```

#### Paso 2: Crear el Usuario y Otorgar Permisos

1. En el Query Tool, ejecuta:
   ```sql
   CREATE USER admin WITH PASSWORD 'admin';
   GRANT ALL PRIVILEGES ON DATABASE my_db TO admin;
   ```

#### Paso 3: Crear la Tabla

1. Conéctate a `my_db` en el Query Tool y ejecuta:
   ```sql
   CREATE TABLE events (
       id UUID PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       description TEXT,
       location VARCHAR(255),
       start_time TIMESTAMP NOT NULL,
       end_time TIMESTAMP,
       category VARCHAR(50),
       creator_id UUID NOT NULL,
       is_public BOOLEAN NOT NULL DEFAULT true,
       is_recurring BOOLEAN NOT NULL DEFAULT false,
       series_id UUID,
       status VARCHAR(50) NOT NULL
   );
   ```

### Notas Importantes

- Sigue estos pasos en orden para evitar errores. Crear la base de datos, el usuario y la tabla en una sola ejecución puede causar errores debido a restricciones de transacciones en PostgreSQL.

## Ver las Tablas de una Base de Datos en pgAdmin

### Pasos

1. Abre pgAdmin y conéctate al servidor.
2. Expande la base de datos `my_db`.
3. Expande `Schemas` -> `public` -> `Tables` para ver todas las tablas.

Siguiendo estos pasos, podrás ver y gestionar las tablas en tu base de datos usando pgAdmin.


### Logging

pgAdmin muestra logs de forma excesiva. Para silenciarlos, realicé los siguientes pasos:

Para silenciar los logs de pgAdmin (o de cualquier servicio en un contenedor Docker), puedes configurar el nivel de logging o redirigir los logs a `/dev/null`. Existen varias opciones para hacerlo, utilice la siguiente:

#### Redirigir los logs en Docker

Puedes redirigir la salida del contenedor a `/dev/null`. Esto se puede hacer modificando el archivo `docker-compose.yml` o los comandos de `docker run`.

**Usando `docker-compose.yml`**

Edita tu archivo `docker-compose.yml` para que redirija los logs a `/dev/null`:

```yaml
version: '3.7'
services:
  pgadmin:
    image: dpage/pgadmin4
    environment:
      PGADMIN_DEFAULT_EMAIL: "admin@example.com"
      PGADMIN_DEFAULT_PASSWORD: "admin"
    ports:
      - "8081:80"
    logging:
      driver: "none" # <--- esta configuración
```
### Otras opciomes

1. Configurar los niveles de logging si es posible desde pgAdmin.
2. Configurar Nginx para desactivar los logs de acceso.
3. Configurar el logging globalmente en Docker.
