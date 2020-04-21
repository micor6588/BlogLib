//往终端上写日志

package mylog

import (
	"fmt"
	"time"
)

type LevelLog uint16 //定义日志级别
const (
	UNKNOWN LevelLog = iota // 定义日志级别
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger 日志结构体
type ConsoleLogger struct {
	LogLevel LevelLog
}

// NewLogger 构造函数
func NewLogger(levalStr string) ConsoleLogger {
	level, err := ParseLogLevel(levalStr)
	if err != nil {
		fmt.Println("输出日志级别失败，err=", err)

	}
	return ConsoleLogger{
		LogLevel: level,
	}
}

// Enable 比较日志等级
func (c ConsoleLogger) Enable(logLevel LevelLog) bool {
	return logLevel >= c.LogLevel
}

//  输出日志级别信息
func (c ConsoleLogger) logRemand(level LevelLog, format string, a ...interface{}) {
	if c.Enable(level) {
		now := time.Now()
		nowTime := now.Format("2006-01-02 15:04:05")
		funcName, fileName, line := GetInfo(3)
		//得到表示日志等级的字符串
		msg := fmt.Sprintf(format, a...)
		levelStr := GetLoggerString(level)
		fmt.Printf("输出日志信息：[%s:][日志等级：%s][日志生成时间:%s][函数名%s:,文件名%s:,行号%d:]\n,", msg, levelStr, nowTime, funcName, fileName, line)
	}
}

// Debug 级别日志输出
func (c ConsoleLogger) Debug(msg string, a ...interface{}) {
	c.logRemand(DEBUG, msg, a...)

}

// Trace 级别日志输出
func (c ConsoleLogger) Trace(msg string, a ...interface{}) {
	c.logRemand(TRACE, msg, a...)
}

// Info 级别日志输出
func (c ConsoleLogger) Info(msg string, a ...interface{}) {
	c.logRemand(INFO, msg, a...)
}

// Warning 级别日志输出
func (c ConsoleLogger) Warning(msg string, a ...interface{}) {
	c.logRemand(WARNING, msg, a...)
}

// Error 级别日志输出
func (c ConsoleLogger) Error(msg string, a ...interface{}) {
	c.logRemand(ERROR, msg, a...)
}

// Fatal 级别日志输出
func (c ConsoleLogger) Fatal(msg string, a ...interface{}) {
	c.logRemand(FATAL, msg, a...)
}
