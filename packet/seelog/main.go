package main

import (
    log "github.com/cihub/seelog"
)

func main()  {
    defer log.Flush()
    logger, err := log.LoggerFromConfigAsFile("config.xml");

    if err != nil {
        log.Critical("err parsing config log file", err)
        return
    }
    log.ReplaceLogger(logger)

    log.Trace("我是trace级别的日志，用来跟踪数据输出哦")
    log.Debug("我是debug级别的日志，用来输出调试信息哦")
    log.Info("我是info级别的日志，用来输出提示哦")
    log.Warn("我是warn级别的日志，用来输出警告信息")
    log.Error("我是error级别的日志，哈，程序出错了！")
    log.Critical("我是critical级别的日志，我是最高级别错误")


}