package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
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

	w.Header().Set("Content-Type", "text/json")
	col := myDB.Use("posts")
	vars := mux.Vars(r)
	idString := vars["id"]
	id, _ := strconv.Atoi(idString)
	fmt.Printf("ID: %v, is of type %T\n", id, id)

	post, _ := col.Read(int(id))

	response, _ := json.Marshal(post)
	fmt.Fprint(w, string(response))
	fmt.Println(string(response))
}

func Posts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/json")
	col := myDB.Use("posts")
	vars := mux.Vars(r)
	numString := vars["number"]
	num, _ := strconv.Atoi(numString)
	count := 0
	var ids []string

	col.ForEachDoc(func(id int, doc []byte) bool {

		ids = append(ids, fmt.Sprint(id))
		count++
		if count == num {
			return false
		}
		return true
	})

	response, _ := json.Marshal(ids)

	fmt.Fprint(w, string(response))
	fmt.Println(string(response))
}

func NewPost(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var pd PostData
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		http.Error(w,
			"Please send a valid body including a PostData object encoded in JSON format",
			http.StatusBadRequest)
		return
	}

	col := myDB.Use("posts")

	id, err := col.Insert(structs.Map(pd))
	if err != nil {
		http.Error(w, "Had trouble adding record to the database",
			http.StatusInternalServerError)
		return
	}
	fmt.Printf("ID: %v, is of type %T\n", id, id)

	response := fmt.Sprintf(`{"Success":true,"newID":"%v"}`, id)
	fmt.Fprint(w, response)
	fmt.Println(response)
	count := col.ApproxDocCount()

	fmt.Println(count)
}

var myDB *db.DB

type PostData struct {
	Author string `json:"author"`
	Body   string `json:"body"`
	Date   string `json:"date,omitempty"`
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

	col, err := db.OpenCol(myDB, "posts")
	count := col.ApproxDocCount()

	fmt.Print(count, "\n")

	m := mux.NewRouter()
	m.HandleFunc("/api/newpost", NewPost)
	m.HandleFunc("/api/post/{id}", Post).Methods("GET")
	m.HandleFunc("/api/posts/{number}", Posts).Methods("GET")
	m.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	fmt.Println("Starting server @ http//:localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", m))
}
