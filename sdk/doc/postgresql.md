Para crear bases de datos en PostgreSQL mediante scripts de inicialización:

Hay varias formas de solucionarlo, una de ellas es:

1. **Usar PSQL con template1**
```sql
\connect template1;
CREATE DATABASE users_db;
```

Por ejemplo:

`001_create_users_db.sql`:
```sql
\connect template1;
CREATE DATABASE users_db;
```

`002_create_tables.sql`:
```sql
\connect users_db;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    ...
);
```

`003_insert_initial_data.sql`:
```sql
\connect users_db;
INSERT INTO users (name) VALUES ('John Doe');
```

De esta manera, aseguras que:
1. Primero se crea la base de datos (conectándote a template1)
2. Los scripts posteriores se ejecutan dentro de la base de datos que acabas de crear

Si necesitas trabajar con múltiples bases de datos, puedes crear todas las que necesites en el primer script y luego conectarte a la base de datos correspondiente en cada script posterior. Por ejemplo:

`001_create_databases.sql`:
```sql
\connect template1;
CREATE DATABASE users_db;
CREATE DATABASE products_db;
```

`002_create_users_tables.sql`:
```sql
\connect users_db;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100)
);
```

`003_create_products_tables.sql`:
```sql
\connect products_db;
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    price DECIMAL(10,2)
);
```

`004_insert_users_data.sql`:
```sql
\connect users_db;
INSERT INTO users (name) VALUES ('John Doe');
```

`005_insert_products_data.sql`:
```sql
\connect products_db;
INSERT INTO products (name, price) VALUES ('Product 1', 99.99);
```

Los scripts se ejecutarán en orden alfabético/numérico, por lo que:
1. Primero se crean todas las bases de datos
2. Luego puedes crear las tablas en cada base de datos
3. Finalmente puedes insertar los datos iniciales

Recuerda que puedes organizar tus scripts en subdirectorios dentro de `postgres-init` para mantenerlo más ordenado, por ejemplo:
```
postgres-init/
├── 001_create_databases.sql
├── users/
│   ├── 002_create_tables.sql
│   └── 004_insert_data.sql
└── products/
    ├── 003_create_tables.sql
    └── 005_insert_data.sql
```
No necesariamente deben ir en el mismo script. Puedes separarlas siempre y cuando mantengas el orden correcto de ejecución usando la numeración. Por ejemplo:

`001_create_users_db.sql`:
```sql
\connect template1;
CREATE DATABASE users_db;
```

`002_create_products_db.sql`:
```sql
\connect template1;
CREATE DATABASE products_db;
```

`003_create_users_tables.sql`:
```sql
\connect users_db;
CREATE TABLE users (...);
```

`004_create_products_tables.sql`:
```sql
\connect products_db;
CREATE TABLE products (...);
```

La ventaja de separarlas es:
- Mejor organización
- Más fácil de mantener
- Más claro qué script corresponde a cada base de datos
- Más fácil de desactivar una base de datos si temporalmente no la necesitas (simplemente renombrando o moviendo su script)

Una estructura de directorios podría ser:

```
postgres-init/
├── users/
│   ├── 001_create_users_db.sql
│   ├── 003_create_tables.sql
│   └── 005_insert_data.sql
└── products/
    ├── 002_create_products_db.sql
    ├── 004_create_tables.sql
    └── 006_insert_data.sql
```

Otra forma sera:

Es necesario agregar el prefijo alfabético a los directorios:

```
postgres-init/
├── a_users/
│   ├── 001_create_db.sql
│   ├── 002_create_tables.sql
│   └── 003_insert_data.sql
└── b_products/
    ├── 001_create_db.sql
    ├── 002_create_tables.sql
    └── 003_insert_data.sql
```

De esta manera, PostgreSQL ejecutará:
1. Todos los scripts de `a_users` en orden
2. Luego todos los scripts de `b_products` en orden

El prefijo alfabético (`a_`, `b_`) es necesario para controlar el orden de ejecución entre directorios.

O con números `1_` y `2_` en los directorios asegurarán que:

1. Primero se ejecuten todos los scripts de `1_users` en orden:
   - 001_create_db.sql
   - 002_create_tables.sql
   - 003_insert_data.sql

2. Luego todos los scripts de `2_products` en orden:
   - 001_create_db.sql
   - 002_create_tables.sql
   - 003_insert_data.sql

Esta estructura es:
- Clara y fácil de entender
- Mantiene todos los scripts de cada base de datos juntos
- El orden de ejecución es predecible
- Fácil de mantener y expandir si necesitas agregar más bases de datos (3_orders, 4_inventory, etc.)