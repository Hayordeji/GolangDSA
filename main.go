package main

import "fmt"

func main() {

	data := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23

	result := BinarySearch(data, target)
	if result != -1 {
		fmt.Printf("Element found at index: %d\n", result)
	} else {
		fmt.Println("Element not found in slice")
	}
}

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		// Calculate the middle index.
		// Using low + (high-low)/2 prevents potential overflow in other languages,
		// though Go's int type is usually 64-bit.
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid // Target found
		}

		if nums[mid] < target {
			low = mid + 1 // Search the right half
		} else {
			high = mid - 1 // Search the left half
		}
	}

	return -1 // Target not found
}
