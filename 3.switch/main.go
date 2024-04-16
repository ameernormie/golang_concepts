package main

import "fmt"

func main()  {
	today := 7

	switch  {
		case today == 1:
			fmt.Println("Today is 1")
		case today == 2:
			fmt.Println("Today is 2")
		case today == 7:
			fmt.Println("Today is 7")
		default:
			fmt.Println("Today is invalid")
	}

	var val interface{} = "Go programming language"

	switch a := val.(type) {
	case int64:
		fmt.Println("Type is integer: ", a)
	case float64:
		fmt.Println("Type is float: ", a)
	case string:
		fmt.Println("Type is string: ", a)
	default: 
		fmt.Println("Type is unknown: ", a)
	}
	
}