package main

import (
	"fmt"
	"net/url"
	"os"
)

func handleError(err error) {
	if e, ok := err.(*url.Error); ok {
		fmt.Printf("Operation: %v, URL: %v, Error: %v\n", e.Op, e.URL, e.Err)
	}
	os.Exit(1)
}

func main() {
	const invalidUrl string = "https://a b.com/"

	_, err := url.Parse(invalidUrl)
	if err != nil {
		handleError(err)
	}
}