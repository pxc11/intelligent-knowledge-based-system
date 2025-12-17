package logger

import (
	"fmt"
	"ikbs/lib/basic"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type currentLogger struct {
	currentFilename string
	file            *os.File
	logger          *log.Logger
}

var currentLogger1 = &currentLogger{
	currentFilename: "",
	file:            nil,
	logger:          nil,
}

var lock sync.Mutex

func writeLog(level string, msg ...any) {
	lock.Lock()
	defer lock.Unlock()
	logFileName := time.Now().Format("2006-01-02") + ".log"
	if currentLogger1.currentFilename != logFileName {
		if currentLogger1.file != nil {
			err := currentLogger1.file.Close()
			if err != nil {
				log.Panic(err)
			}

		}
		err := os.MkdirAll(basic.GetRootPath()+"/logs", 0755)
		if err != nil {
			log.Panic(err)
		}
		fileWriter, err := os.OpenFile(basic.GetRootPath()+"/"+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Panic(err)
		}
		currentLogger1.file = fileWriter

		currentLogger1.logger = log.New(
			io.MultiWriter(os.Stdout, fileWriter),
			"",
			log.Ldate|log.Ltime|log.Lshortfile,
		)
		currentLogger1.currentFilename = logFileName
	}

	sendMsg := fmt.Sprintf("%+v", msg)

	currentLogger1.logger.Print("[" + level + "] " + sendMsg)

}

func Warning(msg ...any) {
	writeLog("warning", msg...)
}
func Info(msg ...any) {
	writeLog("info", msg...)
}
func Error(msg ...any) {
	writeLog("error", msg...)
}
