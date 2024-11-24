package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name string = "4"
	fmt.Printf("The data type of num is %T\n", name)
	var newint, _ = strconv.Atoi(name)
	fmt.Printf("The data type of num is %T\n", newint)
	fmt.Printf("The new value of num is %d\n", newint)

	// 	//fmt.Println("hey for conversion")

	// 	// var numm float64 = float64(num)
	// 	// fmt.Printf("The new data typefor num is %T and its value is %f", numm, numm)
	// 	var newnum string = strconv.Itoa(num)
	// 	fmt.Printf("The new data type for num is %T and its value is %s\n", newnum, newnum)

	// 	var newFloat, _ = strconv.ParseFloat(newnum, 64)
	// 	fmt.Printf("The new data type for num is %T and its value is %.2f", newFloat, newFloat)
}

// func main() {
// 	fmt.Println("This is for pointers")
// 	var point *int
// 	num := 2
// 	point = &num

// 	fmt.Println(point, num)
// 	fmt.Println(*point, num)

// 	num := "name"
// 	point := &num
// 	fmt.Println(point, num)
// 	fmt.Println(*point, num)
// 	num1 := 3
// 	num2 := 4
// 	fmt.Println(sum(&num1, &num2))
// 	fmt.Println("The new value of num1 is ", num1)
// }

// func sum(a, b *int) int {
// 	var sum int
// 	sum = *a + *b
// 	*a = *a + 20
// 	return sum
// }

// func main() {
// 	fmt.Println("This is for structure and data conversion")

// 	type data struct {
// 		what  string
// 		is    int
// 		thiss float64
// 	}

// 	type address struct {
// 		number int
// 		street string
// 	}
// 	type phone struct {
// 		phoneno int
// 	}

// 	type employee struct {
// 		x address
// 		y phone
// 	}

// 	emp1 := employee{
// 		x: address{number: 1, street: "new"},
// 		y: phone{phoneno: 34342},
// 	}

// 	fmt.Println(emp1.y.phoneno)
// 	var dat data
// 	fmt.Println("the structure is", dat)

// 	datt := data{
// 		what:  "Harsh",
// 		is:    2,
// 		thiss: 3.4,
// 	}
// 	fmt.Println("the structure is", datt)
// 	datt.what = "jfldns"
// 	fmt.Println("the structure is", datt)

// }

//for for loop
// func main() {
// 	fmt.Println("hey")

// 	var num = []int{1, 2, 3, 4, 5}
// 	for index, value := range num {
// 		fmt.Printf("The index is %d and value is %d\n", index, value)
// 	}
// }

//func main() {
// 	//fmt.Println("Hi")

// 	// myMap := make(map[string]int)

// 	// myMap["one"] = 3
// 	// myMap["two"] = 4
// 	// myMap["three"] = 5
// 	// myMap["four"] = 6

// 	// fmt.Println("the value of one is", myMap["one"])
// 	// fmt.Println("the value of one is", myMap["two"])
// 	// fmt.Println("the value of one is", myMap["three"])
// 	// fmt.Println("the value of one is", myMap["four"])
// 	//fmt.Println("the value of one is", myMap["jdne"])

// 	// value, exist := myMap["on"]
// 	// fmt.Println("the value for one is", value)
// 	// fmt.Println("does the value for one exists", exist)

// 	// fmt.Println("the value of four is", myMap["four"])
// 	// fmt.Println("after deletion")
// 	// delete(myMap, "four")
// 	// fmt.Println("the value of four is", myMap["four"])

// 	//another way of initialising maps

// 	// mpp := map[string]int{
// 	// 	"first":  1,
// 	// 	"second": 2,
// 	// 	"third":  3,
// 	// }
// 	// fmt.Println("the value of first is", mpp["first"])
// }
