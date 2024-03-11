package main

import (
	_ "bufio"
	_ "fmt"
	"net/http"
	_ "os"
	_ "strings"
	"sync"

	_ "github.com/yhat/scrape"
	"golang.org/x/net/html"
	_ "golang.org/x/net/html/atom"
)

const (
	urlRoot = "https://www.wanted.co.kr"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

var wg sync.WaitGroup

func main() {
	res, err := http.Get(urlRoot)
	errorCheck(err)

	defer res.Body.Close()

	root, err := html.Parse(res.Body)
	errorCheck(err)

}
