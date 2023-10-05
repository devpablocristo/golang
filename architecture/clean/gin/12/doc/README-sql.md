# Golang y SQL

Es una buena práctica tener una carpeta "db" y otra carpeta "migrations" en un proyecto Go para separar claramente el esquema original de la base de datos y los archivos de migración.

La carpeta "db" es donde se almacena el archivo schema.sql que contiene el esquema original de la base de datos. Este archivo generalmente se usa para crear la base de datos inicial y se ejecuta una sola vez al inicio del proyecto.

La carpeta "migrations", por otro lado, es donde se almacenan los archivos de migración que contienen los cambios en el esquema de la base de datos que se aplican a medida que se desarrolla el proyecto. Estos archivos de migración se utilizan para actualizar el esquema de la base de datos a nuevas versiones o para realizar cambios en la estructura de la base de datos.

La convención de tener una carpeta "db" separada de una carpeta "migrations" es común en aplicaciones Go que utilizan un ORM (Mapeador objeto-relacional) como GORM, ya que permite automatizar la tarea de modificar el esquema de la base de datos de manera programática.

En resumen, se recomienda tener una carpeta "db" para almacenar el archivo schema.sql que contiene el esquema original de la base de datos y una carpeta "migrations" para almacenar los archivos de migración que contienen los cambios en el esquema de la base de datos en un proyecto Go. Esta convención de organización de proyectos ayuda a mantener una estructura clara y organizada para la base de datos del proyecto.
