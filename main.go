package main

func main() {

}

func bubbleSort(arr []int) {
	n := len(arr)

	// Outer loop: n-1 passes
	for i := 0; i < n-1; i++ {
		swapped := false

		// Inner loop: compare adjacent elements
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// Optimization: if no swaps, array is sorted
		if !swapped {
			break
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)

	// Move boundary of unsorted subarray
	for i := 0; i < n-1; i++ {
		// Find minimum in unsorted portion
		minIndex := i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Swap minimum with first unsorted element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func insertionSort(arr []int) {
	n := len(arr)

	// Start from second element
	for i := 1; i < n; i++ {
		key := arr[i] // Element to insert
		j := i - 1

		// Shift elements greater than key to the right
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Insert key at correct position
		arr[j+1] = key
	}
}
