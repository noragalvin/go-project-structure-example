package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger var
var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	Logger.Out = os.Stdout
	// switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
	// case "debug":
	// 	Logger.Level = logrus.DebugLevel
	// case "info":
	// 	Logger.Level = logrus.InfoLevel
	// case "warn":
	// 	Logger.Level = logrus.WarnLevel
	// case "error":
	// 	Logger.Level = logrus.ErrorLevel
	// case "fatal":
	// 	Logger.Level = logrus.FatalLevel
	// case "panic":
	// 	Logger.Level = logrus.PanicLevel
	// default:
	// 	Logger.Level = logrus.InfoLevel
	// }
	Logger.Formatter = &logrus.TextFormatter{ForceColors: true}
}
