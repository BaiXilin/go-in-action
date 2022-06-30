package main

import "fmt"

func main() {
	// array that stores pointers
	arr_ptr := [3]*int{0: new(int), 1: new(int), 2: new(int)}
	fmt.Printf("array that contains pointer: %v\n", arr_ptr)
	arr_ptr[1] = new(int)

	// modify value pointed by pointer
	*arr_ptr[0] = 10
	fmt.Printf("array that contains pointer after modification: %v\n", arr_ptr)

	// printing out value
	fmt.Printf("Actual value stored in pointer array: ")
	for i := range arr_ptr {
		fmt.Printf("%d ", *arr_ptr[i])
	}
	fmt.Println()
}
