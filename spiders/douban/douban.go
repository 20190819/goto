package douban

import (
	"fmt"
	"goto/helpers"
	"net/http"
	"sync"
	"time"

	"github.com/spf13/viper"

	"golang.org/x/net/html"

	"github.com/antchfx/htmlquery"
	"github.com/sirupsen/logrus"
)

var parseRes []map[string]string
var wg sync.WaitGroup



func Crawl() {
	defer func() {
		if recover() != nil {
			logrus.Fatalln("系统错误：", recover() == nil)
			return
		}
	}()

	for page := 1; page <= 3; page++ {
		urlStr := viper.GetString("DOUBAN_BOOK_URL")
		urlStr = fmt.Sprintf(urlStr, 20*(page-1))
		wg.Add(1)
		go getHtml(urlStr)
	}

	wg.Wait()
}

func getHtml(url string) {
	defer func() {
		wg.Done()
	}()
	doc := new(html.Node)
	logrus.Info("handle url is :", url)
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		logrus.Fatalln("http get err", err)
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		doc, err = htmlquery.Parse(response.Body)
		if err != nil {
			logrus.Fatalln("html query parse err:", err)
		}
		time.Sleep(1)
		parseDoc(doc)
	} else {
		logrus.Warn("response.StatusCode: ", response.StatusCode)
	}

}

func parseDoc(doc *html.Node) {
	expr := "//div[@class='info']"
	nodes, err := htmlquery.QueryAll(doc, expr)

	if err != nil {
		logrus.Error(err)
	}
	exprTitle := "./h2/a/text()"
	exprStar := "div[contains(@class,'star')]/span[2]/text()"
	exprAuthor := "/div[@class='pub']/text()"
	for _, node := range nodes {
		titleNode := htmlquery.FindOne(node, exprTitle)
		title := htmlquery.InnerText(titleNode)
		starNode := htmlquery.FindOne(node, exprStar)
		star := htmlquery.InnerText(starNode)
		authorNode := htmlquery.FindOne(node, exprAuthor)
		author := htmlquery.InnerText(authorNode)
		parseRes = append(parseRes, map[string]string{
			"title":  helpers.StrTrimSpace(title),
			"star":   helpers.StrTrimSpace(star),
			"author": helpers.StrTrimSpace(author),
		})
	}
	logrus.Infof("parseRes len =%d", len(parseRes))
}
