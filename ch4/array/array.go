package main

import "fmt"

func main() {
	// declaration: method 1
	var arr1 [5]int // initialize to zero value
	fmt.Printf("declaration 1: %v\n", arr1)

	// declaration: method 2
	arr2 := [5]int{1, 2, 3, 4, 5} // initialize using literal
	fmt.Printf("declaration 2: %v\n", arr2)

	// declaration: method 2-1
	arr2_1 := [...]int{1, 2, 3, 4, 5} // using ... as array length
	fmt.Printf("declaration 3: %v\n", arr2_1)
}
