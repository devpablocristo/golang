\connect users_db;

-- Primero las tablas sin dependencias
CREATE TABLE IF NOT EXISTS persons (
    uuid UUID PRIMARY KEY,
    cuil VARCHAR(20) NOT NULL UNIQUE,
    dni VARCHAR(20) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    nationality VARCHAR(50),
    email VARCHAR(255),
    phone VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS companies (
    uuid UUID PRIMARY KEY,
    cuit VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255),
    address VARCHAR(255),
    phone VARCHAR(50),
    email VARCHAR(255),
    industry VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Luego las tablas con dependencias
CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    person_uuid UUID,
    company_uuid UUID,
    user_type VARCHAR(50) NOT NULL CHECK (user_type IN ('person', 'company')),
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    logged_at TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_person FOREIGN KEY(person_uuid) REFERENCES persons(uuid) ON DELETE SET NULL,
    CONSTRAINT fk_company FOREIGN KEY(company_uuid) REFERENCES companies(uuid) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_uuid UUID,
    role_id INT,
    PRIMARY KEY(user_uuid, role_id),
    CONSTRAINT fk_user FOREIGN KEY(user_uuid) REFERENCES users(uuid) ON DELETE CASCADE,
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT,
    permission_id INT,
    PRIMARY KEY(role_id, permission_id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id) ON DELETE CASCADE,
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);