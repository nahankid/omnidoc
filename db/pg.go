package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // comment
)

// PGConn is used to open the database connection
type PGConn struct{}

// GetConnection returns database connection
func (p *PGConn) GetConnection() (db *gorm.DB, err error) {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	return gorm.Open("postgres", dbURI)
}
