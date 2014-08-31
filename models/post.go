package models

type Post struct {
  Id       int64  `db:"id"`
  Title    string `db:"title"`
  Body     string `db:"body"`
  AuthorId int    `db:"author_id"`
}

type JSONPost struct {
  Id       int     `json:"id"`
  Title    string  `json:"title"`
  Body     string  `json:"body"`
  AuthorId int     `json:"author_id"`
}
