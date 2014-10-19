package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"

	"github.com/koleksiuk/goblog/models"
)

func GetPosts(args martini.Params, r render.Render, dbmap *gorp.DbMap, req *http.Request) {
	var posts []models.Post

	page, perPage := HandlePageParams(req)

	postsSql := &models.AllPosts{Page: page, PerPage: perPage}

	_, err := dbmap.Select(&posts, postsSql.GetAllPosts())

	log.Printf("%#v", posts)

	if err != nil {
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
	var post models.Post

	err := dbmap.SelectOne(&post, "SELECT * FROM posts WHERE id = ?", args["id"])

	if err != nil {
		jsonError := &models.JSONError{
			Error: "Post not found",
		}

		r.JSON(404, jsonError)
	} else {
		r.JSON(200, post)
	}
}

func CreatePost(post models.Post, r render.Render, dbmap *gorp.DbMap) {
	log.Printf("%#v", post)

	err := dbmap.Insert(&post)
	if err != nil {
		log.Fatal("Something went wrong", err)
		jsonError := &models.JSONError{
			Error: "err",
		}

		r.JSON(404, jsonError)
	} else {
		r.JSON(200, post)
	}
}

func HandlePageParams(req *http.Request) (int, int) {
	var perPage int
	var page int
	var err error

	args := req.URL.Query()

	if args["per_page"] != nil {
		perPage, err = strconv.Atoi(args["per_page"][0])
		if err != nil {
			perPage = models.PostPerPage
		}
	} else {
		perPage = models.PostPerPage
	}

	if args["page"] != nil {
		page, err = strconv.Atoi(args["page"][0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	return page, perPage
}
