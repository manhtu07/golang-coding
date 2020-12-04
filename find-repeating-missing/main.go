package main

import (
	"fmt"
	"sort"
)

// Check next number in the following cases
// - if next number is equal previous number then it's repeating number
// - if next number is different previous number and previous number + 1 then it's missing number
// Time Complexity: O(nlogn)
func findRepeatingMissingWithSorting(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var missing int
	var nextNum int
	var preNum int
	var repeating int
	for i, num := range arr {
		if i+1 < len(arr) {
			nextNum = arr[i+1]
			preNum = num
		} else {
			nextNum = arr[i]
			preNum = arr[i-1]
		}
		if nextNum == preNum {
			repeating = nextNum
		} else if nextNum != preNum && preNum+1 != nextNum {
			missing = preNum + 1
		}
	}

	fmt.Println("missing = ", missing)

	fmt.Println("repeating = ", repeating)

}

// Create a temp array temp[] of size n with all initial values as 0.
// Traverse the input array arr[], and do following for each arr[i]
//		if(temp[arr[i]] == 0) temp[arr[i]] = 1;
//		if(temp[arr[i]] == 1) output “arr[i]” //repeating
// Traverse temp[] and output the array element having value as 0 (This is the missing element)
// Time Complexity: O(n)
func findRepeatingMissingWithTemp(arr []int) {
	temp := make([]int, len(arr))
	var repeating int
	for i, _ := range arr {
		if temp[arr[i]-1] == 0 {
			temp[arr[i]-1] = 1
		} else {
			repeating = arr[i]
		}
	}

	var missing int

	for i, _ := range temp {
		if temp[i] == 0 {
			missing = i + 1
		}
	}

	fmt.Println("missing = ", missing)

	fmt.Println("repeating = ", repeating)
}

func main() {
	arr := []int{4, 3, 6, 2, 1, 1}
	findRepeatingMissingWithSorting(arr)

	findRepeatingMissingWithTemp(arr)
}
