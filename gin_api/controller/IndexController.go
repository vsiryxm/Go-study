package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func Index(c *gin.Context)  {
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "msg": "Hello API！",
        "data": nil, //nil返回前端显示null
    })
}













