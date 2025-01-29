/* Problem Overview:

Given two arrays, arr1[] and arr2[], you need to find the intersection of the two arrays. The intersection of two arrays is defined as the set of elements that appear in both arrays. You need to return an array containing these common elements, without duplicates.

Input:
Two arrays of integers, arr1[] and arr2[].
Both arrays can have any number of elements, and the integers can range from -10^4 to 10^4.
Output:
An array containing the elements that are present in both arrays.
The output array should not contain duplicate elements.

examples :
arr1 = [1, 2, 2, 1]
arr2 = [2, 2]
o/p : [2]

arr1 = [4, 9, 5]
arr2 = [9, 4, 9, 8, 4]
o/p : [9, 4]*/

package basics

import "fmt"

func intersection(arr1, arr2 []int) []int {

    map1 := make(map[int]bool)
    result := []int{} 
    
    for _, num := range arr1 {
        map1[num] = true
    }

    for _, num := range arr2 {
        if _, found := map1[num]; found {
            result = append(result, num)
            delete(map1, num) 
        }
    }

    return result
}

func prob6() {
	arr1 := []int{1, 2, 2, 1}
    arr2 := []int{2, 2}
    fmt.Println(intersection(arr1, arr2)) 

    // Test case 2
    arr1 = []int{4, 9, 5}
    arr2 = []int{9, 4, 9, 8, 4}
    fmt.Println(intersection(arr1, arr2)) 
}

