### Golang-migrate: Herramienta de Migración de Bases de Datos

**Golang-migrate** es una herramienta de migración de bases de datos escrita en Go. Permite gestionar las versiones de la estructura de una base de datos de manera eficiente y controlada. Con golang-migrate, puedes aplicar y revertir cambios estructurales en tu base de datos de forma segura, asegurando que todas las migraciones se ejecuten de manera consistente en todos los entornos.

Realiza un seguimiento de las migraciones aplicadas utilizando una tabla de control en la base de datos llamada schema_migrations. Esta tabla guarda el estado de cada migración aplicada, lo que permite a Golang Migrate saber qué migraciones ya se han ejecutado y cuáles faltan por aplicar.

### Creación de Archivos de Migración

Si no tienes ningún archivo de migraciones todavía, debes crear al menos uno para iniciar el proceso de migración de tu base de datos. Aquí tienes un conjunto de pasos y ejemplos para crear y manejar tus primeras migraciones usando golang-migrate y PostgreSQL en un contenedor Docker.

Es una buena práctica tener un archivo `up` y otro `down` para cada migración. El nombre del archivo debe seguir la nomenclatura definida o no será encontrado.

### Características Clave de Golang-migrate

1. **Soporte para Múltiples Bases de Datos**:
   - PostgreSQL
   - MySQL
   - SQLite
   - SQL Server
   - Cassandra
   - y muchas más.

2. **Control de Versiones**:
   - Cada migración tiene un número de versión asociado que permite llevar un seguimiento de los cambios aplicados a la base de datos.

3. **Compatibilidad con Diferentes Fuentes de Migración**:
   - Archivos locales
   - AWS S3
   - Google Cloud Storage
   - y otros.

4. **Comandos Básicos**:
   - `up`: Aplica todas las migraciones pendientes.
   - `down`: Revierte las migraciones aplicadas.
   - `goto`: Aplica o revierte migraciones hasta alcanzar una versión específica.
   - `migrate`: Aplica o revierte migraciones necesarias para alcanzar una versión específica.

### Ejemplo de Uso

Para utilizar golang-migrate, debes tener archivos de migración que definan los cambios estructurales de la base de datos. Estos archivos tienen una convención de nombres específica, como `0001_create_users_table.up.sql` y `0001_create_users_table.down.sql`, donde `up` define cómo aplicar el cambio y `down` cómo revertirlo.

#### Ejemplo de Archivo de Migración

**0001_create_users_table.up.sql**:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

**0001_create_users_table.down.sql**:
```sql
DROP TABLE users;
```

### Ejecución de Migraciones

Supongamos que tienes una base de datos PostgreSQL y quieres aplicar las migraciones:

1. **Instala golang-migrate**:
   ```sh
   brew install golang-migrate  # macOS
   # o para Linux
   curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz
   sudo mv migrate.linux-amd64 /usr/local/bin/migrate
   ```

2. **Aplica las Migraciones**:
   ```sh
   migrate -path ./migrations -database postgres://user:password@localhost:5432/mydb?sslmode=disable up
   ```

3. **Revierte las Migraciones**:
   ```sh
   migrate -path ./migrations -database postgres://user:password@localhost:5432/mydb?sslmode=disable down
   ```

### Beneficios de Usar Golang-migrate

- **Consistencia**: Asegura que todas las migraciones se ejecuten en el orden correcto y en todos los entornos.
- **Facilidad de Uso**: Simplifica la gestión de versiones de la base de datos.
- **Flexibilidad**: Compatible con múltiples bases de datos y fuentes de migración.

En resumen, golang-migrate es una herramienta poderosa y flexible para manejar migraciones de bases de datos en proyectos de desarrollo, asegurando que los cambios estructurales se apliquen de manera controlada y segura.