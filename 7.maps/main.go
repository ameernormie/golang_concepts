package main

import (
	"fmt"
	"sort"
)

func main()  {
	// Key value pairs
	// Implemented hash table
	// most data types can be used for keys
	var emptyMap = map[int]string{}

	fmt.Println("Empty map: ", emptyMap)
	fmt.Println("Empty map length: ", len(emptyMap))

	var mymap = map[int]string{
		1: "ameer",
		2: "hamza",
		3: "talha",
		4: "wow",
	}
	fmt.Println("mymap: ", mymap)

	fmt.Println("**********Intializing using make**********")
	var makemap = make(map[string]float64)
	makemap["ameer"] = 80.55
	makemap["hamza"] = 845.55
	fmt.Println("map using make: ", makemap)
	fmt.Println("length of map: ", len(makemap))

	
	fmt.Println("********** Adding Deleting Updating map items  **********")
	var studentGrades = map[string]string{
		"ameer": "A+",
		"hamza": "F",
		"moeed": "D+",
		"ali": "C+",
		"talha": "B+",
	}
	fmt.Println("Student Grades: ", studentGrades)

	studentGrades["ameer"] = "B"
	fmt.Println("Student Grades after update: ", studentGrades)

	delete(studentGrades, "ali")
	fmt.Println("Student Grades after delete: ", studentGrades)

	fmt.Println("********** Iterating over map items  **********")
	for name, grade := range studentGrades {
		fmt.Printf("Studend %s has gotten %s\n", name, grade)
	}

	fmt.Println("********** sorting map  **********")

	studentNames := make([]string, 0, len(studentGrades))
	for i := range studentGrades {
		studentNames = append(studentNames, i)
	}

	sort.Strings(studentNames)

	for _, name := range studentNames {
		fmt.Println(name, studentGrades[name])
	}

	fmt.Println("********** truncate map  **********")
	fmt.Println("Map before truncate(delete): ", mymap)

	// for i := range studentGrades {
	// 	delete(studentGrades, i)
	// }
	mymap = make(map[int]string)
	fmt.Println("Map after truncate(delete): ", mymap)

	fmt.Println("********** Merge map  **********")
	newStudents := map[string]string{
		"Adil": "A",
		"Salman": "B-",
	}

	for name, grade := range newStudents {
		studentGrades[name] = grade
	}

	fmt.Println("New studens after merge ", studentGrades)

}