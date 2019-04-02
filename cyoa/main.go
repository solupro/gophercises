package main

import (
	"encoding/json"
	"flag"
	"github.com/fatedier/frp/utils/log"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	var jsonPath string
	flag.StringVar(&jsonPath, "json", "gopher.json", "a json config path")
	flag.Parse()

	jsonFile, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		panic(err)
	}

	story := Story{}
	err = json.Unmarshal(jsonFile, &story)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc := strings.TrimLeft(r.URL.Path, "/")
		if arc == "" {
			arc = "intro"
		}
		log.Info("accept %s", arc)

		if chapter, ok := story[arc]; ok {
			temp, err := template.ParseFiles("index.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, chapter)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)

}

type Story map[string]Chapter

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
