package main

import "fmt"

func main() {
	data := []int{8, 10, 9, 11, 2, 5, 3, 4, 1, 6}
	QuickSort(data, 0, len(data)-1)
	fmt.Println(data)

}

/*
*
* Worst case performance = O(n^2)
* Average case performance = O(nlogn)
* Best case time complexity = O(nlogn)
* Space complexity = O(nlogn)
* Stable sorting = No
*
 */
func QuickSort(data []int, start int, end int) {
	if start >= end {
		return
	}
	partitionIndex := partition(data, start, end)
	QuickSort(data, start, partitionIndex-1)
	QuickSort(data, partitionIndex+1, end)

}

func partition(data []int, start int, end int) int {
	pivot:=data[end]
	partitionIndex:=start
	for i:=start;i<end;i++{
		if data[i]<=pivot{
			data[i],data[partitionIndex]=data[partitionIndex],data[i]
			partitionIndex++
		}
	}
	data[partitionIndex],data[end]=data[end],data[partitionIndex]
	return partitionIndex
}
