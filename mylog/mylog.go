package mylog

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// ParseLogLevel 解析日志信息
func ParseLogLevel(s string) (LevelLog, error) {
	//将对应的日志级别字符串都转化成小写
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效日志级别")
		return UNKNOWN, err
	}
}

// GetInfo 获取日志文件信息
func GetInfo(skip int) (funcName, fileName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)

	if !ok {

		fmt.Println("runtime.Caller failed err = ")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}

// GetLoggerString 将日志等级转化为字符串
func GetLoggerString(level LevelLog) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}
