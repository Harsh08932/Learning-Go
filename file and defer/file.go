package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//open a file
	file, err := os.Open("New file.txt")

	if err != nil {
		fmt.Println("File not read sucessfully")
	}
	//close a file after main fn executed
	defer file.Close()
	//create a buffer to read file content
	buffer := make([]byte, 100)

	//Read the file content using the loop

	for {
		n, err := file.Read(buffer)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading file", n)
			return
		}
		fmt.Println(string(buffer[:n]))

	}

}

// func main() {
// 	file, err := os.Create("New file.txt")
// 	if err != nil {
// 		fmt.Println("File not crated succesfully")
// 		return
// 	}
// 	defer file.Close()

// 	content := "Hello World"
// 	_, errors := io.WriteString(file, content+"\n")
// 	if errors != nil {
// 		fmt.Println("Data not entered")
// 	}

// 	fmt.Println("File created succesfully")
// }
