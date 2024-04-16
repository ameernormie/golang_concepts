package main

import (
	"fmt"
	"sort"
)

func main()  {
	slice1 := []string{"Java", "Go", "Perl", "Scala", "Rust"}
	slice2 := []int{1,45,12,45,1,4,56,25,6,7}
	fmt.Println("Before sorting:")
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)

	sort.Strings(slice1)
	sort.Ints(slice2)

	fmt.Println("After sorting:")
	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
}
