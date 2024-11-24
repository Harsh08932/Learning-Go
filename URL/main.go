package main

import (
	"fmt"
	"net/url"
)

func main() {
	//fmt.Println("New URL")

	myURL := "https://jsonplaceholder.typicode.com/"

	ParsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("URL was not parsed correctly")
		return
	}
	fmt.Println("scheme:", ParsedURL.Scheme)
	fmt.Println("host:", ParsedURL.Host)
	fmt.Println("path:", ParsedURL.Path)
	fmt.Println("RawQuery", ParsedURL.RawQuery)

	ParsedURL.Host = "fwfsefes"
	//changing the given URL/modifying the URL
	newURL := ParsedURL.String()
	fmt.Println("newURL", newURL)

}
