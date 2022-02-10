package logger

import (
	"fmt"
	"os"

	"github.com/alvinarthas/myrepo/config"
	"github.com/sirupsen/logrus"
)

const (
	LOG_LEVEL_TRACE   = "TRACE"
	LOG_LEVEL_DEBUG   = "DEBUG"
	LOG_LEVEL_INFO    = "INFO"
	LOG_LEVEL_WARNING = "WARNING"
	LOG_LEVEL_FATAL   = "FATAL"
	LOG_LEVEL_ERROR   = "ERROR"
)

var logLevel = map[string]int{
	LOG_LEVEL_TRACE:   6,
	LOG_LEVEL_DEBUG:   5,
	LOG_LEVEL_INFO:    4,
	LOG_LEVEL_WARNING: 3,
	LOG_LEVEL_FATAL:   2,
	LOG_LEVEL_ERROR:   1,
}

func init() {
	logrus.SetOutput(os.Stdout)
}

func Trace(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_TRACE] {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Trace(message)
	}
}

func Debug(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_DEBUG] {
		logrus.SetReportCaller(true)
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: true,
		})
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Debug(message)
	}
}

func Info(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_INFO] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Info(message)
	}
}

func Warning(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_WARNING] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Warning(message)
	}
}

func Fatal(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_FATAL] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Fatal(message)
	}
}

func Error(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[LOG_LEVEL_ERROR] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Error(message)
	}
}

func getActiveLogLevel() int {
	return logLevel[config.CONFIG.Log.Level]
}

func parseData(data interface{}) string {
	strData := ""
	if data != nil {
		strData = fmt.Sprintf("%v", data)
	}

	return strData
}
