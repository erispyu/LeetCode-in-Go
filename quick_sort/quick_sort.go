package main

import "fmt"

func partition(list []int, low, high int) int {
	pivot := low
	for low < high {
		for low < high && list[high] >= list[pivot] {
			high--
		}
		for low < high && list[low] <= list[pivot] {
			low++
		}
		list[low], list[high] = list[high], list[low]
	}
	list[pivot], list[low] = list[low], list[pivot]
	return low
}

func QuickSort(list []int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(list, low, high)
	QuickSort(list, low, pivot-1)
	QuickSort(list, pivot+1, high)

}

func main() {
	list := []int{3, 2, 1, 4, 5}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)

	list = []int{6, 5, 4, 3, 2, 1}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
