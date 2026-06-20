package main

import (
	"fmt"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	// Test find insert position
	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	l := len(arr)
	i := l / 2

	ll := 0     // последний индекс слева
	lr := l - 1 // последний индекс справа

	for {
		el := arr[i]
		if target == el {
			return i
		}

		if ll == lr {
			return -1
		}

		if target < el {
			if i == ll {
				return -1 // только если ll и lr это индексы
			}
			lr = i - 1
		} else {
			if i == lr {
				return -1
			}
			ll = i + 1
		}

		currLen := lr - ll // является разницей между индексами. по сути, на один меньше
		i = currLen/2 + ll
	}
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	wl := len(arr) // whole arr len
	if wl == 0 {
		return -1
	}

	l := right - left

	i := l/2 + left

	el := arr[i]
	if target == el {
		return i
	}

	if left == right {
		return -1
	}

	if target < el {
		if i == left {
			return -1 // только если ll и lr это индексы
		}
		right = i - 1
	} else {
		if i == right {
			return -1
		}
		left = i + 1
	}

	return BinarySearchRecursive(arr, target, left, right)
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
	l := len(arr)
	if len(arr) == 0 {
		return 0
	}

	i1 := l / 2

	left := 0
	right := l - 1

	for {

		i2 := i1 + 1
		el1 := arr[i1]
		el2 := arr[i2]
		if target == el1 {
			return i1
		}
		if target == el2 {
			return i2
		}

		if el1 < target {
			if target < el2 {
				return i2
			} else {
				if i2 == l-1 {
					return l
				}

				left = i2
			}
		} else {
			if i1 == 0 {
				return 0
			}
			right = i1
		}

		i1 = (right-left)/2 + left
	}
}
