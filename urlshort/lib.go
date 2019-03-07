package main

import (
	"github.com/fatedier/frp/utils/log"
	yaml2 "gopkg.in/yaml.v2"
	"net/http"
)

func YAMLHandler(yaml []byte, fallback *http.ServeMux) (*http.ServeMux, error) {
	yamlConfig := []Config{}
	err := yaml2.Unmarshal(yaml, &yamlConfig)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(yamlConfig)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(c []Config) map[string]string {
	m := make(map[string]string)
	for _, v := range c {
		m[v.Path] = v.Url
	}

	return m
}

func MapHandler(m map[string]string, mux *http.ServeMux) *http.ServeMux {
	for path, url := range m {
		toUrl := url
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			log.Info("redirect to %s", toUrl)
			http.Redirect(w, r, toUrl, http.StatusTemporaryRedirect)
		})
	}

	return mux
}