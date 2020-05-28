package log

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

var (
	logger *log.Logger

	DefaultPrefix     = ""
	DefaultCallerPath = 2

	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

func init() {
	filepath := getLogFileFullPath()
	F := openLogFile(filepath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerPath)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}
