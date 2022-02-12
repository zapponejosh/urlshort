package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joshzappone/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	// mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlUrls, err := os.ReadFile("urls.yaml")
	if err != nil {
		os.Exit(2)
		fmt.Printf("Oh no! Error: %s", err)
	}

	// json urls

	yamlHandler, err := urlshort.YAMLHandler(yamlUrls, mux)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
