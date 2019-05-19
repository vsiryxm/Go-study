package controller

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    . "../model"
    "log"
    "strconv"
)
//Index GetHelp GetHelpById AddHelp

//获取帮助列表
func GetHelpList(c *gin.Context)  {
    limit, _ := strconv.Atoi(c.Param("limit"))
    page, _  := strconv.Atoi(c.Param("page"))

    println("limit=", limit)
    println("page=", page)

    var h Help
    helps, err := h.GetList(page, limit)
    if err != nil {
        log.Fatalln(err)
    }

    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "data": helps,
        "msg":  "success",
    })

}

//获取详情
func GetHelpDetail(c *gin.Context)  {
    //content := c.PostForm("content")
    id, _ := strconv.Atoi(c.Param("id"))
    h := Help{Id: id}
    h.GetDetail()
    if h.Id > 0 {
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "data": h,
            "msg":  "success",
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "code": 0,
            "data": nil,
            "msg":  "user not found",
        })
    }

}

//添加数据
func AddHelp(c *gin.Context) {
    //content := c.PostForm("content") //获取单个参数
    //now := time.Now().UnixNano() / 1e9
    //id := c.Param("id")  //Post请求获取不到值，只适合Get
    title := c.Request.FormValue("title")
    content := c.Request.FormValue("content")

    h := Help{Title: title, Content: content}
    rs, err := h.Add()
    if err != nil {
        log.Fatalln(err)
        return
    }

    msg := fmt.Sprintf("insert success %d", rs)
    c.JSON(http.StatusOK, gin.H{
        "code": 1,
        "data": true,
        "msg":  msg,
    })


    //stmt, err := db.Prepare("INSERT vs_help SET title=?,content=?,brief=?,keywords=?,hits=?,is_publish=?,publish_time=?,add_time=?")
    //res, err := stmt.Exec("我是标题", "我是内容", "我是摘要", "我是关键词", 1, 0, now, now)

}

//修改数据
func UpdateHelp(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id")) //Param方法用于Get
    title := c.Request.FormValue("title") //FormValue方法用于Post
    content := c.Request.FormValue("content")
    
    if err != nil {
        log.Fatalln(err)
    }

    h := Help{Id: id}
    h.GetDetail()
    if h.Id > 0 {
        h.Title = title
        h.Content = content
        ra, err := h.Update()
        if err != nil {
            log.Fatalln(err)
        }

        msg := fmt.Sprintf("update success %d", ra)
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "data": true,
            "msg":  msg,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "code": 0,
            "data": "",
            "msg":  "not found",
        })
    }

}

//删除数据
func DelHelp(c *gin.Context) {
    id, _ := strconv.Atoi(c.Request.FormValue("id"))
    h := Help{Id: id}
    h.GetDetail()
    if h.Id > 0 {
        rs, _err := h.Del()
        if _err != nil {
            log.Fatalln(_err)
        }

        msg := fmt.Sprintf("delete success %d", rs)
        c.JSON(http.StatusOK, gin.H{
            "code": 1,
            "data": true,
            "msg":  msg,
        })
    } else {
        c.JSON(http.StatusOK, gin.H{
            "code": 0,
            "data": false,
            "msg":  "delete faild",
        })
    }

}