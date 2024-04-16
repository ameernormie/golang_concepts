package main

import "fmt"

func main()  {
	// 1. Simple loop
	for i:=0; i<5; i++ {
		fmt.Printf("i = %d \t Go programming language\n", i)
	}

	// 2. Infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// 3. loop as while loop
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Break 
	for i:=0; i<5; i++ {
		fmt.Printf("i = %d \t Go break \n", i)
		if(i == 3) {
			break
		}
	}
	
	// Continue 
	for i:=0; i<5; i++ {
		if(i == 3) {
			continue
		}
		fmt.Printf("i = %d \t Go contineu \n", i)
	}

	// Goto 
	var x int = 0;
	label1:
		for x < 8 {
			if x == 5 {
				x +=1
				goto label1;
			}
			fmt.Printf("value of X: %d\n", x)

			x++
		}

}