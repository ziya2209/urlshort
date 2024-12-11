package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"urlshort/url"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := url.MapHandler(pathsToUrls, mux) 

	ymlData, err := os.ReadFile("path.yaml")
	if err != nil {
		log.Fatal(err)
	}
	

	ymlHandler, err := url.YAMLHandler(ymlData, mapHandler)

	err = http.ListenAndServe(":8080", ymlHandler) 
	if err != nil {
		fmt.Println("error is found", err)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
