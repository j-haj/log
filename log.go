package log

import (
	"fmt"
	"io"
	golog "log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
)

// logLevel is the current logging level. Default logging level is INFO
var logLevel = getEnvLogVar()

// getEnvLogVar looks for environment variables used to specify the logging level
func getEnvLogVar() LogLevel {
	_, isDebug := os.LookupEnv("DEBUG")
	_, isError := os.LookupEnv("ERROR")
	if isDebug {
		return DEBUG
	} else if isError {
		return ERROR
	}
	// INFO is the default logging level
	return INFO
}

func SetLogLevel(level LogLevel) {
	logLevel = level
}

// Loggers
var debugLogger = golog.New(os.Stderr, "[DEBUG] ", golog.LstdFlags)
var infoLogger = golog.New(os.Stderr, "[INFO]  ", golog.LstdFlags)
var errorLogger = golog.New(os.Stderr, "[ERROR] ", golog.LstdFlags)

func SetOutput(out io.Writer) {
	debugLogger = golog.New(out, "[DEBUG] ", golog.LstdFlags)
	infoLogger = golog.New(out, "[INFO]  ", golog.LstdFlags)
	errorLogger = golog.New(out, "[ERROR] ", golog.LstdFlags)
}

// Logf performs the actual logging. All other log functions are simply helper
// functions
func Logf(level LogLevel, format string, v ...interface{}) {
	if level < logLevel {
		return
	}
	output := fmt.Sprintf(format, v...)
	switch level {
	case DEBUG:
		debugLogger.Print(output)
	case INFO:
		infoLogger.Print(output)
	case ERROR:
		errorLogger.Print(output)
	}
}

func Debug(v ...interface{}) {
	Logf(DEBUG, "%v", v...)
}

func Debugf(format string, v ...interface{}) {
	Logf(DEBUG, format, v...)
}

func Debugln(v ...interface{}) {
	Logf(DEBUG, "%v\n", v...)
}

func Info(v ...interface{}) {
	Logf(INFO, "%v", v...)
}

func Infof(format string, v ...interface{}) {
	Logf(INFO, format, v...)
}

func Infoln(v ...interface{}) {
	Logf(INFO, "%v\n", v...)
}

func Error(v ...interface{}) {
	Logf(ERROR, "%v", v...)
}

func Errorf(format string, v ...interface{}) {
	Logf(ERROR, format, v...)
}

func Errorln(v ...interface{}) {
	Logf(ERROR, "%v\n", v...)
}
