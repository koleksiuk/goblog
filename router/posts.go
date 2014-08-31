package router

import (
  "../models"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "github.com/coopernurse/gorp"
  _ "github.com/go-sql-driver/mysql"
)

func HandlePosts(r render.Render) {
  r.JSON(200, nil)
}

func HandlePost(args martini.Params, r render.Render, dbmap *gorp.DbMap) {
  var post     models.Post

  err := dbmap.SelectOne(&post, "SELECT * FROM posts WHERE id = ?", args["id"])

  if(err != nil) {
    jsonError := &models.JSONError{
      Error: "Post not found",
    }

    r.JSON(404, jsonError)
  } else {
    id := int(post.Id)

    jsonPost := &models.JSONPost{
      Id:       id,
      Title:    post.Title,
      Body:     post.Body,
      AuthorId: post.AuthorId,
    }

    r.JSON(200, jsonPost)
  }
}
