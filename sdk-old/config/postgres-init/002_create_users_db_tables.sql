-- Conexión a la base de datos
\connect users_db;

-- Tabla `persons`
CREATE TABLE IF NOT EXISTS persons (
    uuid UUID PRIMARY KEY,
    cuil VARCHAR(20) NOT NULL UNIQUE,
    dni VARCHAR(20) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    nationality VARCHAR(50),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(50)
);

-- Índices adicionales
CREATE INDEX IF NOT EXISTS idx_persons_cuil ON persons(cuil);
CREATE INDEX IF NOT EXISTS idx_persons_email ON persons(email);

-- Tabla `permissions`
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

-- Tabla `roles`
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT
);

-- Tabla `users`
CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    person_uuid UUID,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    email_validated BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    logged_at TIMESTAMPTZ DEFAULT NULL,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT fk_person FOREIGN KEY(person_uuid) REFERENCES persons(uuid) ON DELETE SET NULL
);

-- Tabla `user_roles`
CREATE TABLE IF NOT EXISTS user_roles (
    user_uuid UUID,
    role_id INT,
    PRIMARY KEY(user_uuid, role_id),
    CONSTRAINT fk_user_roles_user FOREIGN KEY(user_uuid) REFERENCES users(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_user_roles_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE
);

-- Tabla `role_permissions`
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT,
    permission_id INT,
    PRIMARY KEY(role_id, permission_id),
    CONSTRAINT fk_role_permissions_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_role_permissions_permission FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);
