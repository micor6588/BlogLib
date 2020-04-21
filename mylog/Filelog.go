// 在文件当中写日志的相关代码

package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 日志文件的结构体
type FileLogger struct {
	FileLogLevel LevelLog
	FilePath     string //日志文件信息
	FileName     string // 日志文件保存的文件名
	FileObj      *os.File
	errFileObj   *os.File
	MaxFileSize  int64 //日志文件切割
}

//对文件进行相关操作
func (f *FileLogger) initFile() error {
	//1.首先对文件进行拼接
	FullPath := path.Join(f.FilePath, f.FileName)
	fileObj, err := os.OpenFile(FullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Open log file failed err=%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(FullPath+".err", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Open err log file failed err=%v\n", err)
		return err
	}
	// 日志文件都打开了
	f.errFileObj = errfileObj
	f.FileObj = fileObj
	return nil
}

// NewFileLogger 将文件日志结构体实例化
func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int64) *FileLogger {
	logLevel, err := ParseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f1 := &FileLogger{
		FileLogLevel: logLevel,
		FilePath:     filePath,
		FileName:     fileName,
		MaxFileSize:  maxFileSize,
	}
	err = f1.initFile() //按照路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return f1

}

// Enable 比较日志等级
func (f FileLogger) Enable(logLevel LevelLog) bool {
	return logLevel >= f.FileLogLevel
}

//  输出日志级别信息
func (f FileLogger) logRemand(level LevelLog, format string, a ...interface{}) {
	if f.Enable(level) {
		now := time.Now()
		nowTime := now.Format("2006-01-02 15:04:05")
		funcName, fileName, line := GetInfo(3)
		//得到表示日志等级的字符串
		msg := fmt.Sprintf(format, a...)
		levelStr := GetLoggerString(level)
		fmt.Fprintf(f.FileObj, "输出日志信息 ：[%s:][日志等级：%s][日志生成时间:%s][函数名%s:,文件名%s:,行号%d:]\n,", msg, levelStr, nowTime, funcName, fileName, line)
		if level > ERROR {
			//如果日志级别大于EROOR级别，还要再err文件当中记录一边
			fmt.Fprintf(f.errFileObj, "输出日志信息：[%s:][日志等级：%s][日志生成时间:%s][函数名%s:,文件名%s:,行号%d:]\n,", msg, levelStr, nowTime, funcName, fileName, line)
		}
	}
}

// Debug 级别日志输出
func (f FileLogger) Debug(msg string, a ...interface{}) {
	f.logRemand(DEBUG, msg, a...)
}

// Trace 级别日志输出
func (f FileLogger) Trace(msg string, a ...interface{}) {
	f.logRemand(TRACE, msg, a...)
}

// Info 级别日志输出
func (f FileLogger) Info(msg string, a ...interface{}) {
	f.logRemand(INFO, msg, a...)
}

// Warning 级别日志输出
func (f FileLogger) Warning(msg string, a ...interface{}) {
	f.logRemand(WARNING, msg, a...)
}

// Error 级别日志输出
func (f FileLogger) Error(msg string, a ...interface{}) {
	f.logRemand(ERROR, msg, a...)
}

// Fatal 级别日志输出
func (f FileLogger) Fatal(msg string, a ...interface{}) {
	f.logRemand(FATAL, msg, a...)
}
