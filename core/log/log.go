//Package log is used for doing custom log.
package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/syariatifaris/arkeus/core/errors"
)

const logFormat = `date=%s, method=%s, url=%s,  response_time=%s`

// Config of log
type Config struct {
	LogLevel     string
	ErrorLogPath string
}

var (
	// DefaultConfig of log
	DefaultConfig = Config{
		LogLevel:     logrus.InfoLevel.String(),
		ErrorLogPath: "/var/log/transactionapp/transactionapp.error.log",
	}
)

// SetConfig for log object
func SetConfig(config Config) {
	// set default config if settings not there
	if config.LogLevel == "" {
		config.LogLevel = DefaultConfig.LogLevel
	}
	if config.ErrorLogPath == "" {
		config.ErrorLogPath = DefaultConfig.ErrorLogPath
	}

	SetLevel(config.LogLevel)
	if f := reopen(1, config.ErrorLogPath); f != nil {
		logrus.SetOutput(f)
	}
}

func reopen(fd int, filename string) *os.File {
	if filename == "" {
		return nil
	}

	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// logrus.Error("Error in opening ", filename, err)
		return nil
	}
	return logFile
}

// Fields for logrus fields
type Fields logrus.Fields

// SetLevel of logs
func SetLevel(level string) {
	switch level {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warning":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

// GetLevel get log level
func GetLevel() string {
	return strings.ToUpper(logrus.GetLevel().String())
}

func getFileAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		i := strings.Index(file, "app")
		file = file[i:]
	}

	return file, line
}

func formatFilePath(f string) string {
	slash := strings.LastIndex(f, "/")
	return f[slash+1:]
}

// Info log
func Info(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Infoln log
func Infoln(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infoln(args...)
}

// Infof log
func Infof(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Print log
func Print(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Info(args...)
}

// Println log
func Println(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infoln(args...)
}

// Printf log
func Printf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Infof(format, args...)
}

// Debug log
func Debug(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debug(args...)
}

// Debugln log
func Debugln(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debugln(args...)
}

// Debugf log
func Debugf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Debugf(format, args...)
}

// Warn log
func Warn(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warn(args...)
}

// Warnln log
func Warnln(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warnln(args...)
}

// Warnf log
func Warnf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Warnf(format, args...)
}

// Error log
func Error(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Error(args...)
}

// Errorln log
func Errorln(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Errorln(args...)
}

// Errorf log
func Errorf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Errorf(format, args...)
}

// Errors should be called by using errors package
// errors package have special error fields to add more context in error
func Errors(err error) {
	var (
		errFields errors.Fields
		file      string
		line      int
	)
	switch err.(type) {
	case *errors.Errs:
		errs := err.(*errors.Errs)
		errFields = errs.GetFields()
		file, line = errs.GetFileAndLine()
	}
	// transform error fields to log fields
	logFields := logrus.Fields(errFields)
	// check if file and line is exists
	if line != 0 {
		logFields["err_file"] = formatFilePath(file)
		logFields["err_line"] = line
	}
	logrus.WithFields(logFields).Error(err.Error())
}

// Fatal log
func Fatal(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatal(args...)
}

// Fatalln log
func Fatalln(args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatalln(args...)
}

// Fatalf log
func Fatalf(format string, args ...interface{}) {
	file, line := getFileAndLine()
	logrus.WithField("source", fmt.Sprintf("%s:%d", file, line)).Fatalf(format, args...)
}

// WithFields log
func WithFields(fields Fields) *logrus.Entry {
	file, line := getFileAndLine()
	fields["source"] = fmt.Sprintf("%s:%d", file, line)
	return logrus.WithFields(logrus.Fields(fields))
}
