package main

import (
  "log"
  "net/http"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/binding"
  "github.com/martini-contrib/render"

  "./config"
  "./router"

  "./models"

  "github.com/coopernurse/gorp"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options{
    IndentJSON: true,
  }))
  var dbmap *gorp.DbMap

  dbmap = config.InitDb()
  defer dbmap.Db.Close()

  m.Map(dbmap)

  m.Group("/posts", func(r martini.Router) {
    r.Get("", router.GetPosts)
    r.Get("/:id", router.GetPost)
    r.Post("", binding.Json(models.Post{}), router.CreatePost)
  })

  log.Fatal(http.ListenAndServe("localhost:3000", m))
}
