package models

import (
  "fmt"
  "log"
)

type Post struct {
  Id       int64  `db:"id"        json:"id"`
  Title    string `db:"title"     json:"title"`
  Body     string `db:"body"      json:"body"`
  AuthorId int    `db:"author_id" json:"author_id"`
}

type AllPosts struct {
  Page    int
  PerPage int
}

var PostPerPage = 10

func (posts *AllPosts) GetAllPosts() string {
  var sql string

  sql = fmt.Sprintf("SELECT * FROM posts LIMIT %v", posts.PerPage)

  if(posts.Page != 1) {
    offset := (posts.Page - 1) * posts.PerPage

    temp_sql := fmt.Sprintf("%v OFFSET %v", sql, offset)
    sql = temp_sql
  }

  log.Println(sql)

  return sql
}
