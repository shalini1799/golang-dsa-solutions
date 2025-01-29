/*  Find the non-repeating element in an array where every other element appears twice 
Problem Overview:

You are given an array where every element appears twice, except for one element that appears only once. Your task is to find the non-repeating element in the array.

The array is unsorted, and it can contain both positive and negative integers.
The solution should be efficient, ideally O(n) time complexity and O(1) space complexity.

eg : nums = [4, 1, 2, 1, 2]
o/p := 4

*/

package basics

import "fmt"

func findNonRepeatingElement(arr []int) int {
	var elem int
	n := len(arr)

	for i:=0; i<n ; i++ {
		elem ^= arr[i]
	} 
	return elem
}

func prob1 () {

	num := []int{4, 1, 2, 1, 2}
	element := findNonRepeatingElement(num)
	fmt.Printf("the non-repeating element in the given array : %d", element)
}
 

