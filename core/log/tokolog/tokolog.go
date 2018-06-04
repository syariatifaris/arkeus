package tokolog

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"

	"fmt"

	"github.com/agtorre/gocolorize"
	"github.com/syariatifaris/arkeus/util"
)

//Log Level All Flag
const logLevelAll = "all"

//tokolog tag
type Tag struct {
	Name, Value interface{}
}

var (
	//Global Log Level
	globalLogLevels []string

	//Colorize Map
	colors = map[string]gocolorize.Colorize{
		"debug": gocolorize.NewColor("magenta"),
		"info":  gocolorize.NewColor("green"),
		"warn":  gocolorize.NewColor("yellow"),
		"error": gocolorize.NewColor("red"),
		"trace": gocolorize.NewColor("blue"),
	}

	//Log with level
	warnLog  = tokopediaLogs{c: colors["warn"], w: os.Stdout, k: "warn"}
	infoLog  = tokopediaLogs{c: colors["info"], w: os.Stdout, k: "info"}
	errorLog = tokopediaLogs{c: colors["error"], w: os.Stderr, k: "error"}
	debugLog = tokopediaLogs{c: colors["debug"], w: os.Stdout, k: "debug"}
	traceLog = tokopediaLogs{c: colors["trace"], w: os.Stdout, k: "trace"}

	//Print calls with level
	DEBUG = log.New(&debugLog, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	INFO  = log.New(&infoLog, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	WARN  = log.New(&warnLog, "WARN ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(&errorLog, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)
	TRACE = log.New(&traceLog, "TRACE ", log.Ldate|log.Ltime|log.Lshortfile)
)

//Config tokolog config structure
type Config struct {
	LogLevels,
	ErrorLogPath,
	AccessLogPath string
}

//tokopediaLogs structure
type tokopediaLogs struct {
	c gocolorize.Colorize
	w io.Writer
	f *os.File
	k string
}

//Write write the log
func (r *tokopediaLogs) Write(p []byte) (n int, err error) {
	if util.InArray(globalLogLevels, r.k) || util.InArray(globalLogLevels, logLevelAll) || len(globalLogLevels) == 0 {
		//write both file and on console
		colorizedBytes := []byte(r.c.Paint(string(p)))
		r.f.Write(colorizedBytes)
		return r.w.Write(colorizedBytes)
	}

	return 0, errors.New("message wont be printed")
}

//Init initialize tokopedia log level
func Init(cfg *Config) {
	globalLogLevels = strings.Split(cfg.LogLevels, ",")

	//set file for error and warn
	if f := reopen(cfg.ErrorLogPath); f != nil {
		errorLog.f = f
	}

	//set log location for access
	if f := reopen(cfg.AccessLogPath); f != nil {
		debugLog.f = f
		traceLog.f = f
		infoLog.f = f
		warnLog.f = f
	}
}

//reopen opens the existing files as a log destination
func reopen(filename string) *os.File {
	if filename == "" {
		return nil
	}

	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}

	return logFile
}

//SetInfo sets the information of log such as prefixes, and tags
func SetInfo(tags []Tag, prefixes ...string) string {
	var mi string
	for _, arg := range prefixes {
		mi = fmt.Sprintf("%s[%s]", mi, arg)
	}

	if tags != nil {
		for _, tag := range tags {
			mi = fmt.Sprintf("%s[%s:%+v]", mi, tag.Name, tag.Value)
		}
	}

	return mi
}
