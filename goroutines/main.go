package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Execution Completed")
}

func sayHi() {
	fmt.Println("Hi Harsh")
	time.Sleep(2000 * time.Millisecond)

}

func main() {
	fmt.Println("learning GoRoutines")
	go sayHello()
	go sayHi()

	time.Sleep(1000 * time.Millisecond)
}
