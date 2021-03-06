package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))

}

// Article struct
type Article struct {
	Id string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles array global
var Articles []Article

func main() {

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Desc", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article desc 2", Content: "Article Content 2"},
	}
	handleRequests()
}
