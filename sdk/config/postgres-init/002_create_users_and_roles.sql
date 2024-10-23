
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

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
