package repositories

import "database/sql"

// Repository handles DB operations relative to the role repository.
type Repository struct {
	DB *sql.DB
}
