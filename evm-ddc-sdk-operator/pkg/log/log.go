package log

import (
	"io"
	"log"
	"os"
)

//默认的log文件输出路径
const defaultLogFilePath = "./defaultLog.log"

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// RegisterLog 用户自定义log文件输出路径
func RegisterLog(filePath string) {
	if len(filePath) == 0 {
		filePath = defaultLogFilePath
	}
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0766)
	if err != nil {
		panic(err.Error())
	}

	multiWriter := io.MultiWriter(os.Stderr, logFile)

	Debug = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}
