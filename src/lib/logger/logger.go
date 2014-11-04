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
func (this *Logger) SetNewLogger() error {
	var err error = nil
	fileNameTemplate := "%s-%s.log"
	fileName := fmt.Sprintf(fileNameTemplate, consts.NAME, time.Now().Format("20060102150405"))
	this.logf, err = os.OpenFile(consts.DIR_LOG+fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	this.logs = log.New(this.logf, "\n", log.Ldate|log.Ltime)
	this.logs.SetPrefix(consts.NAME + "\t")
	return err
}

//Stop Logger
func (this *Logger) CloseLogger() {
	this.logf.Close()
}

//For testing use.
func (this *Logger) logTest(argv ...interface{}) {
	this.logs.Print("TEST: ", argv)
}

//For general use.
func (this *Logger) LogInfo(argv ...interface{}) {
	this.logs.Print("INFO: ", argv)
}

//For warning(s) use.
func (this *Logger) LogWarning(argv ...interface{}) {
	this.logs.Print("WARN: ", argv)
}

//For error(s) occured.
func (this *Logger) LogError(argv ...interface{}) {
	this.logs.Panic("ERROR: ", argv)
}

//For fatal error(s).
func (this *Logger) LogCritical(argv ...interface{}) {
	this.logs.Fatal("CRITICAL: ", argv)
}

//Reimplement of log.SetPrefix()
func (this *Logger) SetPrefix(prefix string) {
	this.logs.SetPrefix(prefix + "\t")
}
