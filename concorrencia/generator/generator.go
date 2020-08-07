package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulos(urls ...string) (ch chan string) {
	ch = make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return
}

func main() {
	tit1 := titulos("http://www.google.com.br", "http://www.kabum.com.br")
	tit2 := titulos("http://www.amazon.com.br", "http://www.udemy.com")

	fmt.Printf("Primeiros: %s e %s\n", <-tit1, <-tit2)
	fmt.Printf("Segundos: %s e %s", <-tit1, <-tit2)

}
