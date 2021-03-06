package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var yamlPath string
	var jsonPath string
	flag.StringVar(&yamlPath, "yaml", "config.yaml", "a yaml config path")
	flag.StringVar(&jsonPath, "json", "config.json", "a json config path")
	flag.Parse()

	yaml, err := ioutil.ReadFile(yamlPath)
	if err == nil {
		mapHandler, err = YAMLHandler(yaml, mapHandler)
		if err != nil {
			panic(err)
		}
	}
	json, err := ioutil.ReadFile(jsonPath)
	if err == nil {
		mapHandler, err = JSONHandler(json, mapHandler)
		if err != nil {
			panic(err)
		}
	}

	//	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
