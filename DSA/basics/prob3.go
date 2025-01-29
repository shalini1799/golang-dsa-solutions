/* Problem overview : Given an array of integers, your task is to check if the array is a palindrome. An array is considered a palindrome if it reads the same from both ends. In other words, for an array arr of size n, the element at index i must be equal to the element at index n - i - 1 for all valid i.

Detailed Problem Description:

You are given an array arr[] of size n (where n is the length of the array).
You need to determine whether the array is a palindrome or not.

Palindrome Definition: An array is a palindrome if:
arr[i]==arr[n−i−1]for all valid i
This means that the first element must match the last element, the second element must match the second last element, and so on.

examples :

arr = [1, 2, 3, 2, 1]
o/p : true

arr = [1, 2, 3, 4, 5]
o/p : false */

package basics

import "fmt"

func isPalindrome(arr []int) bool{
	j := (len(arr)) -1 

	for i:=0; i<len(arr) ; i++ {
		if arr[i] != arr[j] {
			return false
		}
		j--
	}
	return true

}

func prob3() {

	arr1 := []int {1, 2, 3, 2, 1}
	fmt.Printf("%t", isPalindrome(arr1))

	arr2 := []int {1, 2, 3, 4, 5}
	fmt.Printf("%t", isPalindrome(arr2))

}