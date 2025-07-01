package main

import "fmt"

func printArray(arr *[]int) {
	for index, value := range *arr {
		fmt.Println(index, value)
	}
}

func main() {

	// a := 2

	// ptr := &a

	// fmt.Println("value of : ", *ptr)

	// array := []int{1, 2, 3, 4, 5}
	slice := make([]int, 0, 10)
	slice = append(slice, 1)
	slice = append(slice, 2)

	printArray(&slice)
}
