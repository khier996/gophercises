package main

import (
  "net/http"
  "net/url"
  "log"
  "fmt"
  "golang.org/x/net/html"
  "flag"
  "strings"
  "encoding/xml"
  "os"

  linkParser "github.com/khier996/gophercises/exercise-4-html-link-parser/link-parser"
)

type Url struct {
  Loc string `xml:"loc"`
}

type UrlSet struct {
  XMLName   xml.Name `xml:"http://www.sitemaps.org/schemas/sitemap/0.9 urlset"`
  UrlSet []Url `xml:"url"`
}

var hostname string
var parsedLinks = map[string]bool{"/": true}
var allLinks []linkParser.Link

func main() {
  urlString := flag.String("url", "https://www.calhoun.io/", "url of a website")
  flag.Parse()

  if urlObj, err := url.Parse(*urlString); err != nil {
    log.Fatal("Error parsing url")
  } else {
    hostname = urlObj.Hostname()
  }

  landingPageUrl := "http://" + hostname
  scrapeHref(landingPageUrl)

  fmt.Println("all links", allLinks)
  writeXmlFile()
}

func writeXmlFile() {
  urlMap := buildUrlMap()

  if file, fileErr := os.Create("data.xml"); fileErr != nil {
    log.Fatal("error creating file", fileErr)
  } else {
    encoder := xml.NewEncoder(file)
    if encErr := encoder.Encode(urlMap); encErr != nil {
      log.Fatal("error encoding xml", encErr)
    }
  }
}

func scrapeHref(href string) {
  fmt.Println("scraping url", href)

  doc := getDoc(href)
  links := linkParser.Parse(doc)
  links = cleanLinks(links)

  links = chooseNewDomainLinks(links)
  allLinks = append(allLinks, links...)

  traverseLinks(links)
}

func cleanLinks(links []linkParser.Link) []linkParser.Link{
  var cleanLinks []linkParser.Link
  for _, link := range links {
    if link.Href[0:1] == "#" { continue } // links that start with # point to the current page
    cleanLink(&link)
    cleanLinks = append(cleanLinks, link)
  }
  return cleanLinks
}

func cleanLink(link *linkParser.Link) {
  if len(link.Href) > 4 && link.Href[0:4] == "https" {
    link.Href = link.Href[0:4] + link.Href[5:] // substitute https with http
  } else if link.Href[0:1] == "/" {
    link.Href = "http://" + hostname + link.Href
  } else if link.Href[0:3] == "www" {
    link.Href = "http://" + link.Href
  } else if link.Href[0:4] != "http" {
    link.Href = "http://" + link.Href
  }

  link.Href = strings.Split(link.Href, "#")[0]
}

func traverseLinks(links []linkParser.Link) {
  for _, link := range links {
    scrapeHref(link.Href)
  }
}

func getDoc(url string) *html.Node {
  res, err := http.Get(url)
  if err != nil {
    log.Fatal("error with request", err)
  }
  doc, err := html.Parse(res.Body)
  if err != nil {
    log.Fatal("Error parsing html document", err)
  }
  return doc
}

func chooseNewDomainLinks(links []linkParser.Link) []linkParser.Link {
  var filteredLinks []linkParser.Link
  for _, link := range links {
    urlObj, err := url.Parse(link.Href)
    if err != nil {
      fmt.Println("error parsing url")
      continue
    }

    linkHostname := urlObj.Hostname()
    if linkHostname == hostname && !parsedLinks[link.Href] {
      parsedLinks[link.Href] = true
      filteredLinks = append(filteredLinks, link)
    }
  }

  return filteredLinks
}


func buildUrlMap() UrlSet {
  var urlMap []Url
  for _, link := range allLinks {
    newUrl := Url{link.Href}
    urlMap = append(urlMap, newUrl)
  }
  return UrlSet{UrlSet: urlMap}
}
