package link_parser

import (
  "golang.org/x/net/html"
  "errors"
  "strings"
)

type Link struct {
  Href string
  Text string
}

func Parse(doc *html.Node) []Link {
  var links []Link
  traverseNode(doc, &links)

  return links
}

func traverseNode(node *html.Node, links *[]Link) {
  if node.Type == 3 && node.Data == "a" {
    rememberLink(node, links)
  }

  if node.FirstChild != nil {
    traverseNode(node.FirstChild, links)
  }

  if node.NextSibling != nil {
    traverseNode(node.NextSibling, links)
  }
}

func rememberLink(linkNode *html.Node, links *[]Link) {
  if href, err := findHref(linkNode); err != nil {
    return // do not remember links that do not have href attribute
  } else {
    linkText := findLinkText(linkNode)
    link := Link{ href, linkText[1:] } // linkText[1:] gets rid of the first space created by extractLinkChildrenText method
    *links = append(*links, link)
  }
}

func findHref(linkNode *html.Node) (string, error) {
  for _, attr := range linkNode.Attr {
    if attr.Key == "href" {
      return attr.Val, nil
    }
  }
  return "", errors.New("Could not find href attribute")
}

func findLinkText(linkNode *html.Node) string {
  if linkNode.FirstChild == nil {
    return "" // <a> element has no children, so it's empty
  } else {
    text := ""
    extractLinkChildrenText(linkNode.FirstChild, &text)
    return text
  }
}

func extractLinkChildrenText(childNode *html.Node, totalText *string) {
  if childNode.Type == 1 {
    *totalText = *totalText + " " + strings.Trim(childNode.Data, " \n")
  }

  if childNode.FirstChild != nil {
    extractLinkChildrenText(childNode.FirstChild, totalText)
  }

  if childNode.NextSibling != nil {
    extractLinkChildrenText(childNode.NextSibling, totalText)
  }
}
