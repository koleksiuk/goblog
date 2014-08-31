package models

type Post struct {
  Id       int64  `db:"id"        json:"id"`
  Title    string `db:"title"     json:"title"`
  Body     string `db:"body"      json:"body"`
  AuthorId int    `db:"author_id" json:"author_id"`
}
