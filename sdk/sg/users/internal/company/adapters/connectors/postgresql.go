package companyconn

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	sdkpg "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool"
	sdkpgports "github.com/devpablocristo/golang/sdk/pkg/databases/sql/postgresql/pgxpool/ports"

	entities "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/entities"
	ports "github.com/devpablocristo/golang/sdk/sg/users/internal/company/core/ports"
)

type PostgreSQL struct {
	repository sdkpgports.Repository
}

func NewPostgreSQL() (ports.Repository, error) {
	r, err := sdkpg.Bootstrap("USERS_DB")
	if err != nil {
		return nil, fmt.Errorf("bootstrap error: %w", err)
	}
	return &PostgreSQL{repository: r}, nil
}

// Create inserts a new company into the database
func (r *PostgreSQL) Create(ctx context.Context, company *entities.Company) error {
	query := `
        INSERT INTO companies (
            uuid, cuit, name, legal_name, address, phone, email, industry, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
        )
    `
	_, err := r.repository.Pool().Exec(ctx, query,
		company.UUID, company.Cuit, company.Name, company.LegalName,
		company.Address, company.Phone, company.Email, company.Industry,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" { // unique_violation
			return fmt.Errorf("company already exists")
		}
		return err
	}
	return nil
}

// FindByID searches for a company by their UUID
func (r *PostgreSQL) FindByID(ctx context.Context, id string) (*entities.Company, error) {
	company := &entities.Company{}
	query := `
        SELECT uuid, cuit, name, legal_name, address, phone, email, industry, created_at, updated_at, deleted_at
        FROM companies 
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	err := r.repository.Pool().QueryRow(ctx, query, id).Scan(
		&company.UUID,
		&company.Cuit,
		&company.Name,
		&company.LegalName,
		&company.Address,
		&company.Phone,
		&company.Email,
		&company.Industry,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return company, nil
}

// Update updates an existing company
func (r *PostgreSQL) Update(ctx context.Context, company *entities.Company) error {
	query := `
        UPDATE companies 
        SET cuit = $2, name = $3, legal_name = $4, address = $5, phone = $6, email = $7, industry = $8, updated_at = CURRENT_TIMESTAMP
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	result, err := r.repository.Pool().Exec(ctx, query,
		company.UUID, company.Cuit, company.Name, company.LegalName, company.Address,
		company.Phone, company.Email, company.Industry,
	)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("company not found")
	}

	return nil
}

// SoftDelete marks a company as deleted
func (r *PostgreSQL) SoftDelete(ctx context.Context, id string) error {
	query := `
        UPDATE companies 
        SET deleted_at = CURRENT_TIMESTAMP
        WHERE uuid = $1 AND deleted_at IS NULL
    `
	result, err := r.repository.Pool().Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("company not found")
	}

	return nil
}

// FindByCuit searches for a company by their CUIT
func (r *PostgreSQL) FindByCuit(ctx context.Context, cuit string) (*entities.Company, error) {
	company := &entities.Company{}
	query := `
        SELECT uuid, cuit, name, legal_name, address, phone, email, industry, created_at, updated_at, deleted_at
        FROM companies 
        WHERE cuit = $1 AND deleted_at IS NULL
    `
	err := r.repository.Pool().QueryRow(ctx, query, cuit).Scan(
		&company.UUID,
		&company.Cuit,
		&company.Name,
		&company.LegalName,
		&company.Address,
		&company.Phone,
		&company.Email,
		&company.Industry,
		&company.CreatedAt,
		&company.UpdatedAt,
		&company.DeletedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return company, nil
}
