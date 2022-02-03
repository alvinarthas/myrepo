package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	routes "github.com/alvinarthas/myrepo/handlers/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var (
		port   = config.API_PORT
		router = routes.GetRouter()
	)

	var connectionString string

	log.Println("INIT MYSQL CONNECTION")
	connectionString = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s",
		"root",
		"password",
		"localhost",
		"myrepo",
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Connection error: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Println("Failed verifies a connection to the database", nil)
	}

	defer db.Close()

	log.Printf("Service started. Listening on port: %s", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Printf("Failed running service without TLS. Listening on port:%s bind: address already in use", port)
	}
}
