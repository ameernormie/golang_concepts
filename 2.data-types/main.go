package main

import "fmt"

func main()  {
	// Basic Type
	// 1. Numbers
	// 		a. Integers
	// 				uint (platform dependent), uint8(8-bit integers), uint16, uint32, uint64, byte (alias for uint8) 
	// 						byte (used for ascii as go doesn't have char type, also 0-255) 
	// 				int (platform dependent), int8(8-bit integers), int16, int32, int64, rune (alias for uint32) 
	// 						rune (used represent unicode characters) 
	var a uint8 = 200
	fmt.Println(a)
	var b int16 = -15432
	fmt.Println(b)

	// 		b. Floating numbers
	//				float32 (32 bit or 4 bytes)
	//				float64 (64 bit or 8 bytes)
	fVal1 := 35.78
	fVal2 := 46.77
	fCal := fVal2 - fVal1
	fmt.Printf("fVal1 = %f\n", fVal1)
	fmt.Printf("fVal2 = %f\n", fVal2)
	fmt.Printf("%f - %f = %f\n", fVal2, fVal1, fCal)
	fmt.Printf("type of fCal: %T\n",  fCal)

	// 		c. Complex numbers (float32 and float64 is also part of complex numbers)
	//				complex64 (contain float32 as a real or imaginary component)
	//				complex128 (default) (contain float64 as a real or imaginary component)
	var f1 = 3.25
	var f2 = 34.25

	var complexNum = complex(f1, f2)
	fmt.Println(complexNum)
	fmt.Printf("Type of complexNum: %T\n", complexNum)

	var c1 = 3+5i
	var c2 = 2+3i
	var complexSum = c1 + c2;
	var complexSub = c2 - c1;
	var complexMul = c2 * c1;
	var complexDiv = c1/c2;
	fmt.Println(complexSum, complexSub, complexMul, complexDiv)

	// 2. Booleans
	s1 := "Go programming"
	s2 := "Go Programming"
	s3 := "Go programming"
	strNotEqual := s1 == s2;
	strEqual := s1 == s3;

	fmt.Println("String not equal: ", strNotEqual)
	fmt.Println("String equal: ", strEqual)

	// 3. Strings
	
	// Aggregate Type
	// 1. Arrays
	// 2. Structs

	// Reference Type
	// 1. Pointers
	// 2. Slices
	// 3. Maps
	// 4. Functions
	// 5. Channels

	// Interface Type
}