package main

import (
	"fmt"
)

func main() {
	//INITIALIZE ARRAY WITHOUT VALUES
	var array1 [5]int
	//INITIALIZE ARRAY WITH VALUES
	var array2 = [3]string{"Ayodeji", "Golang", "DSA"}

	fmt.Println(array1)
	fmt.Println(array2)

	//MODIFY ARRAY
	array1[0] = 1
	array1[1] = 2

	array2[2] = "Golang is Great"

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(len(array2))

	/* SLICES */

	//INITIALIZE SLICE WITH VALUES
	var slice1 = []int{10, 15, 567, 42, 42, 345, 543, 23, 111, 20, 30, 40}
	result := containsDuplicate(slice1)
	fmt.Println("The result for finding duplicates is", result)

	fmt.Println(len(slice1), cap(slice1), slice1)

	//INITIALIZE SLICE WITH EMPTY VALUES
	var slice2 = make([]string, 3)
	fmt.Println(slice2)

	//INITIALIZE SLICE WITH EMPTY VALUES AND A FIXED CAPACITY
	var slice3 = make([]int, 3, 7)
	fmt.Println(len(slice3), cap(slice3), slice3)

	//ACCESS BY INDEX
	fmt.Println(slice1[2])

	//MODIFY VALUE
	//slice2[1] = "Software"
	fmt.Println(slice2)

	//CREATING SUB SLICES
	var sliceSplit = slice1[0:2]
	fmt.Println(sliceSplit)

	//APPEND SLICES
	slice1 = append(slice1, 1)
	slice1 = append(slice1, 23, 234, 2345, 23456, 234567)
	fmt.Printf("len=%d, cap=%d, %v\n", len(slice1), cap(slice1), slice1)

}

func containsDuplicate(nums []int) bool {
	set := make(map[int]int)

	for _, number := range nums {
		set[number]++

		if set[number] > 1 {
			return true
		}
	}
	return false
}
