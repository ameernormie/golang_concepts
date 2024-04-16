package main

import "fmt"

func main()  {
	fmt.Println("************Call by value by default***************")
	a, b := 100, 200

	fmt.Println("before swap")
	fmt.Printf("a is %d and b is %d \n", a, b)
	swap(&a, &b)
	fmt.Println("after swap")
	fmt.Printf("a is %d and b is %d \n", a, b)

	fmt.Println("\n************ Anonymous function ***************")
	func () {
		fmt.Println("Welcome from anonymous")
	}() 

	fmt.Printf("Rectangle area of 10 is = %d\n", func(w int) int {
		return w * w
	}(10))

}

func swap(one *int, two *int)  {
	var temp int

	temp = *one

	*one = *two
	*two = temp
}
