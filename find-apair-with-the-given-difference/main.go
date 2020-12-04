package main

import (
	"fmt"
	"sort"
)

func FindByBinarySearch(arr []int, l, r, num int) int {
	if l == r {
		if arr[l] == num {
			return l
		}
		return -1
	}
	mid := (l + r) / 2
	if arr[mid] == num {
		return mid
	}
	if arr[mid] <= num {
		return FindByBinarySearch(arr, mid+1, r, num)
	} else {
		return FindByBinarySearch(arr, l, mid-1, num)
	}
}

func main() {
	arr := []int{5, 20, 3, 2, 50, 80}
	n := 78
	// sorting array by ascending order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	a := -1
	b := -1
	for i, value := range arr {
		if i == len(arr)-1 {
			break
		}
		num := value + n
		j := FindByBinarySearch(arr, i+1, len(arr)-1, num)
		if j == -1 {
			continue
		} else {
			b = j
			a = i
			break
		}
	}
	if b == -1 {
		fmt.Println("No Such Pair")
		return
	}
	fmt.Printf("Pair Found: (%d, %d)", arr[a], arr[b])

	c := make(chan string)


	go func() {
		c <- "phong"
	}()

	fmt.Println(<-c)
}
