package main

import "fmt"

func main() {
	data := []int{8, 10, 9, 11, 2, 5, 3, 4, 1, 6}
	SelectionSort(data)
	fmt.Println(data)
}

/*
*
* Worst case performance = O(n^2)
* Average case performance = O(n^2)
* Best case time complexity = O(n^2)
* Space complexity = O(1)
* Stable sorting = No
*
 */
func SelectionSort(arr []int) {
	size := len(arr)
	var min int
	for i := 0; i < size-2; i++ {
		min = i
		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

}
