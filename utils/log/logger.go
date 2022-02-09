package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	logLevelTrace   = "TRACE"
	logLevelDebug   = "DEBUG"
	logLevelInfo    = "INFO"
	logLevelWarning = "WARNING"
	logLevelFatal   = "FATAL"
	logLevelError   = "ERROR"
)

var logLevel = map[string]int{
	logLevelTrace:   6,
	logLevelDebug:   5,
	logLevelInfo:    4,
	logLevelWarning: 3,
	logLevelFatal:   2,
	logLevelError:   1,
}

var activeLogLevel = "TRACE"

func init() {
	logrus.SetOutput(os.Stdout)
}

func Trace(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
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
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
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
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Info(message)
	}
}

func Warning(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Warning(message)
	}
}

func Fatal(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Fatal(message)
	}
}

func Error(message string, data interface{}) {
	if getActiveLogLevel() >= logLevel[logLevelWarning] {
		logrus.WithFields(logrus.Fields{
			"data":         parseData(data),
			"service_name": "serviceName",
		}).Error(message)
	}
}

func getActiveLogLevel() int {
	return logLevel[activeLogLevel]
}

func parseData(data interface{}) string {
	strData := ""
	if data != nil {
		strData = fmt.Sprintf("%v", data)
	}

	return strData
}
