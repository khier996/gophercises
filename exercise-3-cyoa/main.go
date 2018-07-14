package main

import (
  "net/http"
  "html/template"
  "encoding/json"
  "io/ioutil"
  "fmt"
)

type Story map[string]Chapter

type Chapter struct {
  Title string `json:"title"`
  Paragraphs []string `json:"story"`
  Options []Option `json:"options"`
}

type Option struct {
  Text string `json:"text"`
  Arc string `json:"arc"`
}

var chapters map[string]Chapter
var templates *template.Template

func main() {
  downloadChapters()
  downloadTemplates()

  handler := newHandler()
  http.ListenAndServe(":8080", handler)
}

func downloadChapters() {
  if jsonData, err := ioutil.ReadFile("gopher.json"); err != nil {
    fmt.Println("error reading json file:", err)
    return
  } else {
    err := json.Unmarshal(jsonData, &chapters)
    if err != nil {
      fmt.Println("error Unmarshaling", err)
    }
  }
}

func downloadTemplates() {
  var err error
  if templates, err = template.ParseFiles("templates/main.html"); err != nil {
    fmt.Println("error downloading templates", err)
  }
}

func newHandler() http.Handler {
  return handler{chapters, defaultPathFn}
}

type handler struct {
  chapters map[string]Chapter
  defaultPathFn func(r *http.Request) string
}

func defaultPathFn(r *http.Request) string {
  path := r.URL.Path
  if path == "" || path == "/" {
    path = "/intro"
  }

  path = path[1:]
  return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := h.defaultPathFn(r)

  if chapter, ok := chapters[path]; ok {
    if err := templates.ExecuteTemplate(w, "main.html", chapter); err != nil {
      fmt.Println("error executing template", err)
    }
  } else {
    http.Error(w, "Chapter not found...", http.StatusNotFound)
  }
}





