package main

import (
	"log"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/connection/mysql"
	routes "github.com/alvinarthas/myrepo/handlers/routes"
)

func main() {

	var (
		port   = config.API_PORT
		router = routes.GetRouter()
	)

	defer mysql.CloseConnection()

	log.Printf("Service started. Listening on port: %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Printf("Failed running service without TLS. Listening on port:%s bind: address already in use", port)
	}
}
