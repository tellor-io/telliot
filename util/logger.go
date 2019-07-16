package util

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

//LogLevel represents possible log levels
type LogLevel uint

const (
	//ErrorLogLevel will only log error messages
	ErrorLogLevel LogLevel = iota + 1

	//WarningLogLevel will only log warnings
	WarningLogLevel

	//InfoLogLevel will only log infos
	InfoLogLevel

	//DebugLogLevel will only log debugs
	DebugLogLevel
)

//Logger is the log implementation that wraps underlying logging mechanism
type Logger struct {
	pkg       string
	component string
	log       *logrus.Entry
}

//NewLogger creates a log instance for the given component with the given log level enabled
func NewLogger(pkg string, component string) *Logger {
	levels, err := GetLoggingConfig()
	if err != nil {
		panic(err)
	}
	level := levels.GetLevel(pkg, component)
	ll := xlateLevel(level)
	log := logrus.New()
	log.SetLevel(ll)
	log.SetFormatter(&logrus.TextFormatter{ForceColors: true, PadLevelText: true})

	e := log.WithFields(logrus.Fields{"component": component, "package": pkg})
	return &Logger{pkg: pkg, component: component, log: e}
}

//Error log an error message
func (l *Logger) Error(msg string, args ...interface{}) {
	l._logIt(l.log.Errorf, msg, args...)
}

//Warn log a warning message
func (l *Logger) Warn(msg string, args ...interface{}) {
	l._logIt(l.log.Warnf, msg, args...)
}

//Info message logging
func (l *Logger) Info(msg string, args ...interface{}) {
	l._logIt(l.log.Infof, msg, args...)
}

//Debug log a debug message
func (l *Logger) Debug(msg string, args ...interface{}) {
	l._logIt(l.log.Debugf, msg, args...)
}

func (l *Logger) _logIt(fn func(msg string, args ...interface{}), msg string, args ...interface{}) {
	fn(msg, args...)
}

//StringToLevel converts a string named log level into a level ot use for setting log levels
func StringToLevel(level string) (LogLevel, error) {
	l := strings.ToUpper(level)
	switch l {
	case "ERROR":
		{
			return ErrorLogLevel, nil
		}
	case "WARN":
		fallthrough
	case "WARNING":
		{
			return WarningLogLevel, nil
		}
	case "INFO":
		{
			return InfoLogLevel, nil
		}
	case "DEBUG":
		{
			return DebugLogLevel, nil
		}

	}
	return 0, fmt.Errorf("Invalid level name: %s", level)
}

func xlateLevel(level LogLevel) logrus.Level {
	switch level {
	case ErrorLogLevel:
		{
			return logrus.ErrorLevel
		}
	case WarningLogLevel:
		{
			return logrus.WarnLevel
		}
	case InfoLogLevel:
		{
			return logrus.InfoLevel
		}
	case DebugLogLevel:
		{
			return logrus.DebugLevel
		}
	default:
		return logrus.InfoLevel
	}

}
