package main

import "fmt"

func main()  {
	// 1. Arrays (Fixed length)
	var names [3]string
	names[0] = "ameer"
	names[1] = "ali"
	names[2] = "moeed"

	fmt.Println("Arrays Elements: ")
	fmt.Println("Element 1: ", names[0])
	fmt.Println("Element 2: ", names[1])
	fmt.Println("Element 3: ", names[2])

	colors := [4]string{"red", "yellow", "blue", "gold"}
	fmt.Println("Colors of Array: ")

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}


	// Initializing using ellipsis
	x := [...] int {10, 20, 30}
	fmt.Println("Lenght of array: ", len(x))

	// Initializing specific elements in array
	y := [5]int{1:10, 3:30}
	fmt.Println("Dynamic array with specific elements: ")
	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}

	// Filter elements
	friendNames := [5]string{"ameer", "adil", "moeed", "asfar", "islam"}
	fmt.Println(friendNames[:])
	fmt.Println(friendNames[:3])


	// Iterating
	intArray := [5]int{1,2,3,4,5}

	fmt.Println("Using for loop")
	// 1. For loop
	for i:=0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
	fmt.Println("Using range")
	// 2. Range
	for index, val := range intArray{
		fmt.Println(index, " => ", val)
	}
	for _, val := range intArray{
		fmt.Println( "value => ", val)
	}

	for range intArray{
		fmt.Println("perform action n number of times: ")
	}


	// Copy array by value or reference into another array
	arr1 := [3]string{"Kim", "min", "jae"}

	copyArr1 := arr1
	copyRefArr1 :=  &arr1;

	fmt.Printf("Array 1: %v\n", arr1)
	fmt.Printf("copy of Array 1: %v\n", copyArr1)
	fmt.Printf("referece copy of Array 1: %v\n", *copyRefArr1)

	arr1[0] = "ameer"

	fmt.Printf("After mutation Array 1: %v\n", arr1)
	fmt.Printf("After mutation copy of Array 1: %v\n", copyArr1)
	fmt.Printf("After mutation referece copy of Array 1: %v\n", *copyRefArr1)


	fmt.Println("********* Multi dimensional array *********")

	mulArr1 := [3][3]string {
		{"Go", "Java", "Rust"},
		{"C", "Scala", "python"},
		{"Javascript", "ruby", "c++"},
	}

	fmt.Println("Looping 2 dimensional array")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s\t",mulArr1[i][j])
		}
		fmt.Printf("\n")
	}

	var mulArr2 [2][2]int

	mulArr2[0][0] = 0
	mulArr2[0][1] = 1
	mulArr2[1][0] = 2
	mulArr2[1][1] = 3

	fmt.Println("Looping 2 dimensional array")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d\t",mulArr2[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Println("Array equality can be done using ==")
	firstArr := [2]string {"ameer", "hamza"}
	secondArr := [2]string {"ameer", "hamza"}
	thirdArr := [2]string {"talha", "ali"}

	fmt.Printf("First Arr: %v\n", firstArr)
	fmt.Printf("Second Arr: %v\n", secondArr)
	fmt.Printf("Third Arr: %v\n", thirdArr)

	fmt.Printf("First and second are equal? %v\n", firstArr == secondArr)
	fmt.Printf("First and third are equal? %v\n", firstArr == thirdArr)
}