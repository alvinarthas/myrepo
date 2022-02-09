package mysql

import (
	"database/sql"
	"fmt"

	logger "github.com/alvinarthas/myrepo/utils/log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {

	logger.Info("INIT MYSQL CONNECTION", nil)
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		"root",
		"password",
		"localhost",
		"myrepo",
	)

	conn, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatal("Failed on connecting to MySQL", err)
	}

	if err = conn.Ping(); err != nil {
		logger.Fatal("Failed verifies a connection to the database", err)
	}

	// conn.SetConnMaxLifetime(time.Minute * 3)
	// conn.SetMaxOpenConns(10)
	// conn.SetMaxIdleConns(10)

	db = conn

}

func CloseConnection() {
	db.Close()
}

func Connect() *sql.DB {
	return db
}
