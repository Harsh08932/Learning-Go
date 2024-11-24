package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning Web handling")
	// fetch the data and get response object and error
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("data not fetched correctly")
		return
	}
	//for freeing the resouces
	defer res.Body.Close()

	//read response from body

	n, errr := ioutil.ReadAll(res.Body)

	if errr != nil {
		fmt.Println("data not read")
		return
	}
	fmt.Println("response is:", string(n))

}
