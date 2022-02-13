package main

import (
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/connection/mysql"
	routes "github.com/alvinarthas/myrepo/handlers/routes"
	"github.com/alvinarthas/myrepo/utils/logger"
)

// @title MyRepo API Documentation
// @version 1.0
// @description This is a sample Repo for service

// @contact.name Alvin Khair Arthas
// @contact.url https://github.com/alvinarthas
// @contact.email work.alvinkhair@gmail.com

// @host localhost:80
// @BasePath /api/v1
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
