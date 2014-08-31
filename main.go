package main

import (
  "log"
  "net/http"
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"

  "./config"
  "./router"

  "github.com/coopernurse/gorp"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  var dbmap *gorp.DbMap

  dbmap = config.InitDb()
  defer dbmap.Db.Close()

  m.Map(dbmap)

  m.Get("/posts", router.HandlePosts)
  m.Get("/posts/:id", router.HandlePost)

  log.Fatal(http.ListenAndServe("localhost:3000", m))
}
