package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

// GET
// Home page.
// URL: http://localhost:10000
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// GET
// Return all articles.
// URL: http://localhost:10000/all
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// GET
// Return an article with give id, else do nothing.
// URL: http://localhost:10000/article/{id}
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	fmt.Fprintf(w, "Article with id: %s is not found.", key)
}

// POST
// Create a new article.
// URL: http://localhost:10000/article
// Payload:
//
//	{
//	   "Id": "3",
//	   "Title": "Newly Created Post",
//	   "desc": "The description for my new post",
//	   "content": "my articles content"
//	}
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := io.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

// DELETE
// Delete an article with give id, else do nothing.
// URL: http://localhost:10000/article/{id}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	for idx, article := range Articles {
		if article.Id == key {
			Articles = append(Articles[:idx], Articles[idx+1:]...)
			return
		}
	}
	fmt.Fprintf(w, "Article with id: %s is not found.", key)
}

// PUT
// Update an article with given id, else do nothing.
// URL: http://localhost:10000/article/{id}
// Payload:
//
//	{
//	   "Id": "3",
//	   "Title": "Newly Created Post",
//	   "desc": "The description for my new post",
//	   "content": "my articles content"
//	}
func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	var updateArt Article
	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updateArt)
	for idx, article := range Articles {
		if article.Id == key {
			article.Title = updateArt.Title
			article.Desc = updateArt.Desc
			article.Content = updateArt.Content
			Articles[idx] = article
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	fmt.Fprintf(w, "Article with id: %s is not found.", key)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/all", returnAllArticles)
	router.HandleFunc("/article", createNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	Articles = append(Articles, Article{Id: "1", Title: "Hello1", Desc: "Article Description", Content: "Article Content"})
	Articles = append(Articles, Article{Id: "2", Title: "Hello2", Desc: "Article Description", Content: "Article Content"})
	handleRequests()
}
