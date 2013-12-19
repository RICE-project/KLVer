package logger

import(
        "log"
        "lib/consts"
        "os"
        "time"
        "fmt"
)

type Logger struct{
        logs *log.Logger
        logf *os.File
}

func (this *Logger) SetNewLogger() error {
        var err error = nil
        fileNameTemplate := "glvsadm-%s.log"
        fileName := fmt.Sprintf(fileNameTemplate, time.Now().Format("20060102150405"))
        this.logf, err = os.OpenFile(consts.DIR_LOG + fileName, os.O_RDWR | os.O_CREATE, 0666)
        if err != nil{
                this.CloseLogger()
                return err
        }

        this.logs = log.New(this.logf, "\n", log.Ldate | log.Ltime | log.Lshortfile)
        this.logs.SetPrefix("gLVSAdm\t")
        return err
}

func (this *Logger) CloseLogger() {
        this.logf.Close()
}

func (this *Logger) LogTest(argv ...interface{}) {
        this.logs.Print("TEST: ", argv)
}

func (this *Logger) LogInfo(argv ...interface{}) {
        this.logs.Print("INFO: ", argv)
}

func (this *Logger) LogError(argv ...interface{}) {
        this.logs.Panic("ERROR: ", argv)
}

func (this *Logger) LogCritical(argv ...interface{}) {
        this.logs.Fatal("CRITICAL: ", argv)
}

func (this *Logger) SetPrefix(prefix string) {
        this.logs.SetPrefix(prefix + "\t")
}
