package main

import "fmt"

func main() {
	data := []int{8, 10, 9, 11, 2, 5, 3, 4, 1, 6}
	InsertionSort(data, more)
	fmt.Println(data)
}

/*
*
* Worst case performance = O(n^2)
* Average case performance = O(n^2)
* Best case time complexity = O(n)
* Space complexity = O(1) as we need only one temp variable
* Stable sorting = Yes
*
*/
func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var i, j, temp int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func more(val1, val2 int) bool {
	return val1 > val2
}
