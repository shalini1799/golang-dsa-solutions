package main

import ()


func SelectionSort(arr []int) {
	for i:=0; i< len(arr); i++ {
		min := i
		for j:= i+1; j<len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}