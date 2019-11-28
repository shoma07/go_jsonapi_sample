package main

import (
	"fmt"
	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func PostsIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := gorm.Open("sqlite3", "db/development.sqlite3")
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	defer db.Close()

	posts := []Post{}
	db.Find(&posts)

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := gorm.Open("sqlite3", "db/development.sqlite3")
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	defer db.Close()

	post := &Post{}

	if db.First(&post, ps.ByName("id")).RecordNotFound() {
		NotFoundHandler(w, r)
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := gorm.Open("sqlite3", "db/development.sqlite3")
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	defer db.Close()

	post := &Post{}

	if err := jsonapi.UnmarshalPayload(r.Body, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db.Create(post)

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusCreated)

	if err := jsonapi.MarshalPayload(w, post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsUpdate")
}

func PostsDestroy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsDestroy")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusNotFound)

	if err := jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Title:  "Not Resource Found",
		Detail: "Not Resource Found",
		Code:   "404",
	}}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
