package main

import "fmt"

func main() {
	var arr_source1, arr_target1 [5]string
	arr_source1 = [5]string{"Red", "Blue", "Green", "White", "Purple"}
	// copying arrays are deep copy!
	arr_target1 = arr_source1
	fmt.Printf("%p %p\n", &arr_source1, &arr_target1)

	// copying arrays that contain pointers
	var arr_source2, arr_target2 [3]*string
	arr_source2 = [3]*string{new(string), new(string), new(string)}
	*arr_source2[0] = "One"
	*arr_source2[1] = "Two"
	*arr_source2[2] = "Three"

	arr_target2 = arr_source2
	fmt.Printf("%p %p\n", arr_source2[1], arr_target2[1])
}
