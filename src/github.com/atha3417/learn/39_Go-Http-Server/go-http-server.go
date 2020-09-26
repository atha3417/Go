package main;

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// "io/ioutil"
);

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "Judul pertama", Desc: "Deskripsi pertama"},
	Article{Title: "Judul kedua", Desc: "Deskripsi kedua"},
};

func main() {
	http.HandleFunc("/", getHome);
	http.HandleFunc("/articles", getArticles);
	http.HandleFunc("/articles/post", withLogging(postArticle));
	http.ListenAndServe(":3000", nil);
}

func getHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Testing kamu berada di home"));
}

func getArticles(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(articles);
}

func postArticle(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var newArticle Article;
		var error = json.NewDecoder(req.Body).Decode(&newArticle);

		if error != nil {
			fmt.Printf("Ada error", error);
		}
		articles = append(articles, newArticle);
		json.NewEncoder(res).Encode(articles);
	} else {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed);
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func (res http.ResponseWriter, req *http.Request) {
		log.Printf("Logged koneksi dari ", req.RemoteAddr);
		next.ServeHTTP(res, req);
	}
}
