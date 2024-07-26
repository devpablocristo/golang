## Documentación PostgreSQL

En PostgreSQL, no hay un nombre de usuario y contraseña por defecto. Sin embargo, cuando instalas PostgreSQL, suele crearse un usuario administrador llamado postgres sin contraseña, o bien con una contraseña que el usuario debe definir durante la instalación. En esta API se configura la contraseña
"root" para el super usuario "postgres".

Aquí hay algunos puntos clave:

### 1. Preparación del Entorno

1. **Configura tu archivo `.env`**: Asegúrate de que tu archivo `.env` en la raíz de tu proyecto contiene todas las variables necesarias para configurar la base de datos.

2. **Instala PostgreSQL**: Asegúrate de que PostgreSQL está instalado y funcionando en tu máquina. En caso crearlo con docker-compose.

### 2. Creación Manual de la Base de Datos y la Tabla

1. **Crea el script de bash `setup_db.sh`**: Coloca este script en el directorio `scripts` para crear la base de datos y la tabla utilizando las variables del archivo `.env`.
   
2. **Ejecuta el script**: Dale permisos de ejecución y ejecútalo para crear la base de datos y la tabla.

### 3. Configuración e Inicialización de la API

1. **Asegúrate de que tu archivo de configuración de la API está correcto**: Configura tu archivo `config.go` para leer las variables de entorno del archivo `.env`.

2. **Inicialización de la conexión a la base de datos**: Configura la inicialización de la conexión a la base de datos y la aplicación de migraciones en `postgres.go`.

### 4. Inicio de la API

1. **Compila y ejecuta tu aplicación**: Asegúrate de que tu aplicación esté configurada para usar el paquete `db` y `config` correctamente.

2. **Ejecuta la aplicación**: Asegúrate de tener las migraciones en el directorio `migrations` y ejecuta tu aplicación.
