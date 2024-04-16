package main

import "fmt"

func main()  {
	arr := [5]int{1,2,3,4,5}

	slice1 := arr[1:2]
	slice2 := arr[0:]
	slice3 := arr[:2]
	slice4 := arr[:]

	fmt.Println("Array: ", arr)
	fmt.Println("slice1 arr[1:2]: ", slice1)
	fmt.Println("slice2 arr[0:]: ", slice2)
	fmt.Println("slice3 arr[:2]: ", slice3)
	fmt.Println("slice4 arr[:]: ", slice4)

	// Slice from a slice

	sliceofSlice1 := slice1[:]
	sliceofSlice2 := slice2[2:]
	fmt.Println("slice of slice 1: ", sliceofSlice1)
	fmt.Println("slice of slice 2: ", sliceofSlice2)
}
