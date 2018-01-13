package main

import (
	"fmt"
	"log"
	"net/http"

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

var myDB *db.DB

func main() {
	// (Create if not exist) open a database
	var err error
	myDB, err = db.OpenDB("tmp/MyDatabase")
	if err != nil {
		panic(err)
	}
	defer myDB.Close()

	myDB.Create("names")

	m := mux.NewRouter()
	m.HandleFunc("/api", Index)
	m.HandleFunc("/api/hello/{name}", Hello)
	m.HandleFunc("/api/names", Names)
	m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":8080", m))
}
