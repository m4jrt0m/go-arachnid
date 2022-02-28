package main

import (
	"flag"
	"fmt"

	curl "github.com/andelf/go-curl"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "The url to search for")
	flag.Parse()

	fmt.Println(url)

	page := get_page(url)
	fmt.Println(page)
}

func get_page(url string) string {
	easy := curl.EasyInit()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, url)

	err := easy.Perform()
	if err != nil {
		//fmt.Println("ERROR: ${err}")
	}

	return ""
}
