package model

import . "../database"

//表结构
type Help struct {
    Id           int    `json:"id" form "id"`
    Title        string `json:"title" form:"title"`
    Content      string `json:"content" form:"content"`
    Brief        string `json:"brief" form:"brief"`
    Keywords     string `json:"keywords" form:"keywords"`
    Hits         string `json:"hits" form:"hits"`
    IsPublish    string `json:"is_publish" form:"is_publish"`
    PublishTime  string `json:"publish_time" form:"publish_time"`
    AddTime      string `json:"add_time" form:"add_time"`
}

//表名
const tableName  = "vs_help"

//获取列表
func (h *Help) GetList(page int, limit int) (helps []Help, err error) {
    //创建一个切片
    helps = make([]Help, 0)
    rows, err := SqlDB.Query("select * from "+ tableName +" order by id desc")
    if err != nil {
        return
    }

    for rows.Next() {
        var help Help
        rows.Scan(&help.Id, &help.Title, &help.Content, &help.Brief, &help.Keywords, &help.Hits, &help.IsPublish, &help.PublishTime, &help.AddTime)
        helps = append(helps, help)
    }

    if err = rows.Err(); err != nil {
        return
    }

    return
}

//获取一条记录
func (h *Help) GetDetail() (err error) {

    SqlDB.QueryRow("select * from "+ tableName +" where `id` = ?", h.Id).Scan(&h.Id, &h.Title, &h.Content, &h.Brief, &h.Keywords, &h.Hits, &h.IsPublish, &h.PublishTime, &h.AddTime)
    return
}

//插入
func (h *Help) Add() (lastId int64, err error) {
    rs, err := SqlDB.Exec("INSERT INTO "+ tableName +"(`title`,`content`)VALUES(?, ?)", h.Title, h.Content)
    if err != nil {
        return
    }
    lastId, err = rs.LastInsertId()
    return lastId, err
}

//修改记录
func (h *Help) Update() (effect_num int64, err error) {
    rs, err := SqlDB.Exec("UPDATE "+ tableName + " SET `title`= ?,`content`= ? where `id`= ?", h.Title, h.Content, h.Id)
    if err != nil {
        return
    }
    effect_num, err = rs.RowsAffected()
    return effect_num, err
}

//删除记录
func (h *Help) Del() (effort_num int64, err error) {
    rs, err := SqlDB.Exec("DELETE FROM "+ tableName +" where id= ?", h.Id)
    if err != nil {
        return
    }
    effort_num, err = rs.RowsAffected()
    return
}
