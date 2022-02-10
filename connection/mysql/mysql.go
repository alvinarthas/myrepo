package mysql

import (
	"database/sql"
	"fmt"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {

	logger.Info("INIT MYSQL CONNECTION", nil)
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.CONFIG.Connection.MySQL.Username,
		config.CONFIG.Connection.MySQL.Password,
		config.CONFIG.Connection.MySQL.Host,
		config.CONFIG.Connection.MySQL.Port,
		config.CONFIG.Connection.MySQL.DB,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		logger.Fatal("Failed on connecting to MySQL", err)
	}

	if err = db.Ping(); err != nil {
		logger.Fatal("Failed verifies a connection to the database", err)
	}

	// conn.SetConnMaxLifetime(time.Minute * 3)
	// conn.SetMaxOpenConns(10)
	// conn.SetMaxIdleConns(10)
}

func CloseConnection() {
	db.Close()
}

func Connect() *sql.DB {
	return db
}
