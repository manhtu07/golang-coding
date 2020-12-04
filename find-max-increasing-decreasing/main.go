package main

import "fmt"

func findPivot(arr []int, l, r int) int {
	fmt.Printf("l=%d, r=%d ", l, r)
	if r < l {
		return -1
	}
	if r == l {
		return l
	}
	mid := (l + r) / 2
	fmt.Printf("mid=%d\n", mid)
	if mid < r && arr[mid] > arr[mid+1] {
		return mid
	}
	if mid > l && arr[mid] < arr[mid-1] {
		return mid - 1
	}
	if arr[l] >= arr[mid] {
		return findPivot(arr, l, mid-1)
	}
	return findPivot(arr, mid+1, r)
}

func findMax(arr []int, l, r int) int {
	fmt.Printf("l=%d, r=%d ", l, r)
	if l == r {
		return l
	}
	mid := (l + r) / 2
	fmt.Printf("mid=%d\n", mid)
	if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
		return mid
	}
	if arr[mid] < arr[mid-1] {
		return findMax(arr, l, mid-1)
	} else {
		return findMax(arr, mid+1, r)
	}
}

func main() {
	arr := []int{8, 10, 20, 80, 100, 200, 400, 500, 3, 2, 1}

	i := findMax(arr, 0, len(arr)-1)

	fmt.Println(i)
}
