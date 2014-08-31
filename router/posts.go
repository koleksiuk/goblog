package router

import (
  "log"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "github.com/coopernurse/gorp"
  _ "github.com/go-sql-driver/mysql"

  "../models"
)

func GetPosts(r render.Render, dbmap *gorp.DbMap) {
  var posts []models.Post

  _, err := dbmap.Select(&posts, "SELECT * FROM posts")

  log.Println(posts)

  if(err != nil) {
    log.Fatal("Something went wrong")

    jsonError := &models.JSONError{
      Error: "Something went wrong",
    }

    r.JSON(500, jsonError)
  } else {
    r.JSON(200, posts)
  }
}

func GetPost(args martini.Params, r render.Render, dbmap *gorp.DbMap) {
  var post     models.Post

  err := dbmap.SelectOne(&post, "SELECT * FROM posts WHERE id = ?", args["id"])

  if(err != nil) {
    jsonError := &models.JSONError{
      Error: "Post not found",
    }

    r.JSON(404, jsonError)
  } else {
    r.JSON(200, post)
  }
}
