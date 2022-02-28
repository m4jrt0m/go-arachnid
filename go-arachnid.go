package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "The url to search for")
	flag.Parse()

	fmt.Println(url)

	page := get_page(sanitize_url(url))

	get_urls(page)

	fmt.Println(page)
	fmt.Printf("%s\n", page)
}

func get_urls(html string) [1]string {
	index := strings.Index(html, "http")
	fmt.Printf("%c\n", index)
	array := [1]string{""}
	return array
}

func sanitize_url(url string) string {
	sanitized := url

	if !strings.Contains(url, "http") {
		sanitized = "http://" + url
	}

	return sanitized
}

func get_page(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(html)
}
