package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := defaultMux()

	// pathToURLs := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }

	// handler := mapHandler(pathToURLs, mux)
	handler, _ := YAMLHandler(mux)

	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		fmt.Println(err)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
