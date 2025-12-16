package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/************** level **************/

type Level string

const (
	INFO    Level = "INFO"
	WARNING Level = "WARNING"
	ERROR   Level = "ERROR"
)

/************** rotating writer **************/

type rotatingWriter struct {
	mu       sync.Mutex
	dir      string
	currDate string
	file     *os.File
}

func newRotatingWriter(dir string) (*rotatingWriter, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	w := &rotatingWriter{
		dir: dir,
	}
	if err := w.rotateIfNeeded(); err != nil {
		return nil, err
	}
	return w, nil
}

func (w *rotatingWriter) rotateIfNeeded() error {
	date := time.Now().Format("2006-01-02")

	if date == w.currDate && w.file != nil {
		return nil
	}

	if w.file != nil {
		_ = w.file.Close()
	}

	filename := filepath.Join(w.dir, "log-"+date+".log")
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	w.currDate = date
	w.file = file
	return nil
}

func (w *rotatingWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.rotateIfNeeded(); err != nil {
		return 0, err
	}
	return w.file.Write(p)
}

/************** logger **************/

var (
	std  *log.Logger
	once sync.Once
)

func Init(logDir string) error {
	var err error

	once.Do(func() {
		var writer *rotatingWriter
		writer, err = newRotatingWriter(logDir)
		if err != nil {
			return
		}

		mw := io.MultiWriter(os.Stdout, writer)

		std = log.New(
			mw,
			"",
			log.Ldate|log.Ltime|log.Lshortfile,
		)
	})

	return err
}

func logWithLevel(level Level, v ...any) {
	if std == nil {
		panic("logger not initialized")
	}

	msg := fmt.Sprint(v...)
	std.Print("[" + string(level) + "] " + msg)
}

/************** public API **************/

func Info(v ...any) {
	logWithLevel(INFO, v...)
}

func Warning(v ...any) {
	logWithLevel(WARNING, v...)
}

func Error(v ...any) {
	logWithLevel(ERROR, v...)
}
