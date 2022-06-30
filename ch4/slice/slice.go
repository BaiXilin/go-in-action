package main

import "fmt"

func main() {
	slice1 := make([]string, 5)
	slice1[3] = "apple"

	fmt.Printf("value:%v  len:%d  cap:%d\n", slice1, len(slice1), cap(slice1))
}
