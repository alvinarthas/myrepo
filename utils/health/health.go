package health

import (
	"net/http"

	"github.com/alvinarthas/myrepo/connection/mysql"
	"github.com/alvinarthas/myrepo/utils/logger"
	"github.com/alvinarthas/myrepo/utils/response"
)

func HealthStatus(res http.ResponseWriter, req *http.Request) {

	mysqlStatus := true
	mysqlConnection := mysql.Connect()
	if err := mysqlConnection.Ping(); err != nil {
		mysqlStatus = false
		logger.Info("Failed verifies a connection to the database", err)
	}

	var result = map[string]bool{
		"mysql": mysqlStatus,
	}

	response.Success(res, http.StatusOK, result)
}
