package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Static files dirを定義
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.HandleFunc("/", indexHandler)
	http.HandleFunc("/", templatesHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// indexHandler 単純なhandler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "index handler!!")
}

// templatesHandler 単純なTemplatesParse handler
func templatesHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() { // 実行速度
		end := time.Now()
		log.Println("exec time: ", end.Sub(start))
	}()

	// templates dir 以下をparseし、下記 "layout"を読み関連defineを挿入する
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.ExecuteTemplate(w, "layout", nil)
}
