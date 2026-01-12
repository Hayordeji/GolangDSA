package main

func main() {

}
func maxSum(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}

	// Build initial window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// Slide the window
	for right := k; right < len(nums); right++ {
		windowSum += nums[right]   // Add new element
		windowSum -= nums[right-k] // Remove leftmost element

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[rune]int) // character → last seen index
	maxLength := 0
	windowStart := 0

	for windowEnd, char := range s {
		// If character seen before and inside current window
		if lastIndex, found := charIndex[char]; found && lastIndex >= windowStart {
			// Shrink window: move start past the duplicate
			windowStart = lastIndex + 1
		}

		// Update last seen index
		charIndex[char] = windowEnd

		// Update max length
		currentLength := windowEnd - windowStart + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
