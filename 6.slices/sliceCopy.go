package main

import (
	"fmt"
)

func main()  {
	slice1 := []string{"Java", "Go", "Perl", "Scala", "Rust"}
	fmt.Printf("Slice 1 = %v\tLength = %d\tCapacity = %d\n", slice1, len(slice1), cap(slice1))

	slice2 := make([]string, 5, 10)
	copy(slice2, slice1)
	slice2[0] = "update"
	fmt.Printf("Slice 2 = %v\tLength = %d\tCapacity = %d\n", slice2, len(slice2), cap(slice2))
}
