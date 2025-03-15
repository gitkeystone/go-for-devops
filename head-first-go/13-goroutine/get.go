package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, ch chan<- Page) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	ch <- Page{URL: url, Size: len(bytes)}
}

func main() {
	start := time.Now()

	urls := []string{
		"https://chat.deepseek.com/",
		"https://go.dev/",
		"https://go.dev/doc",
	}

	ch := make(chan Page)
	for _, url := range urls {
		go responseSize(url, ch)
	}

	for range urls {
		p := <-ch
		fmt.Printf("%s: %d\n", p.URL, p.Size)
	}

	fmt.Println(time.Since(start))
}
