package main

import "fmt"

func main()  {
	var slice1 = make([]int, 2, 5)
	slice1[0] = 1
	slice1[1] = 2

	fmt.Println("Slice A: ", slice1)
	fmt.Printf("Slice A = %v\tLength = %d\tCapacity = %d\n", slice1, len(slice1), cap(slice1))

	slice1 = append(slice1, 3,4,5,6)
	fmt.Printf("Slice A after append = %v\tLength = %d\tCapacity = %d\n", slice1, len(slice1), cap(slice1))

	a := []string{"Java", "Go", "Perl", "Scala", "Rust"}
	b := []string{"Javascript", "Test", "GOLANG"}
	b = append(b, a...)

	fmt.Println("Slice a: ", a)
	fmt.Println("Slice b: ", b)
}
