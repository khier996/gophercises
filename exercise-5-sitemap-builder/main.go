package main

import (
  "net/http"
  "net/url"
  "log"
  "fmt"
  "golang.org/x/net/html"
  "flag"
  "strings"

  linkParser "github.com/khier996/gophercises/exercise-4-html-link-parser/link-parser"
)

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
}

func scrapeHref(href string) {
  href = strings.Split(href, "#")[0] // links with different parts after "#" actually point to the same page
  parsedLinks[href] = true
  urlString := buildUrlFromHref(href)

  fmt.Println("scraping url", urlString)

  doc := getDoc(urlString)
  links := linkParser.Parse(doc)
  links = chooseDomainLinks(links)
  allLinks = append(allLinks, links...)

  traverseLinks(links)
}

func buildUrlFromHref(href string) string {
  var urlString string
  if href[0:1] == "/" {
    urlString = "http://" + hostname + href
  } else if href[0:4] != "http" {
    urlString = "http://" + href
  } else {
    urlString = href
  }
  return urlString
}


func traverseLinks(links []linkParser.Link) {
  for _, link := range links {
    if !parsedLinks[link.Href] {
      scrapeHref(link.Href)
    }
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

func chooseDomainLinks(links []linkParser.Link) []linkParser.Link {
  var filteredLinks []linkParser.Link
  for _, link := range links {
    link.Href = strings.Split(link.Href, "#")[0] // links with different parts after "#" actually point to the same page
    if link.Href == "" { link.Href = "/" }

    if link.Href[0:1] == "/" && !parsedLinks[link.Href] {
      filteredLinks = append(filteredLinks, link)
    } else {
      urlObj, err := url.Parse(link.Href)
      if err != nil {
        fmt.Println("error parsing url")
        continue
      }

      linkHostname := urlObj.Hostname()
      if len(linkHostname) > 2 && linkHostname[0:3] == "www" { linkHostname = linkHostname[3:] }
      if linkHostname == hostname && !parsedLinks[linkHostname] {
        filteredLinks = append(filteredLinks, link)
      }
    }
  }

  return filteredLinks
}

