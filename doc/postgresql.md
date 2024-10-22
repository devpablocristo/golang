## Documentación PostgreSQL

### Resumen:

1. Correr container
2. Entrar al container por la terminal
3. Ejecutar: `$ psql -U postgres`
4. En mis contenedores, siempre uso:
    - **Superusuario**: `admin`
    - **Contraseña**: `admin`
    - **Usuario común**: `user`
    - **Contraseña**: `user`
5. Crear usuario y darle permisos:
    ```sql
    CREATE USER admin WITH PASSWORD 'admin';
    ```
6. Convertir en superusuario:
    ```sql
    ALTER USER admin WITH SUPERUSER;
    ```


En PostgreSQL, el nombre de usuario y la contraseña por defecto pueden ser configurados al crear los contenedores de Docker. Comúnmente, se usa el usuario `postgres` como superusuario por defecto, pero se pueden agregar usuarios adicionales con privilegios específicos.

### 1. Preparación del Entorno

1. **Configura tu archivo `config/.env`**: 
   Asegúrate de que el archivo `.env` en el directorio `config` contiene las variables necesarias para configurar la base de datos PostgreSQL. Por ejemplo:

   ```env
   POSTGRES_HOST=postgres
   POSTGRES_PORT=5432
   POSTGRES_DATABASE=my_database
   POSTGRES_USERNAME=admin
   POSTGRES_PASSWORD=admin
   ```

2. **Instala PostgreSQL usando Docker Compose**: 
   Define el servicio de PostgreSQL en tu archivo `docker-compose.yml` y asegúrate de incluir las credenciales en las variables de entorno. Por ejemplo:

   ```yaml
   postgres:
     image: postgres:latest
     container_name: postgres
     environment:
       POSTGRES_USER: admin
       POSTGRES_PASSWORD: admin
       POSTGRES_DB: my_database
     ports:
       - "5432:5432"
     volumes:
       - ./postgres_data:/var/lib/postgresql/data
     networks:
       - app-network
   ```

3. **Levanta el contenedor de PostgreSQL**: 
   Ejecuta el siguiente comando para iniciar el servicio de PostgreSQL con Docker Compose:

   ```bash
   docker-compose up -d
   ```

### 2. Creación Manual de la Base de Datos y Usuarios

1. **Accede al contenedor de PostgreSQL**:
   Una vez que el contenedor esté en ejecución, puedes acceder al contenedor de PostgreSQL con el siguiente comando:

   ```bash
   docker exec -it postgres bash
   ```

2. **Accede a la consola de PostgreSQL**:
   Dentro del contenedor, usa `psql` para conectarte a PostgreSQL como el superusuario `postgres`:

   ```bash
   psql -U postgres
   ```

3. **Crear usuario y asignar permisos**:
   Puedes crear un usuario normal y un superusuario dentro de PostgreSQL con los siguientes comandos:

   - Crear un nuevo usuario:
     ```sql
     CREATE USER user WITH PASSWORD 'user';
     ```

   - Convertir a un usuario en superusuario:
     ```sql
     ALTER USER admin WITH SUPERUSER;
     ```

4. **Verificar los usuarios existentes**:
   Usa el comando `\du` para listar los usuarios actuales en PostgreSQL:

   ```sql
   \du
   ```

### 3. Flujo para Administrar Usuarios y Permisos en PostgreSQL

1. **Correr el contenedor de PostgreSQL**:
   Asegúrate de que el contenedor de PostgreSQL esté corriendo. Si no lo está, inicia el contenedor con Docker Compose.

   ```bash
   docker-compose up -d
   ```

2. **Entrar al contenedor por la terminal**:
   Usa el siguiente comando para acceder al contenedor de PostgreSQL:

   ```bash
   docker exec -it postgres bash
   ```

3. **Conectar a PostgreSQL**:
   Una vez dentro del contenedor, conéctate a PostgreSQL usando el siguiente comando:

   ```bash
   psql -U postgres
   ```

4. **Gestión de usuarios**:
   - **Superusuario**: `admin`
     ```sql
     CREATE USER admin WITH PASSWORD 'admin';
     ALTER USER admin WITH SUPERUSER;
     ```

   - **Usuario común**: `user`
     ```sql
     CREATE USER user WITH PASSWORD 'user';
     ```

5. **Verificación y gestión**:
   - **Ver usuarios**:
     ```sql
     \du
     ```

6. **Salir de PostgreSQL y del contenedor**:
   - Para salir de PostgreSQL, usa `\q`.
   - Para salir del contenedor, usa `exit`. 