package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		file, err := ioutil.ReadFile("static/index.html")
		if err != nil {
			log.Fatal(err)
		}

		writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		writer.Write(file)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
