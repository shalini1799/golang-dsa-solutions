/* Problem : Find the majority element in an array (Element that appears more than n/2 times)

You are given an array of integers, and your task is to find the majority element in the array. The majority element is defined as the element that appears more than n / 2 times in the array, where n is the length of the array.

If such an element exists, return it.
If no majority element exists (i.e., no element appears more than n / 2 times), return -1.

example :

nums = [1, 2, 3, 4, 5]
o/p : -1

nums = [1, 1, 2, 1, 3, 1]
o/p : 1

*/
package basics

import "fmt"

func findMajorityElem(arr []int) int{
	arrMap := make(map[int]int)
	maxOccurance := (len(arr)) / 2

	for i:=0; i<len(arr); i++ {
		arrMap[arr[i]] ++
		if arrMap[arr[i]] >= maxOccurance{
			return arr[i]
		}
	}

	return -1
}

func prob2() {
	nums1 := []int {1, 1, 2, 1, 3, 1}
	majElem1 := findMajorityElem(nums1)
	fmt.Printf("the majority element is :%d", majElem1)

	nums2 := []int {1, 2, 3, 4, 5}
	majElem2 := findMajorityElem(nums2)
	fmt.Printf("the majority element is :%d", majElem2)
}

