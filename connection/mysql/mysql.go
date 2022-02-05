package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {

	var connectionString string
	log.Println("INIT MYSQL CONNECTION")
	connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		"root",
		"password",
		"localhost",
		"myrepo",
	)

	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Connection error: ", err)
	}

	if err = conn.Ping(); err != nil {
		log.Println("Failed verifies a connection to the database", nil)
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	db = conn

}

func CloseConnection() {
	db.Close()
}

func Connect() *sql.DB {
	return db
}
