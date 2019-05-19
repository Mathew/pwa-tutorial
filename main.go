package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getStatic(filename string) []byte {
	file, err := ioutil.ReadFile(fmt.Sprintf("static/%v", filename))
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func htmlFileHandler(filename string, contentType string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Serving: %v\n", filename)
		content := getStatic(filename)

		if contentType == "" {
			contentType = "text/html; charset=utf-8"
		}
		w.Header().Set("Content-Type", contentType)
		w.Write(content)
	}
}

func homeHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
		} else {
			handler(w, r)
		}
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/offline", htmlFileHandler("offline.html", ""))
	http.HandleFunc("/service-worker.js", htmlFileHandler("service-worker.js", "text/javascript"))
	http.HandleFunc("/", homeHandler(htmlFileHandler("index.html", "")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
