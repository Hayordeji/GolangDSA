package main

import (
	"fmt"
	"unicode"
)

func main() {
	slice1 := []int{1, 3, 3, 5, 2, 2, 6, 7, 8, 8}
	removeDuplicates(slice1)

	//PRINT RESULT
	fmt.Println(slice1)
}

// REMOVE DUPLICATES IN A SLICE
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}

func CheckPalindrome(word string) bool {
	//CONVERT TO LOWER CASE AND RUNE
	var chars []rune
	for _, char := range word {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			chars = append(chars, unicode.ToLower(char))
		}
	}

	left := 0
	right := len(chars) - 1

	for left < right {
		if chars[left] != chars[right] {
			return false
		} else {
			left++
			right--
		}
		return true
	}
	return false

}

func TwoSumBruteForce(nums []int, target int) []int {
	n := len(nums)

	// Outer loop: pick first number
	for i := 0; i < n; i++ {
		// Inner loop: pick second number
		for j := i + 1; j < n; j++ {
			// Check if they sum to target
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

// TwoSum : FIND THE SUM OF TWO ELEMENTS IN THE SLICE THAT EQUALS THE TARGET
func TwoSum(numbers []int, target int) []int {
	var left = 0
	var right = len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	//SUM NOT FOUND
	return []int{-1, -1}
}
