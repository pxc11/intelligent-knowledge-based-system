package logger

import (
	"fmt"
	"ikbs/lib/basic"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type currentLogger struct {
	currentFilename string
	file            *os.File
	logger          *log.Logger
}

type logMessage struct {
	level   string
	message []any
	file    string
	line    int
}

func Init() {
	writingLog()
}

var channel = make(chan *logMessage, 1000)

var currentLogger1 = &currentLogger{
	currentFilename: "",
	file:            nil,
	logger:          nil,
}

func pushMsg(level string, msg ...any) {
	_, file, line, _ := runtime.Caller(2)
	message := &logMessage{
		level:   level,
		message: msg,
		file:    file,
		line:    line,
	}
	channel <- message
}

var once sync.Once

func writingLog() {
	once.Do(func() {
		go func() {
			for msg := range channel {
				writeLog(msg)
			}
		}()

	})
}

func writeLog(message *logMessage) {
	level := message.level
	msg := message.message
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
		fileWriter, err := os.OpenFile(basic.GetRootPath()+"/logs/"+logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Panic(err)
		}
		currentLogger1.file = fileWriter

		currentLogger1.logger = log.New(
			io.MultiWriter(os.Stdout, fileWriter),
			"",
			log.Ldate|log.Ltime,
		)
		currentLogger1.currentFilename = logFileName
	}

	if currentLogger1.logger == nil {
		log.Panic("logger 不存在")
	}
	currentLogger1.logger.Print("[" + level + "][ " + message.file + ":" + strconv.Itoa(message.line) + " ]" + formatArgs(msg...))

}

func formatArgs(args ...any) string {
	var sb strings.Builder
	for i, a := range args {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprintf("%+v", a))
	}
	return sb.String()
}

func Warning(msg ...any) {
	pushMsg("warning", msg...)

}
func Info(msg ...any) {
	pushMsg("info", msg...)

}
func Error(msg ...any) {
	pushMsg("error", msg...)

}
