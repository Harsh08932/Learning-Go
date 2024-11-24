package main

import "fmt"

func main() {
	num := 5
	fmt.Println("Hi The program runs normally as the number is ", num)
}

// func main() {
// 	// nums := [8]int{}
// 	// fmt.Printf("%d\n", nums)

// 	// how to use slice in GO
// 	// num := []int{}
// 	// num = append(num, 2)
// 	// num = append(num, 2, 1, 2, 3, 4)
// 	// num = append(num, 2)
// 	// fmt.Println("The slice will be", num)
// 	// fmt.Println("Size is ", len(num))
// 	// fmt.Println("capacity is ", cap(num))

// 	//make fuction in slice

// 	// num := make([]int, 5, 10)
// 	// num = append(num, 6, 7, 8, 9, 10, 11)
// 	// fmt.Println("The slice will be", num)
// 	// fmt.Println("Size is ", len(num))
// 	// fmt.Println("capacity is ", cap(num))

// }

//{

// num := 4

// cse := 3

// switch {
// case cse < 8:
// 	fmt.Println("case1")
// case cse > 4:
// 	fmt.Println("case2")
// case cse == 9:
// 	fmt.Println("case3")
// }

// if num < 3 {
// 	fmt.Println("yee we hit the if block")
// } else if num == 4 {
// 	fmt.Println("yee we hit the 2nd if block")
// } else {
// 	fmt.Println("yee we hit the else block")
// }

//}
// func main() {
// 	fmt.Println("Lets start arrays")
// 	// var myName [5]int
// 	// myName[0] = 1
// 	// myName[1] = 2
// 	// myName[2] = 3
// 	// myName[3] = 4
// 	// myName[4] = 5

// 	var myName = [5]int{1, 2, 3, 4, 5}
// 	var name [8]int
// 	var key = [8]string{"heheehe"}
// 	fmt.Println("the array is ", name)
// 	fmt.Printf("Print with %q\n", key)
// 	//fmt.Println("the array is ", key)
// 	fmt.Println("the length of array is", len(myName))

// }

// error handling
// func main() {

// 	fmt.Println("Lets start with Error handling")
// 	// fmt.Println("The value returned by fn call is ", divide(10, 2))

// 	var result = divide(10, 0)
// 	fmt.Println("The value returned by fn call is", result)
// }

// func divide(a, b float64) (float64, error) {
// 	return a / b, fmt.Errorf("The denominator is 0")
// }

// func divide(a, b float64) float64 {
// 	if b == 0 {
// 		return 0
// 	}

// 	return a / b
// }
