package main

import (
	"fmt"
)

func waitingTime(tickets []int, p int)  {
	n := len(tickets) - 1
	i := 0
	sum := 0
	for true {
		if i > n {
			i = 0
		}
		if tickets[i] == 0 {
			i += 1
			continue
		}
		tickets[i]--
		sum++
		if i == p && tickets[p] == 0 {
			break
		}
		i++
	}
	fmt.Println(sum)
}

func main() {
	arr := []int{2, 6, 3, 4, 5}
	waitingTime(arr, 4)
}
