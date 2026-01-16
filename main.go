package main

import "fmt"

func main() {

	// Method 1: Declare and initialize
	ages := map[string]int{
		"Ayodeji": 25,
		"Peter":   30,
	}

	// Method 2: Using make
	scores := make(map[string]int)
	scores["Ayodeji"] = 95
	scores["Peter"] = 87

	fmt.Println(ages)   // map[Ayodeji:25 Peter:30]
	fmt.Println(scores) // map[Ayodeji:95 Peter:87]

	// Insert/Update - O(1)
	scores["John"] = 23
	scores["Esther"] = 47
	scores["Clement"] = 70 // Updates existing key

	// Lookup - O(1)
	value := scores["Clement"]
	fmt.Println(value) // 70

	// Check if key exists - IMPORTANT PATTERN
	value, exists := scores["John"]
	if exists {
		fmt.Println("Found:", value)
	} else {
		fmt.Println("Not found")
	}

	// Delete - O(1)
	delete(scores, "John")

	// Length
	fmt.Println(len(scores)) // 1

	// Iterate
	for key, value := range scores {
		fmt.Printf("%s: %d\n", key, value)
	}
}
