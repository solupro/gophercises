package main

import (
	"encoding/json"
	"github.com/fatedier/frp/utils/log"
	"gopkg.in/yaml.v2"
	"net/http"
)

func YAMLHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yamlConfig := []Config{}
	err := yaml.Unmarshal(content, &yamlConfig)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(yamlConfig)
	return MapHandler(pathMap, fallback), nil
}

func JSONHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	jsonConfig := []Config{}
	err := json.Unmarshal(content, &jsonConfig)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(jsonConfig)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(c []Config) map[string]string {
	m := make(map[string]string)
	for _, v := range c {
		m[v.Path] = v.Url
	}

	return m
}

func MapHandler(m map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := m[r.URL.Path]; ok {
			log.Info("redirect to %s", url)
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
