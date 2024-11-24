package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	fmt.Println("Practice JSON in GO")
	person := Person{Name: "Harsh", Age: 24, IsAdult: true}
	fmt.Println("The struct is", person)

	//use marshalling technique JSONENCODING
	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println("Error in Marshalling", err)
		return
	}
	fmt.Println("JSON object is", string(jsonData))

	//decoding/unmarshalling
	var persondata Person
	error := json.Unmarshal(jsonData, &persondata)
	if error != nil {
		fmt.Println("Error unmarshalling")
	}
	fmt.Println("The new GO struct is:", persondata)
}
