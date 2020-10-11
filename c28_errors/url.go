package main

import (
	"fmt"
	"net/url"
)

func handleError(err URL.Error) {

}

func main() {
	const invalidUrl string = "https://a b.com/"

	parsedUrl, err := url.Parse(invalidUrl)
	if err != nil {
		handleError(err)
	}
}