package main

import (
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/connection/mysql"
	routes "github.com/alvinarthas/myrepo/handlers/routes"
	"github.com/alvinarthas/myrepo/utils/logger"
)

func main() {
	var (
		port   = config.CONFIG.Server.Port
		router = routes.GetRouter()
	)

	defer mysql.CloseConnection()

	logger.Info(fmt.Sprintf("Service started. Listening on port: %s", port), nil)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed running service without TLS. Listening on port:%s bind: address already in use", port), err)
	}
}
