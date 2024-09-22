package main

import (
	"fmt"
	"net/http"
)

func mapHandler(pathToURLs map[string]string, fallback http.Handler) http.HandlerFunc {

	// fmt.Println("map handler")

	// for k, v := range pathToURLs {
	// 	if k == "/yaml-godoc" {
	// 		fmt.Println("pathtourl", k)
	// 		handler := http.RedirectHandler(v, 400)
	// 		return handler.ServeHTTP
	// 	}
	// }
	// return fallback.ServeHTTP

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.URL.Path: %v\n", r.URL.Path)
		if url, ok := pathToURLs[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}

	}
}

func YamlHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	return nil, nil

}

// func redirectUrl() {

// }
