package logger

/** Desacoplar y simplificar la libreria Logrus Logs */
import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

const eventName = "x_event"

var logger *logrus.Logger
var once sync.Once

func ConfigLogs(level string) {
	once.Do(func() {
		logger = &logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{PrettyPrint: false},
			Hooks:     make(logrus.LevelHooks),
			Level:     logrus.DebugLevel,
		}

	})
	SetLogLevel(level)
}

func SetLogLevel(logLevel string) {
	if level, error := logrus.ParseLevel(logLevel); error != nil {
		logger.Level = logrus.TraceLevel
	} else {
		logger.Level = level
	}
}

func Trace(msg string) {
	logger.Trace(msg)
}
func Info(msg string) {
	logger.Info(msg)
}
func Debug(msg string) {
	logger.Debug(msg)
}
func Error(msg string) {
	logger.Error(msg)
}
func Fatal(msg string) {
	logger.Fatal(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func TraceF(msg string, logFields interface{}) {

	logger.WithFields(logrus.Fields{eventName: logFields}).Trace(msg)
}
func InfoF(msg string, logFields interface{}) {
	logger.WithFields(logrus.Fields{eventName: logFields}).Info(msg)
}
func DebugF(msg string, logFields interface{}) {
	logger.WithFields(logrus.Fields{eventName: logFields}).Debug(msg)
}
func ErrorF(msg string, logFields interface{}) {
	logger.WithFields(logrus.Fields{eventName: logFields}).Error(msg)
}
func FatalF(msg string, logFields interface{}) {
	logger.WithFields(logrus.Fields{eventName: logFields}).Fatal(msg)
}

func WarnF(msg string, logFields interface{}) {
	logger.WithFields(logrus.Fields{eventName: logFields}).Warn(msg)
}
