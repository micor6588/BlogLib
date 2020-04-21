package main

import "BlogLib/mylog"

// 测试我们自己写的日志
func main() {
	// log := mylog.NewLogger("INFO")
	log := mylog.NewFileLogger("Info", "./", "micor.log", 10*1024)
	for {
		id := 10102066
		name := "micor"
		log.Debug("这是一条Debug日志")
		log.Debug("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Error("这是一条Error级别日志,id:%d,name:%s", id, name)
		log.Warning("这是一条Warning级别日志")
		log.Fatal("这是一条Fatal级别日志")
		// time.Sleep(time.Second * 2)
	}
}
