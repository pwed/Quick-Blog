package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func posts(w http.ResponseWriter, r *http.Request) {

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
