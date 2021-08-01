package common

import "golang.org/x/net/html"

type Spider interface {
	Crawl()
	GetHtml(url string)
	ParseDoc(doc *html.Node)
}
