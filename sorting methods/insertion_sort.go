package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Generates a random array of the given size
func input(size int) []int {
	arr := make([]int, size)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(999) - rand.Intn(999)
	}
	return arr
}

// Implements the insertion sort algorithm. use i only for iteration , use j for comparision n swap
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j = j - 1
		}
	}
}

func selectionSort(arr []int) {
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

func bubbleSort(arr []int) {
	for i:=0; i< len(arr); i++ {
		for j:=0; j< len(arr); j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	//insertion sort
	arr := input(8)
	fmt.Printf("INSERTION SORT :Unsorted array: %v\n", arr)
	insertionSort(arr)
	fmt.Printf("INSERTION SORT :Sorted array: %v\n", arr)
	//selection sort 
	arr1 := input(8)
	fmt.Printf("SELECTION SORT :Unsorted array: %v\n", arr1)
	selectionSort(arr1)
	fmt.Printf("SELECTION SORT :Sorted array: %v\n", arr1)
	//bubble sort
	arr2 := input(8)
	fmt.Printf("BUBBLE SORT :Unsorted array: %v\n", arr1)
	bubbleSort(arr2)
	fmt.Printf("BUBBLE SORT :Sorted array: %v\n", arr1)

}
