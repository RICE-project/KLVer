/*
Reimplements of go native log module.
*/
package logger

import (
	"fmt"
	"lib/consts"
	"log"
	"os"
	"time"
)

//A Class of logger.
type Logger struct {
	logs *log.Logger
	logf *os.File
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
func (l *Logger) logTest(argv ...interface{}) {
	l.logs.Print("TEST: ", argv)
}

//For general use.
func (l *Logger) LogInfo(argv ...interface{}) {
	l.logs.Print("INFO: ", argv)
}

//For warning(s) use.
func (l *Logger) LogWarning(argv ...interface{}) {
	l.logs.Print("WARN: ", argv)
}

//For error(s) occured.
func (l *Logger) LogError(argv ...interface{}) {
	l.logs.Panic("ERROR: ", argv)
}

//For fatal error(s).
func (l *Logger) LogCritical(argv ...interface{}) {
	l.logs.Fatal("CRITICAL: ", argv)
}

//Reimplement of log.SetPrefix()
func (l *Logger) SetPrefix(prefix string) {
	l.logs.SetPrefix(prefix + "\t")
}
