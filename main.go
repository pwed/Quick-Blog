package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	bolt "github.com/coreos/bbolt"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("names"))
		err := b.Put([]byte("name"), []byte(ps.ByName("name")))
		return err
	})
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Names(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var names []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("names"))
		if b == nil {
			return fmt.Errorf("Bucket %q not found!", "names")
		}
		v := b.Get([]byte("name"))
		fmt.Print(string(v))
		names = v
		return nil
	})
	fmt.Fprintf(w, "names =, %s!\n", string(names))
}

// Open the my.db data file in your current directory.
// It will be created if it doesn't exist.
// Process will timeout after 1 second if file can't be opened.
var db, err = bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

func main() {
	db.Update(func(tx *bolt.Tx) error {

		tx.CreateBucketIfNotExists([]byte("names"))

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/names", Names)

	log.Fatal(http.ListenAndServe(":8080", router))
}
