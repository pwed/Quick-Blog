package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func post(w http.ResponseWriter, r *http.Request) {

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
