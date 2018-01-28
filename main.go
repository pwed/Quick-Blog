package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/HouzuoGuo/tiedot/db"
)

var myDB *db.DB

// PostData cantains all data relating to a post
type PostData struct {
	Author string `json:"author"`
	Body   struct {
		MD   string `json:md,omitempty`
		HTML string `json:html,omitempty`
	} `json:body,omitempty`
	Date string `json:"date,omitempty"`
	ID   int    `json:"id,omitempty"`
}

func main() {
	// (Create if not exist) open a database
	var err error
	myDB, err = db.OpenDB("data/db")
	if err != nil {
		panic(err)
	}
	defer myDB.Close()

	// Create tables if not exist
	myDB.Create("users")
	myDB.Create("posts")
	myDB.Create("comments")

	// Create http router
	m := mux.NewRouter()
	m.HandleFunc("/api/newpost", newPost).Methods("POST")
	m.HandleFunc("/api/post/{id}", post).Methods("GET")
	m.HandleFunc("/api/posts/{number}", posts).Methods("GET")

	// if in development mode, will serve static folder
	// else serve embedded compiled static files
	if len(os.Args) > 1 && os.Args[1] == "dev" {
		m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	} else {
		m.PathPrefix("/").Handler(http.FileServer(assetFS()))
	}

	fmt.Println("Starting server @ http//:localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", m))
}
