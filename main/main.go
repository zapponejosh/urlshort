package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshzappone/urlshort"
)

func main() {
	mux := defaultMux()

	yamlUrls, err := os.ReadFile("urls.yaml")
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(yamlUrls, mux)
	if err != nil {
		panic(err)
	}

	// json urls
	jsonUrls, err := os.ReadFile("urls.json")
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler(jsonUrls, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
