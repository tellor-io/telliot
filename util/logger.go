package util

import (
	"fmt"
	"strings"
	"sync"

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

//people want to declare singleton logger instances as global vars
//but the logging config won't be loaded yet
//so we hand out a pointer to a logger instance, but leave it uninitalized and keep a copy for ourselves
//then when the logging config is loaded, we can init all the instances we gave out earlier
//since none of these modules should be using the logger before the config is loaded, this works
var loggersToInit []*Logger
var loggerMutex sync.Mutex


//NewLogger creates a log instance for the given component with the given log level enabled
func NewLogger(pkg string, component string) *Logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	l := &Logger{pkg: pkg, component: component}
	levels := GetLoggingConfig()
	if levels != nil {
		initLogger(levels, l)
	} else {
		loggersToInit = append(loggersToInit, l)
	}
	return l
}

func initLogger(levels *LogConfig, l *Logger) {
	level := levels.GetLevel(l.pkg, l.component)
	ll := xlateLevel(level)
	log := logrus.New()
	log.SetLevel(ll)
	log.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	l.log = log.WithFields(logrus.Fields{"component": l.component, "package": l.pkg})
}

func initLoggers(levels *LogConfig) {
	for _,l := range loggersToInit {
		initLogger(levels, l)
	}
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
