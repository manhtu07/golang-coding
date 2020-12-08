package main

func find(arr []int) []int  {
	var subArray []int
	var subArrayFar []int
	max, maxFar := 0, 0
	for _, v := range arr {
		max = v + max
		if max < 0 {
			subArray = []int{}
			max = 0
		} else {
			subArray = append(subArray, v)
		}
		if max > maxFar {
			maxFar = max
			subArrayFar = subArray
		}
	}
	return subArrayFar
}

func main()  {
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	find(arr)
}
