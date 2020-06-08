package main

import "fmt"

func main() {
	data := []int{8, 10, 9, 11, 2, 5, 3, 4, 1, 6}
	BubbleSort(data, more)
	fmt.Println(data)

	data = []int{1, 2, 3, 6, 5, 7, 8, 9, 10}
	BubbleSortImproved(data, more)
	fmt.Println(data)
}

/*
*
* Worst case performance = O(n^2)
* Average case performance = O(n^2)
* Space complexity = O(1) as we need only one temp variable
* Stable sorting = Yes
*
 */
func BubbleSort(arr []int, comp func(int, int) bool) {
	size := len(arr)

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if comp(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

}

/*
*
* Worst case performance = O(n^2)
* Average case performance = O(n^2)
* Space complexity = O(1)
* Stable sorting = Yes
* Adaptive when array is nearly sorted = O(n)
*
 */
func BubbleSortImproved(arr []int, more func(int, int) bool) {
	size := len(arr)
	sorted := true
	for i := 0; i < (size-1) && sorted; i++ {
		sorted = false
		for j := 0; j < size-i-1; j++ {
			if more(arr[j], arr[j+1]) {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				sorted = true
			}
		}
	}
}

func more(val1, val2 int) bool {
	return val1 > val2
}
