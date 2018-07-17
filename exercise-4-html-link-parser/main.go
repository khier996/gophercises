package main

import (
  "golang.org/x/net/html"
  "os"
  "fmt"
  "log"
  "flag"

  linkParser "github.com/khier996/gophercises/exercise-4-html-link-parser/link-parser"
)

func main() {
  filePath := flag.String("file-path", "html/ex1.html", "Path to html file")
  flag.Parse()

  file, err := os.Open(*filePath)
  if err != nil {
    log.Fatal("Error opening html file", err)
  }

  doc, err := html.Parse(file)
  if err != nil {
    log.Fatal("Error parsing html document", err)
  }

  links := linkParser.Parse(doc)
  fmt.Printf("These are links: %+v", links)
}

