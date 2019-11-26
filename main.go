package main

import (
	"flag"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/netutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time
}

func PostsIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var posts []*Post
	db, err := gorm.Open("sqlite3", "db/development.sqlite3")
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	defer db.Close()

	db.Find(&posts)

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	if err := jsonapi.MarshalPayload(w, posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostsShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsShow")
}

func PostsCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsCreate")
}

func PostsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsUpdate")
}

func PostsDestroy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "PostsDestroy")
}

func main() {
	router := httprouter.New()
	router.GET("/posts", PostsIndex)
	router.GET("/posts/:id", PostsShow)
	router.POST("/posts", PostsCreate)
	router.PUT("/posts/:id", PostsUpdate)
	router.PATCH("/posts/:id", PostsUpdate)
	router.DELETE("/posts/:id", PostsDestroy)

	var (
		host     string
		port_str string
	)

	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.StringVar(&port_str, "p", "80", "port")
	flag.Parse()

	port, err := strconv.ParseUint(port_str, 10, 16)

	if err != nil {
		log.Fatalln(err)
	}

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		log.Fatalln(err)
	}

	limit_ln := netutil.LimitListener(ln, 100)
	defer limit_ln.Close()

	server := &http.Server{
		Handler: router,
	}

	err = server.Serve(limit_ln)

	if err != nil {
		log.Fatalln(err)
	}
}
