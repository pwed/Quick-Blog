package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/HouzuoGuo/tiedot/db"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Println("Welcome to the homepage")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	n := v["name"]
	fmt.Println("About to request collection")
	col := myDB.Use("names")
	fmt.Println("collection requested")
	col.Insert(map[string]interface{}{
		"name": n,
	})
	fmt.Println("Data added")
	fmt.Fprintf(w, "hello, %s!\n", n)
}

func Names(w http.ResponseWriter, r *http.Request) {
	var names string
	col := myDB.Use("names")
	col.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		names = names + fmt.Sprintln("Document", id, "is", string(docContent))
		return true  // move on to the next document OR
		return false // do not move on to the next document
	})
	fmt.Fprintf(w, "names =, %s!\n", string(names))
}

func Post(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	col := myDB.Use("posts")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)

	post, _ := col.Read(id)

	response, _ := json.Marshal(post)
	fmt.Fprint(w, string(response))
}

var myDB *db.DB

type PostData struct {
	Author string
	Body   string
}

func main() {
	// (Create if not exist) open a database
	var err error
	myDB, err = db.OpenDB("tmp/MyDatabase")
	if err != nil {
		panic(err)
	}
	defer myDB.Close()

	myDB.Create("names")
	myDB.Create("users")
	myDB.Create("posts")
	myDB.Create("comments")
	posts := myDB.Use("posts")
	posts.InsertRecovery(1, map[string]interface{}{
		"Author": "Pwed",
		"Body":   "hello",
		"Title":  "Test",
	})

	m := mux.NewRouter()
	m.HandleFunc("/api", Index)
	m.HandleFunc("/api/hello/{name}", Hello)
	m.HandleFunc("/api/names", Names)
	m.HandleFunc("/api/post/{id:[1-9]+}", Post)
	m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":8080", m))
}
