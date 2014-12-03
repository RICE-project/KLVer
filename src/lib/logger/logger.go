/*
Reimplements of go native log module.
*/
package logger

import (
	"fmt"
	"lib/consts"
	"log"
	"os"
	"strings"
	"time"
)

const (
	LOG_LEVEL_TEST = iota
	LOG_LEVEL_INFO
	LOG_LEVEL_WARNING
	LOG_LEVEL_ERROR
	LOG_LEVEL_CRIT
)

//A Class of logger.
type Logger struct {
	logs     *log.Logger
	logf     *os.File
	logLevel int
}

//Init Logger
func (l *Logger) SetNewLogger() error {
	var err error = nil
	fileNameTemplate := "%s-%s.log"
	fileName := fmt.Sprintf(fileNameTemplate, consts.NAME, time.Now().Format("20060102150405"))
	l.logf, err = os.OpenFile(consts.DIR_LOG+fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	l.logs = log.New(l.logf, "\n", log.Ldate|log.Ltime)
	l.logs.SetPrefix(consts.NAME + "\t")
	return err
}

//Stop Logger
func (l *Logger) CloseLogger() {
	l.logf.Close()
}

//For testing use.
func (l *Logger) LogTest(argv ...interface{}) {
	if l.logLevel <= LOG_LEVEL_TEST {
		l.logs.Print("TEST: ", argv)
	}
}

//For general use.
func (l *Logger) LogInfo(argv ...interface{}) {
	if l.logLevel <= LOG_LEVEL_INFO {
		l.logs.Print("INFO: ", argv)
	}
}

//For warning(s) use.
func (l *Logger) LogWarning(argv ...interface{}) {
	if l.logLevel <= LOG_LEVEL_WARNING {
		fmt.Fprintln(os.Stderr, "\033[31;1mWARNING:\033[0m", argv)
		l.logs.Print("WARN: ", argv)
	}
}

//For error(s) occured.
func (l *Logger) LogError(argv ...interface{}) {
	if l.logLevel <= LOG_LEVEL_ERROR {
		l.logs.Panic("ERROR: ", argv)
	}
}

//For fatal error(s).
func (l *Logger) LogCritical(argv ...interface{}) {
	if l.logLevel <= LOG_LEVEL_CRIT {
		l.logs.Fatal("CRITICAL: ", argv)
	}
}

//Reimplement of log.SetPrefix()
func (l *Logger) SetPrefix(prefix string) {
	l.logs.SetPrefix(prefix + "\t")
}

func (l *Logger) SetLevel(level int) {
	l.logLevel = level
}

func (l *Logger) GetLevel() int {
	return l.logLevel
}

func GetLevel(str string) int {
	str = strings.ToLower(str)
	switch str {
	case "deb":
		fallthrough
	case "test":
		fallthrough
	case "debug":
		return LOG_LEVEL_TEST
	case "warn":
		fallthrough
	case "warning":
		return LOG_LEVEL_WARNING
	case "err":
		fallthrough
	case "error":
		return LOG_LEVEL_ERROR
	case "fatal":
		fallthrough
	case "critical":
		fallthrough
	case "crit":
		return LOG_LEVEL_CRIT
	case "info":
		fallthrough
	default:
		return LOG_LEVEL_INFO
	}
}
