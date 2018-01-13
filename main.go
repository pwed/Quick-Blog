package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/HouzuoGuo/tiedot/db"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
	fmt.Println("Welcome to the homepage")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("About to request collection")
	col := myDB.Use("names")
	fmt.Println("collection requested")
	col.Insert(map[string]interface{}{
		"name": ps.ByName("name"),
	})
	fmt.Println("Data added")
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Names(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	col := myDB.Use("names")
	fmt.Println("collection requested")
	col.Insert(map[string]interface{}{
		"name": "pwed",
	})
	fmt.Println("Data added")

	col = myDB.Use("names")
	col.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
		fmt.Println("Document", id, "is", string(docContent))
		return true  // move on to the next document OR
		return false // do not move on to the next document
	})

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/names", Names)

	log.Fatal(http.ListenAndServe(":8080", router))
}
