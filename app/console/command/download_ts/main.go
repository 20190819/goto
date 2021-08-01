package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type DownloadTs struct {
	start   int
	end     int
	version int
	url     string
	suffix  string
}

var wg sync.WaitGroup

var ts = DownloadTs{
	start:   0,
	end:     40,
	version: 501407,
	url:     "https://cdn.91p07.com//m3u8/501407/",
	suffix:  ".ts",
}

func (*DownloadTs) Intro() string {
	return "下载 .ts 视频文件"
}

func (*DownloadTs) Query(i int, body io.Reader) {
	defer func() {
		wg.Done()
	}()

	// 创建文件
	file, fInfo := ts.createPath(i)

	defer file.Close()
	// 写文件
	if fInfo.Size() == 0 {
		url := ts.url + fInfo.Name()
		// 请求
		response := ts.httpQuery(url, nil)
		defer response.Body.Close()
		written, err := io.Copy(file, response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(" download done,", "file size(byte)=", written)
	} else {
		fmt.Println("original file size=", fInfo.Size())
	}
}

func (*DownloadTs) createPath(i int) (*os.File, os.FileInfo) {
	var err error
	dir, _ := os.Getwd()
	path := dir + "/" + strconv.Itoa(ts.version)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
	fileStr := fmt.Sprintf("%s/%d%d%s", path, ts.version, i, ts.suffix)
	fInfo, err := os.Stat(fileStr)
	var file *os.File
	if os.IsNotExist(err) {
		file, _ = os.Create(fileStr)
	} else {
		file, _ = os.OpenFile(fileStr, os.O_RDWR, 0777)
	}
	return file, fInfo

}

func (*DownloadTs) httpQuery(url string, body io.Reader) *http.Response {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, body)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	return response
}

func (*DownloadTs) Handle() {
	defer func() {
		if recover() != nil {
			fmt.Println(recover())
		}
	}()

	for i := ts.start; i <= ts.end; i++ {
		wg.Add(1)
		go ts.Query(i, nil)
	}
}

func main() {
	ts.Handle()
	wg.Wait()
	fmt.Println("get .ts ok!")
}
