package main

import (
	"goto/config"
	"goto/spiders/douban"
)

func main(){
	config.InitConfig("../")
	// 采集豆瓣读书
	douban.Crawl()
}