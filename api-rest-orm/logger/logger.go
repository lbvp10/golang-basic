package log

/**Desacoplar y simplificar la libreria logrus Logs*/
import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var logger *logrus.Logger
var once sync.Once

type FieldsLogs map[string]interface{}

func ConfigLogs(level string) {
	once.Do(func() {
		logger = &logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{},
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

func TraceF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Trace(msg)
}
func InfoF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Info(msg)
}
func DebugF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Debug(msg)
}
func ErrorF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Error(msg)
}
func FatalF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Fatal(msg)
}

func WarnF(msg string, logFields FieldsLogs) {
	logger.WithFields(logrus.Fields(logFields)).Warn(msg)
}
