package config

import (
	"database/sql"
	"fmt"
	// github.com/lib/pq is part of connecting to postgres
	_ "github.com/lib/pq"
	"log"
)

// Database is the database connection for the system to use
var Database *sql.DB

// NewDatabaseConnection creates a connection to the database for the system
func NewDatabaseConnection() (*sql.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Config.Database.Host,
		Config.Database.Port,
		Config.Database.User,
		Config.Database.Pass,
		Config.Database.DBname)

	Database, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Panic("Can't Connect to Database", err)
	}

	err = Database.Ping()
	if err != nil {
		log.Panic("Can't Ping Database", err)
	}

	return Database, err
}
