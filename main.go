package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// DB stuff here
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Names(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var names []byte
	// DB stuff here
	fmt.Fprintf(w, "names =, %s!\n", string(names))
}

func main() {

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/names", Names)

	log.Fatal(http.ListenAndServe(":8080", router))
}
