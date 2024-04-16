package main

import "fmt"

func main()  {
	// & denotes an address in memory
	// & dereferencing operator (access value stored in the address)

	var a int = 20

	// Pointer variable declaration
	var p *int
	// Store address of a variable in pointer variable
	p = &a

	fmt.Printf("Address of A variable: %x\n", &a)
	fmt.Printf("Address stored in pointer variable: %x\n", p)
	fmt.Printf("Access value using dereferncing: %d\n ", *p)

	fmt.Println("*************** Nil pointer *****************")

	var po *int
	if po == nil {
		fmt.Printf("The value of po is: %x\n", po)
	}
	fmt.Println("*************** Change pointer value *****************")
	var val = 458
	fmt.Printf("Value : %d\n", val)
	var pointerVal = &val

	*pointerVal = 400
	fmt.Printf("Value updated via reference: %d\n", val)
	
	fmt.Println("*************** Pointer to pointer (Double pointer) chain pointer *****************")

	var x int = 100
	var pt1 *int = &x
	var pt2 **int = &pt1

	fmt.Println("The value of x: ", x)
	fmt.Println("The value of pt1: ", pt1)
	fmt.Println("The value of pt2: ", pt2)
	fmt.Println("The value at address of pt2 or *pt2: ", *pt2)
	fmt.Println("The value at address of pt2 or **pt2: ", **pt2)
}