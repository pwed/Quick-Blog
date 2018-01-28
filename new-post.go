package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func newPost(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var np NewPost

	var pd PostData
	err := json.NewDecoder(r.Body).Decode(&np)
	if err != nil {
		http.Error(w,
			"Please send a valid body including a PostData object encoded in JSON format",
			http.StatusBadRequest)
		return
	}

	pd = np.createPostData()

	pd.mdToHTML(true)

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

func (pd *PostData) mdToHTML(safe bool) {
	result := []byte(pd.Body.MD)
	result = blackfriday.Run(result, blackfriday.WithExtensions(blackfriday.CommonExtensions))
	if safe {
		result = bluemonday.UGCPolicy().SanitizeBytes(result)
	}
	pd.Body.HTML = string(result)
}

type NewPost struct {
	Author string `json:author`
	Body   string `json:body`
}

func (np *NewPost) createPostData() PostData {
	var pd PostData
	pd.Author = np.Author
	pd.Body.MD = np.Body
	return pd
}
