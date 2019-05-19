package database

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

//初始化
func init() {
    var err error
    //https://www.cnblogs.com/zhja/p/5604553.html
    //db, err := sql.Open("mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")
    SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_vcstock?charset=utf8")
    // SqlDB, err = sql.Open("mysql", "beta-mysql:nopass.2@tcp(10.255.72.27:3306)/test?parseTime=true")
    if err != nil {
        log.Fatal(err.Error())
    }

    err = SqlDB.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
}