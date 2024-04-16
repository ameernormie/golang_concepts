package main

import "fmt"

func main()  {
	// points to underlying array. internally represented by sliceheader
	// Represented by 3 things
	// 1. Pointers: Pointer to underlying array
	// 2. Length: Current length
	// 3. Capacity: total capacity to which underlying array can expand

	var slice1 = []string{"Go", "Java", "C#","Perl"}
	fmt.Println("Slice 1: ", slice1)
	fmt.Printf("Type of slice1: %T\n", slice1)

	slice2 := []int{1,3,5,7,9}
	fmt.Println("Slice 2: ", slice2)
	fmt.Printf("Type of slice2: %T\n", slice2)

}
