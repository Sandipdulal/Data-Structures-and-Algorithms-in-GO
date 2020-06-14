package main

import "fmt"

func main() {
	data := []int{8, 10, 9, 11, 2, 5, 3, 4, 1, 6}
	size := len(data)
	temp := make([]int, size)
	MergeSort(data, temp, 0, size-1)
	fmt.Println(data)
}

/*
*
* Worst case performance = O(nlogn)
* Average case performance = O(nlogn)
* Best case time complexity = O(nlogn)
* Space complexity = O( n)
* Stable sorting = Yes
*
 */
func MergeSort(arr, temp []int, lowerIndex, upperIndex int) {
	if lowerIndex >= upperIndex {
		return
	}
	midIndex := (lowerIndex + upperIndex) / 2
	MergeSort(arr, temp, lowerIndex, midIndex)
	MergeSort(arr, temp, midIndex+1, upperIndex)
	merge(arr, temp, lowerIndex, midIndex, upperIndex)

}

func merge(arr []int, temp []int, lowerIndex int, midIndex int, upperIndex int) {
	lowerStart := lowerIndex
	lowerEnd := midIndex
	upperStart := midIndex + 1
	upperEnd := upperIndex
	tempIndex := lowerIndex

	for lowerStart <= lowerEnd && upperStart <= upperEnd {
		if arr[lowerStart] < arr[upperStart] {
			temp[tempIndex] = arr[lowerStart]
			lowerStart++
		} else {
			temp[tempIndex] = arr[upperStart]
			upperStart++
		}
		tempIndex++
	}
	for lowerStart <= lowerEnd {
		temp[tempIndex] = arr[lowerStart]
		lowerStart++
		tempIndex++
	}
	for upperStart <= upperEnd {
		temp[tempIndex] = arr[upperStart]
		upperStart++
		tempIndex++
	}
	for i := lowerIndex; i <= upperIndex; i++ {
		arr[i] = temp[i]
	}
}
