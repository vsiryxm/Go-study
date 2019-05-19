package main

import (
    . "./database"
    . "./controller"
    "github.com/gin-gonic/gin"
)

func main() {
    defer SqlDB.Close()
    router := initRouter()
    router.Run(":8080")
}

func initRouter() *gin.Engine {
    router := gin.Default()

    router.GET("/", Index)
    router.GET("/help", GetHelpList) //获取列表数据
    router.GET("/help/:id", GetHelpDetail) //获取一条数据
    router.POST("/help", AddHelp) //添加数据
    router.POST("/help/:id", UpdateHelp) //更新数据
    router.DELETE("/help", DelHelp) //删除数据
    return router
}