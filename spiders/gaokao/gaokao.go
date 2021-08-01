package gaokao

import (
	"fmt"

	"golang.org/x/net/html"
)

type Universities struct {
}

func (u *Universities) Crawl() {
	fmt.Println("crawl")
}

func (u *Universities) GetHtml(url string) {

}

func (u *Universities) ParseDoc(doc *html.Node) {

}
