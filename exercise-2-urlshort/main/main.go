package main

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "flag"

	"github.com/khier996/gophercises/exercise-2-urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

  yamlFilePath := flag.String("yaml", "path-urls.yaml", "Path to URLs yaml file")
  jsonFilePath := flag.String("json", "path-urls.json", "Path to URLs json file")
  useFile := flag.String("use-file", "yaml", "Which file to use")
  flag.Parse()

  yamlHandler := buildYamlHandler(yamlFilePath, mapHandler)
  jsonHandler := buildJSONHandler(jsonFilePath, mapHandler)

	fmt.Println("Starting the server on :8080")

  if *useFile == "json" {
    http.ListenAndServe(":8080", jsonHandler)
  } else {
    http.ListenAndServe(":8080", yamlHandler)
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

func buildYamlHandler(yamlFilePath *string, fallback http.Handler) http.Handler {
  yaml, yamlErr := ioutil.ReadFile(*yamlFilePath)
  if yamlErr != nil {
    fmt.Println("Error opening yaml file", yamlErr)
  }

  yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), fallback)
  if err != nil {
    panic(err)
  }
  return yamlHandler
}

func buildJSONHandler(jsonFilePath *string, fallback http.Handler) http.Handler {
  json, jsonErr := ioutil.ReadFile(*jsonFilePath)
  if jsonErr != nil {
    fmt.Println("error with opening json file", jsonErr)
  }
  jsonHandler, err := urlshort.JSONHandler(json, fallback)
  if err != nil {
    panic(err)
  }
  return jsonHandler
}
