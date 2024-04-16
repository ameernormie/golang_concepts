package main

import "fmt"

var globalVar int = 1000

func main()  {
	fmt.Println("test")
	var variable1  = 100;
	var variable2  = "go programming language";
	var variable3  = 2.45;

	fmt.Printf("The  value of variable 1 is: %d\n", variable1)
	fmt.Printf("The  type of variable 1 is: %T\n", variable1)
	fmt.Printf("The  value of variable 2 is: %s\n", variable2)
	fmt.Printf("The  type of variable 2 is: %T\n", variable2)
	fmt.Printf("The  value of variable 3 is: %f\n", variable3)
	fmt.Printf("The  type of variable 3 is: %T\n", variable3)

	var variable4 int
	var variable5 string
	var variable6 float64

	fmt.Printf("The  value of unintialized variable 4 is: %d\n", variable4)
	fmt.Printf("The  value of unintialized variable 5 is: %s\n", variable5)
	fmt.Printf("The  value of unintialized variable 6 is: %f\n", variable6)

	var var7, var8, var9 int=2, 23, 123
	fmt.Printf("The  value of unintialized var7 4 is: %d\n", var7)
	fmt.Printf("The  value of unintialized var8 5 is: %d\n", var8)
	fmt.Printf("The  value of unintialized var9 6 is: %d\n", var9)


	var (
		firstName = "ameer"
		num1 = 300
		num2 = 12.45
		married = true
	)

	fmt.Println(firstName)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(married)

	var10 := 1
	var11 := "go programming"
	var12 := 1.23
	fmt.Printf("The  value of var10 is: %d\n", var10)
	fmt.Printf("The  value of var11 is: %s\n", var11)
	fmt.Printf("The  value of var12 is: %f\n", var12)

	var13, var14, var15 := 1, "go", 1.12
	fmt.Printf("The  value of var13 is: %d\n", var13)
	fmt.Printf("The  value of var14 is: %s\n", var14)
	fmt.Printf("The  value of var15 is: %f\n", var15)


	display()
}

func display()  {
	fmt.Printf("I can access global variable here: %d\n", globalVar)
}