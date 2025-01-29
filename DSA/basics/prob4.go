/* Problem Overview:

You are given an array of integers and an integer k. Your task is to rotate the array to the right by k positions. The elements of the array should be moved to the right by k positions, and the elements that go beyond the array's bounds should wrap around to the beginning.

Rotate Right by k positions means:
The element at index i should move to index (i + k) % n in the array, where n is the size of the array.
After rotating, the element at index i moves to index (i + k) % n and elements that go beyond the last index should wrap around to the start.
Detailed Problem Description:

You are given an array arr[] of size n, and an integer k.
You need to rotate the array to the right by k positions.

examples :
arr = [1, 2, 3, 4, 5, 6, 7], k = 3
o/p : [5, 6, 7, 1, 2, 3, 4]

arr = [-1, -100, 3, 99], k = 2
o/p : [3, 99, -1, -100]*/

package basics

import "fmt"

func rotateArr(arr []int, k int) []int{
	n := len(arr)
	// Handle the case where k is larger than n
	k = k % n

	// Create a new slice for the rotated array
	rotated_arr := append([]int{}, arr[n-k:n]...) 
	rotated_arr = append(rotated_arr, arr[0:n-k]...)

	return rotated_arr

}

func prob4() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotated_arr := rotateArr(arr1, k)
	fmt.Printf("rotated array: %v", rotated_arr)

}

/* alternate approach :
For the array [1, 2, 3, 4, 5, 6, 7] and k = 3:

Initial Array: [1, 2, 3, 4, 5, 6, 7]
Reverse the entire array:
After reversing: [7, 6, 5, 4, 3, 2, 1]
Reverse the first k = 3 elements:
After reversing the first 3 elements: [5, 6, 7, 4, 3, 2, 1]
Reverse the remaining n-k = 4 elements:
After reversing the last 4 elements: [5, 6, 7, 1, 2, 3, 4]
The final result is [5, 6, 7, 1, 2, 3, 4], which is the array rotated by 3 positions to the right.
*/