package main

import (
	"github.com/julienschmidt/httprouter"
)

func GenerateRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/posts", PostsIndex)
	r.GET("/posts/:id", PostsShow)
	r.POST("/posts", PostsCreate)
	r.PUT("/posts/:id", PostsUpdate)
	r.PATCH("/posts/:id", PostsUpdate)
	r.DELETE("/posts/:id", PostsDestroy)

	return r
}
