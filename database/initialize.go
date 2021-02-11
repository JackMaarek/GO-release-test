package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// DBDetails contains the necessary informations to connect to the database.
type DBDetails struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     uint16
	Engine   string
}

// Connect opens an SQL connection to the DB and pings it.
func Connect(details *DBDetails) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		details.Username,
		details.Password,
		details.Host,
		details.Port,
		details.Name,
	)

	db, err := sql.Open(details.Engine, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
