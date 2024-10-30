package authconn

import "database/sql"

// Función auxiliar para manejar sql.NullString
func getStringFromNullString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}
