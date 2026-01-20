package main

func main() {

}
func mergeSort(arr []int) []int {
	// Base case: array of 0 or 1 element is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Divide: split array in half
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // Recursively sort left half
	right := mergeSort(arr[mid:]) // Recursively sort right half

	// Conquer: merge the sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements from left and right, add smaller to result
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements (one array will be exhausted)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// QUICK SORT

func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get pivot index
		pivotIndex := partition(arr, low, high)

		// Recursively sort elements before and after partition
		quickSort(arr, low, pivotIndex-1)
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// Choose rightmost element as pivot
	pivot := arr[high]

	// Index of smaller element
	i := low - 1

	for j := low; j < high; j++ {
		// If current element is smaller than pivot
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap
		}
	}

	// Place pivot in correct position
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
