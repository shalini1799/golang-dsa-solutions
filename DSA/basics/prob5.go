/* 
Problem Overview:

You are given an array of integers and you need to check whether the array contains any duplicates. If there are any duplicates, return true; otherwise, return false.

Detailed Problem Description:

Input:
An array of integers arr[] where the length of the array n is between 1 and 10^5.
The integers can be both positive and negative, and they can also include zero.
Output:
Return a boolean value:
true if the array contains at least one duplicate element.
false if all elements in the array are distinct (i.e., no duplicates).

examples :
arr = [1, 2, 3, 4, 5]
o/p : false 

arr = [1, 2, 3, 4, 1]
o/p : true

*/

package basics

import "fmt"

func findDuplicates(arr []int) bool{
	arrMap := make(map[int]bool)
	for i:=0; i<len(arr); i++ {
		if arrMap[arr[i]] {
			return true
		}
		arrMap[arr[i]] = true
	}

	return false

}
func prob5() {
	arr1 := []int{1, 2, 3, 4, 1}
	d := findDuplicates(arr1)
	fmt.Printf("%t",d)

}