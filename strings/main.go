package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hi")
	var data string = "apple,banana,lemon"
	parts := strings.Split(data, ",")
	fmt.Println(parts[0])

	data = "one two three one one one"

	count := strings.Count(data, "one")
	fmt.Println(count)

	newstr := "          jscdndsc      fdsf      "
	trimmer := strings.Trim(newstr, " ")
	fmt.Println(trimmer)

}
