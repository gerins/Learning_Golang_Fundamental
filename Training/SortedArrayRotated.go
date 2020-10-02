package main

import "log"

func main() {
	var myArray = []int{1, 2, 3, 4, 5, 8, 9, 12, 15, 16}
	result := binarySearch(myArray, 3, 0, len(myArray)-1, -1)
	log.Println(`found at index `, result)
}

func binarySearch(array []int, key, min, max, result int) int {
	if min > max {
		return result
	}

	mid := (min + max) / 2
	if array[mid] == key {
		return mid
	} else if key < mid {
		return binarySearch(array, key, min, mid-1, result)
	} else {
		return binarySearch(array, key, mid+1, max, result)
	}
}
