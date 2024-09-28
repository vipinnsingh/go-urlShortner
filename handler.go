package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Url struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type Config struct {
	Urls []Url
}

func mapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {

	fmt.Println("map handler", pathToURLs)

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.URL.Path: %v\n", r.URL.Path)
		if url, ok := pathToURLs[r.URL.Path]; ok {
			fmt.Println("redirecting")
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fmt.Println("not found")
			fallback.ServeHTTP(w, r)
		}

	}
}

func YAMLHandler(fallback http.Handler) (http.HandlerFunc, error) {

	file, err := os.ReadFile("url.yaml")
	if err != nil {
		fmt.Println(err)
	}

	config, err := parseYAML(file)
	if err != nil {
		return nil, err
	}

	pathMap := buildMap(config)
	return mapHandler(pathMap, fallback), nil
}

func parseYAML(file []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(file, &config)
	if err != nil {
		fmt.Println("error in unmarshalling", err)
	}

	return config, nil
}

func buildMap(config Config) map[string]string {
	pathMap := make(map[string]string)

	for _, v := range config.Urls {
		pathMap[v.Path] = v.Url
	}

	return pathMap
}
