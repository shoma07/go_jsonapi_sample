package main

import (
	"time"
)

type Post struct {
	ID        int       `jsonapi:"primary,post"`
	Title     string    `jsonapi:"attr,title"`
	Body      string    `jsonapi:"attr,body"`
	CreatedAt time.Time `jsonapi:"attr,created_at"`
	UpdatedAt time.Time `jsonapi:"attr,updated_at"`
}
