package main

import (
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/connection/mysql"
	routes "github.com/alvinarthas/myrepo/handlers/routes"
	"github.com/alvinarthas/myrepo/utils/logger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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
